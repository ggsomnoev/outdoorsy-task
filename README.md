# Outdoorsy Task

A simple go microservice which serves rentals data based on some criteria

## Prerequisites

- [Go](https://golang.org/) (version 1.21 or higher)

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

To run the project, use the following command:

```shell
go run main.go
```

This will start the application, and you can access it at [http://localhost:8080](http://localhost:8080).
Few examples:
[http://localhost:8080/rentals/12](http://localhost:8080/rentals/12)
[http://localhost:8080/rentals?price_min=9000&price_max=75000](http://localhost:8080/rentals?price_min=9000&price_max=75000)
[http://localhost:8080/rentals?limit=3&offset=6](http://localhost:8080/rentals?limit=3&offset=6)


### Running Unit Tests

To run unit tests, execute the following command:

```shell
go test ./...
```

This will run all unit tests in the project. You can use additional flags or specify specific test files or directories as needed.

## Configuration

Explain how to configure the project, including environment variables, configuration files, or any other relevant settings.

## Contributing

If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: `git checkout -b feature-name`
3. Make your changes and commit them: `git commit -m 'Add new feature'`
4. Push to the branch: `git push origin feature-name`
5. Create a pull request.