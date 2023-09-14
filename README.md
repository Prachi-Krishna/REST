# REST

For understanding,the project structure is as follows:
cmd - contains the main.go file
config - contains the configuration files
db - contains the db connection file
pkg - the pkg folder contains the api,models and server folder
    api - contains the handler,service,repository and helper files (the CRUD services for blog are inside the service.go file)
   models - contains the db structure for blog table
   server - contains the binding.go file to bind handlers,service and repository files & the server.go file contains the endpoints of the services implemented
