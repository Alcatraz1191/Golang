##Base Image
FROM golang:1.12.0-alpine


##Create /app directory within image to hold the appplication

RUN mkdir /app

##Copy everything in the directory

ADD . /app

## Further execute any other commands in /app directory

WORKDIR /app

EXPOSE 8080
## Add git to alpine 1.12.0
RUN apk add --no-cache git

## Compile binary
RUN go build -o main .

## Command to kick off  newly created binary

CMD ["/app/main"]