package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseStatusCodes(codes string) ([]int, error) {
	if codes == "" {
		return nil, nil
	}

	var result []int
	parts := strings.Split(codes, ",")
	
	for _, part := range parts {
		code, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid status code: %s", part)
		}
		if !validateStatusCode(code) {
			return nil, fmt.Errorf("status code must be between 100 and 599: %d", code)
		}
		result = append(result, code)
	}
	
	return result, nil
}

func ReadURLsFromStdin() []string {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

	return urls
}

func validateStatusCode(code int) bool {
	return code >= 100 && code < 600
} 