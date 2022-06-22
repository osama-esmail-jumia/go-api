# The arch
### `/cmd`
* Main applications for this project
  * api command to start the API server
  * command to run workers
  * command for migrations
* Don't put a lot of code in the application directory.
  * If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory.
  * If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory.
* Commands can be grouped in packages for example
* ```
  api.go
  worker/main.go
  worker/blog/main.go
  worker/blog/create.go
  worker/blog/update.go
  worker/comment/main.go
  worker/comment/create.go
  ```
  * * It is domain arch(`blog` package & `comment` package) because each worker will not share the logic with other workers and the shared logic can be in `internal` dir
  * `*/main.go` to group them 
    * `worker blog create`
    * `worker blog update`
    * `worker comment create`

### `/boot` or `/init`
* It works as definitions or container that initialize all dependencies to be injected to each other.
* It is mainly called in `/boot`

### `/configs`
* Contains each app instance config
* Better to create a new config struct for each command because in most cases each command require fewer parameters than other, and the command maybe fail because we have to fill all config params. for example
  * If we use the same app config in `blogWorker` and the app config requires redis client then we have to fill it even with random data to be able to run the `blogWorker`
* Mainly called in `/cmd` because it is the first package called and maybe some configs need there to start the command.

### `/router`
* Contains the routes for each group

### `/definitions`
* Contains all models/entities & interfaces.
* `Models` have to be inside `definitions` to avoid circular dependencies because `models` are accessed by the interfaces & the internal packages at the same time `interfaces` are accessed in internal package.

### `/internals`
* This is the code you don't want others importing in their applications or libraries.
* Contains the application logic(controller, handler, service, etc)
* Subdirectories divided to layers in which they can be access each other in unidirectional to avoid circular dependencies.
  * `Controller` can call `Handler` & `Service`
  * `Service` can call `Mapper` & `Repository` & `Clients` & other services.


### `/pkg`
* Library code that's ok to use by external applications.
* Contains all helper packages that is not related to the business logic.
* They could be published to another repo to be accessed by all go applications.

### `functional arch` vs `domain arch`
* `functional arch`
  * pros
    * Avoid circular dependencies.
      * `blogResponse` has a list of `commentResponse` as param & `commentResponse` has `blogResponse` param
      * `blogMapper` calls `commentMapper` & `commentMapper` calls `blogMapper`
      * `blogService` calls `commentRepo` & `commentService` calls `blogRepo`
    * Each package has the same functionality.
  * cons
    * Some applications don't follow the same arch.(grpc, graphql, etc)
* `domain arch`
  * pros
    * All domain functionality will be in the same package.
  * cons
    * circular dependency.
      * `blogResponse` has a list of `commentResponse` as param & `commentResponse` has `blogResponse` param
      * `blogMapper` calls `commentMapper` & `commentMapper` calls `blogMapper`
      * `blogService` calls `commentRepo` & `commentService` calls `blogRepo`

### `request/response` vs `Http_model`
* `request/response`
  * pros
    * The models separated to two packages
    * Small struct name without prefixing it
      * `http_model.BlogCreateRequest` vs `request.BlogCreate`
    * Can be used in swagger yaml as `request.BlogCreate`
  * cons
    * Maybe other http models with not `request` or `response` like subrequest(query param) for example.
* `Http_model`
  * pros & cons are the opposite of the `request/response`

### `definitions dir` vs `interface with struct`
* `definitions dir`
  * pros
    * All definitions will be in a separate dir and can be exported as pkg
    * The interface not tied to one struct (`CRUDController` interface, `Producer` interface, etc)
  * cons
    * We have to create the same dir hierarchy as in `internal` for the interfaces that are tied to structs(`controllers`, `service`, etc)
    * We may import `pkg/*` structs instead of importing `defenitions/*` 
* `interface with struct`
  * The opposite of `definitions dir` cons


## `TODO`
* Migrations files
* Reading the config from exported env without reading `.env` file
* Swagger documentation
* Use postgresSQL & containerize it
* Unit tests

