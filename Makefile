dev:
	@go run main.go
db-start:
	@docker-compose -f docker/mongo.yml up -d
db-stop:
	@docker-compose -f docker/mongo.yml down