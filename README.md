# ⚖️ TCP Load Balancer in Go

This project is a simple TCP Load Balancer written in Go. It forwards client requests to a set of backend servers using a Round Robin algorithm.


## 🚀 How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/chaithanya99/LoadBalancer.git
cd LoadBalancer
```
### 2. Open Backend Server Repository
In the terminal one use these commands
```bash
cd backend-server
go run .
```
### 3. Open Load balancer Repository
In the terminal two use these commands
```bash
cd LoadBalancer
go run .
```
### 4. Test Load Balancer
In the terminal three use these commands to test
```bash
nc localhost 8080
```
If you run this command multiple times you will see responses like this
hello from 5000
hello from 5001
hello from 5002
Each request goes to a different backend — thanks to the Round Robin algorithm.


## 📖 Read the Full Story

Want to dive deeper into the design, the challenges, and how to improve this further?

📰 Check out the full article on Medium:  
👉 [Building a TCP Load Balancer in Go](https://medium.com/@chaithanya970/load-balancer-9e33c2b647f0)
