IMAGE=registry.cn-shanghai.aliyuncs.com/antshome/CloudSilk:1.0.0
run:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn MOM_DISABLE_AUTH=true DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" go run main.go 
run-sqlite:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn MOM_DISABLE_AUTH=true DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" go run main.go --ui ./web/dist --db_type="sqlite" 
run-all:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn MOM_DISABLE_AUTH=true DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" SERVICE_MODE="ALL" go run main.go --ui ./web/dist --service_mode="ALL" --port=48089 --single_db=false
run-all-single-db:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn MOM_DISABLE_AUTH=true DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" SERVICE_MODE="ALL" go run main.go --ui ./web/dist --service_mode="ALL" --port=48089 --single_db=true
build-image:
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o CloudSilk main.go
	sudo docker build -f local.Dockerfile -t ${IMAGE} .
	rm CloudSilk
test-image:
	docker run -v `pwd`:/workspace/code -p 48081:48081 --env DUBBO_GO_CONFIG_PATH="./code/dubbogo.yaml" --rm  ${IMAGE}
push-image:
	sudo docker push ${IMAGE}
gen-doc:
	swag init --parseDependency --parseInternal --parseDepth 1

start-web:
	cd web && WEB_BASE=/web yarn start
build-web:
	cd web && WEB_BASE=/web yarn build