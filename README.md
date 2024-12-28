# Human-Friendly File Search Tool for Linux

## Overview
The Human-Friendly File Search Tool is a Linux-based application that lets users search for files using natural language queries such as:

```bash
"Find workflow files under 100KB modified last week"
```

This project leverages Go and text-processing libraries to make file searches intuitive and user-friendly, eliminating the need for complex command-line syntax.

---

## Features

- **Natural Language Queries**: Perform file searches in plain language.
- **Powerful Filters**: Search by file type, size, modification date, and more.
- **Text Processing with Go**: Uses advanced text-processing libraries like `Prose` to parse queries.
- **Optimized for Linux**: Seamlessly integrates with Linux file systems.
- **Extensibility**: Built to easily add support for additional search parameters.

---

## Installation

### Prerequisites

- Go (version 1.19 or later)
- A Linux environment

### Clone the Repository

```bash
git clone https://github.com/HimanshuMishir/human-friendly-file-search.git
cd human-friendly-file-search
```

### Build the Project

```bash
go build -o file-search
```

### Run the Tool

```bash
./file-search "Find workflow files under 100KB modified last week"
```

---

## Usage

1. Build and run the application.
2. Pass your query as a command-line argument, e.g.:

```bash
./file-search "Find log files larger than 1MB created yesterday"
```

### Supported Query Examples:
- "Find all PDF files under 500KB"
- "Search for image files modified in the last 7 days"
- "Workflow files under 100KB modified last week"

---

## Planned Features

- **Advanced Filters**: Search by file owner, permissions, and more.
- **OpenAI Integration**: Use AI to understand even more complex queries.
- **Interactive CLI**: Include autocomplete and suggestions for query building.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch:

   ```bash
   git checkout -b feature-name
   ```

3. Commit your changes and push to your fork:

   ```bash
   git commit -m "Add new feature"
   git push origin feature-name
   ```

4. Submit a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgments

- **Prose**: For enabling text processing in Go.
- Open-source contributors and the Linux community for inspiration and tools.

---

Feel free to reach out if you have questions, suggestions, or ideas for improvement!
