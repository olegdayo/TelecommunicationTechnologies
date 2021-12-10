package main

import (
	"ProgramExamples/cryptoAlgos"
	"fmt"
)

func main() {
	var power int = cryptoAlgos.GetPower(3, 2, 5)
	fmt.Println("The power is", power)

	alice := cryptoAlgos.NewAlice(6, 23, 5)
	bob := cryptoAlgos.NewBob(15, 23, 5)
	key1, key2 := cryptoAlgos.GetKeys(alice, bob)

	fmt.Println("Are keys equal:", key1.Int64() == key2.Int64())
	fmt.Println("Alice's key:", key1)
	fmt.Println("Bob's key:", key2)
}
