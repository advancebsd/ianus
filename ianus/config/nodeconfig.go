package config

type network_info struct {
	ip string
	port int16
}

type gemini struct {
	root string
	net_info []network_info
}

type http struct {
	root string
	net_info []network_info
}

type Node struct {
	hostname string
	alias []string
	gemini gemini
	http http
}

type Nodes struct {
	nodes map[string]Node
}

func CreateNodeMap() *Nodes {
	return new(Nodes)
}


