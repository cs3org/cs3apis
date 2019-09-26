# Docker file to build the cs3apis
FROM golang
RUN apt-get update
RUN apt-get install build-essential curl unzip sudo -y
RUN apt-get install python-pip -y
ADD . /root/cs3apis
RUN go run /root/cs3apis/build.go -deps-proto -deps-go
RUN sudo python -m pip install grpcio grpcio-tools --ignore-installed
WORKDIR /root/cs3apis
