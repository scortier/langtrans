# langtrans

## Overview

"langtrans" is a command-line interface (CLI) tool that allows users to translate text from one language to another using Google Translate. With "langtrans," you can easily specify the source language, target language, and the text you want to translate, and the tool will provide you with the translated text in your chosen target language.

## Features

- Translate text from one language to another.
- Specify the source language and target language as command-line flags.
- Simple and intuitive command-line interface.
- Utilizes the power of Google Translate for accurate translations.

## Getting Started

### Prerequisites

Before you can use "langtrans," make sure you have the following prerequisites installed on your system:

- Go programming language (Golang)

### Installation

To install "langtrans," you can use the following Go command:

```bash
go get -u github.com/scortier/langtrans
```

### Usage

To translate text using "langtrans," follow these steps:

1. Open your terminal.

2. Run the "langtrans" command with the following flags:

   - `-source`: Specify the source language (e.g., `-source en` for English).
   - `-target`: Specify the target language (e.g., `-target fr` for French).
   - `-text`: Provide the text you want to translate (e.g., `-text "Hello, World!"`).

   Example:

   ```bash
   langtrans -source en -target fr -text "Hello, World!"
   ```

3. "langtrans" will send your request to Google Translate, and you will receive the translated text in the target language.

### Example

```bash
$ langtrans -source en -target es -text "Good morning!"
Source: English
Target: Spanish
Translated Text: ¡Buenos días!
```

## Contributing

Contributions to "langtrans" are welcome! If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.

2. Create a new branch for your feature or bug fix.

3. Make your changes and ensure they are well-tested.

4. Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

"langtrans" is powered by Google Translate.

## Contact

If you have any questions, suggestions, or feedback, please feel free to contact us at [onlytoaditya@gmail.com](mailto:your-email@example.com).
