
# gopix

gopix is a command-line interface (CLI) tool written in Go that allows users to convert images into ASCII art. It offers features like scaling and detailed ASCII character usage to enhance the ASCII art experience.

## Features

- Convert images to ASCII art.
- Customize the scale of the ASCII art output.
- Option to use detailed ASCII characters for a more intricate representation.

## Installation

To install gopix, you need to have Go installed on your system. Follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/[your-username]/gopix.git
   ```

2. Navigate to the cloned directory:

   ```sh
   cd gopix
   ```

3. Build the project:

   ```sh
   ./build.sh
   ```

This will create an executable file in your directory.

## Usage

After installing, you can use gopix to convert images to ASCII art. Here's how to use it:

```sh
./bin/gopix ascii -i [path-to-your-image] -o [output-file-path] -s [scale] -d
```

- `-i` or `--input`: Path to the input image file.
- `-o` or `--output`: Path to the output file (optional). If not specified, the ASCII art will be printed on the console.
- `-s` or `--scale`: Scale of the ASCII image (optional).
- `-d` or `--detailed`: Use detailed ASCII characters for the output (optional).

## Contributing

Contributions to gopix are welcome! Here are some ways you can contribute:

- Reporting bugs
- Suggesting enhancements
- Submitting pull requests

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the [MIT License](LICENSE.md) - see the LICENSE.md file for details.
