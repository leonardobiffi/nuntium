FROM alpine:3.17 as app

RUN apk --no-cache add ca-certificates openssl curl bash jq

WORKDIR /root/

COPY scripts/ scripts/
RUN sh scripts/install.sh

ENTRYPOINT ["nuntium"]
