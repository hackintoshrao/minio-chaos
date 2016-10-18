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
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	// TODO: parse all the flags here.
	endpointStr := flag.String("endpoints", "", "RPC endpoints of workers.")
	flag.Parse()

	// TODO: collect all the end points into `endPoints`.

	// TODO: Iterate through the endPoints and create `ChaosTest` instance.
	for i, endPoint := range endPoint {
		worker := ChaosWorker{}
		// TODO: Init the struct here.
		// push all the workers into the array.
		chaosWorkers = append(chaosWorkers, worker)
	}
	// Create `ChaosTest` instance here.
	chaosTest := ChaosTest{
		ChaosWorkers: chaosWorkers,
	}
	// TODO: Start the Server and thed RPC service correctly below.

	// Register the chaos test reporting end point.
	sh := statusHandler{status: &t.status}

	http.Handle("/status", sh)

	// Register the Prometheus metrics endPoint,
	// Currently used only for simple counters.
	http.Handle("/metrics", prometheus.Handler())
	// Start the server.
	go func() {
		http.ListenAndServe(":9998", nil)
	}()

	// Initialize all the workers on remote nodes.
	// also confirms that minio server instances are running on the remote nodes.
	chaosTest.InitChaosTest()
	// Unleash the chaos test.
	chaosTest.UnleashChaos()
}
