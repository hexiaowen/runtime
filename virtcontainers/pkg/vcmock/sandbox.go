// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vcmock

import (
	vc "github.com/kata-containers/runtime/virtcontainers"
)

// ID implements the VCSandbox function of the same name.
func (p *Sandbox) ID() string {
	return p.MockID
}

// Annotations implements the VCSandbox function of the same name.
func (p *Sandbox) Annotations(key string) (string, error) {
	return p.MockAnnotations[key], nil
}

// SetAnnotations implements the VCSandbox function of the same name.
func (p *Sandbox) SetAnnotations(annotations map[string]string) error {
	return nil
}

// GetAnnotations implements the VCSandbox function of the same name.
func (p *Sandbox) GetAnnotations() map[string]string {
	return p.MockAnnotations
}

// GetAllContainers implements the VCSandbox function of the same name.
func (p *Sandbox) GetAllContainers() []vc.VCContainer {
	var ifa = make([]vc.VCContainer, len(p.MockContainers))

	for i, v := range p.MockContainers {
		ifa[i] = v
	}

	return ifa
}

// GetContainer implements the VCSandbox function of the same name.
func (p *Sandbox) GetContainer(containerID string) vc.VCContainer {
	for _, c := range p.MockContainers {
		if c.MockID == containerID {
			return c
		}
	}
	return &Container{}
}