docker build -t gcr.io/confab-cloud/echo-server:latest .

docker build -t gcr.io/confab-cloud/echojs:latest .

docker push gcr.io/confab-cloud/echo-server:latest

docker push gcr.io/confab-cloud/echojs:latest


protoc -I proto/ proto/service.proto --go_out=plugins=grpc:proto

protoc -I proto/ service.proto --js_out=import_style=commonjs:protojs

protoc -I build/ service.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:protojs



go get github.com/grpc/grpc-web

cd $GOPATH/src/github.com/grpc/grpc-web

git submodule update --init --recursive


kubectl apply -f deploy/echo.yaml

kubectl apply -f deploy/gateway.yaml