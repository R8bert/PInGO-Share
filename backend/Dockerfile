
FROM golang:1.22
WORKDIR /app
COPY . .
RUN go mod init backend && \
    go get github.com/gin-gonic/gin github.com/google/uuid && \
    go build -o main .
CMD ["./main"]
