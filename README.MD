## Description

The package that I use to set my envs in Golang.

The application can receive environment variables through the -envs flag like this: **`go run cmd/main.go -envs "NAME=john LASTNAME=doe"`**.

Additionally, it can read environment variables from a .env file or any other file you specify by passing the file path to it as a parameter to the env.InitEnv() function.

