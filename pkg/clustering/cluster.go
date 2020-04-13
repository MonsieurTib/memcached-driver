package clustering

type ICluster interface {
	Locate(key string) INode
}

type EndPoint struct {
	Address string
	Port    string
}

type Cluster struct {
	Nodes []INode
}

func (cluster Cluster) Locate(key string) INode {
	return cluster.Nodes[0]
}

func CreateCluster(endPoints ...EndPoint) ICluster {
	nodes := make([]INode, len(endPoints))
	for i, endPoint := range endPoints {
		var node INode
		node = &Node{}
		node.Start(endPoint)
		nodes[i] = node
	}
	var cluster ICluster
	cluster = Cluster{
		Nodes: nodes,
	}
	return cluster
}
