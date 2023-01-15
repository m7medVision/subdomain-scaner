# Subdomain Scanner

This repository contains a tool written in Go that can scan a given domain and find its subdomains. It uses a wordlist to perform this process.

## Installation

To use this tool, you will need to have Go installed on your machine. If you don't have it, you can download it from the official website https://golang.org/dl/

Once you have Go installed, clone this repository to your local machine using the following command:

```sh
git clone https://github.com/majhcc/subdomain-scaner.git
```
Navigate to the repository's directory and run the following command to build the tool:

```sh
go build
````

## Usage

The tool can be run from the command line using the following syntax:
```sh
./main
Enter DOMAIN: majhcc.com
Starting...
```

The tool will then begin scanning the domain and will output any discovered subdomains to the console.

## Note
This is a proof of concept tool, use it responsibly and make sure you have permission to scan the domain.

## Contribution
If you want to contribute to this project, feel free to open a pull request or issue.

## License
This tool is under MIT License.



