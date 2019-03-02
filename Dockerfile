FROM golang:latest

WORKDIR $GOPATH/src/go-pubg

COPY . .
# Install Iris Framework
RUN go get -u github.com/kataras/iris

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the "go run" command
ENTRYPOINT ["go", "run", "/go/src/go-pubg"]