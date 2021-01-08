/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	nsdelete "github.com/storageos/api-manager/controllers/namespace-delete"
	nodedelete "github.com/storageos/api-manager/controllers/node-delete"
	"github.com/storageos/api-manager/internal/controllers/sharedvolume"
	"github.com/storageos/api-manager/internal/pkg/storageos"
	apimetrics "github.com/storageos/api-manager/internal/pkg/storageos/metrics"
	// +kubebuilder:scaffold:imports
)

const (
	// EventSourceName is added to Kubernetes events generated by the api
	// manager.  It can be used for filtering events.
	EventSourceName = "storageos-api-manager"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("api-manager")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var apiSecretPath string
	var apiEndpoint string
	var apiPollInterval time.Duration
	var apiRefreshInterval time.Duration
	var apiRetryInterval time.Duration
	var cacheExpiryInterval time.Duration
	var k8sCreatePollInterval time.Duration
	var k8sCreateWaitDuration time.Duration
	var gcNamespaceDeleteInterval time.Duration
	var gcNodeDeleteInterval time.Duration
	var nsDeleteWorkers int
	var nodeDeleteWorkers int
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&apiSecretPath, "api-secret-path", "/etc/storageos/secrets/api", "Path where the StorageOS api secret is mounted.  The secret must have \"username\" and \"password\" set.")
	flag.StringVar(&apiEndpoint, "api-endpoint", "storageos", "The StorageOS api endpoint address.")
	flag.DurationVar(&apiPollInterval, "api-poll-interval", 5*time.Second, "Frequency of StorageOS api polling.")
	flag.DurationVar(&apiRefreshInterval, "api-refresh-interval", time.Minute, "Frequency of StorageOS api authentication token refresh.")
	flag.DurationVar(&apiRetryInterval, "api-retry-interval", 5*time.Second, "Frequency of StorageOS api retries on failure.")
	flag.DurationVar(&cacheExpiryInterval, "cache-expiry-interval", time.Minute, "Frequency of cached volume re-validation.")
	flag.DurationVar(&k8sCreatePollInterval, "k8s-create-poll-interval", 1*time.Second, "Frequency of Kubernetes api polling for new objects to appear once created.")
	flag.DurationVar(&k8sCreateWaitDuration, "k8s-create-wait-duration", 20*time.Second, "Maximum time to wait for new Kubernetes objects to appear.")
	flag.DurationVar(&gcNamespaceDeleteInterval, "namespace-delete-gc-interval", 1*time.Hour, "Frequency of node garbage collection.")
	flag.DurationVar(&gcNodeDeleteInterval, "node-delete-gc-interval", 1*time.Hour, "Frequency of node garbage collection.")
	flag.IntVar(&nodeDeleteWorkers, "node-delete-workers", 5, "Maximum concurrent node delete operations.")
	flag.IntVar(&nsDeleteWorkers, "namespace-delete-workers", 5, "Maximum concurrent namespace delete operations.")

	flag.Parse()

	ctrl.SetLogger(zap.New(zap.StacktraceLevel(zapcore.FatalLevel)))

	// Block startup until there is a working StorageOS API connection.  Unless
	// we loop here, we'll get a number of failures on cold cluster start as it
	// takes longer for the api to be ready than the api-manager to start.
	var api *storageos.Client
	var err error
	for {
		username, password, err := storageos.ReadCredsFromMountedSecret(apiSecretPath)
		if err != nil {
			setupLog.Info(fmt.Sprintf("unable to read storageos api secret, retrying in %s", apiRetryInterval), "msg", err)
			apimetrics.Errors.Increment("setup", err)
			time.Sleep(apiRetryInterval)
			continue
		}
		api, err = storageos.NewTracedClient(username, password, apiEndpoint)
		if err == nil {
			apimetrics.Errors.Increment("setup", nil)
			break
		}
		setupLog.Info(fmt.Sprintf("unable to connect to storageos api, retrying in %s", apiRetryInterval), "msg", err)
		apimetrics.Errors.Increment("setup", storageos.GetAPIErrorRootCause(err))
		time.Sleep(apiRetryInterval)
	}
	setupLog.V(1).Info("connected to the storageos api")

	// Only attempt to grab leader lock once we have an API connection.
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "d73494fd.storageos.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	// Events sent on apiReset channel will trigger the api client to re-initialise.
	apiReset := make(chan struct{})

	// Events sent on errCh will trigger a graceful shutdown.  Another instance
	// will take over while the pod restarts.
	errCh := make(chan error, 3)

	// Parent context will be closed on interrupt or sigterm.
	ctx, cancel := context.WithCancel(ctrl.SetupSignalHandler())

	var wg sync.WaitGroup

	shutdown := func() {
		// Stop other goroutines.
		cancel()
		// Wait for the other goroutines to finish.
		wg.Wait()
		setupLog.Info("shutdown complete")
		os.Exit(0)
	}

	// Goroutine to handle api credential refreshes and client reconnects
	// whenever events are received on the apiReset channel.
	wg.Add(1)
	go func() {
		err := api.Refresh(ctx, apiSecretPath, apiEndpoint, apiReset, apiRefreshInterval, apimetrics.Errors, setupLog)
		errCh <- fmt.Errorf("api token refresh error: %w", err)
		wg.Done()
	}()

	// Goroutine to run the kubebuilder manager.
	wg.Add(1)
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			errCh <- fmt.Errorf("manager error: %w", err)
		} else {
			errCh <- fmt.Errorf("manager stopped")
		}
		wg.Done()
	}()

	// Wait for this instance to be elected leader.
	select {
	case <-mgr.Elected():
	case err := <-errCh:
		setupLog.Info("exiting", "reason", err)
		shutdown()
	}

	// Goroutine to poll StorageOS api for shared volumes and create/update
	// Kubernetes services as needed.
	wg.Add(1)
	go func() {
		setupLog.Info("starting shared volume controller ")
		err := sharedvolume.NewReconciler(api, apiReset, mgr.GetClient(), mgr.GetEventRecorderFor(EventSourceName)).Reconcile(ctx, apiPollInterval, cacheExpiryInterval, k8sCreatePollInterval, k8sCreateWaitDuration)
		errCh <- fmt.Errorf("shared volume reconciler error: %w", err)
		wg.Done()
	}()

	// Additional controllers go here.
	setupLog.Info("starting node delete controller ")
	if err := nodedelete.NewReconciler(api, mgr.GetClient(), gcNodeDeleteInterval).SetupWithManager(mgr, nodeDeleteWorkers); err != nil {
		errCh <- fmt.Errorf("node delete reconciler error: %w", err)
	}
	setupLog.Info("starting namespace delete controller ")
	if err := nsdelete.NewReconciler(api, mgr.GetClient(), gcNamespaceDeleteInterval).SetupWithManager(mgr, nsDeleteWorkers); err != nil {
		errCh <- fmt.Errorf("namespace delete reconciler error: %w", err)
	}

	// Wait until a goroutine sends an error or a shutdown signal received, then
	// cancel the others.
	select {
	case err = <-errCh:
		setupLog.Info("exiting", "reason", err)
	case <-ctx.Done():
		setupLog.Info("shutdown requested")
	}

	shutdown()
}
