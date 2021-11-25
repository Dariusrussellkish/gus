# Gus - Thrift Protocol Testing

On one computer:
```
go run ./main.go --server=true --P=compact --addr="localhost:9999" #on one terminal
go run ./main.go --P=compact --addr="localhost:9999" #on another terminal
```

With two computers on [cloudlab](https://www.cloudlab.us/p/f08bff16177c0602efcaf211be7034939fc86d06):
```
go run ./main.go --server=true --P=compact --addr="10.10.1.1:9999" #on node0
go run ./main.go --P=compact --addr="10.10.1.1:9999" #on node1
```

On a single computer, expect to wait around 5 minutes for 100,000 requests. The client will exit when all the requests have been dealt with successfully.