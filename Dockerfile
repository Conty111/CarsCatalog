# image for compiling binary
ARG BUILDER_IMAGE="golang:1.22"
# here we'll run binary app
ARG RUNNER_IMAGE="alpine:latest"

FROM ${BUILDER_IMAGE} as builder

ENV GO111MODULE on
#ENV GOPRIVATE ${GOPRIVATE}

RUN mkdir source
WORKDIR /source
COPY go.mod go.sum ./
# Get dependencies. Also will be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# creates build/main files
RUN make build

FROM ${RUNNER_IMAGE}

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories &&\
    apk update &&\
    apk add --no-cache\
    ca-certificates \

RUN apk add musl-dev && apk add libc6-compat

RUN mkdir -p ./api
RUN mkdir -p ./db/migrations

COPY --from=builder /source/docs/api ./docs/api
COPY --from=builder /source/db/migrations ./db/migrations

COPY --from=builder /source/build/app .

RUN chmod +x app