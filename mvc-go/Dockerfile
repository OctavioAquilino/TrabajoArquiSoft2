FROM golang:1.17

#RUN mkdir /mvc-go
ADD . /mvc-go

WORKDIR /mvc-go
RUN go mod init mvc-go
RUN go mod tidy
RUN go build -o mvc-go .
RUN chmod +x /mvc-go
#RUN chmod +x /mvc-go/mvc-go

# opcion 1 el entrypoint
ENTRYPOINT ["/mvc-go"]
#ENTRYPOINT ["/mvc-go/mvc-go"]

# opcion 2 el expose con el cmd
#EXPOSE 8090

#CMD ["go","run","main.go"]