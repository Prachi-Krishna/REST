# REST

For understanding,the project structure is as follows:

cmd - contains the main.go file <br>
config - contains the configuration files<br>
db - contains the db connection file<br>
pkg - the pkg folder contains the api,models and server folder<br>
  api - contains the handler,service,repository and helper files (the CRUD services for blog are inside the service.go file)<br>
  models - contains the db structure for blog table<br>
  server - contains the binding.go file to bind handlers,service and repository files & the server.go file contains the endpoints of the services implemented<br>
