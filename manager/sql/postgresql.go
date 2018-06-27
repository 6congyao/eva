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
	"eva/agent"
	"eva/policy"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

type dbPolicy struct {
	Id        string         `db:"policy_id"`
	Key       string         `db:"entity_urn"`
	Statement types.JSONText `db:"statement"`
}

// PgSqlManager is a postgres implementation for Manager to fetch policies persistently.
type PgSqlManager struct {
	db *sqlx.DB
}

// NewPgSqlManager initializes a new SQLManager for given db instance.
func NewPgSqlManager(db *sqlx.DB) *PgSqlManager {
	return &PgSqlManager{
		db: db,
	}
}

// Create a new policy to NewPgSqlManager.
func (m *PgSqlManager) Create(policy policy.Policy) error {
	//tx, err := m.db.Beginx()
	//
	//if err != nil {
	//	return err
	//}
	//
	//
	//if err = tx.Commit(); err != nil {
	//	return err
	//}
	//tx := m.db.MustBegin()
	//
	//for _, v := range policies {
	//	tx.MustExec("INSERT INTO iam_policy (statement) VALUES ($1)", v)
	//}
	//
	//tx.Commit()

	return nil
}

func (m *PgSqlManager) FindCandidates(keys []string) (policy.Policies, error) {
	query, args, err := sqlx.In(findCandidatesQuery, keys)
	if err != nil {
		return nil, err
	}
	query = m.db.Rebind(query)
	rows, err := m.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	if rows != nil {
		defer rows.Close()
	}
	return scanRows(rows)
}

// Get retrieves a policy.
func (m *PgSqlManager) Get(id string) (policy.Policy, error) {
	query := m.db.Rebind(getOneQuery)
	row, err := m.db.Queryx(query,id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	data, err := scanRows(row)
	if len(data)==0{
		return nil ,err
	}
	return data[0], err
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
	query := m.db.Rebind(getAllQuery)

	rows, err := m.db.Queryx(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRows(rows)
}

func scanRows(rows *sqlx.Rows) (policy.Policies, error) {
	var ps [][]byte = nil

	for rows.Next() {
		var dp dbPolicy
		if err := rows.StructScan(&dp); err != nil {
			return nil, err
		}

		ps = append(ps, dp.Statement)

		//		if len(dp.Key) > 1 {
		////todo: cache
		//		}
	}

	pa := agent.NewPolAgent(ps)

	return pa.NormalizePolicies()
}
