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

package main

import (
	"eva"
	"eva/agent"
	"eva/manager/memory"
	"eva/manager/sql"
	"eva/mock"
	"eva/policy"
	"eva/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"fmt"
)

var hostname string
var warden *eva.Eva00

func main() {
	iamInit()
	router := gin.Default()

	router.GET("/hi", greeting)
	router.POST("/evaluation", auth)
	router.POST("/policy", createPolicy)
	router.GET("/policy", getPolicy)

	router.Run()
}

func iamInit() {
	hostname, _ = os.Hostname()

	warden = &eva.Eva00{
		Manager: sqlInit(),
	}

	//for _, pol := range mock.Policies {
	//	warden.Manager.Create(pol)
	//}
	pa := agent.NewPolAgent(mock.Jps)
	a, err := pa.NormalizePolicies()
	if err != nil {
		fmt.Println(err)
	}
	for _, pol := range a {
		warden.Manager.Create(pol)
	}
}

func memoryInit() *memory.MemoryManager {
	return memory.NewMemoryManager()
}

func sqlInit() *sql.PgSqlManager {
	return sql.NewPgSqlManager()
}

func auth(c *gin.Context) {
	rag := agent.NewReqAgent()
	if err := c.ShouldBindJSON(rag.Payload()); err == nil {
		keys, rcs := rag.NormalizeRequests()
		err := warden.Authorize(rcs, keys)
		if err != nil {
			switch e := err.(type) {

			case utils.ErrorWithContext:
				c.JSON(e.Code(), gin.H{
					"type":    e.Causer(),
					"status":  e.Error(),
					"source":  e.Source(),
					"decider": e.Decider(),
					"from":    hostname,
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Allow", "from": hostname})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getPolicy(c *gin.Context) {
	policies, err := warden.Manager.GetAll(100, 0)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"policies": policies})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func createPolicy(c *gin.Context) {
	dp := &policy.DefaultPolicy{}

	if err := c.ShouldBindJSON(dp); err == nil {
		warden.Manager.Create(dp)
		c.JSON(http.StatusOK, gin.H{"status": "create successfully", "from": hostname})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func greeting(c *gin.Context) {
	c.String(http.StatusOK, "Greetings! This is from %s \n", hostname)
}
