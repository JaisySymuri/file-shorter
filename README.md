# File Shorter

File Shorter is a simple command-line utility that helps you organize and tidy up your files into appropriate categories based on predefined rules. This tool is ideal for managing your "Downloads" and "Academic" folders by moving files to their respective destinations.

## Features

- Organize files in your Downloads and Academic directories.
- Define categories and rules in a YAML configuration file.
- Automatically create destination directories if they don't exist.
- Validate configuration to avoid duplicate rules.

## Prerequisites

- Go 1.16 or later
- [Viper](https://github.com/spf13/viper) - configuration with fangs!

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/file-shorter.git
cd file-shorter
```

2. Install dependencies:

```sh
go get -u github.com/spf13/viper
```

3. Build the application:

```sh
go build -o file-shorter
```

## Usage

1. Create a `config.yml` file in the root directory with the following structure:

```yaml
categories:
  - name: "Documents"
    rules:
      - "doc"
      - "pdf"
    destination: "C:\\Users\\bests\\Documents\\"
  - name: "Images"
    rules:
      - "jpg"
      - "png"
    destination: "C:\\Users\\bests\\Pictures\\"
  - name: "Videos"
    rules:
      - "mp4"
      - "avi"
    destination: "C:\\Users\\bests\\Videos\\"
```

2. Run the application:

```sh
./file-shorter
```

3. Follow the on-screen prompts to choose the directory you want to tidy up.

## Configuration

The configuration file (`config.yml`) defines the categories and their respective rules. Each category must have a unique name, a list of rules (file extensions or patterns), and a destination directory.

Example `config.yml`:

```yaml
categories:
  - name: "Documents"
    rules:
      - "doc"
      - "pdf"
    destination: "C:\\Users\\bests\\Documents\\"
  - name: "Images"
    rules:
      - "jpg"
      - "png"
    destination: "C:\\Users\\bests\\Pictures\\"
  - name: "Videos"
    rules:
      - "mp4"
      - "avi"
    destination: "C:\\Users\\bests\\Videos\\"
```

## Contributing

Contributions are welcome! Please open an issue or a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Viper](https://github.com/spf13/viper) for making configuration management easy.
- The Go programming language for providing a powerful and efficient toolset.

Happy organizing!