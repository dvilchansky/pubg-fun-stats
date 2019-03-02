FROM golang:latest

WORKDIR $GOPATH/src/go-pubg

COPY . .
# Install Iris Framework
RUN go get -u github.com/kataras/iris

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD go run /go/src/go-pubg