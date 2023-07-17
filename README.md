##### A simple API to get negatives of base64 images
## Usage
- GUI is on http://localhost
- GIN API is on http://0.0.0.0:8000
- API documentation(swagger) is on http://0.0.0.0:8000/docs/index.html
- DB is on http://0.0.0.0:5432

## Run
To install and run perform these commands
```sh
git clone https://github.com/Jabca/GoTest
cd GoTest
docker-compose up
```

## Test
To run test you need already running server via docker-compose
```sh
cd tests
go get .
go test main_test.go -v
```
