package main

import (
	"log"

	"block-chain-inspection/block"
)

func main() {
	log.Println("START")
	defer log.Println("END")

	block1 := block.NewBlock(block.NewData("Hello"), nil)

	block1Hash, err := block1.Sha256()
	if err != nil {
		log.Fatal(err)
	}

	block2 := block.NewBlock(block.NewData("Hello"), block1Hash)

	log.Printf("%+v\n", block2)
}
