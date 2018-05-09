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

package sql

import (
	"eva/policy"
)

// SQLManager is a postgres implementation for Manager to store policies persistently.
type PgSqlManager struct {
	dbstr string
}

// NewSQLManager initializes a new SQLManager for given db instance.
func NewPgSqlManager() *PgSqlManager {

	return nil
}

// Create a new policy to SQLManager.
func (m *PgSqlManager) Create(policy policy.Policy) error {

	return nil
}

func (m *PgSqlManager) FindCandidates(keys []string) (policy.Policies, error) {


	return nil, nil
}

// Get retrieves a policy.
func (m *PgSqlManager) Get(id string) (policy.Policy, error) {

	return nil, nil
}

// Delete removes a policy.
func (m *PgSqlManager) Delete(id string) error {

	return nil
}

// Update updates an existing policy.
func (m *PgSqlManager) Update(policy policy.Policy) error {

	return nil
}

// GetAll returns all policies.
func (m *PgSqlManager) GetAll(limit, offset int64) (policy.Policies, error) {

	return nil, nil
}

