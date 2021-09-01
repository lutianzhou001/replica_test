FROM golang:1.15.6

ENV GO111MODULE="on"

ENV GOPROXY="https://goproxy.cn"

RUN mkdir application

COPY . ./application

WORKDIR "application"

RUN  go build -o main main.go

EXPOSE 1926

CMD ["./main"]
