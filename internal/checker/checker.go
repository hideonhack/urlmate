package checker

import (
	"net/http"
	"strings"
	"time"
)

type URLStatus struct {
	URL      string
	Status   int
	Error    error
	Response time.Duration
}

type Checker struct {
	statusCodes []int
	detailed    bool
	timeout     time.Duration
}

func New(statusCodes []int, detailed bool, timeoutSeconds int) *Checker {
	return &Checker{
		statusCodes: statusCodes,
		detailed:    detailed,
		timeout:     time.Duration(timeoutSeconds) * time.Second,
	}
}

func (c *Checker) Check(url string) URLStatus {
	url = normalizeURL(url)
	start := time.Now()
	
	client := &http.Client{
		Timeout: c.timeout,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return URLStatus{
			URL:      url,
			Status:   0,
			Error:    err,
			Response: time.Since(start),
		}
	}
	defer resp.Body.Close()
	
	return URLStatus{
		URL:      url,
		Status:   resp.StatusCode,
		Error:    nil,
		Response: time.Since(start),
	}
}

func (c *Checker) ShouldShow(status int) bool {
	if len(c.statusCodes) == 0 {
		return true
	}
	
	for _, code := range c.statusCodes {
		if status == code {
			return true
		}
	}
	return false
}

func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return strings.TrimSpace(url)
} 