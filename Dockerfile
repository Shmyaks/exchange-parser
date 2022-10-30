from golang:1.19.2-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main app/cmd/main.go
EXPOSE 5050
CMD [ "/app/main" ]