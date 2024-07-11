package env

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// LoadEnv loads environment variables from a file and command-line flags.
func LoadEnv(filePath string) ([]string, error) {
	keys, err := readEnvFile(filePath)
	if err != nil {
		return nil, err
	}
	keys2, err := readEnvFromFlags()
	if err != nil {
		return nil, err
	}

	mergedKeys := append(keys, keys2...)

	return mergedKeys, nil
}

// readEnvFile reads environment variables from a .env file and sets them in the OS environment.
func readEnvFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		env := scanner.Text()
		// Ignore blank lines and comments
		if len(env) == 0 || strings.HasPrefix(env, "#") {
			continue
		}
		// Split the variable from its value
		pair := strings.SplitN(env, "=", 2)
		// If there aren't exactly two elements, it's an error
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid .env file format: %s", env)
		}
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])
		os.Setenv(key, value)
		keys = append(keys, key)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return keys, nil
}

// readEnvFromFlags reads environment variables from command-line flags.
func readEnvFromFlags() ([]string, error) {
	envs := flag.String("envs", "", "Environment variables in the format 'KEY1=value1 KEY2=value2'")
	flag.Parse()
	var keys []string
	if *envs != "" {
		pairs := strings.Split(*envs, " ")
		for _, pair := range pairs {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				os.Setenv(key, value)
				keys = append(keys, key)
			} else {
				return nil, fmt.Errorf("invalid environment variable format: %s", pair)
			}
		}
	}
	return keys, nil
}
