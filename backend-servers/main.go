package main

func main() {
	ports := []string{
		":5000",
		":5001",
		":5002",
	}
	for i := 0; i < len(ports); i++ {
		s := Server{
			port: ports[i],
		}
		go s.CreateServer()
	}
	select {}
}
