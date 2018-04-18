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
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"eva/policy"
)

var hostname string
var warden *eva.Eva00

// Binding from JSON
type AuthRequestInput struct {
	Principal string `json:"principal"`
	// Resource is the resource that access is requested to.
	Resource string `json:"resource" binding:"required"`

	// Action is the action that is requested on the resource.
	Action string `json:"action" binding:"required"`

	// Subejct is the subject that is requesting access.
	Subject []string `json:"subject" binding:"required"`
}

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
	}

	for _, pol := range mock.Polices {
		warden.Manager.Create(pol)
	}
}

func memoryInit() *memory.MemoryManager {
	return memory.NewMemoryManager()
}

func auth(c *gin.Context) {
	json := &AuthRequestInput{}

	if err := c.ShouldBindJSON(json); err == nil {
		requests := []*eva.RequestContext{}

		request := &eva.RequestContext{
			Principal: json.Principal,
			Action:    json.Action,
			Resource:  json.Resource,
		}
		requests = append(requests, request)

		err := warden.Authorize(requests, json.Subject)
		if err != nil {
			//ret := errors.Cause(err)
			//ret2 := ladon.ErrRequestDenied
			//if ret == ret2 {
			//	fmt.Printf("Type: %T\n", errors.Cause(err))
			//}
			//switch et := errors.Cause(err).(type) {
			//case *ladon.errorWithContext:
			//	// handle specifically
			c.JSON(http.StatusForbidden, gin.H{"status": err.Error(), "from": hostname})
			//default:
			//	// unknown error
			//	c.JSON(http.StatusForbidden, gin.H{"status": et})
			//}

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
