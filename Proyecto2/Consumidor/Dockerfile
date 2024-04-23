FROM golang:1.22
WORKDIR /appgo
COPY . .
RUN go mod download
RUN go build -o consumer
EXPOSE 9092
CMD ["./consumer"]