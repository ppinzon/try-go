# Lets Go

### Part A: CRUD web server

###### Description: implement basic web server with endpoints to create/get/delete users

1. Endpoints example:
   
         /users

(creates user, HTTP POST request) -> POST returns ID of new user
   
      /users/:id 

(DELETE / GET requests) -> GET returns JSON user
   
Choose response status by yourself (e.g. StatusCreated, StatusOK, StatusDeleted etc.)

2. Store data in memory :heavy_check_mark:
3. Concurrently safe (sync package)
4. JSON format on endpoints (json package) :heavy_check_mark:
5. Use echo framework (or any other of your choice) :heavy_check_mark:
6. Test endpoints with POSTMAN :heavy_check_mark:

Resources:


https://echo.labstack.com/guide


https://www.postman.com/

### Part B: Middleware, testing and linting

###### Description: implement a basic middleware for the web server, write unit tests

1. Centralised logging of incoming request / outgoing response :heavy_check_mark:
2. Require basic auth for all requests (use authorisation header). :heavy_check_mark:
3. Write unit tests for storage
4. Write unit tests for CRUD endpoints :heavy_check_mark:
5. Install and run golangci-lint (linter)

OPTIONAL:
1. DO NOT USE ready middleware solution, write it yourself
2. Add option to log also to file

Resources:

https://echo.labstack.com/middleware

https://echo.labstack.com/guide/testing

https://medium.com/wesionary-team/introduction-to-linters-in-go-and-know-about-golangci-lint-6486fb2864b3

https://golangci-lint.run/

### Part C: Configuration

###### Description: configure the application and dockerise it (optional)

1. Parameters: port, basic auth user acc/pass, logging enable/disable
2. Configuring the application (yaml, env) (use yaml parsing package, environment variables parsing implement yourself or use any available package)
3. Validate the configuration (require port, user and pass; make sure port is free; make sure pass is not from the list of forbidden passwords) (use validator v9 or implement simple checks yourself)

OPTIONAL:
1. Dockerising application (writing Dockerfile). Separate steps for build and run


Resources:

https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64

https://pkg.go.dev/github.com/go-playground/validator

https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

https://docs.docker.com/engine/reference/builder/


### (OPTIONAL) PART D: Adding database

1. Run your docker image from docker-compose file 
2. Add any docker database image to docker-compose
3. Write Go code to connect to this storage from app
4. Implement Storage interface using this database connection

Resources:

https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/

https://docs.docker.com/compose/compose-file/compose-file-v3/



