FROM golang:1.18-bullseye as dev
WORKDIR /go/src/work

FROM golang:1.18-bullseye as builder
WORKDIR /go/src/work
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build .

FROM gcr.io/distroless/base-debian11 as deployer
WORKDIR /go/src/work
COPY --from=builder /go/src/work/departures-time-api ./
CMD [ "./departures-time-api" ]
