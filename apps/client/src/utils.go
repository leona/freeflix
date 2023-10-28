package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func base64ToHex(input string) string {
	p, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		log.Println("Failed to decode:", input)
		log.Panic(err)
	}

	return hex.EncodeToString(p)
}

func DefaultString(input string, defaultValue string) string {
	if input == "" {
		return defaultValue
	}
	return input
}

func DefaultInt(input string, defaultValue int) int {
	if input == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(input)

	if err != nil {
		return defaultValue
	}

	return i
}

func DefaultSlice(input string, defaultValue []string) []string {
	if input == "" {
		return defaultValue
	}

	split := strings.Split(input, ",")

	for i, item := range split {
		split[i] = strings.ToLower(strings.TrimSpace(item))
	}

	return split
}

func stringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

func roundFloat64(input float64, places int) float64 {
	rounding := 1.0

	for i := 0; i < places; i++ {
		rounding *= 10.0
	}

	return float64(int(input*rounding)) / rounding
}

func GetRandomFile(path string, extension string) (string, error) {
	var files []string

	// read the files in the directory
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() { // skip if it is a directory
			if filepath.Ext(path) == "."+extension {
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", nil
	}

	// select a random file
	rand.Seed(time.Now().Unix())
	randomFile := files[rand.Intn(len(files))]

	return randomFile, nil
}
