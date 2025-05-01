# Contributing to Podlog

Thank you for your interest in contributing to Podlog! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct.

## How to Contribute

1. Fork the repository
2. Create a new branch for your feature/fix
3. Make your changes
4. Add tests if applicable
5. Update documentation
6. Submit a Pull Request

## Development Setup

1. Install Go (version 1.21 or later)
2. Clone your fork:
   ```bash
   git clone https://github.com/Grim-R3ap3r/podlog.git
   cd podlog
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Build the project:
   ```bash
   go build -o podlog
   ```

## Pull Request Process

1. Ensure your code follows the project's style guidelines
2. Update the README.md with details of changes if applicable
3. Update the version number in the version package
4. The PR must pass all CI checks
5. The PR must be reviewed and approved by at least one maintainer

## Style Guidelines

- Use meaningful variable and function names
- Add comments for complex logic
- Follow Go's standard formatting
- Write tests for new features
- Update documentation for changes

## Testing

Run the tests:

```bash
go test ./...
```

## Documentation

- Update README.md for user-facing changes
- Add comments for code changes
- Update any relevant documentation

## Questions?

Feel free to open an issue for any questions or concerns.
