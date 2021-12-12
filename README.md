# swag_gin_detect_changes

Sample code to fail a CI build in case Swagger isn't kept up to date. Commands to run are:

```
docker build -t detect_swagger_changes .

docker run -v $(pwd):/go/src detect_swagger_changes
```
