# Swagger
SWAGGER=$(shell which swagger)

gen_code:
	@echo "generate code"
	@$(shell if [ ! -e src/pkg/v1 ];then mkdir -p src/pkg/v1; fi)
	@$(SWAGGER) generate model -q -t src/pkg/v1 -f apis/swagger.yaml
	@$(SWAGGER) generate server -q -A ced -t src/pkg/v1 -f ./apis/swagger.yaml --existing-models=my-swagger/src/pkg/v1/models --model-package=models --skip-models --server-package=server
	@$(SWAGGER) generate client -q -A ced -t src/pkg/v1 -f ./apis/swagger.yaml --existing-models=my-swagger/src/pkg/v1/models --model-package=models --skip-models --client-package=client
