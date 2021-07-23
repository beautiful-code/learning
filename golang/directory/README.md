# Employees Directory 

LEVEL 1 - database modelling

* Do data modelling for employees directory.
* Add a db package to connect to a postgres database. Use CLI arguments to read database name.
* Add necessary migrations and seed data.
* Use SendWelcomeEmail from directory/alerting package to send welcome notification on employee creation.


LEVEL 2 - GRPC service

* Add GRPC enpoints to make this as a CRUD service for employees directory.
* Add a GRPC client outside this project to consume the above developed gRPC service.