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
   go run ./... [flags]
   ```

## Updating Dependencies

Execute the following commands:

1. Update the dependencies:

   ```shell
   go mod tidy -compat=1.19
   ```

## Testing

Execute the following commands:

1. Run all tests:

   ```shell
   go test -v ./...
   ```

## References

* [vrachieru/asuswrt-api](https://github.com/vrachieru/asuswrt-api)
* [Vaskivskyi/asusrouter](https://github.com/Vaskivskyi/asusrouter)
* [kennedyshead/aioasuswrt](https://github.com/kennedyshead/aioasuswrt)
