FROM node:18-alpine AS webbuilder
WORKDIR /web
ENV NODE_OPTIONS=--max-old-space-size=4096
COPY web/package.json web/yarn.lock ./
RUN yarn install --frozen-lockfile --network-timeout 600000
COPY web/ .
RUN WEB_BASE=/web yarn build

FROM registry.cn-shanghai.aliyuncs.com/swtsoft/golang-build:1.20.0-alpine3.17 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
COPY . .

RUN go env && go build -o CloudSilk main.go

FROM registry.cn-shanghai.aliyuncs.com/swtsoft/golang-run:alpine-3.16.0
LABEL MAINTAINER="guoxf@swtsoft.com"

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yaml"

WORKDIR /workspace
COPY --from=builder /workspace/CloudSilk ./
COPY --from=builder /workspace/config.yaml ./
COPY --from=builder /workspace/docs/swagger.json ./docs
COPY --from=builder /workspace/docs/swagger.yaml ./docs
COPY --from=webbuilder /web/dist ./web/dist

EXPOSE 20000

ENTRYPOINT ["./CloudSilk", "--ui", "./web/dist", "--port=20000", "--service_mode=ALL", "--single_db=true"]
