# URLMate

A powerful Go-based tool for checking URL availability and accessibility. URLMate helps you monitor and validate whether web resources are accessible and responding properly.

## Features

- ✨ Fast and efficient URL accessibility checking
- ⏱️ Precise response time measurement
- 🔍 HTTP status code validation
- ⚡ Timeout handling and error reporting
- 📦 Batch URL processing
- 🔒 Support for both HTTP and HTTPS protocols

## Configuration

Default timeout is set to 10 seconds. You can modify this in the code according to your needs.

## Technical Specifications

- Modern web protocols support (HTTP/1.1, HTTP/2)
- Comprehensive HTTP status code handling
- Configurable timeout settings
- Detailed error reporting
- Clean and intuitive output format

<a href="https://github.com/hideonhack/"><img align="right" src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/30c621a657fb4a0bf4234e1f20f7ce91333fd712.png" style="width: 200px; position:" alt="gopher" title="gopher" /></a>

## **How to download ?**
If you have Go installed:
```
go install github.com/hideonhack/urlmate@latest
```
## **Usage**
```
echo "example.com" | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]

cat linklist.txt | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]
```
## **Parameters**
```
URLMate - URL Availability Checker

Options:
  -o            Output file path
  -status-code  Filter results by HTTP status codes (comma-separated, e.g., 200,404)
  -detailed     Show detailed output including status symbols and response time
  -timeout      Set timeout in seconds (default: 10)

Usage:
  echo "example.com" | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]
  cat linklist.txt | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]

Examples:
  cat urls.txt | urlmate                                    # Basic output
  cat urls.txt | urlmate -detailed                         # Detailed output
  cat urls.txt | urlmate -status-code 200,404             # Filter by status codes
  cat urls.txt | urlmate -timeout 30                      # Set timeout to 30 seconds
  echo "example.com" | urlmate -detailed -o results.txt -timeout 30   # Full example
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

<div style="text-align: center; justify-content: center;">
<hr/>
<a style="margin-left: 15px;" href="https://www.linkedin.com/in/boradogru/" target="_blank"><img src="https://img.shields.io/badge/Linkedin-blue.svg"/ target="_blank"></a>
<a style="margin-left: 5px;" href="https://www.instagram.com/hideonhack/" target="_blank"><img src="https://img.shields.io/badge/Instagram-pink.svg"/></a>
</div>