FROM golang:1.17

ADD . /mvc-go

WORKDIR /mvc-go
RUN go mod init mvc-go
RUN go mod tidy
RUN go build -o mvc-go .
RUN chmod +x /mvc-go

# opcion 1 el entrypoint
ENTRYPOINT ["/mvc-go"]

# opcion 2 el expose con el cmd
#EXPOSE 8090

#CMD ["go","run","main.go"]