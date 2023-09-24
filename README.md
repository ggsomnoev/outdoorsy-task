# Outdoorsy Task

A simple go microservice which serves rentals data based on some criteria

## Prerequisites

- [Go](https://golang.org/) (version 1.21 or higher) or/and Docker

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ggsomnoev/outdoorsy-task.git
   ```

2. Change into the project directory:

   ```shell
   cd outdoorsy-task
   ```

3. Install project dependencies:

   ```shell
   go mod tidy
   ```

## Usage

### Running the Project

To run the project with docker:

```shell
docker-compose up -d simple-rentals-api
docker exec -it simple-rentals-api bash
```

Inside the container use the following command:

```shell
go run main.go
```

This will start the application, and you can access it at [http://localhost/rentals](http://localhost/rentals).

You can add to your machine host file:

```shell
127.0.0.1 api.outdoorsy.com
```

Few examples:
1. [http://api.outdoorsy.com/rentals/12](http://api.outdoorsy.com/rentals/12)
2. [http://api.outdoorsy.com/rentals?limit=3&offset=6](http://api.outdoorsy.com/rentals?limit=3&offset=6)
3. [http://api.outdoorsy.com/rentals?price_min=9000&price_max=75000](http://api.outdoorsy.com/rentals?price_min=9000&price_max=75000)


### Running Unit Tests

To run unit tests, execute the following command:

```shell
go test ./...
```

This will run all unit tests in the project. You can use additional flags or specify specific test files or directories as needed.

### Running a load test

Please navigate to the loadtest dir and execute the following command.

```shell
docker run -p 8089:8089 -v $PWD:/mnt/locust/ locustio/locust -f /mnt/locust/locust.py
```

## Contributing

If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: `git checkout -b feature-name`
3. Make your changes and commit them: `git commit -m 'Add new feature'`
4. Push to the branch: `git push origin feature-name`
5. Create a pull request.