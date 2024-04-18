FROM golang

WORKDIR /servidor

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *go ./

RUN go build -o /server

EXPOSE 3000

CMD [ "/server"]