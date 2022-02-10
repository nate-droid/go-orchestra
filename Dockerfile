############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# WORKDIR $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/app/

COPY . .

# Fetch dependencies.
RUN go mod download

# Build the binary.
# TODO too slow...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-orchestra
############################
# STEP 2 build from scratch
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /go/bin/go-orchestra /go/bin/go-orchestra

COPY ./client/build /go/client/build
# Use an unprivileged user.
USER appuser:appuser
# Run the hello binary.
ENTRYPOINT ["/go/bin/go-orchestra"]