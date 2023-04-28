.PHONY : build
all: build generate_swagger_doc
build:
	scripts/init.sh "cmd/app"

generate_swagger_doc:
	@echo "Generate swagger doc"
	swag init -g cmd/app/app.go
