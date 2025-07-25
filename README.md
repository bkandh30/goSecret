# goSecret

A secure command-line tool for managing encrypted secrets and sensitive data locally. It is written in Go using AES-GCM encryption for maximum security.

## Installation

### Build from Source

```bash
git clone https://github.com/bkandh30/goSecret.git
cd goSecret
go build -o goSecret
```

### Install Globally

```bash
go install
```

## Usage

### Basic Commands

#### Store a Secret

```bash
./goSecret set mykey "my secret value"
```

#### Retrieve a Secret

```bash
./goSecret get mykey
```

#### List Available Commands

```bash
./goSecret help
```

### Example Workflow

```bash
# Store database credentials
./goSecret set db_password "super_secure_password123"

# Store API keys
./goSecret set api_key "sk-1234567890abcdef"

# Retrieve when needed
./goSecret get db_password
# Output: db_password = super_secure_password123
```

## Development

### Building

```bash
# Build for current platform
go build -o goSecret
```

## Configuration

goSecret stores encrypted secrets in a `.secrets` file in your project directory. This file is automatically created when you store your first secret.

**Important**: Keep your `.secrets` file secure and consider backing it up. If you lose this file, your encrypted secrets cannot be recovered.
