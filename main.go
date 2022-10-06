package main

import (
	"context"
	"fmt"
"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    fmt.Println("Hello, World!")
}
