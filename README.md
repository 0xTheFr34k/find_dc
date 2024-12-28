# find_dc

`find_dc` is a Go-based tool designed to process the output of the `nxc` command, identify domain controllers, and generate the necessary `/etc/hosts` entries. It automatically fetches the domain controller for each domain.

## Features

- Parses `nxc` output to extract relevant IP, hostname, and domain information.
- Uses `nslookup` to find the domain controller (DC) for each domain.
- Generates and prints `/etc/hosts` entries, formatted in a tabular structure.
- Supports both long (`--help`) and short (`-h`) flags for help information.
- Output is suitable for direct inclusion in your `/etc/hosts` file.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Help Flag](#help-flag)
- [Examples](#examples)
- [Flags](#flags)
- [Contributing](#contributing)

## Installation

To use `find_dc`, you need to have Go installed on your machine. You can install Go from [here](https://golang.org/doc/install).

1. Clone the repository:
   ```bash
   git clone https://github.com/0xTheFr34k/find_dc.git
   cd find_dc
   ```

2. Build the project:
   ```bash
   go build -o find_dc .
   ```
3. Installation using go cli
    ```bash
    go install github.com/0xTheFr34k/find_dc@latest
    ```

3. Run the binary directly or use it as part of a pipeline with `nxc` output.

## Usage

### Basic Usage

After building the tool, you can use `find_dc` in the following way:

1. **Pipe `nxc` output** to `find_dc`:
   ```bash
   nxc smb 192.168.56.10-12 | find_dc
   ```

2. The tool will process the output, lookup the domain controller, and print the `/etc/hosts` entries.

### Help Flag

You can display help information using the `--help` or `-h` flag:

```bash
find_dc --help
```

This will show the following output:

```
Usage: find_dc [OPTIONS]

Description:
A Go tool to parse 'nxc' output, find domain controllers, and generate /etc/hosts entries.

Options:
-h, --help    Display this help message and exit

Examples:
nxc smb 192.168.56.10-12 | find_dc
nxc smb 192.168.56.10-12 | find_dc --help

Output:
The tool prints /etc/hosts entries in a table format.
```

### Input and Output

- The input should come from the `nxc` command, specifically from the `smb` module that outputs information like IP addresses, hostnames, and domains.
- The output is a list of `/etc/hosts` entries formatted as follows:

```
192.168.56.12 meereen.essos.local meereen essos.local
192.168.56.23 braavos.essos.local braavos
192.168.56.11 winterfell.north.sevenkingdoms.local winterfell north.sevenkingdoms.local
192.168.56.10 kingslanding.sevenkingdoms.local kingslanding sevenkingdoms.local
192.168.56.22 castelblack.north.sevenkingdoms.local castelblack
```

This output can be copied directly into your `/etc/hosts` file.

## Examples

### Example 1: Basic usage with IP range

```bash
nxc smb 192.168.56.10-12 | find_dc
```

**Output:**

```
Add the following lines to /etc/hosts:
192.168.56.12 meereen.essos.local meereen essos.local
192.168.56.23 braavos.essos.local braavos
```

### Example 2: Using help flag

```bash
nxc smb 192.168.56.10-12 | find_dc --help
```

**Output:**

```
Usage: find_dc [OPTIONS]

Description:
A Go tool to parse 'nxc' output, find domain controllers, and generate /etc/hosts entries.

Options:
-h, --help    Display this help message and exit

Examples:
nxc smb 192.168.56.10-12 | find_dc
nxc smb 192.168.56.10-12 | find_dc --help
```

### Example 3: Multiple IP ranges

```bash
nxc smb 192.168.56.10-12 192.168.56.22-23 | find_dc
```

**Output:**

```
Add the following lines to /etc/hosts:
192.168.56.12 meereen.essos.local meereen essos.local
192.168.56.23 braavos.essos.local braavos
192.168.56.11 winterfell.north.sevenkingdoms.local winterfell north.sevenkingdoms.local
192.168.56.10 kingslanding.sevenkingdoms.local kingslanding sevenkingdoms.local
192.168.56.22 castelblack.north.sevenkingdoms.local castelblack
```

## Flags

- `-h, --help`: Display help information.
  
## Contributing

If you'd like to contribute to this project, feel free to open an issue or submit a pull request. Contributions are welcome!

### Steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

---

