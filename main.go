package main

import (
        "fmt"
        "math/rand"
        "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func generatePassword(length int) string {
        var password string
        rand.Seed(time.Now().UnixNano())

        for i := 0; i < length; i++ {
                randomIndex := rand.Intn(len(charset))
                password += string(charset[randomIndex])
        }

        return password
}

func main() {
        var passwordLength int

        fmt.Println("Enter desired password length: ")
        fmt.Scanln(&passwordLength)

        password := generatePassword(passwordLength)
        fmt.Println("Generated password:", password)
}

