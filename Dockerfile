FROM golang:1.20.3-bullseye

# Set the Working Directory inside the container
WORKDIR /app

# Cache and install dependencies
COPY go.mod ./
RUN go mod download

# Copy app files
COPY . .

RUN go build -o go-app

EXPOSE 8000

CMD ["./go-app"]
