FROM node:20-alpine AS webbuilder
WORKDIR /web
ENV NODE_OPTIONS=--max-old-space-size=4096
COPY web/package.json web/yarn.lock ./
# --ignore-engines：传递依赖 @testing-library/jest-dom 声明 node>=22
RUN yarn install --frozen-lockfile --network-timeout 600000 --ignore-engines
COPY web/ .
RUN WEB_BASE=/web yarn build

FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
COPY . .

RUN go env && go build -o CloudSilk main.go

FROM registry.cn-shanghai.aliyuncs.com/swtsoft/golang-run:alpine-3.16.0
LABEL MAINTAINER="guoxf@swtsoft.com"

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yaml"
# dubbogo v3.1 的 dubbo3 描述符注册与既有 pb 存在同名文件重复注册（内容一致，无实际冲突），
# 按 protobuf 官方 FAQ 降级为警告（https://protobuf.dev/reference/go/faq#namespace-conflict）
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn

WORKDIR /workspace
COPY --from=builder /workspace/CloudSilk ./
COPY --from=builder /workspace/config.yaml ./
COPY --from=builder /workspace/docs/swagger.json ./docs
COPY --from=builder /workspace/docs/swagger.yaml ./docs
COPY --from=webbuilder /web/dist ./web/dist

EXPOSE 20000

ENTRYPOINT ["./CloudSilk", "--ui", "./web/dist", "--port=20000", "--service_mode=ALL", "--single_db=true"]
