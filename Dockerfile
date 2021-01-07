FROM swaggerapi/swagger-codegen-cli-v3

WORKDIR /action

COPY codegen.sh .

RUN ["chmod", "+x", "codegen.sh"]

FROM golang
WORKDIR /action
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN go build -o main

ENTRYPOINT ["./main"]
