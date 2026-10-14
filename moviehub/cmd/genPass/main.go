package main

import (
    "fmt"
    //"strings"
    //"time"

    //"golang.org/x/crypto/bcrypt"

    "moviehub/auth"
)

func main() {
    // 1) รหัสเดียวกัน บดสองครั้ง
    a, _ := auth.HashPassword("ISIS")
    b, _ := auth.HashPassword("ISIS")
    fmt.Println(a)
    fmt.Println(b)
}