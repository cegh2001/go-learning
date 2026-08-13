// Guiding Question: Imagine you are using a third-party library where a single function requires
//  7 different bool arguments in a row,
// like Config(true, false, false, true, true, false, true).
// Why is that a nightmare for readability and a massive invitation for code misuse?

// R) Because arguments don't communicate or express the intention of the function clearly and easily

package main

import "fmt"

// 1. We group all configuration options into a descriptive struct 📦
type ServerConfig struct {
	EnableSSL    bool
	IsProduction bool
	DebugMode    bool
}

// 2. The function accepts the struct instead of a mess of booleans 🔌
func NewServer(config ServerConfig) {
	if config.EnableSSL {
		fmt.Println("Secure server initialized.")
	}
}

func main() {
	// The developer is forced to name the fields. Clear and safe! ✅
	NewServer(ServerConfig{
		EnableSSL:    true,
		IsProduction: false,
		DebugMode:    true,
	})
}
