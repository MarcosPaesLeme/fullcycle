# Imersão FullCycle - Go

## Description

Repository to reproduce the motion of vehicle

**Attention**: The Apache Kafka container needs to be run previous.

## Run application

Follow this steps:

Execute Apache Kafka

```
# Access the folder
cd .docker/kafka
# Start container
docker-compose up -d
```

Execute Simulator:

```
# Start container
docker-compose up -d
# Enter in container
docker-compose exec simulator bash
# Execute Golang
go run main.go
```
