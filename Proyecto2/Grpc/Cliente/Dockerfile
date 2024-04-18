FROM golang:1.22-alpine

WORKDIR /cliente-go

COPY . .

RUN go mod download
RUN go build -o main

EXPOSE 3000

CMD [ "./main"]