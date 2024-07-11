package env

import (
	"fmt"
	"os"
	"strconv"
)

// Envs is a struct that holds the environment variables.
type Envs struct {
	Name   string
	Age    int
	Member bool
	Salary float64
}

// InitEnv initializes the environment variables.
func InitEnv(filePath string) *Envs {
	// If the file path is empty, use the default file name.
	if filePath == "" {
		filePath = ".env"
	}

	keys, err := LoadEnv(filePath)
	if err != nil {
		panic(err)
	}
	//Log all the keys of the envs that are loaded. feel free to remove this line and the return of keys.
	fmt.Println(keys)

	//TO INT
	age, err := strconv.Atoi(os.Getenv("AGE"))
	if err != nil {
		panic(err)
	}
	//TO BOOL
	member, err := strconv.ParseBool(os.Getenv("MEMBER"))
	if err != nil {
		panic(err)
	}

	//TO FLOAT
	salary, err := strconv.ParseFloat(os.Getenv("SALARY"), 64)
	if err != nil {
		panic(err)
	}

	return &Envs{
		Name:   os.Getenv("NAME"),
		Age:    age,
		Member: member,
		Salary: salary,
	}
}
