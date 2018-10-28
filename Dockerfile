FROM golang:1.11.1-alpine

ENV GOPATH=/go:/Orgpa

COPY . /Orgpa/src/orgpa-frontend

WORKDIR /Orgpa/src/orgpa-frontend

EXPOSE 80

CMD [ "go", "run", "main.go" ]
