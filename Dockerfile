# Build the project
FROM golang:alpine AS build
ADD . /code
RUN cd /code && go build -o goChat

# Package up the built application
FROM alpine
WORKDIR /app
COPY --from=build /code/goChat /app/
ENTRYPOINT ./goChat
