package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// The function `LoadEnv` reads an environment file, splits each line into key-value pairs, and sets
// the corresponding environment variables.
func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

// The GetTag function in Go splits a tag string into a tag name and a map of tag attributes.
func GetTag(tag string) (string, map[string]string) {
	parts := strings.Split(tag, ",")
	tagName := parts[0]
	tagAttrs := make(map[string]string)

	for _, attr := range parts[1:] {
		keyValue := strings.Split(attr, ":")
		if len(keyValue) == 2 {
			tagAttrs[keyValue[0]] = keyValue[1]
		}
	}

	return tagName, tagAttrs
}
