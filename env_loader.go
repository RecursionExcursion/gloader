package gloader

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const defaultEnv = ".env"

type EnvLoader struct {
	loaded bool
}

// returns value associated with key or an error if no key is found
func (el *EnvLoader) Get(key string) (string, error) {
	if !el.loaded {
		if err := el.LoadEnv(); err != nil {
			return "", err
		}
	}

	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("env key %v not set", key)
	}
	return val, nil
}

// returns value associated with key or panics if no key is found
func (el *EnvLoader) MustGet(key string) string {
	if !el.loaded {
		if err := el.LoadEnv(); err != nil {
			panic(err)
		}
	}

	val := os.Getenv(key)
	if val == "" {
		log.Panicf("Env key %v not set", key)
	}
	return val
}

// returns value associated with key or fallback if no key is found
func (el *EnvLoader) GetOrFallback(key string, fallback string) string {
	if !el.loaded {
		if err := el.LoadEnv(); err != nil {
			return fallback
		}
	}

	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// returns value associated with key or zero value ("") if no key is found
func (el *EnvLoader) GetOrDefault(key string) string {
	if !el.loaded {
		if err := el.LoadEnv(); err != nil {
			return ""
		}
	}
	return os.Getenv(key)
}

// LoadEnv loads environment variables from one or more files into the process's environment.
// Files should contain lines in the format KEY=VALUE.
// If no files are provided, it defaults to loading from ".env".
func (el *EnvLoader) LoadEnv(readers ...io.Reader) error {

	//If no readers are passed, fallback onto default env file (.env)
	if len(readers) == 0 {
		f, err := os.Open(defaultEnv)
		if err != nil {
			return err
		}
		defer f.Close()
		readers = []io.Reader{f}
	}

	for _, r := range readers {
		if err := parseEnv(r); err != nil {
			return err
		}
	}
	el.loaded = true
	return nil
}

func parseEnv(r io.Reader) error {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		/* TODO Need to impl logic to cut out trailing comments
		 * '#' should only be permitted inside quotes
		 * if not then everything after needs to be treated as a comment
		 */
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)
		os.Setenv(key, val)
	}

	return scanner.Err()
}
