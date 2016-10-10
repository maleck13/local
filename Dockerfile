FROM golang:1.6

ADD . /go/src/github.com/maleck13/locals

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd /go/src/github.com/maleck13/locals &&  go install .

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/locals serve --config=/go/src/github.com/maleck13/locals/config/config.json

# Document that the service listens on port 8080.
EXPOSE 3000
