//go:build windows
// +build windows

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//      http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package manageddaemon

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cihub/seelog"
)

const (

	// This is the image tag that will be used for the ManagedDaemon image
	imageTagDefault = "latest"
	// These two constants store the names of the different mounts that will be used for the Windows ManagedDaemon
	defaultAgentCommunicationMount = "agentCommunicationMount"
	defaultApplicationLogMount     = "applicationLogMount"
	// The following two host and container paths are for communicating with the CSIProxy disk named pipe for the CSIDriverManaged Daemon
	csiProxyDiskMountHostPath      = "\\\\.\\pipe\\csi-proxy-disk-v1"
	csiProxyDiskMountContainerPath = "\\\\.\\pipe\\csi-proxy-disk-v1"
	// The following two host and container paths are for communicating with the CSIProxy filesystem named pipe for the CSIDriverManaged Daemon
	csiProxyFileSystemMountHostPath      = "\\\\.\\pipe\\csi-proxy-filesystem-v1"
	csiProxyFileSystemMountContainerPath = "\\\\.\\pipe\\csi-proxy-filesystem-v1"
	// The following two host and container paths are for communicating with the CSIProxy disk named pipe for the CSIDriverManaged Daemon
	csiProxyVolumeMountHostPath      = "\\\\.\\pipe\\csi-proxy-volume-v1"
	csiProxyVolumeMountContainerPath = "\\\\.\\pipe\\csi-proxy-volume-v1"
	// This is the path to the CSIDriver listener socket. The agent communicates with the CSIDriver ManagedDaemon on this path
	csiDriverListenerSocketPath = "unix://ebs-csi-driver/csi-driver.sock"
)

var (
	envProgramFiles = defaultIfBlank(os.Getenv("ProgramFiles"), `C:\Program Files`)
	envProgramData  = defaultIfBlank(os.Getenv("ProgramData"), `C:\ProgramData`)

	AmazonProgramFiles = filepath.Join(envProgramFiles, "Amazon")
	AmazonProgramData  = filepath.Join(envProgramData, "Amazon")

	AmazonECSProgramFiles = filepath.Join(envProgramFiles, "Amazon", "ECS")
	AmazonECSProgramData  = filepath.Join(AmazonProgramData, "ECS")

	CsiDriverContainerWorkingPath    = "C:\\csi-driver\\"
	CsiDriverSocketCommunicationPath = "C:\\ebs-csi-driver\\"
	// This is the path to where the CSIDriver ManagedDaemon image tar file is
	imageTarPath = filepath.Join(AmazonECSProgramData, "images")

	// The following two paths are used for agent communication and log paths to be shared with the ManagedDaemon
	defaultAgentCommunicationPathHostRoot = AmazonECSProgramData
	defaultApplicationLogPathHostRoot     = filepath.Join(AmazonECSProgramData, "log")

	// The following two host and container paths are for the agent communication with the CSIDriver ManagedDaemon
	agentCommunicationMountHostPath      = filepath.Join(AmazonECSProgramData, "ebs-csi-driver")
	agentCommunicationMountContainerPath = CsiDriverSocketCommunicationPath

	// The following two host and container paths are for the log storage for the CSIDriverManaged Daemon
	csiDriverLogMountHostPath      = filepath.Join(defaultApplicationLogPathHostRoot, "daemons")
	csiDriverLogMountContainerPath = filepath.Join(CsiDriverContainerWorkingPath, "log")

	// The following two host and container paths are for the shared mount for the CSIDriverManaged Daemon.
	sharedMountsHostPath      = filepath.Join(AmazonECSProgramData, "ebs")
	sharedMountsContainerPath = filepath.Join(CsiDriverContainerWorkingPath, "ebs")

	// This is the path to the log file that is generated by the CSIDriver ManagedDaemon
	csiDriverLogFilePath = filepath.Join(csiDriverLogMountContainerPath, "csi.log")
)

// defaultImportAll function will parse/validate all managed daemon definitions
// defined and will return an array of valid ManagedDaemon objects
func defaultImportAll() ([]*ManagedDaemon, error) {
	ebsCsiTarFile := filepath.Join(imageTarPath, EbsCsiDriver, imageFileName)
	if _, err := os.Stat(ebsCsiTarFile); err != nil {
		seelog.Warnf("Unable to locate the managed daemon image tar file at: %s", ebsCsiTarFile)
		return []*ManagedDaemon{}, nil
	}
	// locate the EBS CSI tar file -- import
	ebsManagedDaemon := NewManagedDaemon(EbsCsiDriver, "latest")
	// add required mounts
	ebsMounts := []*MountPoint{
		&MountPoint{
			SourceVolumeID:       defaultAgentCommunicationMount,
			SourceVolume:         defaultAgentCommunicationMount,
			SourceVolumeType:     "host",
			SourceVolumeHostPath: agentCommunicationMountHostPath,
			ContainerPath:        agentCommunicationMountContainerPath,
		},
		&MountPoint{
			SourceVolumeID:       defaultApplicationLogMount,
			SourceVolume:         defaultApplicationLogMount,
			SourceVolumeType:     "host",
			SourceVolumeHostPath: csiDriverLogMountHostPath,
			ContainerPath:        csiDriverLogMountContainerPath,
		},
		&MountPoint{
			SourceVolumeID:       "sharedMounts",
			SourceVolume:         "sharedMounts",
			SourceVolumeType:     "host",
			SourceVolumeHostPath: sharedMountsHostPath,
			ContainerPath:        sharedMountsContainerPath,
			PropagationShared:    false,
		},
		// the following three mount points are for connecting to the CSIProxy server that is running
		// as a Windows Service on the host
		&MountPoint{
			SourceVolumeID:       "csiProxyDiskMount",
			SourceVolume:         "csiProxyDiskMount",
			SourceVolumeType:     "npipe",
			SourceVolumeHostPath: csiProxyDiskMountHostPath,
			ContainerPath:        csiProxyDiskMountContainerPath,
			PropagationShared:    false,
		},
		&MountPoint{
			SourceVolumeID:       "csiProxyFileSystemMount",
			SourceVolume:         "csiProxyFileSystemMount",
			SourceVolumeType:     "npipe",
			SourceVolumeHostPath: csiProxyFileSystemMountHostPath,
			ContainerPath:        csiProxyFileSystemMountContainerPath,
			PropagationShared:    false,
		},
		&MountPoint{
			SourceVolumeID:       "csiProxyVolumeMount",
			SourceVolume:         "csiProxyVolumeMount",
			SourceVolumeType:     "npipe",
			SourceVolumeHostPath: csiProxyVolumeMountHostPath,
			ContainerPath:        csiProxyVolumeMountContainerPath,
			PropagationShared:    false,
		},
	}
	if err := ebsManagedDaemon.SetMountPoints(ebsMounts); err != nil {
		return nil, fmt.Errorf("Unable to import EBS ManagedDaemon: %s", err)
	}
	var thisCommand []string
	thisCommand = append(thisCommand, fmt.Sprintf("-endpoint=%s", csiDriverListenerSocketPath))
	thisCommand = append(thisCommand, fmt.Sprintf("-log_file=%s", csiDriverLogFilePath))
	thisCommand = append(thisCommand, "-log_file_max_size=20")
	thisCommand = append(thisCommand, "-logtostderr=false")

	ebsManagedDaemon.command = thisCommand
	ebsManagedDaemon.privileged = false
	return []*ManagedDaemon{ebsManagedDaemon}, nil
}

func defaultIfBlank(str string, default_value string) string {
	if len(str) == 0 {
		return default_value
	}
	return str
}
