# minio-chaos
Choas framework for testing Minio's fault tolerance capability.


# Initial Design 

 - The master program pings the chaos workers across nodes and ensures
    that they are running and the choas workers in turn verify that the Minio
    server is running on the remote node at the specified port.
 - On the event of any chaos worker not reachable or if the Minio
    process is not running on the remote node the test fails.
 - The status of the choas test while its running can be viewed at
    `/report` endPoint.
 - Prometheus experimental usage is done to keep count of nodes which
    are down. The use of prometheus will be extended as we proceed
    further with the project.
 - Once the chaos workers are in place and Minio server is running on
    the remote nodes the master randomly picks the nodes and stops the
    server for specified/random time interval.
 - Stopping and Starting the server process will be the only kind of failure for the initial release.
    As we proceed further different failures like bad network will also be simulated.
 - After the specified/random time interval the Minio node will be
    started again and the operation will be logged and the status can be
    viewed on the report endPoint.


