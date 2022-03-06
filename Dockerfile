FROM golang:1.17.8-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ENV EINHORN_FDS="3"
ENV ISLOCAL="1"
RUN go get github.com/zimbatm/socketmaster 
EXPOSE 10002
RUN go build -o antre-app /app/cmd/

CMD ["socketmaster", "-listen=tcp://:10002", "-command=/app/antre-app"]
