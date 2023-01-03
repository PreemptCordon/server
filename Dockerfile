FROM golang as builder 
RUN mkdir src/app
WORKDIR /src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build
FROM scratch
COPY --from=builder /src/app/server /server
