FROM golang:1.19-alpine as builder

# Install all required dependencies for building
RUN apk update
RUN apk add make
WORKDIR /service

COPY . .
RUN make build

FROM alpine:3.16


COPY --from=builder ./service/out/stakefish /service/stakefish
COPY --from=builder ./service/internal/db/migrations/scripts /service/migrations
WORKDIR /service
ENTRYPOINT ["./stakefish"]