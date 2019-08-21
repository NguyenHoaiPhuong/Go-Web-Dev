FROM golang:alpine

# ENV APP_DIR $GOPATH/src/chatbot

# Copy the current directory contents into the container at /app
COPY . /app

# Set the working directory to /app
WORKDIR /app/src

RUN apk add --update --no-cache git && \
    apk add --no-cache bash && \
    go get -u github.com/gorilla/websocket

RUN go build main.go

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]