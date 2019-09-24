# Docker file to build the cs3apis
FROM golang
RUN apt-get update
RUN apt-get install build-essential curl unzip sudo -y
ADD . /root/cs3apis
WORKDIR /root/cs3apis
ENTRYPOINT ["make", "build-all"]
