package utils

import (
    "encoding/json"
    "log"
    "os"
)

func ReadJsonFile(filename string, v interface{}) {
    file, err := os.Open(filename)
    if err != nil {
        log.Printf("Error opening file %s: %v", filename, err)
        return
    }
    defer file.Close()

    if err := json.NewDecoder(file).Decode(v); err != nil {
        log.Printf("Error decoding JSON from file %s: %v", filename, err)
    }
}
