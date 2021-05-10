FROM golang:1.16-alpine AS compiler

WORKDIR /src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o ./a.out .

FROM alpine

COPY --from=compiler /src/app/a.out /server
COPY --from=compiler /src/app/assets /assets
ENTRYPOINT ["/server"]
