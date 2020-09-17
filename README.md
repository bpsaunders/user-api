# user-api

An API offering functionality with which to create and fetch users' data.

### Prerequisites
- [Golang 1.12+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [MongoDB](https://www.mongodb.com/try/download)

### Environment variables
The following environment variables are used for program execution

Variable         |Required  |Example                    |Default |Notes
-----------------|----------|---------------------------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------
MONGODB_URL      | &#x2713; | mongodb://localhost:27017/|        | This variable must follow the standardised [MongoDB connection string format](https://docs.mongodb.com/manual/reference/connection-string/)
MONGODB_DATABASE | &#x2713; | users_application         |        |
LOG_LEVEL        | &#x2717; | debug                     | info   | A lower case representation of the standard log level enumerations. Possible values can be found [here](https://github.com/sirupsen/logrus/blob/master/logrus.go#L25)

### Building and running

#### Natively
A `Makefile` has been added for convenience, offering the following targets:

- fmt: runs `go fmt` recursively
- clean: tidies up built resources
- build: produced a compiled binary called `main` at the root of the project
- lint: runs a linter over the project and outputs warnings / issues to a `lint.txt` file at the root of the project
- test: runs tests within the project and generates a coverage report

Once built, export any environment variables required for execution and execute the `main` binary;
the app listens at port `8888` and will connect to MongoDB on startup.

#### Docker
Bake a Docker image using the following command at the base of the project:

```
docker build --build-arg mongodb_url="<your_mongo_connection_string>" --build-arg mongodb_database="<your_mongo_db_name> -t <image_name> .
```

The `--build-arg` flags are not required; if not provided, values will default to:
- mongodb_url: mongodb://localhost:27017/
- mongodb_database: user_application

An additional `--build-arg` flag with key log_level can be optionally set to configure the log level of the environment.

Once built, run the image using:

```
docker run -p <port_of_choice>:8888 <image_name>
```

The API will be available at `<port_of_choice>` and will connect to MongoDB on startup.

### Endpoints

4 endpoints are exposed by the application:

#### Health check
```
(GET) /health-check
```
A simple health check endpoint which will return an `OK` response to indicate the app is running and available.

#### Fetch all users
```
(GET) /users
```
Fetch an array of all users in the database.

Possible response codes:
- `OK`: a successful response, accompanied by an array of users (empty array if none exist)

#### Create a user
```
(POST) /users
```
Create a user with data provided in the following shape:
```
{
	"first_name": "",
	"last_name": "",
	"email": "",
	"country": ""
}
```

Possible response codes:
- `Created`: user created successfully
- `Bad Request`: the request was invalid, be it from malformed JSON, or from validation errors
- `Conflict`: an attempt was made to create a user with an email which already exists

#### Fetch a user
```
(GET) /users/{id}
```
Fetch an individual user according to an id.

Possible response codes:
- `OK`: a successful response accompanied by a user
- `Not Found`: no user was found for the given id.

#### Errors

Any application errors are handled gracefully, and an `Internal Server Error` response is returned to the user.

#### Validation

On creation of a user, validation is performed to assure the integrity of the data;

##### Name fields

Both first and last names must:
- not be empty
- be between 2 and 30 characters
- not contain invalid characters

##### Email field

Email must:
- not be blank
- not be longer than 120 characters
- follow a valid email format

##### Country field

Country must:
- not be blank
- be a 2 character, upper-cased country code

### Logging

Different log levels offer different levels of verbosity in program output; 
currently this application uses `debug`, `info`, and `error` levels.

### Tech Test Addendum

Firstly, thanks for reading! This application displays *most* of the things I
consider important in programming.

If I had more time, there are a few things I would like to do:
- improve granularity of logging, making use of other log levels to improve debugging - 
perhaps adding some mechanism by which to configure log level during program execution
- create proper API specs, since putting them in a README feels all wrong!
- fix the `TODO` under country validation - currently the app asserts the shape of the
data, but does not verify that it's a real country
