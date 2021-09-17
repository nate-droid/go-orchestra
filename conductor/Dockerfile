FROM golang:1.17-buster AS builder

WORKDIR /conductor



COPY ../go.sum ./
COPY ../go.mod ./

# RUN go mod download \
#     && go mod verify
RUN go mod download

COPY *.go ./


RUN go build -o bin/conductor

# step 2 deploy
FROM gcr.io/distroless/base-debian10

# Copy our static executable.
COPY --from=builder ./build/conductor /conductor

# Use an unprivileged user.
USER nonroot:nonroot

EXPOSE 8111
ENTRYPOINT ["/conductor"]
# CMD ["server"]


