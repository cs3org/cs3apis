# Docker file to build the cs3apis
FROM golang
RUN apt-get update
RUN apt-get install build-essential curl unzip sudo -y
RUN apt-get install python-pip -y
#RUN  sudo python -m pip install grpcio grpcio-tools --ignore-installed
ADD . /root/cs3apis
WORKDIR /root/cs3apis
