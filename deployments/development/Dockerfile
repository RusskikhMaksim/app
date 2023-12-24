FROM golang

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest
COPY ./go.mod ./
RUN go mod download

COPY --chown=app:app . .

EXPOSE ${APPLICATION_PORT}

CMD go run ./cmd/api/main.go
CMD CompileDaemon \
  -build="go build -ldflags=-X=main.BranchName=$(git rev-parse --abbrev-ref HEAD) -o ./bin/api ./cmd/api" \
  -command="./bin/api" \
  -exclude-dir=".git" \
  -color \
  -log-prefix=false
