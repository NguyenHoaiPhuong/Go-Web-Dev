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

    In order to consume tasks, you need to have one or more workers running. All you need to run a worker is a Server instance with registered tasks.

    ```
    worker := server.NewWorker("worker_name", 10)
    err := worker.Launch()
    if err != nil {
    // do something with the error
    }
    ```

    Each worker will only consume registered tasks. For each task on the queue the Worker.Process() method will be run in a goroutine. Use the second parameter of server.NewWorker to limit the number of concurrently running Worker.Process() calls (per worker). Example: 1 will serialize task execution while 0 makes the number of concurrently executed tasks unlimited (default).

4. Tasks

    Tasks are a building block of Machinery applications. A task is a function which defines what happens when a worker receives a message.

    Each task needs to return an error as a last return value. In addition to error tasks can now return any number of arguments.

    - Registering tasks

        Before your workers can consume a task, you need to register it with the server. This is done by assigning a task a unique name:

        ```
        server.RegisterTasks(map[string]interface{}{
            "add":      Add,
            "multiply": Multiply,
        })
        ```

        Tasks can also be registered one by one:

        ```
        server.RegisterTask("add", Add)
        server.RegisterTask("multiply", Multiply)
        ```

        Simply put, when a worker receives a message like this:

        ```
        {
            "UUID": "48760a1a-8576-4536-973b-da09048c2ac5",
            "Name": "add",
            "RoutingKey": "",
            "ETA": null,
            "GroupUUID": "",
            "GroupTaskCount": 0,
            "Args": [
                {
                "Type": "int64",
                "Value": 1,
                },
                {
                "Type": "int64",
                "Value": 1,
                }
            ],
            "Immutable": false,
            "RetryCount": 0,
            "RetryTimeout": 0,
            "OnSuccess": null,
            "OnError": null,
            "ChordCallback": null
        }
        ```

        It will call Add(1, 1). Each task should return an error as well so we can handle failures.

    - Signatures
    - Supported types
    - Sending tasks
    - Delayed tasks
    - Retry tasks
    - Get pending tasks
    - Keeping results

5. Workflows

Running a single asynchronous task is fine but often you will want to design a workflow of tasks to be executed in an orchestrated way. There are couple of useful functions to help you design workflows.

- Groups

    `Group` is a set of tasks which will be executed in parallel, independent of each other.

- Chords

    `Chord` allows you to define a callback to be executed after all tasks in a group finished processing.

- Chains

    `Chain` is simply a set of tasks which will be executed one by one, each successful task triggering the next task in the chain.

## Remarks
- 01_Example and 02_Example : send task
- 03_Example : send a group of task
- 04_Example : send a chord
- 05_Example : send a chain

Below is the content of request need to be sent to 

```
POST /tasks HTTP/1.1
Host: localhost:9000
Content-Type: application/javascript
User-Agent: PostmanRuntime/7.20.1
Accept: */*
Cache-Control: no-cache
Postman-Token: 06933167-de70-4dac-9e24-ddba448c751a,8057dec4-4fa5-4bd6-9fdc-f74cccbd609e
Host: localhost:9000
Accept-Encoding: gzip, deflate
Content-Length: 127
Connection: keep-alive
cache-control: no-cache

[
	{
		"taskName": "add",
		"args": [1, 2, 3]
	},
	{
		"taskName": "add",
		"args": [1, 4]
	},
	{
		"taskName": "multiply"
	}
]
```

## References

1. https://github.com/RichardKnop
