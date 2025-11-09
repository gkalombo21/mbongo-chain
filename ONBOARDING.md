# 🚀 Mbongo Chain - Developer Onboarding Guide

Welcome to Mbongo Chain! This guide will help you get started with development and contribute to the project.

## Table of Contents

- [Development Environment Setup](#development-environment-setup)
- [Running Mbongo Chain Locally](#running-mbongo-chain-locally)
- [Testing the Code](#testing-the-code)
- [Contributing New Modules](#contributing-new-modules)
- [Project Philosophy](#project-philosophy)
- [Next Steps](#next-steps)

## Development Environment Setup

### 1. Install Go

Mbongo Chain requires **Go 1.21 or later**.

#### Windows

1. Download Go from [golang.org/dl](https://golang.org/dl/)
2. Run the installer and follow the setup wizard
3. Verify installation:
   ```bash
   go version
   ```
   You should see: `go version go1.21.x` or higher

#### macOS

```bash
# Using Homebrew
brew install go

# Verify installation
go version
```

#### Linux

```bash
# Download and install
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Verify installation
go version
```

### 2. Install Git

Git is required for version control and cloning the repository.

#### Windows

1. Download from [git-scm.com](https://git-scm.com/download/win)
2. Run the installer with default settings
3. Verify installation:
   ```bash
   git --version
   ```

#### macOS

```bash
# Using Homebrew
brew install git

# Verify installation
git --version
```

#### Linux

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install git

# Verify installation
git --version
```

### 3. Install Cursor (Recommended)

Cursor is an AI-powered code editor that enhances development with intelligent assistance.

1. Download from [cursor.sh](https://cursor.sh/)
2. Install and open Cursor
3. Open the Mbongo Chain project:
   ```bash
   cursor .
   ```
4. The project includes `.cursor/rules` that guide AI assistance

**Why Cursor?**
- AI-powered code generation and suggestions
- Understands project-specific conventions
- Automates repetitive tasks
- Helps with code review and refactoring

### 4. Install Docker (Optional)

Docker is optional but useful for:
- Testing API endpoints
- Running database containers
- Containerized development environments

#### Windows/macOS

1. Download [Docker Desktop](https://www.docker.com/products/docker-desktop)
2. Install and start Docker Desktop
3. Verify installation:
   ```bash
   docker --version
   docker-compose --version
   ```

#### Linux

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install docker.io docker-compose

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Verify installation
docker --version
docker-compose --version
```

### 5. Clone the Repository

```bash
# Clone the repository
git clone https://github.com/gkalombo21/mbongo-chain.git

# Navigate to the project directory
cd mbongo-chain

# Install dependencies
go mod download

# Verify everything is set up correctly
go mod verify
```

### 6. Verify Your Setup

Run these commands to ensure everything is configured correctly:

```bash
# Check Go version (should be 1.21+)
go version

# Check Git version
git --version

# Verify Go module
go mod verify

# Check if the project builds
go build ./cmd/mbongo-chain

# Run tests to ensure everything works
go test ./...
```

## Running Mbongo Chain Locally

### Quick Start

The easiest way to run Mbongo Chain locally is using `go run`:

```bash
go run ./cmd/mbongo-chain
```

This command will:
- Compile the application
- Run demonstration scenarios
- Display output showing:
  - Bank account creation and transactions
  - User management operations
  - Transaction history

### Expected Output

When you run the application, you should see output similar to:

```
MBongo node started
Created account #1 for Alice at 2025-11-08T21:45:00-05:00 with initial balance: 100.00
After deposit, balance: 150.00
After withdraw, balance: 120.00
Transactions:
- #1 create 100.00 at 2025-11-08T21:45:00-05:00
- #2 deposit 50.00 at 2025-11-08T21:45:00-05:00
- #3 withdraw 30.00 at 2025-11-08T21:45:00-05:00
User #1: Bob <bob@example.com> (created 2025-11-08T21:45:00-05:00)
```

### Building the Executable

To build a standalone executable:

```bash
# Windows
go build -o mbongod.exe ./cmd/mbongo-chain

# Linux/macOS
go build -o mbongod ./cmd/mbongo-chain

# Run the executable
./mbongod.exe  # Windows
./mbongod      # Linux/macOS
```

### Development Mode

For development, you can modify the code and rerun:

```bash
# Run and watch for changes (requires air tool)
go install github.com/cosmtrek/air@latest
air

# Or simply rerun
go run ./cmd/mbongo-chain
```

## Testing the Code

### Running All Tests

To run all tests in the project:

```bash
go test ./...
```

### Running Tests with Verbose Output

To see detailed test output:

```bash
go test -v ./...
```

### Running Tests for a Specific Package

To test a specific module:

```bash
# Test bank module
go test ./internal/bank/... -v

# Test user module
go test ./internal/user/... -v

# Test a specific test file
go test ./internal/bank/tests -v
```

### Running Tests with Coverage

To generate test coverage reports:

```bash
# Run tests with coverage
go test -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./...

# View coverage in browser (HTML)
go tool cover -html=coverage.out
```

### Writing Tests

When writing new code, always include tests. Follow these guidelines:

1. **Test Naming**: Use `TestFunctionName` or `TestStruct_MethodName`
2. **Test Files**: Place tests in the same package with `_test.go` suffix
3. **Table-Driven Tests**: Use table-driven tests for multiple test cases
4. **Coverage**: Aim for high test coverage, especially for keeper packages

**Example Test:**

```go
func TestAccountCreation(t *testing.T) {
    tests := []struct {
        name          string
        accountID     string
        initialBalance float64
        wantErr       bool
    }{
        {
            name:          "valid account",
            accountID:     "acc1",
            initialBalance: 100.0,
            wantErr:       false,
        },
        {
            name:          "zero balance",
            accountID:     "acc2",
            initialBalance: 0.0,
            wantErr:       false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
            account, err := createAccount(tt.accountID, tt.initialBalance)
            if (err != nil) != tt.wantErr {
                t.Errorf("createAccount() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && account == nil {
                t.Error("createAccount() returned nil account")
            }
        })
    }
}
```

## Contributing New Modules

### Module Structure

When creating a new module, follow this structure:

```
internal/
  └── your-module/
      ├── types.go          # Type definitions
      ├── keeper/
      │   └── keeper.go     # State management
      ├── handler.go        # Business logic
      ├── doc.go            # Package documentation
      └── tests/            # Test files
          └── your-module_test.go
```

### Step-by-Step: Creating a New Module

#### 1. Create Module Directory

```bash
mkdir -p internal/your-module/keeper
mkdir -p internal/your-module/tests
```

#### 2. Create Types File

Create `internal/your-module/types.go`:

```go
package yourmodule

// YourType represents a type in your module
type YourType struct {
    ID   string
    Data string
}

// TODO: Integrate with Cosmos SDK types
```

#### 3. Create Keeper

Create `internal/your-module/keeper/keeper.go`:

```go
package keeper

import (
    "fmt"
    "sync"
)

// Keeper manages the state of your module
type Keeper struct {
    mu    sync.RWMutex
    store map[string]interface{}
}

// NewKeeper creates a new keeper instance
func NewKeeper() *Keeper {
    fmt.Println("Creating new keeper for your-module")
    return &Keeper{
        store: make(map[string]interface{}),
    }
}

// Get retrieves a value from the store
func (k *Keeper) Get(key string) (interface{}, bool) {
    k.mu.RLock()
    defer k.mu.RUnlock()
    value, exists := k.store[key]
    return value, exists
}

// Set stores a value in the store
func (k *Keeper) Set(key string, value interface{}) {
    k.mu.Lock()
    defer k.mu.Unlock()
    fmt.Printf("Setting %s in keeper\n", key)
    k.store[key] = value
}
```

#### 4. Create Handler

Create `internal/your-module/handler.go`:

```go
package yourmodule

import (
    "fmt"
    "github.com/gkalombo21/mbongo-chain/internal/your-module/keeper"
)

// Handler provides business logic for your module
type Handler struct {
    keeper *keeper.Keeper
}

// NewHandler creates a new handler instance
func NewHandler(k *keeper.Keeper) *Handler {
    fmt.Println("Creating new handler for your-module")
    return &Handler{
        keeper: k,
    }
}

// ProcessData processes data using the keeper
func (h *Handler) ProcessData(data string) error {
    fmt.Printf("Processing data: %s\n", data)
    h.keeper.Set("data", data)
    return nil
}
```

#### 5. Write Tests

Create `internal/your-module/tests/your-module_test.go`:

```go
package tests

import (
    "testing"
    "github.com/gkalombo21/mbongo-chain/internal/your-module"
    "github.com/gkalombo21/mbongo-chain/internal/your-module/keeper"
)

func TestHandler_ProcessData(t *testing.T) {
    k := keeper.NewKeeper()
    h := yourmodule.NewHandler(k)
    
    err := h.ProcessData("test data")
    if err != nil {
        t.Errorf("ProcessData() error = %v", err)
    }
    
    value, exists := k.Get("data")
    if !exists {
        t.Error("Data was not stored in keeper")
    }
    
    if value != "test data" {
        t.Errorf("Expected 'test data', got %v", value)
    }
}
```

#### 6. Integration

Wire your module into the main application in `cmd/mbongo-chain/main.go`:

```go
import (
    "github.com/gkalombo21/mbongo-chain/internal/your-module"
    "github.com/gkalombo21/mbongo-chain/internal/your-module/keeper"
)

func main() {
    // Create keeper
    k := keeper.NewKeeper()
    
    // Create handler
    h := yourmodule.NewHandler(k)
    
    // Use handler
    h.ProcessData("example data")
}
```

### Module Guidelines

- **Keep functions small** (<50 lines)
- **Use dependency injection** (pass keepers to handlers)
- **Log key steps** using `fmt.Println` for debugging
- **Include TODOs** for Cosmos SDK integration points
- **Write tests** for all keeper packages
- **Follow naming conventions**:
  - Functions: `camelCase`
  - Structs: `PascalCase`
  - Constants: `SCREAMING_SNAKE_CASE`

## Project Philosophy

### Blockchain + AI = Useful Decentralization

Mbongo Chain is built on a fundamental philosophy: **combining blockchain technology with artificial intelligence to create useful, decentralized computation**.

#### Core Principles

1. **Proof of Useful Work (PoUW)**
   - Traditional blockchain mining wastes computational resources
   - Mbongo Chain uses GPU power for meaningful AI workloads
   - Miners contribute to useful computation, not arbitrary puzzles

2. **Hybrid Consensus (PoS + PoUW)**
   - Proof of Stake (PoS) for energy-efficient consensus
   - Proof of Useful Work (PoUW) for decentralized AI computation
   - Best of both worlds: efficiency + utility

3. **AI-Powered Decentralization**
   - GPU miners perform useful AI tasks
   - Distributed machine learning and inference
   - Democratized access to AI computation

4. **Modular Architecture**
   - Clean separation of concerns
   - Independent, testable modules
   - Easy to extend and maintain

#### Why This Matters

- **Environmental Impact**: Useful work instead of wasted computation
- **Accessibility**: Decentralized AI for everyone
- **Efficiency**: Hybrid consensus combines efficiency and utility
- **Innovation**: New paradigm for blockchain and AI integration

#### Vision

Mbongo Chain aims to be a platform where:
- Blockchain security meets AI intelligence
- Miners contribute to meaningful computation
- Developers build decentralized AI applications
- Users benefit from useful, decentralized services

## Next Steps

Now that you're set up, here's what you can do next:

1. **Explore the Codebase**
   - Read through existing modules (`internal/bank`, `internal/user`)
   - Understand the project structure
   - Review the `.cursor/rules` file

2. **Run the Application**
   - Execute `go run ./cmd/mbongo-chain`
   - Experiment with the code
   - Make small changes and see the results

3. **Write Tests**
   - Add tests to existing modules
   - Improve test coverage
   - Practice table-driven test patterns

4. **Contribute**
   - Check out [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed guidelines
   - Pick an issue from the issue tracker
   - Create a new module or improve existing ones

5. **Join the Community**
   - Ask questions in discussions
   - Share your ideas
   - Help other contributors

## Resources

- **Project Repository**: [github.com/gkalombo21/mbongo-chain](https://github.com/gkalombo21/mbongo-chain)
- **Contributing Guide**: [CONTRIBUTING.md](./CONTRIBUTING.md)
- **Go Documentation**: [golang.org/doc](https://golang.org/doc/)
- **Cosmos SDK**: [docs.cosmos.network](https://docs.cosmos.network/)
- **Cursor IDE**: [cursor.sh](https://cursor.sh/)

---

**💡 Tip: Use the Cursor agent to automate setup and coding tasks.**

