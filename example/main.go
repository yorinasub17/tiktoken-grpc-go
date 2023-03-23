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
