FROM golang:1.19

WORKDIR /app
COPY . /app/
RUN go build

FROM debian:bullseye-slim

WORKDIR /work
COPY --from=0 /app/HottoCoffee /work/
EXPOSE 8080
CMD ["./HottoCoffee"]
