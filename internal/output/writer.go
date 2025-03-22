package output

import (
	"fmt"
	"os"

	"github.com/hideonhack/urlmate/internal/checker"
)

type Writer struct {
	file *os.File
}

func NewWriter(outputFile string) (*Writer, error) {
	var file *os.File
	var err error
	
	if outputFile != "" {
		file, err = os.Create(outputFile)
		if err != nil {
			return nil, fmt.Errorf("failed to create output file: %v", err)
		}
	}
	
	return &Writer{file: file}, nil
}

func (w *Writer) Write(message string) {
	fmt.Print(message)
	if w.file != nil {
		fmt.Fprint(w.file, message)
	}
}

func (w *Writer) WriteResult(result checker.URLStatus, detailed bool) {
	if result.Error != nil {
		if detailed {
			w.Write(fmt.Sprintf("❌ %s - Error: %v\n", result.URL, result.Error))
		} else {
			w.Write(fmt.Sprintf("%s Error\n", result.URL))
		}
		return
	}

	if detailed {
		statusSymbol := "✅"
		if result.Status >= 400 {
			statusSymbol = "❌"
		}
		w.Write(fmt.Sprintf("%s %s (Status: %d, Response Time: %v)\n",
			statusSymbol,
			result.URL,
			result.Status,
			result.Response))
	} else {
		w.Write(fmt.Sprintf("%s %d\n",
			result.URL,
			result.Status))
	}
}

func (w *Writer) Close() {
	if w.file != nil {
		w.file.Close()
	}
} 