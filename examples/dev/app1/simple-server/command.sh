go get github.com/gorilla/context@latest
go get github.com/gorilla/mux@latest

curl -G -d "name=tanaka" http://localhost:8080/
curl http://localhost:8080/

go test

docker pull golang:alpine3.21

docker build -t simple-server .
docker run -it --detach --publish 8080:8080 --name simple-server simple-server