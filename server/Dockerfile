FROM golang:1.17-buster AS builder

WORKDIR /app

COPY go.sum ./
COPY go.mod ./

RUN go mod download

COPY . .


RUN go build -o ./server

# step 2 deploy
# FROM gcr.io/distroless/base-debian10
# FROM scratch

# Copy our static executable.
# COPY --from=builder ./ ./

# Use an unprivileged user.
# USER nonroot:nonroot

ENTRYPOINT ["./server"]
