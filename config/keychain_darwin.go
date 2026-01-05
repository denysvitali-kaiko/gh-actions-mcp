//go:build darwin

package config

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/keybase/go-keychain"
)

// getTokenFromKeychain attempts to get the GitHub token from the macOS keychain
func getTokenFromKeychain() (string, error) {
	if runtime.GOOS != "darwin" {
		return "", fmt.Errorf("keychain is only available on macOS")
	}

	// GitHub CLI stores the token with service "github.com" and account "oauth"
	data, err := keychain.GetGenericPassword("github.com", "oauth", "", "")
	if err != nil {
		return "", fmt.Errorf("failed to get token from keychain: %w", err)
	}

	if data == nil || len(data) == 0 {
		return "", fmt.Errorf("token found but empty in keychain")
	}

	token := string(data)

	// Verify it looks like a GitHub token (starts with gho_)
	if !strings.HasPrefix(token, "gho_") {
		log.Debugf("Warning: token from keychain doesn't start with gho_, may not be a GitHub token")
	}

	return token, nil
}
