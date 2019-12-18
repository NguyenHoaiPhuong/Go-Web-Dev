# Machinery

Machinery is an asynchronous task queue/job queue based on distributed message passing.

## Basic concepts

1. Configuration

    The config package has convenience methods for loading configuration from environment variables or a YAML file:

    ```
    // loading configuration from env variables
    cnf, err := config.NewFromEnvironment(true)

    // loading configuration from YAML file
    cnf, err := config.NewFromYaml("config.yml", true)
    ```

2. Server

A Machinery library must be instantiated before use. The way this is done is by creating a Server instance. Server is a base object which stores Machinery configuration and registered tasks.

3. Worker

4. Tasks

- Registering tasks
- Signatures
- Supported types
- Sending tasks
- Delayed tasks
- Retry tasks
- Get pending tasks
- Keeping results

5. Workflows

- Groups
- Chords
- Chains

## References

1. https://github.com/RichardKnop
