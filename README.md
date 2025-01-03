# Redis Easy Driver

This project provides a simple interface to interact with a Redis database using Go. It includes functions to create, retrieve, update, and delete sets in Redis, making it easier to manage data.

## Features

- **CreateSet**: Store a key-value pair in Redis with an expiration time.
- **FetchValue**: Retrieve a value from Redis based on a key.
- **UpdateValue**: Update the value of an existing key in Redis with an optional new expiration time.
- **Delete**: Delete a key-value pair from Redis.

## Installation

To use this project, you need to have Go and Redis installed on your machine.

1. Install Go: [Go Installation Guide](https://golang.org/doc/install)
2. Install Redis: [Redis Installation Guide](https://redis.io/download)

## Usage

### Create a Redis Client

First, create a Redis client and initialize the `RedisDB` struct:

```go
package main

import (
    "github.com/coles243/red/red"
    "github.com/redis/go-redis/v9"
)

func main() {
    test := red.RedisDB{
        DB: redis.NewClient(&redis.Options{
            Addr: "localhost:6379", // Redis server address
        }),
    }
}
```

### Create a Set

Use the `CreateSet` method to store a key-value pair in Redis with an expiration time:

```go
data, err := test.CreateSet("key", "value", time.Duration(10)*time.Second)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(data)
}
```

### Retrieve a Set

Use the `FetchValue` method to retrieve a value from Redis based on a key:

```go
getdata, err := test.FetchValue("key")
if err != nil {
    fmt.Println(err)
    return
}

if strValue, ok := getdata.(string); ok {
    fmt.Println("Retrieved value:", strValue)
} else {
    fmt.Println("Unexpected type")
}
```

### Update a Set

Use the `UpdateValue` method to update the value of an existing key in Redis. This method also allows you to set a new expiration time:

```go
response, err := test.UpdateValue("key", "newValue", time.Duration(20)*time.Second)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### Delete a Set

Use the `Delete` method to delete a key-value pair from Redis:

```go
response, err := test.Delete("key")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Deleted records:", response)
}
```

### Example

Here's a complete example that demonstrates how to use the `CreateSet`, `FetchValue`, `UpdateValue`, and `Delete` methods:

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/coles243/red/red"
    "github.com/redis/go-redis/v9"
)

func main() {
    test := red.RedisDB{
        DB: redis.NewClient(&redis.Options{
            Addr: "localhost:6379",
        }),
    }

    type Author struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }

    var person Author
    jsonz, err := json.Marshal(Author{Name: "Elliot", Age: 25})
    if err != nil {
        fmt.Println(err)
    }

    test.CreateSet("al", jsonz, time.Duration(10)*time.Second)

    getdata, err := test.FetchValue("al")
    if err != nil {
        fmt.Println(err)
        return
    }

    v, ok := getdata.(string)
    if ok {
        err = json.Unmarshal([]byte(v), &person)
        if err != nil {
            fmt.Println("Error decoding JSON:", err)
            return
        }
        fmt.Printf("Decoded Author: %+v\n", person)
    } else {
        fmt.Println("Unexpected type")
    }

    test.UpdateValue("al", Author{Name: "Elliot", Age: 30}, time.Duration(20)*time.Second)

    response, err := test.Delete("al")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Deleted records:", response)
    }
}
```

## Benefits

- **Simplicity**: Provides a simple interface to interact with Redis.
- **Flexibility**: Supports storing and retrieving various data types.
- **Efficiency**: Uses Redis for fast data storage and retrieval.
