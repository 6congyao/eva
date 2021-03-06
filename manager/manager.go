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

package manager

import (
	"eva/policy"
)

// Manager is responsible for managing and persisting policies.
type Manager interface {

	// Create persists the policy.
	Create(policy policy.Policy) error

	// Update updates an existing policy.
	Update(policy policy.Policy) error

	// Get retrieves a policy.
	Get(id string) (policy.Policy, error)

	// Delete removes a policy.
	Delete(id string) error

	// GetAll retrieves all policies.
	GetAll(limit, offset int64) (policy.Policies, error)

	// Find related policies by entity keys
	FindCandidates(keys []string)(policy.Policies, error)
}
