FROM golang:1.17
RUN go install github.com/DataDog/temporalite/cmd/temporalite@latest
ENTRYPOINT temporalite start --namespace default --ephemeral --ip 0.0.0.0 --log-format pretty
