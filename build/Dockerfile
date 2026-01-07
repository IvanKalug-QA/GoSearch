FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod dowload

WORKDIR /app/cmd

RUN GOOS=linux go build -o app

EXPOSE 8080

CMD [ "./app" ]

# docker build -f .\build\Dockerfile --progress=plain -t kaluginivan/go_search:latest .
# docker run --rm -p 8080:8080 kaluginivan/go_search:latest