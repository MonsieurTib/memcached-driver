package clustering

type ClusterNodeLocator interface {
	Locate(key string) *Node
}

type EndPoint struct {
	Address string
	Port    string
}

type Cluster struct {
	Nodes []Node
}

func (cluster *Cluster) Locate(key string) *Node {
	return &cluster.Nodes[0]
}

func CreateCluster(endPoints ...EndPoint) *Cluster {
	nodes := make([]Node, len(endPoints))
	for i, endPoint := range endPoints {
		node := Node{}
		node.Start(endPoint)
		nodes[i] = node
	}
	return &Cluster{
		Nodes: nodes,
	}
}
