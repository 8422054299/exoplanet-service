FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY handlers/*.go ./handlers/
COPY models/*.go ./models/
COPY database/*.go ./database/

RUN go build -o /exoplanet-service

EXPOSE 8080

CMD [ "/exoplanet-service" ]
