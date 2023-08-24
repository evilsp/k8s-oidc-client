package oidcInstaller

import (
	"fmt"
	osCommand "os"
	"os/exec"
)

type installCommand struct {
	kubectlInstallCom     string
	kubectlKrewInstallCom string
	KubectlOIDCInstallCom string
}

func (insCom installCommand) unixInstall(stage string) {

	fmt.Println(exec.Command("/bin/sh", "-c", "echo", insCom.KubectlOIDCInstallCom, ">", "./kubectlOIDCInstall.sh", "&&", "chmod", "u+x", "./kubectlOIDCInstall.sh").Run())
	fmt.Println(exec.Command("/bin/sh", "-c", "echo", insCom.kubectlKrewInstallCom, ">", "./kubectlKrewInstall.sh", "&&", "chmod", "u+x", "./kubectlKrewInstall.sh").Run())
	fmt.Println(exec.Command("/bin/sh", "-c", "echo", insCom.kubectlInstallCom, ">", "./kubectlInstall.sh", "&&", "chmod", "u+x", "./kubectlInstall.sh").Run())
	switch stage {
	case "oidcCheck":
		if err := exec.Command("/bin/sh", "-c", "./kubectlOIDCInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "krewCheck":
		if err := exec.Command("/bin/sh", "-c", "./kubectlKrewInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", "./kubectlOIDCInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "kubectlCheck":
		if err := exec.Command("/bin/sh", "-c", "./kubectlInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", "./kubectlKrewInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", "./kubectlOIDCInstall.sh").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	}
}

func (insCom installCommand) winInstall(stage string) {

	homeDir, _ := osCommand.UserHomeDir()
	exec.Command("echo", insCom.KubectlOIDCInstallCom, ">", "./kubectlOIDCInstall.bat").Run()
	exec.Command("echo", insCom.kubectlKrewInstallCom, ">", "./kubectlKrewInstall.bat").Run()
	exec.Command("echo", insCom.kubectlInstallCom, ">", "./kubectlInstall.bat").Run()
	switch stage {
	case "oidcCheck":
		if err := exec.Command(homeDir + "/kubectlOIDCInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "krewCheck":
		if err := exec.Command(homeDir + "/kubectlKrewInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command(homeDir + "/kubectlOIDCInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "kubectlCheck":
		if err := exec.Command(homeDir + "/kubectlInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl went error", err)
		}
		if err := exec.Command(homeDir + "/kubectlKrewInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command(homeDir + "/kubectlOIDCInstall.bat").Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	}
}

func Installer() {

	kctlAmd64LinuxInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\""
	kctlArm64LinuxInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl\""
	kctlWindowsInstallComm := "curl.exe -LO \"https://dl.k8s.io/release/v1.28.0/bin/windows/amd64/kubectl.exe\""
	kctlAmd64MacOSInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl\""
	kctlArm64MacOSInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/arm64/kubectl\""

	krewBashInstallComm := `  set -x; cd "$(mktemp -d)" &&
  		OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  		ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  		KREW="krew-\${OS}_\${ARCH}" &&
  		curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/\${KREW}.tar.gz" &&
  		tar zxvf "\${KREW}.tar.gz" &&
		./"\${KREW}" install krew
  		echo 'export PATH="\${KREW_ROOT:-\$HOME/.krew}/bin:\$PATH"' >> ~/.bashrc
		source ~/.bashrc
`

	krewZshInstallComm := `  set -x; cd "$(mktemp -d)" &&
  		OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  		ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  		KREW="krew-\${OS}_\${ARCH}" &&
  		curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/\${KREW}.tar.gz" &&
  		tar zxvf "\${KREW}.tar.gz" &&
		./"\${KREW}" install krew
  		echo 'export PATH="\${KREW_ROOT:-\$HOME/.krew}/bin:\$PATH"' >> /.zshrc
		source ~/.zshrc
	`

	krewWindowsInstallComm := `  wget https://github.com/kubernetes-sigs/krew/releases/download/v0.4.4/krew.exe -o ./krew.exe
		.\krew install krew
		setx PATH "%PATH%;%USERPROFILE%\.krew\bin"
	`

	kctlOidcLoginInstallComm := "kubectl krew install oidc-login"

	linuxAmd64 := installCommand{
		kubectlInstallCom:     kctlAmd64LinuxInstallComm,
		kubectlKrewInstallCom: krewBashInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
	}

	linuxArm64 := installCommand{
		kubectlInstallCom:     kctlArm64LinuxInstallComm,
		kubectlKrewInstallCom: krewBashInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
	}

	windows := installCommand{
		kubectlInstallCom:     kctlWindowsInstallComm,
		kubectlKrewInstallCom: krewWindowsInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
	}

	macOSIntel := installCommand{
		kubectlInstallCom:     kctlAmd64MacOSInstallComm,
		kubectlKrewInstallCom: krewZshInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
	}

	macOSAppleSilicon := installCommand{
		kubectlInstallCom:     kctlArm64MacOSInstallComm,
		kubectlKrewInstallCom: krewZshInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
	}

	os := new(string)
	arch := new(string)
	stage := new(string)

	if resultMap := Checker(); resultMap != nil {
		*os = resultMap["os"]
		*arch = resultMap["arch"]
		*stage = resultMap["exitState"]
	} else {
		fmt.Println("fully Installed!")
		osCommand.Exit(0)
	}

	if *os == "linux" {
		// 进一步判断 Linux 下的体系结构

		switch *arch {
		case "amd64", "386":
			linuxAmd64.unixInstall(*stage)
		case "arm64":
			linuxArm64.unixInstall(*stage)
		default:
			fmt.Println("Unknown architecture on Linux")
		}
	} else if *os == "windows" {
		switch *arch {
		case "amd64", "386":
			windows.winInstall(*stage)
		default:
			fmt.Println("Unknown architecture on Windows")
		}
	} else if *os == "darwin" {
		switch *arch {
		case "amd64":
			macOSIntel.unixInstall(*stage)
		case "arm64":
			macOSAppleSilicon.unixInstall(*stage)
		default:
			fmt.Println("Unknown architecture on macOS")
		}
	} else {
		fmt.Println("Unknown operating system")
	}
}
