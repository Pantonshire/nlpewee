GO := go
GO_BUILD := $(GO) build
PROTOC := protoc

MAIN := cmd/main/main.go
BUILD := build
DIST := $(BUILD)/dist
EXEC := nlpewee

DEFAULT_CONF := config_default.yaml

DIST_TARGETS := linux-amd64 linux-arm linux-arm64 darwin-amd64 windows-amd64 windows-arm
FULL_DIST_TARGETS := $(addprefix dist-,$(DIST_TARGETS))

PROTO_PATH := proto
PROTO_SOURCE := $(wildcard $(PROTO_PATH)/*.proto)
PROTO_BUILD := $(PROTO_SOURCE:.proto=.pb.go)

build: $(BUILD)/$(EXEC) $(EXEC)

$(BUILD)/$(EXEC): buildpath
	$(GO_BUILD) -o $(BUILD)/$(EXEC) $(MAIN)

$(EXEC):
	ln -s $(PWD)/$(BUILD)/$(EXEC) $(EXEC)

dist: $(FULL_DIST_TARGETS)

dist-%: distpath
	@ mkdir -p $(DIST)/$(EXEC)-$*
	@ cp $(DEFAULT_CONF) $(DIST)/$(EXEC)-$*/config.yaml
	GOOS=$(word 1,$(subst -, ,$*)) GOARCH=$(word 2,$(subst -, ,$*)) $(GO_BUILD) -o $(DIST)/$(EXEC)-$*/nlpewee $(MAIN)
	@ cd $(DIST) && zip -rq $(EXEC)-$*.zip $(EXEC)-$*
	@ rm -rf $(DIST)/$(EXEC)-$*

buildpath:
	@ mkdir -p $(BUILD)

distpath: buildpath
	@ mkdir -p $(DIST)

clean:
	rm -rf $(BUILD)

proto: $(PROTO_BUILD)

$(PROTO_PATH)/%.pb.go: $(PROTO_PATH)/%.proto
	$(PROTOC) -I $(PROTO_PATH) --go_out=plugins=grpc,paths=source_relative:$(PROTO_PATH) $<

clean-proto:
	rm $(wildcard $(PROTO_PATH)/*.pb.go)

.PHONY: buildpath distpath build dist dist-% clean proto clean-proto
