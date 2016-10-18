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

type ChaosTest struct {
	ChaosWorkers []chaosWorker
}

// Ping all the workers via net/rpc, make sure they are reachable and running,
// these workers on the nodes will also make sure Minio servers too are running on these nodes,
// In the event of any chaos worker not reachable or Minio server not running on these nodes it'll return error
// and the chaos test will be aborted.
func (chaos *ChaosTest) InitChaosTest() {
	// TODO: Code goes here.
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
