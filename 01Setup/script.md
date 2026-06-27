# Installing Go (Windows & Linux)

This guide walks you through installing Go from scratch, verifying the installation, and preparing your development environment.

# Prerequisites

Before installing Go, make sure you have:

* Administrator (Windows) or sudo (Linux) privileges
* A stable internet connection
* At least 500 MB of free disk space

# Install Go on Windows

## Step 1: Download Go

1. Open your web browser.
2. Visit the official Go downloads page.

> [https://go.dev/dl/](https://go.dev/dl/)

3. Download the latest **Windows Installer (.msi)**.

Example:

```Shell
go1.xx.x.windows-amd64.msi
```

Choose the correct version:

| System                    | Download          |
| ------------------------- | ----------------- |
| 64-bit Windows (Most PCs) | windows-amd64.msi |
| ARM Windows               | windows-arm64.msi |

---

## Step 2: Run the Installer

Locate the downloaded file.

Example:

```
Downloads/
    go1.xx.x.windows-amd64.msi
```

Double-click it.

You will see the Go Setup Wizard.

---

## Step 3: Follow the Installation Wizard

Click through the installer:

```
Welcome
    ↓
Next
    ↓
Accept License
    ↓
Next
    ↓
Choose Installation Folder
```

Default installation location:

```
C:\Program Files\Go
```

It is recommended to keep the default location.

Click:

```
Install
```

Wait for installation to complete.

Finally click:

```
Finish
```

---

## Step 4: Verify PATH Configuration

The installer automatically adds Go to your system `PATH`.

To verify:

Open **Command Prompt**

Press:

```
Windows + R
```

Type:

```
cmd
```

Run:

```bash
go version
```

Example output:

```bash
go version go1.25.0 windows/amd64
```

If you see the version number, Go is installed correctly.

---

## Step 5: Check Go Environment

Run:

```bash
go env
```

Example:

```bash
GOARCH="amd64"
GOOS="windows"
GOROOT="C:\\Program Files\\Go"
GOPATH="C:\\Users\\YourName\\go"
```

This confirms Go is configured properly.

---

## Step 6: Verify the Go Binary

Run:

```bash
where go
```

Expected output:

```text
C:\Program Files\Go\bin\go.exe
```

---

## Step 7: Create Your Workspace (Optional)

Modern Go projects can live anywhere because Go Modules are enabled by default.

Example:

```
Documents/
    GoProjects/
        hello-go/
```

---

# Install Go on Linux

There are two common installation methods:

* Method 1 (Recommended): Install the official Go binary.
* Method 2: Use your Linux distribution's package manager (often provides older versions).

---

# Method 1 (Recommended): Official Go Binary

This method ensures you get the latest stable release.

---

## Step 1: Download Go

Visit:

```
https://go.dev/dl/
```

Or download directly using `wget`:

```bash
wget https://go.dev/dl/go1.xx.x.linux-amd64.tar.gz
```

For ARM64 systems:

```bash
wget https://go.dev/dl/go1.xx.x.linux-arm64.tar.gz
```

---

## Step 2: Remove Old Installation (If Any)

```bash
sudo rm -rf /usr/local/go
```

This prevents conflicts with previous versions.

---

## Step 3: Extract Go

Install Go into `/usr/local`:

```bash
sudo tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
```

This creates:

```
/usr/local/go
```

---

## Step 4: Add Go to PATH

Open your shell configuration file.

For Bash:

```bash
nano ~/.bashrc
```

For Zsh:

```bash
nano ~/.zshrc
```

Add:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Save the file.

Reload it:

```bash
source ~/.bashrc
```

or

```bash
source ~/.zshrc
```

---

## Step 5: Verify Installation

Run:

```bash
go version
```

Example:

```bash
go version go1.25.0 linux/amd64
```

---

## Step 6: Check Environment

```bash
go env
```

Example:

```bash
GOOS="linux"
GOARCH="amd64"
GOROOT="/usr/local/go"
GOPATH="/home/user/go"
```

---

## Step 7: Locate the Go Binary

```bash
which go
```

Example:

```bash
/usr/local/go/bin/go
```

---

# Method 2: Install Using Package Manager

This method is quick but may install an older version.

### Ubuntu / Debian

```bash
sudo apt update
sudo apt install golang-go
```

Verify:

```bash
go version
```

---

### Fedora

```bash
sudo dnf install golang
```

---

### Arch Linux

```bash
sudo pacman -S go
```

---

# Verify Everything Works

Run the following commands on either Windows or Linux:

```bash
go version
```

```bash
go env
```

```bash
go help
```

If all commands execute successfully, your Go installation is complete.

---

# Install Visual Studio Code (Recommended)

Although you can use any text editor, **Visual Studio Code** is one of the best editors for Go development.

1. Download VS Code from:

   [https://code.visualstudio.com/](https://code.visualstudio.com/)
2. Install it using the default options.
3. Launch VS Code.

---

# Install the Go Extension in VS Code

1. Open VS Code.
2. Press:

```
Ctrl + Shift + X
```

3. Search for:

```
Go
```

4. Install the official extension published by the **Go Team at Google**.

The extension provides:

* Syntax highlighting
* Code completion (IntelliSense)
* Automatic formatting
* Error detection
* Debugging support
* Test integration
* Refactoring tools
* Go Modules support

---

# Verify Go in VS Code

Open the integrated terminal:

```
Ctrl + `
```

Run:

```bash
go version
```

If the version is displayed, VS Code can access your Go installation.
