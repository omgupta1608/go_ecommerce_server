FROM golang:1.18

WORKDIR /usr/aftershoot_task

COPY go.mod .
COPY go.sum .

# RUN go mod download

COPY . .

RUN go build -o builds/ ./main.go

CMD [ "/usr/aftershoot_task/builds/main" ]