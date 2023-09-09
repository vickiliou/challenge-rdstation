FROM golang:1.16-alpine

WORKDIR /challenge-rdstation

COPY . .

CMD [ "go", "test" ]