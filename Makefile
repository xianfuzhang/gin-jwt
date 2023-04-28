OUT_DIR = _output

.PHONY : build
all: build generate_swagger_doc
build:
	scripts/init.sh "cmd/app/clean_app.go"

generate_swagger_doc:
	@echo "Generate swagger doc"
	swag init -g cmd/app/clean_app.go

install:
	install -D -m 755 ${OUT_DIR}/clean_app /usr/bin/

clean:
	rm -rf ${OUT_DIR}