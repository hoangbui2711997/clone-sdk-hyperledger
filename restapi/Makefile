gen-swag:
	swag init --dir ./ --generalInfo routes/router.go --propertyStrategy snakecase --output ./routes/docs;
	go build
build:
	sh build.sh
gen-proto:
	buf generate
gen-build: gen-proto build
build-and-deploy: build
	sh ./deploy-tk-qp-api-image.sh
build-and-deploy-grpc: build
	sh ./deploy-tk-qp-grpc-image.sh
build-and-deploy-worker: build
	sh ./deploy-tk-qp-worker-image.sh