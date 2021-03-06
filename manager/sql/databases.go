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

var Schema = `
CREATE TABLE IF NOT EXISTS iam_policy (
  policy_id    VARCHAR(50)                                  NOT NULL
    CONSTRAINT iam_policy_pkey
    PRIMARY KEY,
  policy_name  VARCHAR(255) DEFAULT '' :: CHARACTER VARYING NOT NULL,
  description  TEXT,
  qrn          VARCHAR(255)                                 NOT NULL,
  path         VARCHAR(255) DEFAULT '' :: CHARACTER VARYING NOT NULL,
  statement    TEXT                                         NOT NULL,
  version      VARCHAR(255)                                 NOT NULL,
  create_time  TIMESTAMP DEFAULT now()                      NOT NULL,
  root_user_id VARCHAR(255) DEFAULT '' :: CHARACTER VARYING NOT NULL
);

CREATE TABLE IF NOT EXISTS policy_binding (
  resource_id     VARCHAR(50)                                 NOT NULL,
  entity_urn      VARCHAR(255)                                NOT NULL,
  binding_context VARCHAR(255)                                NOT NULL,
  policy_id       VARCHAR(255)                                NOT NULL,
  resource_type   VARCHAR(50) DEFAULT '' :: CHARACTER VARYING NOT NULL,
  CONSTRAINT policy_binding_pkey
  PRIMARY KEY (policy_id, entity_urn)
);

CREATE INDEX policy_binding_resource_id
  ON policy_binding (resource_id);

CREATE INDEX policy_binding_entity_urn
  ON policy_binding (entity_urn);

CREATE INDEX policy_binding_policy_id
  ON policy_binding (policy_id);`

var findCandidatesQuery = `SELECT 
		p.policy_id, statement, entity_urn
	FROM 
		iam_policy AS p
	INNER JOIN 
		policy_binding AS b
	ON 
		p.policy_id = b.policy_id
	WHERE 
		entity_urn 
	IN (?)`

var getAllQuery = `SELECT
		policy_id, statement
	FROM
		iam_policy LIMIT ? OFFSET ?`

var getOneQuery= `SELECT
		policy_id, statement
	FROM
		iam_policy 
	WHERE
		policy_id= ?`