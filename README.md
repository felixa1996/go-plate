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
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NTMwNSwidXNlcm5hbWUiOiIyNjA0OTciLCJqYWJhdGFuIjoxMCwiaXNzIjoiR3IxSDlDWmpMY1Vob3pjbWZDaG5MeTRyaVplTTRvd1AiLCJpYXQiOjE2Mjc1NzczNTd9.ZUxgYHVYlv63JQ_PhvZ6dUhs27MpRC8k7BdEi57J2NM