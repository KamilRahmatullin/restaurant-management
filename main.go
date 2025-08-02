package main

import (
	"github.com/kamilrahmatullin/restaurant-management/env"
)

func main() {
	port := env.GetValue("PORT", "8080")
}
