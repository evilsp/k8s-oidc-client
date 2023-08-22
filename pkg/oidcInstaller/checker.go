package oidcInstaller

import (
	"fmt"
	"os/exec"
)

func Checker() int {

	// Check kubectl oidc-login version
	if err := exec.Command("kubectl", "oidc-login", "version").Run(); err != nil {
		// Execute command group A
		fmt.Printf("Detected oidc-login not installed")
		return 2 // oidc-login not installed
	}
	if err := exec.Command("kubectl", "krew", "version").Run(); err != nil {
		// Execute command group A
		fmt.Printf("Detected oidc-login not installed")
		return 3 // krew and oidc-login not installed
	}
	if err := exec.Command("kubectl", "version").Run(); err != nil {
		// Execute command group A
		fmt.Printf("Detected oidc-login not installed")
		return 4 // kubectl not installed
	}

	return 1
}
