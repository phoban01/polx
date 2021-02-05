.PHONY: build install test

NAME:=polx
VERSION :=v1.0.0-alpha.1
BUILD := $$(git log -1 --pretty=%h)
BUILD_DATE := $$(date -u +"%Y%m%d.%H%M%S")

build:
	@go build \
		-ldflags "-X main.Version=${VERSION} \
				  -X main.Build=${BUILD} \
				  -X main.BuildDate=${BUILD_DATE}" \
		-o bin/${NAME} ./cmd/${NAME}

install:
	@go install \
		-ldflags "-X main.Version=${VERSION} \
				  -X main.Build=${BUILD} \
				  -X main.BuildDate=${BUILD_DATE}" \
		./cmd/${NAME}
test:
	@gotest -v ./pkg/**
