# asuswrt-api
A Client API for the Asus WRT Router.

## Description

//ToDo

## Building

Execute the following commands:

1. Build the project:

   ```shell
   go build -o build/asuswrt-api
   ```

## Running

Build and run:

1. Run the project:

   ```shell
   go build -o build/asuswrt-api
   ./build/asuswrt-api [flags]
   ```

Or simply use Go Run:

1. Run the project:

   ```shell
   go build -o build/asuswrt-api
   go run ./ [flags]
   ```

## Updating Dependencies

Execute the following commands:

1. Update the dependencies:

   ```shell
   go mod tidy -compat=1.19
   ```
