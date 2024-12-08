#Build Stage
FROM golang:1.22.10-bullseye AS builder 
WORKDIR /app 

# First . copy all folder from current working directory to the current working directory in container
COPY . . 

# Creating executable file
RUN go build -o main main.go 

RUN apt-get update && apt-get install -y curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz -o migrate.tar.gz 
RUN tar -xvzf migrate.tar.gz

FROM debian:bullseye-slim
WORKDIR /app 
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate 
COPY start.sh .
COPY wait-for.sh . 
COPY app.env .
COPY Database/migration ./migration 
RUN ls

# Install netcat in the final image
RUN apt-get update && apt-get install -y netcat

EXPOSE 8080 
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"] 
