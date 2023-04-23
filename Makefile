
generate_swagger_doc:
	@echo "Generate swagger doc"
	swag init -g cmd/app/app.go

clean:
	@echo "Cleaning up ..."

