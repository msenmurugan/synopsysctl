/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package components

import (
	"fmt"

	"github.com/blackducksoftware/horizon/pkg/api"

	"github.com/koki/short/types"

	"k8s.io/apimachinery/pkg/api/resource"
)

// Volume defines the volume component
type Volume struct {
	Name string
	obj  *types.Volume
}

// GetObj will return the volume object in a format the deployer can use
func (v *Volume) GetObj() *types.Volume {
	return v.obj
}

// NewEmptyDirVolume creates an EmptyDir volume object
func NewEmptyDirVolume(config api.EmptyDirVolumeConfig) (*Volume, error) {
	size, err := resource.ParseQuantity(config.SizeLimit)
	if err != nil {
		return nil, fmt.Errorf("invalid size: %v", err)
	}

	v := &types.Volume{
		EmptyDir: &types.EmptyDirVolume{
			SizeLimit: &size,
		},
	}

	switch config.Medium {
	case api.StorageMediumDefault:
		v.EmptyDir.Medium = types.StorageMediumDefault
	case api.StorageMediumMemory:
		v.EmptyDir.Medium = types.StorageMediumMemory
	case api.StorageMediumHugePages:
		v.EmptyDir.Medium = types.StorageMediumHugePages
	}

	return &Volume{Name: config.VolumeName, obj: v}, nil
}

// NewHostPathVolume create a HostPath volume object
func NewHostPathVolume(config api.HostPathVolumeConfig) *Volume {
	v := &types.Volume{
		HostPath: &types.HostPathVolume{
			Path: config.Path,
		},
	}
	switch config.Type {
	case api.HostPathUnset:
		v.HostPath.Type = types.HostPathUnset
	case api.HostPathDirectoryOrCreate:
		v.HostPath.Type = types.HostPathDirectoryOrCreate
	case api.HostPathDirectory:
		v.HostPath.Type = types.HostPathDirectory
	case api.HostPathFileOrCreate:
		v.HostPath.Type = types.HostPathFileOrCreate
	case api.HostPathFile:
		v.HostPath.Type = types.HostPathFile
	case api.HostPathSocket:
		v.HostPath.Type = types.HostPathSocket
	case api.HostPathCharDev:
		v.HostPath.Type = types.HostPathCharDev
	case api.HostPathBlockDev:
		v.HostPath.Type = types.HostPathBlockDev
	}

	return &Volume{Name: config.VolumeName, obj: v}
}

// NewConfigMapVolume creates a ConfigMap volume object
func NewConfigMapVolume(config api.ConfigMapOrSecretVolumeConfig) *Volume {
	dfm := types.FileMode(config.DefaultMode)
	items := map[string]types.KeyAndMode{}
	for k, v := range config.Items {
		fm := types.FileMode(v.Mode)
		items[k] = types.KeyAndMode{
			Key:  v.KeyOrPath,
			Mode: &fm,
		}
	}

	v := &types.Volume{
		ConfigMap: &types.ConfigMapVolume{
			Name:        config.MapOrSecretName,
			DefaultMode: &dfm,
			Items:       items,
			Required:    &config.Required,
		},
	}

	return &Volume{Name: config.VolumeName, obj: v}
}

// NewSecretVolume creates a Secret volume object
func NewSecretVolume(config api.ConfigMapOrSecretVolumeConfig) *Volume {
	dfm := types.FileMode(config.DefaultMode)
	items := map[string]types.KeyAndMode{}
	for k, v := range config.Items {
		fm := types.FileMode(v.Mode)
		items[k] = types.KeyAndMode{
			Key:  v.KeyOrPath,
			Mode: &fm,
		}
	}

	v := &types.Volume{
		Secret: &types.SecretVolume{
			SecretName:  config.MapOrSecretName,
			DefaultMode: &dfm,
			Items:       items,
			Required:    &config.Required,
		},
	}

	return &Volume{Name: config.VolumeName, obj: v}
}

// NewGCEPersistentDiskVolume creates a new GCE Persistent Disk volume object
func NewGCEPersistentDiskVolume(config api.GCEPersistentDiskConfig) *Volume {
	v := &types.Volume{
		GcePD: &types.GcePDVolume{
			PDName:    config.DiskName,
			FSType:    config.FSType,
			Partition: config.Partition,
			ReadOnly:  config.ReadOnly,
		},
	}

	return &Volume{Name: config.VolumeName, obj: v}
}
