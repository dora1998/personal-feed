FROM golang AS build-stage
LABEL build_stage=true
LABEL maintainer="Dora <me@dora1998.net>"
ADD . /go/src/personal-feed

WORKDIR /go/src/personal-feed
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o /go/bin/personal-feed

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build-stage /go/bin/personal-feed /usr/local/bin/personal-feed
ENTRYPOINT ["/usr/local/bin/personal-feed"]