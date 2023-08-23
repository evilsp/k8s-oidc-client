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

	exec.Command("echo", insCom.KubectlOIDCInstallCom, ">", "./kubectlOIDCInstall.sh")
	exec.Command("echo", insCom.kubectlKrewInstallCom, ">", "./kubectlKrewInstall.sh")
	exec.Command("echo", insCom.kubectlInstallCom, ">", "./kubectlInstall.sh")
	switch stage {
	case "oidcCheck":
		exec.Command("sh", "-c", "./kubectlOIDCInstall.sh")
	case "krewCheck":
		exec.Command("sh", "-c", "./kubectlOIDCInstall.sh")
		exec.Command("sh", "-c", "./kubectlKrewInstall.sh")
	case "kubectlCheck":
		exec.Command("sh", "-c", "./kubectlOIDCInstall.sh")
		exec.Command("sh", "-c", "./kubectlKrewInstall.sh")
		exec.Command("sh", "-c", "./kubectlInstall.sh")
	}
}

func (insCom installCommand) winInstall(stage string) {

	homeDir, _ := osCommand.UserHomeDir()
	exec.Command("echo", insCom.KubectlOIDCInstallCom, ">", "./kubectlOIDCInstall.bat")
	exec.Command("echo", insCom.kubectlKrewInstallCom, ">", "./kubectlKrewInstall.bat")
	exec.Command("echo", insCom.kubectlInstallCom, ">", "./kubectlInstall.bat")
	switch stage {
	case "oidcCheck":
		exec.Command(homeDir + "/kubectlOIDCInstall.bat")
	case "krewCheck":
		exec.Command(homeDir + "/kubectlOIDCInstall.bat")
		exec.Command(homeDir + "/kubectlKrewInstall.bat")
	case "kubectlCheck":
		exec.Command(homeDir + "/kubectlOIDCInstall.bat")
		exec.Command(homeDir + "/kubectlKrewInstall.bat")
		exec.Command(homeDir + "/kubectlInstall.bat")
	}
}

func Installer() {

	linuxAmd64 := installCommand{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\"",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.bashrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	linuxArm64 := installCommand{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl\"",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.bashrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	windows := installCommand{
		kubectlInstallCom:     "curl.exe -LO \"https://dl.k8s.io/release/v1.28.0/bin/windows/amd64/kubectl.exe\"",
		kubectlKrewInstallCom: "wget https://github.com/kubernetes-sigs/krew/releases/download/v0.4.4/krew.exe -o ./krew.exe\n  .\\krew install krew\n setx PATH \"%PATH%;%USERPROFILE%\\.krew\\bin\"\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	macOSIntel := installCommand{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl\" \n",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.zshrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login",
	}

	macOSAppleSilicon := installCommand{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/arm64/kubectl\" \n",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.zshrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login",
	}

	resultMap := Checker()
	os := resultMap["os"]
	arch := resultMap["arch"]
	stage := resultMap["exitState"]

	if os == "linux" {
		// 进一步判断 Linux 下的体系结构

		switch arch {
		case "amd64", "386":
			linuxAmd64.unixInstall(stage)
		case "arm64":
			linuxArm64.unixInstall(stage)
		default:
			fmt.Println("Unknown architecture on Linux")
		}
	} else if os == "windows" {
		switch arch {
		case "amd64", "386":
			windows.winInstall(stage)
		default:
			fmt.Println("Unknown architecture on Windows")
		}
	} else if os == "darwin" {
		switch arch {
		case "amd64":
			macOSIntel.unixInstall(stage)
		case "arm64":
			macOSAppleSilicon.unixInstall(stage)
		default:
			fmt.Println("Unknown architecture on macOS")
		}
	} else {
		fmt.Println("Unknown operating system")
	}
}
