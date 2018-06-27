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
	"eva/policy"
	"eva/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var hostname string
var warden eva.Eva
var ready = false
var manager *sql.PgSqlManager

func main() {
	evaInit()
	routerInit()
}

func evaInit() {
	hostname, _ = os.Hostname()

	//warden = &eva.Eva00{
	//	Manager: sqlInit(),
	//}

	manager = sqlInit()
	warden = eva.NewEva(manager)

	//for _, pol := range mock.Policies {
	//	warden.Manager.Create(pol)
	//}

	//pa := agent.NewPolAgent(mock.Jps)
	//a, err := pa.NormalizePolicies()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, pol := range a {
	//	warden.Manager.Create(pol)
	//}
}

func sqlInit() *sql.PgSqlManager {
	dbDrv := os.Getenv(eva.EnvDBDriver)
	dbSrc := os.Getenv(eva.EnvDBSource)

	db, err := sqlx.Connect(dbDrv, dbSrc)

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	return sql.NewPgSqlManager(db)
}

func routerInit() {
	router := gin.Default()

	router.GET("/healthy", healthy)
	router.POST("/evaluation", auth)

	// For testing proposal
	router.POST("/policy", createPolicy)
	router.GET("/policy", getPolicy)

	ready = true
	router.Run()
}

func auth(c *gin.Context) {
	rag := agent.NewReqAgent()
	if err := c.ShouldBindJSON(rag.Payload()); err == nil {
		keys, rcs, _ := rag.NormalizeRequests()

		err := warden.Authorize(rcs, keys)
		if err != nil {
			switch e := err.(type) {

			case utils.ErrorWithContext:
				fmt.Fprintf(os.Stdout, "[EVA] ReqID: %s | Status: Deny | Type: %s | Source: %#v | Decider: %#v \n",
					rag.RequestInput.Id, e.Causer(), e.Source(), e.Decider())
				c.JSON(e.Code(), gin.H{
					"type":    e.Causer(),
					"status":  e.Error(),
					"source":  e.Source(),
					"decider": e.Decider(),
					"from":    hostname,
				})
			default:
				fmt.Fprintf(os.Stdout, "[EVA] ReqID: %s | Status: Deny | Error: %s \n", rag.RequestInput.Id, err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

		} else {
			fmt.Fprintf(os.Stdout, "[EVA] ReqID: %s | Status: Allow | From: %s \n", rag.RequestInput.Id, hostname)
			c.JSON(http.StatusOK, gin.H{"status": "Allow", "from": hostname})
		}

	} else {
		fmt.Fprintf(os.Stdout, "[EVA] ReqID: %s | Status: Deny | Error: %s \n", rag.RequestInput.Id, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func healthy(c *gin.Context) {
	if ready == false {
		c.String(http.StatusServiceUnavailable, "Greetings! This is from %s \n", hostname)
	} else {
		c.String(http.StatusOK, "Greetings! This is from %s \n", hostname)
	}

}

func getPolicy(c *gin.Context) {
	policies, err := manager.GetAll(1000, 0)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"policies": policies})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func createPolicy(c *gin.Context) {
	dp := &policy.DefaultPolicy{}

	if err := c.ShouldBindJSON(dp); err == nil {
		manager.Create(dp)
		c.JSON(http.StatusOK, gin.H{"status": "create successfully", "from": hostname})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func memoryInit() *memory.MemoryManager {
	return memory.NewMemoryManager()
}

func createTables(db *sqlx.DB) {
	db.MustExec(sql.Schema)
}
