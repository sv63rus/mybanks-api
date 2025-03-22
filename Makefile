.PHONY: generate enrich openapi

generate:
	@echo "🔧 Generating Ent client + OpenAPI JSON spec..."
	go generate ./...

enrich: generate
	@echo "➕ Injecting metadata into openapi.json..."
	yq -i ' \
  .info = { \
    "title": "Bank Directory API", \
    "description": "API для справочника банков", \
    "version": "1.0.0" \
  } | \
  .servers = [{ \
    "url": "http://127.0.0.1:3001/api/v1", \
    "description": "Локальный сервер (по умолчанию)" \
  }]' ent/openapi.json

	@echo "✅ openapi.json enriched"

openapi: enrich
	@echo "🔁 Converting enriched JSON to YAML..."
	yq eval -P ent/openapi.json > docs/openapi.generated.yaml
	@echo "✅ docs/openapi.generated.yaml created from enriched JSON"
