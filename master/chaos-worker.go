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

// workerConfig contains info about the chaos workers running on nodes to be tested for fault tolerance.
// Also contains info about Minio server instance running on these machines.
type ChaosWorker struct {
	// Endpoint of the chaos worker.
	WorkerEndpoint string
	// Info of the Minio Server Instance runnign on the node of the chaos worker.
	Node MinioNode
	// Directory in which the chaos worker dumps the report.
	ReportDir string
}

// InitChaos - Pings the Chaos worker on the remote node via RPC and initializes it.
//             Also checks whether Minio server instance is running on the specified port on the remote node.
//             Reports failure if either the chaos-worker on the remote node is not reachable or Minio server
// 	       is not running on the specified port on the remote node.
func (chaos ChaosWorker) ReportStatus() {
	// TODO: Code goes here.
}

// ReportStatus - Obtain the Status of the Minio node and choas-worker using the RPC call.
func (chaos ChaosWorker) ReportStatus() {

	// TODO: Code goes here.
}
