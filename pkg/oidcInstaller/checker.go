package oidcInstaller

import (
	"fmt"
	"os/exec"
	"runtime"
)

type checker struct {
	oidcCheck    string
	krewCheck    string
	kubectlCheck string
}

func (c checker) Check(os string) map[string]string {

	strPtr := new(string)
	resultMap := new(map[string]string)

	switch os {
	case "darwin":
		*strPtr = "sh -C"
	case "windows":
		*strPtr = "cmd /C"
	case "linux":
		*strPtr = "sh -C"
	}
	// Check kubectl oidc-login version
	if err := exec.Command(*strPtr, c.oidcCheck).Run(); err != nil {

		fmt.Printf("Detected oidc-login not installed")
		*resultMap = map[string]string{"exitState": "oidcCheck"}
		// oidc-login not installed
	}
	if err := exec.Command(*strPtr, c.krewCheck).Run(); err != nil {

		fmt.Printf("Detected oidc-login not installed")
		*resultMap = map[string]string{"exitState": "krewCheck"}
		// krew and oidc-login not installed
	}
	if err := exec.Command(*strPtr, c.kubectlCheck).Run(); err != nil {

		fmt.Printf("Detected oidc-login not installed")
		*resultMap = map[string]string{"exitState": "kubectlCheck"}
		// kubectl not installed
	}
	return *resultMap
}

func Checker() map[string]string {

	checkTags := checker{
		"kubectl oidc-login version",
		"kubectl krew version",
		"kubectl version",
	}

	os := runtime.GOOS

	// Get Struct Info
	arch := runtime.GOARCH

	fmt.Printf("Operating System: %s\n", os)
	fmt.Printf("Architecture: %s\n", arch)

	return checkTags.Check(os)
}
