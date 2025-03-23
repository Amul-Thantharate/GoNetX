# 🌐 GoNetX: Your Network and System Utility CLI 🛠️

[![Go](https://img.shields.io/badge/Go-1.20%2B-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

GoNetX is a versatile command-line interface (CLI) tool built with Go, designed to provide quick access to essential network and system information. Whether you need to check your IP addresses, ping a host, or get details about your operating system, RAM, or disk space, GoNetX has you covered!

## ✨ Features

-   **IP Address Display:** Easily view all your network interface IP addresses. 📡
-   **Hostname Lookup:** Quickly find your system's hostname. 🖥️
-   **Ping:** Test network connectivity by pinging a specified host. 🏓
-   **OS Information:** Get detailed information about your operating system. 🐧🍎🪟
-   **RAM Information:** Check your system's RAM usage and capacity. 🧠
-   **Disk Information:** View details about your hard disk space. 💾

## 🚀 Getting Started

### Prerequisites

-   [Go](https://golang.org/dl/) (version 1.20 or higher)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/Amul-Thantharate/GoNetX.git 
    cd GoNetX
    ```

2.  **Build the application:**

    ```bash
    go build -o gonetx main.go
    ```

3.  **Move the binary (optional):**

    ```bash
    sudo mv gonetx /usr/local/bin/
    ```

    This step is optional but recommended to make the `gonetx` command available globally.

### Usage

Once installed, you can use the `gonetx` command followed by the desired subcommand.

```bash
    gonetx [command] [options]
    gonetx address
    gonetx hostname
    gonetx ping [host]
    gonetx os
    gonetx ram
    gonetx disk
```

## Contributing 🤝

We welcome contributions! Here''s how:

1. Fork the repo 🍴
2. Create your branch: `git checkout -b feature/amazing` 🌱
3. Commit changes: `git commit -m ''Add feature''` 💫
4. Push to branch: `git push origin feature/amazing` ⬆️
5. Open a Pull Request 📬

## Support 💪

- 📚 Wiki: [Documentation Link]
- 🐛 Issues: [GitHub Issues]
- 📧 Email: amulthantharate@gmail.com

## License 📄

This project is licensed under the Apache License 2.0.

## Acknowledgments 🙏

- OpenWeatherMap API for weather data 🌤️
- Go community for amazing tools 🛠️
- All our contributors 👥