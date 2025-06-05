# Cheat sheet with commands and queries for project demo

## Start the local DynamoDB instance
```bash
docker-compose up
```

## Create a table in DynamoDB
```bash
aws dynamodb create-table \
    --table-name SwimmersAndSessions \
    --attribute-definitions \
        AttributeName=PK,AttributeType=S \
        AttributeName=SK,AttributeType=S \
    --key-schema \
        AttributeName=PK,KeyType=HASH \
        AttributeName=SK,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
	--endpoint-url http://localhost:8000
```
## Check table exists
```
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

## Add a swimmer to the database
```bash
curl -s localhost:3000/swimmers/add -d '{"name": "Test Testersen", "age": 42}' -H "Content-Type: application/json" | jq
```

## Add a session to the database
```bash
curl -s -X POST http://localhost:3000/sessions/add \
-H "Content-Type: application/json" \
-d '{
  "swimmer_id": "",
  "date": "2025.01.19",
  "distance": 1000,
  "duration": 30,
  "intensity": "moderate",
  "style": "freestyle",
  "notes": "Morning swim"
}' | jq
```

## Get swimmer by ID
```bash
curl -s "localhost:3000/swimmers/find?id=5578ab2b-0ed7-4666-a40a-05c3f795e645" | jq
```

## Find all swimmers
```bash
curl localhost:3000/swimmers
```

## List all items from database
```bash
aws dynamodb scan --table-name SwimmersAndSessions --endpoint-url http://localhost:8000 | jq
```

## Send an event to the lambda function
```bash
sam local invoke aws_lambda_function.my_lambda -e event.json --docker-network dynamodb-local
sam local invoke aws_lambda_function.create_swimmer -e event.json --docker-network dynamodb-local
sam local invoke aws_lambda_function.list_swimmers --docker-network dynamodb-local
sam local invoke aws_lambda_function.update_swimmer -e update_event.json --docker-network dynamodb-local
```
