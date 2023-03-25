<h1 align="center">tiktoken-grpc-go</h1>

<p align="center">
  <a href="https://github.com/yorinasub17/tiktoken-grpc-go/blob/main/LICENSE">
    <img alt="LICENSE" src="https://img.shields.io/github/license/yorinasub17/tiktoken-grpc-go?style=for-the-badge">
  </a>
  <a href="https://goreportcard.com/report/github.com/yorinasub17/tiktoken-grpc-go">
    <img alt="goreportcard" src="https://goreportcard.com/badge/github.com/yorinasub17/tiktoken-grpc-go?style=for-the-badge">
  </a>
  <a href="https://pkg.go.dev/github.com/yorinasub17/tiktoken-grpc-go">
    <img alt="pkg.go reference" src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&style=for-the-badge">
  </a>
  <a href="https://github.com/yorinasub17/tiktoken-grpc-go/blob/main/go.mod">
    <img alt="go.mod version" src="https://img.shields.io/github/go-mod/go-version/yorinasub17/tiktoken-grpc-go?style=for-the-badge&logo=go&label=version">
  </a>
  <a href="https://github.com/yorinasub17/tiktoken-grpc-go/releases/latest">
    <img alt="latest release" src="https://img.shields.io/github/v/release/yorinasub17/tiktoken-grpc-go?style=for-the-badge">
  </a>
</p>


`tiktoken-grpc-go` provides Go client bindings for [the tiktoken gRPC
service](https://github.com/yorinasub17/tiktoken-grpc).

## Usage

> This example targets a tiktoken grpc service running locally. You can modify the host param in the `NewClient` call to
> target a different instance.

```go
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	tiktokengrpc "github.com/yorinasub17/tiktoken-grpc-go"
)

func main() {
	clt, err := tiktokengrpc.NewClient("localhost", 50051, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer clt.Close()

	ctx := context.Background()
	text := "antidisestablishmentarianism"

	cl100k, err := clt.NumTokens(ctx, tiktokengrpc.CL100kBase, text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cl100k_base: %d tokens\n", cl100k)

	p50k, err := clt.NumTokens(ctx, tiktokengrpc.P50kBase, text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p50k_base: %d tokens\n", p50k)

	gpt2, err := clt.NumTokens(ctx, tiktokengrpc.GPT2Encoding, text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("gpt2: %d tokens\n", gpt2)
}
```


## Installation

```
go get github.com/yorinasub17/tiktoken-grpc-go
```
