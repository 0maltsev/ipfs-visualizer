package services

type Container struct {
	Clusters *ClusterService
	Nodes    *NodeService
	Status   *StatusService
	Logs     *LogsService
}
