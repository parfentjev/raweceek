# Rawe Ceek Backend Service

This is a REST API service that provides information about upcoming sessions and Grand Prix weekends.

## Documentation

You can find detailed documentation of the API at: [https://raweceek.eu/api/swagger-ui/index.html#/](https://raweceek.eu/api/swagger-ui/index.html#/)

## Getting Started

To build and run the service, follow these steps:

* Create a new `.env` file based on the `example.env` file.
* Build a Docker image with: `docker build -t rawe-ceek-backend:latest .`
* Before running the service, create the database schema manually using the `spec/schema.sql` file.

Also, see the `Makefile`.
