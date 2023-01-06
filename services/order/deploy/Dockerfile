FROM golang:latest as builder

WORKDIR /src
COPY . .
#RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
#    -ldflags "-w" \
#    -o ./build/order ./services/order/cmd

RUN --mount=type=cache,target=/root/.cache/go-build \
    GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w" \
    -o ./build/order ./services/order/cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /src/build/ ./
ENTRYPOINT [ "./order" ]
