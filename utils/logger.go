package utils

import (
    "log"
    "os"
)

func InitLogger() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
}

func Log(message string) {
    log.Println(message)
}
