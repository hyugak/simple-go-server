# simple-go-server
simple HTTP server with golang
## usage

### 1. build 
`$ docker build --tag simple-go-server .`

#### or pull
`$ docker pull hyugak/simple-go-server`

### 2. run
`$ docker run -id -p 8080:8081 simple-go-server`

### 3. snd request and get response
```
$ curl -XGET -H "Content-Type:application/json" http://localhost:8080/
{"Message": "Hello World"}
```
