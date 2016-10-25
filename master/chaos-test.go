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
	"log"
)

type ChaosTest struct {
	ChaosWorkers []ChaosWorker
}

func (chaos *ChaosTest) InitChaosWorker(args *string, reply *int) error {
	*reply = 0
	return nil
}

// Ping all the workers via net/rpc, make sure they are reachable and running,
// these workers on the nodes will also make sure Minio servers too are running on these nodes,
// In the event of any chaos worker not reachable or Minio server not running on these nodes it'll return error
// and the chaos test will be aborted.
func (chaos *ChaosTest) InitChaosTest() bool {
	// TODO: Code goes here.

	var errorOccured bool
	// Iterate through all the chaos workers on remote nodes.
	// Communicate with them using RPC.
	// Don't stop the process if any of the workers return error on RPC call.
	// Log all the errors.
	// If there's no error RPC client will returned, assign it to worker.Client for
	// any further RPC communication with the workers on remote nodes.
	for _, worker := range chaos.ChaosWorkers {
		log.Println("Initializing worker at: ", worker.WorkerEndpoint)
		// Communicate with remote chaos worker.
		// The worker will also verify whether Minio server is running on their respective nodes and in the specified port.
		rpcClient, err := worker.InitChaos()
		// don't return in the event of an error.
		// log the errors from all nodes before returning.
		if err != nil {
			// flag that an error occured in the remote node.
			errorOccured = true
			// log the error.
			log.Printf("Error from Node %s: <ERROR> %v.", worker.WorkerEndpoint, err)
		}
		worker.Client = rpcClient

	}
	return errorOccured
}

// Use Random picker to select the nodes and signal the chaos workers on the nodes to
// shutdown the Minio server process and restart it after a definate or a random interval,
// as defined while the test process is run.
// The status of the Minio servers can be obtaining by visiting the `/report` endpoint,
// and the chaos event will logged.
// The log contains the info the nodes that were bought down and the time at which they were back.
func (chaos *ChaosTest) UnleashChaos() {
	// TODO: Code goes here.

}
