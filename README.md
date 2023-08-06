# GPGDecryptor

A program to decrypt GPG encrypted files, encrypted using [GoCommon](https://github.com/hibare/GoCommon/blob/main/pkg/crypto/gpg/gpg.go).

Binaries are published on [GitHub releases](https://github.com/hibare/GPGDecryptor/releases). Additionally, pre-build `.deb` and `.rpm` packages are published on [Fury.io](https://gemfury.com/hibare)

# Installing

### Installing .deb

1. Download and import the Hibare repository key:

```sh
curl -fsSL https://apt.hibare.in/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/hibare-keyring.gpg
```

2. Add repository source file

```sh
echo "deb [signed-by=/usr/share/keyrings/hibare-keyring.gpg] https://apt.hibare.in/ * *" | sudo tee /etc/apt/sources.list.d/hibare.list
```

3. Update package list

```sh
sudo apt update
```

4. Install GPGDecryptor

```sh
sudo apt install gpgdecryptor
```

# Usage

Invoke `gpgdecryptor` cli command.

### Help

```sh
gpgdecryptor -h
```

```sh
Program to decrypt GPG files

Usage:
  GPGDecryptor [flags]

Flags:
      --gpg-file string   Input GPG file
  -h, --help              help for GPGDecryptor
      --priv-key string   Private key file path
      --pub-key string    Public key file path
  -v, --version           version for GPGDecryptor

```

### Decrypting file

```sh
gpgdecryptor --gpg-file <GPG file to decrypt> --priv-key <armored private key> --pub-key <armored public key>
```

Example:

```sh
gpgdecryptor --gpg-file db_exports.zip.gpg --priv-key priv.gpg --pub-key pub.key
Enter passphrase:
Decrypting file...
```
