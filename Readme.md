# MTA Server Optimization

## Setup Instructions

1. Open the `app.go` file in your project.

2. Initialize the Go module by running the following command in your terminal or command prompt:

    ```bash
    go mod init mta-server-optimiser
    ```

   This command creates a `go.mod` file that tracks the project's dependencies.

3. Run the following command to ensure the `go.mod` file reflects the correct and updated dependencies based on your code:

    ```bash
    go mod tidy
    ```

4. Use the following command to create a `vendor` directory containing all the dependencies required by your project:

    ```bash
    go mod vendor
    ```

   This step ensures that you have a local copy of the dependencies.

5. Build the project by running:

    ```bash
    go build
    ```

   This command compiles the Go code and generates an executable binary file.

6. Finally, run the server by executing:

    ```bash
    go run app.go
    ```

   This command starts the server and runs your application.

**Note:** Ensure that you are in the correct directory when running these commands. Adjust the commands accordingly if your project structure or file names are different.

## Server Endpoint

After setup, the server will be accessible at: [http://localhost:8080/mta-hosting-optimizer](http://localhost:8080/mta-hosting-optimizer)

## Testing

All unit and integration test cases are written in the `app_test.go` and `helper_test.go` file.

To modify the value of "X," use the `getEnv` function. You can make adjustments either within the `getInstanceName` function or by setting the value directly in the `.env` file.

Here already adjusted for the passing the value by X=1

## Output

The expected output of the server optimization process will be an array:

```json
[ "mta-prod-1", "mta-prod-3" ]
