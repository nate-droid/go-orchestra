FROM golang:1.17-buster AS builder

ADD . /go/src/go-music
WORKDIR /go/src/go-music

# WORKDIR /go-music

# COPY go.sum ./
# COPY go.mod ./

# RUN go mod download \
#     && go mod verify
RUN go mod download

# COPY ./conductor/* .
RUN pwd
RUN ls
RUN cat go.mod

RUN go build -o /bin/conductor

# step 2 deploy
FROM gcr.io/distroless/base-debian10

# Copy our static executable.
COPY --from=builder /bin/conductor /conductor

# Use an unprivileged user.
USER nonroot:nonroot

EXPOSE 8111
ENTRYPOINT ["/conductor"]
CMD ["conductor"]


