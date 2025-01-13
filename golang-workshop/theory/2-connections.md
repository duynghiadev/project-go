# Connecting to MongoDB with Go

In this lesson, we will learn how to connect to a MongoDB database using Go. We'll cover setting up the connection, handling authentication, and ensuring the connection is successful.

## Prerequisites

Before we begin, ensure you have the following:
- MongoDB installed and running locally or on a server.
- Go installed on your machine.
- The MongoDB Go Driver installed. You can install it using:
  ```sh
  go get go.mongodb.org/mongo-driver/mongo
  ```

## Step-by-Step Guide

### 1. Import Required Packages

First, import the necessary packages. These include the MongoDB driver packages and standard Go packages for context, logging, and time management.

```go
import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

### 2. Set MongoDB Client Options

To connect to MongoDB, we need to set client options, including the URI of the MongoDB server and the authentication credentials.

```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
    Username: "root",
    Password: "example",
})
```

- **ApplyURI**: This method sets the MongoDB URI. Here, we are connecting to a MongoDB instance running on `localhost` on the default port `27017`.
- **SetAuth**: This method sets the authentication credentials. Replace `"root"` and `"example"` with your actual username and password.

### 3. Connect to MongoDB

We then create a context with a timeout to avoid indefinitely hanging connections. This context is used for the connection attempt.

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

- **context.WithTimeout**: Creates a context that will automatically cancel after the specified duration (10 seconds in this case).

Now, use the MongoDB client to connect to the database.

```go
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
```

- **mongo.Connect**: Establishes a new connection to the MongoDB server using the specified client options and context.

### 4. Verify the Connection

After connecting, it's crucial to verify that the connection is successful by pinging the MongoDB server.

```go
err = client.Ping(ctx, nil)
if err != nil {
    log.Fatal(err)
}
```

- **client.Ping**: Sends a ping command to the MongoDB server to verify the connection. If the server does not respond, an error is logged.

### Complete Code

Here's the complete code to connect to a MongoDB instance with authentication:

```go
package main

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Set MongoDB client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
        Username: "root",
        Password: "example",
    })

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}
```

### Explanation

1. **Setting Client Options**: We configure the client with the MongoDB URI and authentication credentials.
2. **Creating a Context**: A context with a timeout ensures the connection attempt does not hang indefinitely.
3. **Connecting to MongoDB**: We attempt to establish a connection to the MongoDB server.
4. **Verifying the Connection**: Pinging the server ensures that the connection is successful and the server is reachable.

By following these steps, you can establish a connection to a MongoDB database in Go, handle authentication, and verify the connection. This setup is crucial for any application that relies on MongoDB for data storage and retrieval.