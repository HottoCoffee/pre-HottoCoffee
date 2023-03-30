# HottoCoffee

## Run Swagger(OpenAPI)

```bash
docker run -p 10081:8080 -e SWAGGER_JSON=/app/main.yaml -v ./schema:/app swaggerapi/swagger-ui
```
