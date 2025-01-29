# Introduction to Cred CLI

**Cred** is a simple and secure CLI tool for managing passwords (`pass`) and environment variables (`env`) using Go.
It encrypts sensitive information using **GPG keys**, ensuring safe storage and retrieval.

## 🔹 Features
- **Password Management (`cred pass`)**: Store, retrieve, and manage passwords securely.
- **Environment Variable Management (`cred env`)**: Store encrypted `.env` files and retrieve values on demand.
- **GPG Encryption**: Uses **GNU Privacy Guard (GPG)** to encrypt credentials.
- **Simple CLI Interface**: Easy-to-use commands for adding, viewing, and modifying stored credentials.
- **Cross-Platform Support**: Works on Linux, macOS, and Windows.

## 📌 Example Usage
### 🔑 **Initialize with a GPG key**
```sh
cred init <gpg-key-id>
