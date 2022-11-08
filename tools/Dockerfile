FROM eunomie/dague:0.0.2 as builder
ARG OS
ARG ARCH

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY magefile.go .

RUN mage -compile /go/src/dague -goos $OS -goarch $ARCH

FROM scratch

COPY --from=builder /go/src/dague /
