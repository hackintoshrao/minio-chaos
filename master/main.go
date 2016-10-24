/*
 * Minio Cloud Storage, (C) 2015, 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	//	"net/http"
	//	"net/rpc"
	"log"
	"strings"
	//	"github.com/prometheus/client_golang/prometheus"
)

const (
	MinioDefaultAddr = "http://127.0.0.1:9000"
)

func main() {
	// TODO: parse all the flags here.
	endPointStr := flag.String("endpoints", "", "RPC endpoints of workers.")
	flag.Parse()
	endPoints := strings.Split(*endPointStr, ",")

	chaosWorkers := make([]ChaosWorker, len(endPoints))
	// Iterate through the endPoints and create `ChaosTest` instance.
	for i, endPoint := range endPoints {
		worker := ChaosWorker{
			WorkerEndpoint: endPoint,
			Node: MinioNode{
				Addr: MinioDefaultAddr,
			},
			//TODO: Make use of report Dir.
			ReportDir: "/not-used-yet",
		}
		// push all the workers into the array.
		chaosWorkers[i] = worker
	}
	// Create `ChaosTest` instance here.
	chaosTest := ChaosTest{
		ChaosWorkers: chaosWorkers,
	}
	//	// TODO: Start the Server and thed RPC service correctly below.
	//mux := http.NewServeMux()
	//	sh := statusHandler{status: &t.status}
	//	sh := &reportHandler{}
	//	mux.Handle("/status", sh)
	//stringRpc := new(StringService)
	//	rpcServer := rpc.NewServer()
	//	rpcServer.RegisterName("Master", &chaosTest)
	//	mux.Handle("/master", rpcServer)

	// Initialize all the workers on remote nodes.
	// also confirms that minio server instances are running on the remote nodes.
	if isFailed := chaosTest.InitChaosTest(); isFailed {
		log.Fatal("Iniitalizing of Chaos test failed.")
	}

	//	// Register the chaos test reporting end point.
	//
	//
	//	// Register the Prometheus metrics endPoint,
	//	// Currently used only for simple counters.
	//	http.Handle("/metrics", prometheus.Handler())
	//	// Start the server.
	//http.ListenAndServe(":9998", nil)

	//	// Unleash the chaos test.
	//	chaosTest.UnleashChaos()
}
