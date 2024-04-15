FROM golang:1.22.0-alpine.19

RUN apk update && apk --no-chace add bash curl

ENV HELLO_MESSAGE="Hello, Universe!"
ENV PORT=8080

COPY main-simple.go /app/main-simple.go

CMD ["go", "run", "/app/main-simple,go"]
