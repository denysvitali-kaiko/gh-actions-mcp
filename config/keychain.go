//go:build !darwin

package config

import "fmt"

// getTokenFromKeychain attempts to get the GitHub token from the macOS keychain
func getTokenFromKeychain() (string, error) {
	return "", fmt.Errorf("keychain is only available on macOS")
}
