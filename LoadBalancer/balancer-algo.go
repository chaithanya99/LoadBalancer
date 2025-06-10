package main

type Balancer interface {
	getServer([]Server) Server
}

type RoundRobin struct {
	count int
}

func (r *RoundRobin) getServer(listOfServers []Server) Server {
	s := listOfServers[r.count]
	size := len(listOfServers)
	r.count = (r.count + 1) % size
	return s
}
