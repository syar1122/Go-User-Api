FROM golang as builder

WORKDIR /backend

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine as final

COPY --from=builder /backend/server .
COPY --from=builder /backend/.env .

RUN echo cat .env

RUN mkdir /config

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT [ "/server" ]