# Simple Load Balancer with Go

## How to run

1- Access directory server

```
cd server
```

2- Open 3 terminal and run below commands to start 3 servers:

   Server 1:

   ```
   go run main.go --port=3001

   ```

   Server 2:

   ```
   go run main.go --port=3002

   ```

   Server 3:

   ```
   go run main.go --port=3003

   ```

3- In the root directory, run below command to start the load balancer:

```
go run main.go
```

4- Open any web browser and access to link localhost:3000. The load balancer will route the request to available server.