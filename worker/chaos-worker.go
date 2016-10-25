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
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strings"
)

// MinioNode - info of the Minio node under chaos test.
type MinioNode struct {
	Addr string
}

type Worker struct {
}

func StartMinioServer() {

}

func StopMinioServer() {

}

// Tests whether Minio is running on the node in the specified port.
// Send a GET request and find out whether the header contains string `Minio` in it.
func IsMinioRunning(addr string) error {
	// error to be returned to the master if Minio server is not reachable on the node.
	var errRunMinioServer = fmt.Errorf("Run Minio on %s and start the test again.", addr)
	// send GET request to the specified port.
	resp, err := http.Get(addr)
	// Errors out if server is not running at the specified port.
	// return error to the RPC request.
	if err != nil {
		log.Println(err)
		return errRunMinioServer
	}
	log.Println(resp.Header.Get("Server"))
	// check if the server running is Minio server.
	// this is done by checking for string `Minio` is the `Server` header of the response.
	if !strings.Contains(resp.Header.Get("Server"), "Minio") {
		return errRunMinioServer
	}
	// success, return the error to be `nil` to the RPC request.
	return nil
}

func (w *Worker) InitChaosWorker(args *string, reply *struct{}) error {
	log.Println("control in the init worker")
	log.Println(*args)
	err := IsMinioRunning(*args)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// obtain ServeMux to register the RPC service.
	mux := http.NewServeMux()
	// chaos worker configuration.
	worker := &Worker{}
	// Creating a new instance of the RPC server.
	rpcServer := rpc.NewServer()
	// Registering the RPC handler.
	rpcServer.RegisterName("ChaosWorker", worker)
	// Regsitering the RPC service.
	mux.Handle("/", rpcServer)
	// Run the server.
	http.ListenAndServe(":9997", mux)
}
