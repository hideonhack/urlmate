package urlmate

import (
	"flag"
	"log"
	"os"

	"github.com/hideonhack/urlmate/internal/checker"
	"github.com/hideonhack/urlmate/internal/output"
	"github.com/hideonhack/urlmate/internal/parser"
)

func Execute() {
	flag.Usage = func() {
		writer, err := output.NewWriter("")
		if err != nil {
			log.Fatal(err)
		}
		if writer != nil {
			printUsage(writer)
		}
	}

	var (
		outputFile  = flag.String("o", "", "Output file path")
		statusCodes = flag.String("status-code", "", "Filter results by HTTP status codes (comma-separated, e.g., 200,404)")
		detailed    = flag.Bool("detailed", false, "Show detailed output including status symbols and response time")
		timeout     = flag.Int("timeout", 10, "Timeout in seconds for HTTP requests")
	)
	
	flag.Parse()

	codes, err := parser.ParseStatusCodes(*statusCodes)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	writer, err := output.NewWriter(*outputFile)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer writer.Close()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		urls := parser.ReadURLsFromStdin()
		urlChecker := checker.New(codes, *detailed, *timeout)
		
		for _, url := range urls {
			result := urlChecker.Check(url)
			writer.WriteResult(result, *detailed)
		}
	} else {
		printUsage(writer)
	}
}

func printUsage(writer *output.Writer) {
	writer.Write("URLMate - URL Availability Checker\n")
	writer.Write("\nOptions:\n")
	writer.Write("  -o            Output file path\n")
	writer.Write("  -status-code  Filter results by HTTP status codes (comma-separated, e.g., 200,404)\n")
	writer.Write("  -detailed     Show detailed output including status symbols and response time\n")
	writer.Write("  -timeout      Set timeout in seconds (default: 10)\n")
	writer.Write("\nUsage:\n")
	writer.Write("  echo \"example.com\" | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]\n")
	writer.Write("  cat linklist.txt | urlmate [-o output.txt] [-status-code CODES] [-detailed] [-timeout SECONDS]\n")
	writer.Write("\nExamples:\n")
	writer.Write("  cat urls.txt | urlmate                                    # Basic output\n")
	writer.Write("  cat urls.txt | urlmate -detailed                         # Detailed output\n")
	writer.Write("  cat urls.txt | urlmate -status-code 200,404             # Filter by status codes\n")
	writer.Write("  cat urls.txt | urlmate -timeout 30                      # Set timeout to 30 seconds\n")
	writer.Write("  echo \"example.com\" | urlmate -detailed -o results.txt -timeout 30   # Full example\n")
} 