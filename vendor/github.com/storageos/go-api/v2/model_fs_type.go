/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.4.0-alpha
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// FsType The file system type of a volume. \"block\" is a raw block device (no filesystem). 
type FsType string

// List of FsType
const (
	FSTYPE_EXT2 FsType = "ext2"
	FSTYPE_EXT3 FsType = "ext3"
	FSTYPE_EXT4 FsType = "ext4"
	FSTYPE_XFS FsType = "xfs"
	FSTYPE_BTRFS FsType = "btrfs"
	FSTYPE_BLOCK FsType = "block"
)
