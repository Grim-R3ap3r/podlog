# Podlog
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Website shields.io](https://img.shields.io/website-up-down-green-red/http/shields.io.svg?style=for-the-badge)](http://shields.io/)



A modern CLI tool for viewing Kubernetes pod logs with beautiful formatting and powerful search capabilities.

![image](https://github.com/user-attachments/assets/fc3140f6-c66e-4bda-97f0-52a052d7c9c7)


## Features

- 🎨 Beautiful log formatting with colors
- 🔍 Powerful search capabilities
- 📊 JSON pretty printing
- ⏱️ Precise timestamps
- 🎯 Multiple search terms support
- 🔄 Real-time log following
- 📝 Debug log focus

## Installation

### Quick Install (Linux/macOS)

```bash
curl -sSL https://raw.githubusercontent.com/Grim-R3ap3r/podlog/main/install.sh | bash
```

### Manual Installation

1. Install Go (version 1.21 or later)
2. Clone the repository:
   ```bash
   git clone https://github.com/Grim-R3ap3r/podlog.git
   cd podlog
   ```
3. Build and install:
   ```bash
   go build -o podlog
   sudo cp podlog /usr/local/bin/
   ```

## Usage

### Basic Usage

```bash
podlog -n <namespace> -p <pod-name>
```

### Options

- `-n, --namespace`: Kubernetes namespace (required)
- `-p, --pod`: Pod name (required)
- `-f, --follow`: Follow log output
- `-t, --tail`: Number of lines to show from the end (default: 100)
- `-a, --all`: Show all logs (not just debug)
- `-s, --search`: Search for specific text in logs (comma-separated for multiple terms)

### Examples

1. Show all logs:

   ```bash
   podlog -n stag-2 -p <pod name> -a
   ```

2. Search for specific text:

   ```bash
   podlog -n stag-2 -p <pod name> -a -s "search term"
   ```

3. Search for multiple terms:

   ```bash
   podlog -n stag-2 -p <pod name> -a -s "term1,term2"
   ```

4. Follow logs while searching:
   ```bash
   podlog -n stag-2 -p <pod name> -a -s "search term" -f
   ```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Created by [@Grim-R3ap3r](https://github.com/Grim-R3ap3r)
- Inspired by the need for better Kubernetes log viewing tools
- Built with ❤️ for developers

- [contributors-shield]: https://img.shields.io/github/contributors/Grim-R3ap3r/podlog.svg?style=for-the-badge
[contributors-url]:  https://github.com/Grim-R3ap3r/podlog/graphs/contributors
[forks-shield]: 	https://img.shields.io/github/forks/Grim-R3ap3r/podlog.svg?style=for-the-badge
[forks-url]: https://github.com/Grim-R3ap3r/podlog/
[stars-shield]: https://img.shields.io/github/stars/Grim-R3ap3r/podlog.svg?style=for-the-badge
[stars-url]:  https://github.com/Grim-R3ap3r/podlog/stargazers
[issues-shield]: https://img.shields.io/github/issues/Grim-R3ap3r/podlog.svg?style=for-the-badge
[issues-url]: https://github.com/Grim-R3ap3r/podlog/issues
[license-shield]: https://img.shields.io/github/license/tiwariadarsh/StreamEzy?style=for-the-badge
[license-url]: https://github.com/Grim-R3ap3r/podlog/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

