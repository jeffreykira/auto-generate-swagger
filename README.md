# log-management


## Swagger API Doc
Use [swaggo](https://github.com/swaggo/swag) to auto generate Swagger API doc.

- open API doc
    1. ~/log-management > go run cmd/main.go
    2. http://localhost:5608/swagger/index.html

- How to update API doc
    ```swag init --parseDependency=true -d cmd```