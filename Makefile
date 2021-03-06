GRPC_GATEWAY_DIR := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway 2> /dev/null)
GO_MODULE := $(shell go mod edit -json | grep Path | head -n 1 | cut -d ":" -f 2 | cut -d '"' -f 2)
PROTO_DIR := src/customeraccounts/infrastructure/adapter/grpc/proto
GRPC_TARGET_DIR := src/customeraccounts/infrastructure/adapter/grpc/proto
REST_GW_TARGET_DIR := src/customeraccounts/infrastructure/adapter/rest/proto
REST_GW_OUT_FILE := customer.pb.gw.go
REST_SWAGGER_TARGET_DIR := src/customeraccounts/infrastructure/adapter/rest

generate_proto:
	@protoc \
		-I $(GRPC_TARGET_DIR) \
		-I /usr/local/include \
		-I $(GRPC_GATEWAY_DIR)/third_party/googleapis \
		--go_out=plugins=grpc:$(GRPC_TARGET_DIR) \
		--grpc-gateway_out=logtostderr=true,import_path=customerrest:$(REST_GW_TARGET_DIR) \
		--swagger_out=logtostderr=true:$(REST_SWAGGER_TARGET_DIR) \
		$(PROTO_DIR)/customer.proto

	@# Not possible to split grpc and rest otherwise: https://github.com/grpc-ecosystem/grpc-gateway/issues/353
	@sed -i '/package customerrest/ a \\nimport customergrpcproto "$(GO_MODULE)/$(GRPC_TARGET_DIR)"' $(REST_GW_TARGET_DIR)/$(REST_GW_OUT_FILE)
	@sed -i 's/client CustomerClient/client customergrpcproto.CustomerClient/' $(REST_GW_TARGET_DIR)/$(REST_GW_OUT_FILE)
	@sed -i 's/server CustomerServer/server customergrpcproto.CustomerServer/' $(REST_GW_TARGET_DIR)/$(REST_GW_OUT_FILE)
	@sed -i 's/NewCustomerClient/customergrpcproto.NewCustomerClient/' $(REST_GW_TARGET_DIR)/$(REST_GW_OUT_FILE)
	@sed -i -E 's/var protoReq (.+)/var protoReq customergrpcproto.\1/' $(REST_GW_TARGET_DIR)/$(REST_GW_OUT_FILE)

lint:
	golangci-lint run --build-tags test ./...

# https://github.com/golangci/golangci-lint
install-golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.36.0


# https://github.com/psampaz/go-mod-outdated
outdated-list:
	go list -u -m -json all | go-mod-outdated -update -direct