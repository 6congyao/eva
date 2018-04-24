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
	"eva/manager/memory"
	"eva/mock"
	"eva/policy"
	"eva/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"eva/agent/json"
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
		Manager: memoryInit(),
		Agent: agentInit(),
	}

	for _, pol := range mock.Polices {
		warden.Manager.Create(pol)
	}
}

func memoryInit() *memory.MemoryManager {
	return memory.NewMemoryManager()
}

func agentInit() *json.JsonAgent {
	return json.NewJsonAgent()
}

func auth(c *gin.Context) {

	if err := c.ShouldBindJSON(warden.Agent.Payload()); err == nil {
		keys, rcs, _ := warden.Agent.NormalizeRequests()

		err := warden.Authorize(rcs, keys)
		if err != nil {
			switch e := err.(type) {

			case utils.ErrorWithContext:
				c.JSON(e.StatusCode(), gin.H{"type": e.Details(), "status": e.Error(), "from": hostname})
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
	polices, err := warden.Manager.GetAll(100, 0)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"polices": polices})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func createPolicy(c *gin.Context) {
	jp := &policy.DefaultPolicy{}

	if err := c.ShouldBindJSON(jp); err == nil {
		warden.Manager.Create(jp)
		c.JSON(http.StatusOK, gin.H{"status": "create successfully", "from": hostname})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func greeting(c *gin.Context) {
	c.String(http.StatusOK, "Greetings! This is from %s \n", hostname)
}
