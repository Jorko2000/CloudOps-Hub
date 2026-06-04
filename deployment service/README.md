# Deployment Service

Deployment automation service for CloudOps Hub.

## Responsibilities

* Receive deployment requests
* Validate deployment payloads
* Create Kubernetes resources
* Persist deployment records
* Publish deployment events

## Technology Stack

* Go
* Gin
* PostgreSQL
* Kubernetes Client-Go
* Docker

## API

POST /deployments

Request

{
"appName":"inventory-service",
"image":"nginx:latest",
"replicas":3
}

Response

{
"message":"deployment created"
}
