FROM golang:1.11beta2-alpine3.8 AS build-env

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 rmazur
RUN mkdir /secservicego/ && chown rmazur /secservicego/
USER rmazur

WORKDIR /secservicego/
ADD . /secservicego/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /secservicego/grmz .

# final stage
FROM alpine:3.8

# Secure against running as root
RUN adduser -D -u 10000 rmazur
USER rmazur

WORKDIR /
COPY --from=build-env /secservicego/certs/docker.localhost.* /
COPY --from=build-env /secservicego/grmz /

EXPOSE 8080

CMD ["/grmz"]
