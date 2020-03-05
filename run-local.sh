# docker run -it -p 8092:8080 --rm --name running-patchwork-app patchwork-app
go run *.go && ./patchwork --listenaddr=0.0.0.0:8092
