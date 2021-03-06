/*
Copyright 2017 The Kubernetes Authors.

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

package types

import (
	"github.com/kubernetes-incubator/external-storage/local-volume/provisioner/pkg/cache"
	"github.com/kubernetes-incubator/external-storage/local-volume/provisioner/pkg/util"
	"k8s.io/client-go/kubernetes"
)

const (
	AnnProvisionedBy = "pv.kubernetes.io/provisioned-by"
)

type UserConfig struct {
	// Name of the node this is running on
	// TODO: Need to resolve nodename vs hostname
	NodeName string
	// The hostpath directory
	HostDir string
	// The mount point of the hostpath volume
	MountDir string
	// key = storageclass, value = relative directory to search in
	DiscoveryMap map[string]string
}

type RuntimeConfig struct {
	*UserConfig
	// Unique name of this provisioner
	Name string
	// K8s API client
	Client *kubernetes.Clientset
	// Cache to store PVs managed by this provisioner
	Cache *cache.VolumeCache
	// K8s API layer
	APIUtil util.APIUtil
	// Volume util layer
	VolUtil util.VolumeUtil
}
