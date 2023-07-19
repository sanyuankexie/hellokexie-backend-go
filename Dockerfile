# dockerhub msqt/hello-kexie 🥺
# build stage
FROM golang:1.18-alpine3.17 as build
WORKDIR /app
COPY . .
RUN go build -o main .

# run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 5201
CMD [ "/app/main" ]
