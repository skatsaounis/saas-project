# Source: https://docs.docker.com/language/golang/build-images/

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /askmeanything

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /askmeanything /askmeanything

EXPOSE 80

ENTRYPOINT ["/askmeanything"]
