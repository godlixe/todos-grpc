# todos-grpc
Mini todo service app using gRPC and clean architecture

This app is using docker, make sure you have docker and docker-compose installed on your machine. This app includes :
- Simple todo functionality (create/read/update)
- gRPC calls (unary)
- PostgreSQL database
- Redis caching

Set the env files too (will be updated soon).

Before running, make sure the port `:8080` is not in use.

Run `docker-compose up` to run the service.

The proto files are located in `/pkg/proto` directory.


