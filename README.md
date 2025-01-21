# Currency Converter CLI

A simple command-line interface application for currency conversions using the [Fixer API](https://apilayer.com/marketplace/fixer-api). The API hasn't so good  responsing for the invalid currency codes error. So, i tried to improve it by converting their datas to  csv file format. You can read my comments in the project for the solution. It is built with the beauty of Golang's simplicity. 
## Prerequisites

- Go 1.16 or higher
- Fixer API key (from [API Layer](https://apilayer.com/marketplace/fixer-api))  
## Installation  

```bash
# Clone the repository
git clone https://github.com/muhammedkucukaslan/currency-converter.git
# Navigate to the project directory
cd currency-converter-cli
# Build the application
go build -o main
```
## Usage

```bash
./main -from <base> -to <target> -amount <amount>
```
For explanation of flags:
```bash
./main -help or -h
```

![Screenshot](/public/image.png)
