package trace_writer

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/deependujha/go-chrometrace/trace_event"
)

type TraceWriter struct {
	path       string
	fileHandle *os.File
	open       bool
}

func NewTraceWriter(path string) *TraceWriter {
	return &TraceWriter{path: path, fileHandle: nil, open: false}
}

func (tw *TraceWriter) Open() error {
	if tw.open {
		return nil
	}

	file, err := os.OpenFile(tw.path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString("[")
	if err != nil {
		file.Close()
		return err
	}

	tw.fileHandle = file
	tw.open = true
	return nil
}

func (tw *TraceWriter) Close() error {
	if !tw.open {
		return nil
	}

	if tw.fileHandle == nil {
		return errors.New("invalid file handle")
	}

	// Truncate the last comma and space from the file
	stat, err := tw.fileHandle.Stat()
	if err == nil && stat.Size() > 2 {
		tw.fileHandle.Truncate(stat.Size() - 2)
	}

	_, err = tw.fileHandle.WriteString("]")
	if err != nil {
		return err
	}

	tw.fileHandle.Close()
	tw.fileHandle = nil
	tw.open = false
	return nil
}

func (tw *TraceWriter) Write(events []trace_event.TraceEvent) error {
	if !tw.open || tw.fileHandle == nil {
		return errors.New("I/O operation on closed file")
	}

	data, err := json.Marshal(events)
	if err != nil {
		return err
	}

	_, err = tw.fileHandle.Write(data[1 : len(data)-1]) // Remove opening and closing brackets
	if err != nil {
		return err
	}

	_, err = tw.fileHandle.WriteString(", ")
	return err
}

func (tw *TraceWriter) IsOpen() bool {
	return tw.open
}
