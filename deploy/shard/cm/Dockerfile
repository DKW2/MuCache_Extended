FROM golang:alpine

RUN apk add --no-cache --update gcc musl-dev zeromq-dev

RUN mkdir /app
COPY ./cmd /app/cmd
COPY ./experiments/k8s_cm /app/experiments/k8s_cm
COPY ./internal /app/internal
COPY ./pkg /app/pkg
COPY ./go.mod /app/
COPY ./go.sum /app/
WORKDIR /app

ENV NODE_IDX=$NODE_IDX
ENV CM_ADDS=$CM_ADDS
ENV SHARD_IDX=$SHARD_IDX
ENV SHARD_COUNT=$SHARD_COUNT

CMD CGO_ENABLED=1 GOOS=linux go run -tags k8s,node$NODE_IDX,shard /app/cmd/cm -logtostderr=true --port 5550 --cm_adds $CM_ADDS
#CMD CGO_ENABLED=1 GOOS=linux go test -v github.com/pebbe/zmq4