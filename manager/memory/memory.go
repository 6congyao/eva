/*******************************************************************************
 * Copyright (c) 2018. LuCongyao <6congyao@gmail.com> .
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 ******************************************************************************/

package memory

import (
	"eva/policy"
	"sync"
)

// MemoryManager is an in-memory (non-persistent) implementation of Manager.
type MemoryManager struct {
	Policies map[string]policy.Policies
	sync.RWMutex
}

// NewMemoryManager constructs and initializes new MemoryManager with no policies.
func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		Policies: map[string]policy.Policies{},
	}
}

// Create a new policy to MemoryManager.
func (m *MemoryManager) Create(policy policy.Policy) error {
	m.Lock()
	defer m.Unlock()

	m.Policies[policy.GetID()] = append(m.Policies[policy.GetID()], policy)

	return nil
}

func (m *MemoryManager) FindCandidates(keys []string) (policy.Policies, error) {
	m.RLock()
	defer m.RUnlock()
	var ps policy.Policies
	for _, key := range keys {
		if v, found := m.Policies[key]; found {
			ps = append(ps, v...)
		}
	}

	return ps, nil
}

// Get retrieves a policy.
func (m *MemoryManager) Get(id string) (policy.Policy, error) {
	m.RLock()
	defer m.RUnlock()

	return nil, nil
}

// Delete removes a policy.
func (m *MemoryManager) Delete(id string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.Policies, id)
	return nil
}

// Update updates an existing policy.
func (m *MemoryManager) Update(policy policy.Policy) error {
	m.Lock()
	defer m.Unlock()
	m.Policies[policy.GetID()] = nil
	return nil
}

// GetAll returns all policies.
func (m *MemoryManager) GetAll(limit, offset int64) (policy.Policies, error) {
	ps := policy.Policies{}

	for _, mp := range m.Policies {
		for _, p := range mp {
			ps = append(ps, p)
		}
	}

	return ps, nil
}
