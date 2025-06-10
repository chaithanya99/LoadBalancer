package main

import (
	"fmt"
)

func main() {
	serverList := []Server{
		{ip: "127.0.0.1", port: "5000"},
		{ip: "127.0.0.1", port: "5001"},
		{ip: "127.0.0.1", port: "5002"},
	}
	roundRobin := RoundRobin{
		count: 0,
	}

	port := ":8080"
	lb := LoadBalancer{
		serverList:   serverList,
		balancerAlgo: &roundRobin,
		port:         port,
	}
	err := lb.Start()
	if err != nil {
		fmt.Print(err)
	}
}
