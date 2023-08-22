package oidcInstaller

import (
	"fmt"
	"os/exec"
	"runtime"
)

type install_command struct {
	kubectlInstallCom     string
	kubectlKrewInstallCom string
	KubectlOIDCInstallCom string
}

func installer() {

	linux_amd64 := install_command{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\"",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.bashrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	linux_arm64 := install_command{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl\"",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.bashrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	windows := install_command{
		kubectlInstallCom:     "curl.exe -LO \"https://dl.k8s.io/release/v1.28.0/bin/windows/amd64/kubectl.exe\"",
		kubectlKrewInstallCom: "wget https://github.com/kubernetes-sigs/krew/releases/download/v0.4.4/krew.exe -o ./krew.exe\n  .\\krew install krew\n setx PATH \"%PATH%;%USERPROFILE%\\.krew\\bin\"\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login\n",
	}

	macOS_intel := install_command{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl\" \n",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.zshrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login",
	}

	macOS_apple_silicon := install_command{
		kubectlInstallCom:     "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/arm64/kubectl\" \n",
		kubectlKrewInstallCom: "set -x; cd \"$(mktemp -d)\" &&\n  OS=\"$(uname | tr '[:upper:]' '[:lower:]')\" &&\n  ARCH=\"$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\\(arm\\)\\(64\\)\\?.*/\\1\\2/' -e 's/aarch64$/arm64/')\" &&\n  KREW=\"krew-${OS}_${ARCH}\" &&\n  curl -fsSLO \"https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz\" &&\n  tar zxvf \"${KREW}.tar.gz\" &&\n  ./\"${KREW}\" install krew\n echo \"export PATH=\"${KREW_ROOT:-$HOME/.krew}/bin:$PATH\"\" >> ~/.zshrc\n",
		KubectlOIDCInstallCom: "kubectl krew install oidc-login",
	}

	// Get OS Info
	os := runtime.GOOS

	// Get Struct Info
	arch := runtime.GOARCH

	fmt.Printf("Operating System: %s\n", os)
	fmt.Printf("Architecture: %s\n", arch)

	if os == "linux" {
		// 进一步判断 Linux 下的体系结构
		cmd := exec.Command("sh", "-c", check.checkKubectlOIDCCom)

		switch arch {
		case "amd64", "386":
			cmd := exec.Command("sh", "-c", linux_amd64.kubectlInstallCom)
			fmt.Println("Linux x86-64 or x86")
		case "arm64":
			fmt.Println("Linux ARM64")
		default:
			fmt.Println("Unknown architecture on Linux")
		}
	} else if os == "windows" {
		switch arch {
		case "amd64", "386":
			fmt.Println("Windows x86-64 or x86")
		default:
			fmt.Println("Unknown architecture on Windows")
		}
	} else if os == "darwin" {
		switch arch {
		case "amd64":
			fmt.Println("macOS x86-64")
		case "arm64":
			fmt.Println("macOS ARM64")
		default:
			fmt.Println("Unknown architecture on macOS")
		}
	} else {
		fmt.Println("Unknown operating system")
	}
}
