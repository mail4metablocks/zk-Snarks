package main

import (
	"crypto/rand"
	"fmt"

	"github.com/nadavgolan/zk-snarks"
)

func main() {
	// Define the statement we want to prove
	statement := []byte("This is the statement we want to prove")

	// Generate a random proof
	proof, err := zk_snarks.Prove(statement, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verify the proof
	err = zk_snarks.Verify(statement, proof)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Proof verified successfully")
}
