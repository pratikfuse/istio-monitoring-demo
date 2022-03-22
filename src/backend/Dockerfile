FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/backend

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 4000

# Run the executable
CMD ["backend"]