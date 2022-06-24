# build stage
FROM docker.io/golang:1.18-alpine AS build-env
ENV CGO_ENABLED 0
ADD . .
RUN go build -ldflags '-s' -o greeter main.go

# final stage
FROM scratch
ENV CGO_ENABLED 0
ENV GREETER_NAME John
COPY greeter /
ENTRYPOINT ["/greeter"]