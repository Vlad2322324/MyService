// main.go
package main

import "MyService/internal/repositories"

// Config хранит все настройки приложения
func main() {
	repositories.DBconnect()

}
