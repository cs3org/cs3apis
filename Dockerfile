# Docker file to build the cs3apis
FROM golang
RUN apt-get update
RUN apt-get install build-essential curl unzip sudo -y
RUN apt-get install python-pip -y
ADD . /root/cs3apis
RUN go run /root/cs3apis/build.go -deps-proto -deps-go

# deps for python
RUN sudo python -m pip install grpcio grpcio-tools --ignore-installed

# deps for js
RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/1.0.6/protoc-gen-grpc-web-1.0.6-linux-x86_64 -o /tmp/protoc-gen-grpc-web
RUN sudo mv /tmp/protoc-gen-grpc-web /usr/local/bin/ && sudo chmod u+x /usr/local/bin/protoc-gen-grpc-web

WORKDIR /root/cs3apis
