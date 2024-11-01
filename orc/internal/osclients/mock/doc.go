/*
Copyright 2021 The Kubernetes Authors.

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

package mock

import (
	// Runtime dependency of mockgen, required when using vendoring so go mod knows
	// to pull it in.
	_ "go.uber.org/mock/mockgen/model"
)

//go:generate mockgen -package mock -destination=compute.go -source=../compute.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock ComputeClient
//go:generate /usr/bin/env bash -c "cat ../../../hack/boilerplate.go.txt compute.go > _compute.go && mv _compute.go compute.go"

//go:generate mockgen -package mock -destination=image.go -source=../image.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock ImageClient
//go:generate /usr/bin/env bash -c "cat ../../../hack/boilerplate.go.txt image.go > _image.go && mv _image.go image.go"

//go:generate mockgen -package mock -destination=loadbalancer.go -source=../loadbalancer.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock LbClient
//go:generate /usr/bin/env bash -c "cat ../../../hack/boilerplate.go.txt loadbalancer.go > _loadbalancer.go && mv _loadbalancer.go loadbalancer.go"

//go:generate mockgen -package mock -destination=networking.go -source=../networking.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock NetworkClient
//go:generate /usr/bin/env bash -c "cat ../../../hack/boilerplate.go.txt networking.go > _networking.go && mv _networking.go networking.go"

//go:generate mockgen -package mock -destination=volume.go -source=../volume.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock VolumeClient
//go:generate /usr/bin/env bash -c "cat ../../../hack/boilerplate.go.txt volume.go > _volume.go && mv _volume.go volume.go"