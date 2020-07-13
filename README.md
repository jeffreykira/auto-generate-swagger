# auto-generate-swagger

## Log Configuration

- defaults to showing warnings and above
    ```
    go run cmd/main.go
    ```
- show all logs
    ```
    LOGXI=* go run cmd/main.go
    ```
- auth show debug, other show error log
    ```
    LOGXI=*=ERR,auth=DBG go run cmd/main.go
    ```
- auth close log, other show debug log
    ```
    LOGXI=*=DBG,auth=OFF go run cmd/main.go
    ```
- set all to Error and set data related packages to Debug
    ```
    LOGXI=*=ERR,dat*=DBG go run cmd/main.go
    ```

## Swagger API Doc
Use [swaggo](https://github.com/swaggo/swag) to auto generate Swagger API doc.

- open API doc
    1. ~/log-management > go run cmd/main.go
    2. http://localhost:5608/swagger/index.html

- How to update API doc
    ```
    swag init --parseDependency=true -d cmd
    ```
