package main

import (
	"fmt"

	"github.com/seba000/go-env/env"
)

func main() {
	// Initialize the environment variables, returns a pointer to the Envs struct.
	env := env.InitEnv(".env")
	// Print the environment variables. Feel free to remove this line.
	fmt.Println(env)
}
