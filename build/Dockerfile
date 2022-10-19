# todo: add build stages, make image smaller, run it with less privileges
FROM --platform=linux/amd64 golang:1.19-alpine3.16

WORKDIR /opt/app

COPY . .

RUN cd cmd/http && go build -o /opt/app/kubefirst-api

CMD ["./kubefirst-api"]
