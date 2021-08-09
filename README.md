# **Go Plate** #

Running with Kong API Gateway + PostgreSQL

Microservice Golang + NodeJS

**Switching ORM On Fly**
Currently support raw db,go-pg,gorm

**Swagger**
http://localhost:8000/vhry/data/swagger/index.html

**Install kafka and configure**
cd /server/kafka/
docker-compose up

**Generate Swagger Documentation**
make generate_swagger

go get -v -u github.com/swaggo/swag/cmd/swag

**Default JWT Token**
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlX2Rpc3BsYXlfbmFtZSI6ImFkbWluIiwicm9sZV9pZCI6ImY4MGI5YzQyLTNjM2MtNGMwMy1iZjYxLWRlMzZjN2MwZTMwZSIsImlzcyI6IkhUbXI1UnA3MTFwTE5CSjAzQnNydXdRVGpPSmM3clhGIiwiaWF0IjoxNjI4MzU4NTIyfQ.2-KFIPyku9e80iZ5-vVJzTdCSKX6x3WrVj540UCIgCA
