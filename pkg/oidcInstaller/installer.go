package oidcInstaller

import (
	"fmt"
	"log"
	osCommand "os"
	"os/exec"
)

type installCommands struct {
	kubectlInstallCom     string
	kubectlKrewInstallCom string
	KubectlOIDCInstallCom string
	errorInfo             installErrors
	directory             string
}

type installErrors struct {
	kubectlInstallError     string
	kubectlKrewInstallError string
	KubectlOIDCInstallError string
}

func (insCom installCommands) fileCreator(filename string, content string) {

	/* Create file with Path as Directory/filename */

	filePath := insCom.directory + "/" + filename

	err := osCommand.MkdirAll(insCom.directory, 0755) // 0755表示默认的目录权限
	if err != nil {
		fmt.Println("Menu create failed with error: ", err)
	}

	file, err := osCommand.Create(filePath)
	if err != nil {
		fmt.Println("File create failed with error: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("File write failed with error: ", err)
	}

	err = osCommand.Chmod(filePath, 0755)
	if err != nil {
		fmt.Println("File grant priviledge failed with error: ", err)
	}
	fmt.Printf("File %s created \n", filename)

}

func (insCom installCommands) unixInstall(stage string) {

	insCom.fileCreator("kubectlOIDCInstall.sh", insCom.KubectlOIDCInstallCom)
	insCom.fileCreator("kubectlKrewInstall.sh", insCom.kubectlKrewInstallCom)
	insCom.fileCreator("kubectlInstall.sh", insCom.kubectlInstallCom)

	OIDCFullPath := insCom.directory + "/kubectlOIDCInstall.sh"
	KrewFullPath := insCom.directory + "/kubectlKrewInstall.sh"
	kubectlFullPath := insCom.directory + "/kubectlInstall.sh"

	switch stage {
	case "oidcCheck":
		if err := exec.Command("/bin/sh", "-c", OIDCFullPath).Run(); err != nil {
			log.Println(err)
			fmt.Println(insCom.errorInfo.KubectlOIDCInstallError, err)
		}
	case "krewCheck":
		if err := exec.Command("/bin/sh", "-c", KrewFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", OIDCFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "kubectlCheck":
		if err := exec.Command("/bin/sh", "-c", kubectlFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", KrewFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command("/bin/sh", "-c", OIDCFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	}
	fmt.Println("Successfully installed")
}

func (insCom installCommands) winInstall(stage string) {

	insCom.fileCreator("kubectlOIDCInstall.bat", insCom.KubectlOIDCInstallCom)
	insCom.fileCreator("kubectlKrewInstall.bat", insCom.kubectlKrewInstallCom)
	insCom.fileCreator("kubectlInstall.bat", insCom.kubectlInstallCom)

	OIDCFullPath := insCom.directory + "/kubectlOIDCInstall.bat"
	KrewFullPath := insCom.directory + "/kubectlKrewInstall.bat"
	kubectlFullPath := insCom.directory + "/kubectlInstall.bat"

	switch stage {
	case "oidcCheck":
		if err := exec.Command(OIDCFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "krewCheck":
		if err := exec.Command(KrewFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command(OIDCFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	case "kubectlCheck":
		if err := exec.Command(kubectlFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl went error", err)
		}
		if err := exec.Command(KrewFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl Krew went error", err)
		}
		if err := exec.Command(OIDCFullPath).Run(); err != nil {
			fmt.Println("Installing kubectl OIDC went error", err)
		}
	}
	fmt.Println("Successfully installed")
}

func Installer() {

	var Directory string

	fmt.Print("Print in your Directory the bats settled: \n")
	fmt.Scanln(&Directory)

	errorsSets := installErrors{
		kubectlInstallError:     "Installing kubectl went error",
		kubectlKrewInstallError: "Installing kubectl Krew went error",
		KubectlOIDCInstallError: "Installing kubectl OIDC went error",
	}

	kctlAmd64LinuxInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\"\n"
	kctlArm64LinuxInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl\"\n"
	kctlWindowsInstallComm := "curl.exe -LO \"https://dl.k8s.io/release/v1.28.0/bin/windows/amd64/kubectl.exe\"\n"
	kctlAmd64MacOSInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl\"\n"
	kctlArm64MacOSInstallComm := "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/arm64/kubectl\"\n"

	krewBashInstallComm := `
set -x; cd "$(mktemp -d)" &&
  OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  KREW="krew-${OS}_${ARCH}" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" &&
  tar zxvf "${KREW}.tar.gz" &&
  ./"${KREW}" install krew
  echo 'export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"' >> ~/.bashrc
  source ~/.bashrc

`

	krewZshInstallComm := `
set -x; cd "$(mktemp -d)" &&
  OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  KREW="krew-${OS}_${ARCH}" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" &&
  tar zxvf "${KREW}.tar.gz" &&
  ./"${KREW}" install krew
  echo 'export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"' >> /.zshrc
  source ~/.zshrc

`

	krewWindowsInstallComm := `
wget https://github.com/kubernetes-sigs/krew/releases/download/v0.4.4/krew.exe -o ./krew.exe
.\krew install krew
setx PATH "%PATH%;%USERPROFILE%\.krew\bin"

`

	kctlOidcLoginInstallComm := "kubectl krew install oidc-login\n"

	linuxAmd64 := installCommands{
		kubectlInstallCom:     kctlAmd64LinuxInstallComm,
		kubectlKrewInstallCom: krewBashInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
		directory:             Directory,
		errorInfo:             errorsSets,
	}

	linuxArm64 := installCommands{
		kubectlInstallCom:     kctlArm64LinuxInstallComm,
		kubectlKrewInstallCom: krewBashInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
		directory:             Directory,
		errorInfo:             errorsSets,
	}

	windows := installCommands{
		kubectlInstallCom:     kctlWindowsInstallComm,
		kubectlKrewInstallCom: krewWindowsInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
		directory:             Directory,
		errorInfo:             errorsSets,
	}

	macOSIntel := installCommands{
		kubectlInstallCom:     kctlAmd64MacOSInstallComm,
		kubectlKrewInstallCom: krewZshInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
		directory:             Directory,
		errorInfo:             errorsSets,
	}

	macOSAppleSilicon := installCommands{
		kubectlInstallCom:     kctlArm64MacOSInstallComm,
		kubectlKrewInstallCom: krewZshInstallComm,
		KubectlOIDCInstallCom: kctlOidcLoginInstallComm,
		directory:             Directory,
		errorInfo:             errorsSets,
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
