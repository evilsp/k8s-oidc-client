package oidcInstaller

import (
	"fmt"
	osCommand "os"
	"os/exec"
	"runtime"
	"strings"
)

type checker struct {
	oidcCheck    string
	krewCheck    string
	kubectlCheck string
}

func (c checker) Check(os string, arch string) map[string]string {

	osPtr := new(string)
	var resultMap map[string]string
	homeDir, _ := osCommand.UserHomeDir() /* ~ will not be set to user root directory when it's used in go command*/

	switch os {
	case "darwin":
		*osPtr = ""
	case "windows":
		*osPtr = homeDir + "/"
	case "linux":
		*osPtr = ""
	}

	argsOidc := strings.Fields(*osPtr + c.oidcCheck)
	argsKrew := strings.Fields(*osPtr + c.krewCheck)
	argsKubectl := strings.Fields(*osPtr + c.kubectlCheck)

	// Check kubectl oidc-login version
	if err := exec.Command(argsOidc[0], argsOidc[1:]...).Run(); err != nil {

		fmt.Println("message:", err, "\nDetected oidc-login not installed")
		resultMap = map[string]string{"exitState": "oidcCheck", "arch": arch, "os": os}
		// oidc-login not installed
	}

	if err := exec.Command(argsKrew[0], argsKrew[1:]...).Run(); err != nil {

		fmt.Println("message:", err, "\nDetected krew not installed")
		resultMap = map[string]string{"exitState": "krewCheck", "arch": arch, "os": os}
		// krew and oidc-login not installed
	}
	
	if err := exec.Command(argsKubectl[0], argsKubectl[1:]...).Run(); err != nil {

		fmt.Println("message:", err, "\nDetected kubectl not installed")
		resultMap = map[string]string{"exitState": "kubectlCheck", "arch": arch, "os": os}
		// kubectl not installed
	}
	return resultMap
}

func Checker() map[string]string {

	checkTags := checker{
		"kubectl oidc-login version",
		"kubectl krew version",
		"kubectl",
	}

	os := runtime.GOOS

	// Get Struct Info
	arch := runtime.GOARCH

	fmt.Printf("Operating System: %s\n", os)
	fmt.Printf("Architecture: %s\n", arch)

	return checkTags.Check(os, arch)
}
