# Inventory Application

This is the repository for a simple inventory tracking and management application. The purpose of the application is to track how much customers spend with us, and the available inventory in our warehouse.

The application processes a text file provided as a command line argument, and prints the resulting reports to the console.

## Requirements

- Go version 1.19
- Project uses Go modules

## Usage

```
// run the project
$ go run cmd/inventory/main.go ./path/to/file.txt

// run tests
$ go test -v ./internal/handler ./internal/service
```

## Methodology

- 3 models - Customer, Order and Product. We used the customer and product name as identifiers, but I also added an ID to all models as a best practice. The reason we use name as identifier is because we are provided that by the input, and not an ID. We keep the models simple and lightweight.
- Ideally, the order model would not include the total price of the order, we could just keep the quantity. The decision to include total order price was made to make the report generation more performative. In a more robust project reports would be generated from an OLAP system like a datawarehouse, so we would not need this total price in the order object.
- I implemented a repository interface to abstract away the implementation of our storage. This application uses memory to store objects, but this design allows easy implementation of a DB. The interface also helps with testing, as I can specify expected objects in the testdb.
- Beyond models and repositories, the application is divided into 3 main modules, each with their own single responsibility. Handler is responsible for parsing the input and validating. Service is responsible for the business logic and interacting with models. Report is responsible for generating the formatted reports.
- We use dependency injection in main.go to allow for decoupling of our modules.
- Product Price is set to cents instead of dollars to avoid any precision issues when using floats. The only time a float is used it to calculate average order value.
- All public functions (even obvious ones) have comments for documentation generating.
- I've provided a sample input.txt file in the root directory of the project. To run, open a terminal to the root directory and use command `$ go run cmd/inventory/main.go ./input.txt`
