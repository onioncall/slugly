package slugly

import (
    "errors"
    "fmt"
    "io"
    "net/http"
)

type SlugWriter interface {
    WriteChar(char string) error
    Flush() error
}

type ConsoleWriter struct {
    Writer io.Writer
}

type HTTPWriter struct {
    Writer http.ResponseWriter
    Flusher http.Flusher
}

func (cw *ConsoleWriter) WriteChar(char string) error {
    _, err := fmt.Fprint(cw.Writer, char)
    return err
}

func (cw *ConsoleWriter) Flush() error {
    return nil
}

func NewHTTPWriter(w http.ResponseWriter) (*HTTPWriter, error) {
    flusher, ok := w.(http.Flusher)
    if !ok {
        return nil, errors.New("response writer does not support flushing")
    }
    
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.Header().Set("Transfer-Encoding", "chunked")
    
    return &HTTPWriter{
        Writer:  w,
        Flusher: flusher,
    }, nil
}

func (hw *HTTPWriter) WriteChar(char string) error {
    _, err := fmt.Fprint(hw.Writer, char)
    if err != nil {
        return err
    }
    hw.Flusher.Flush()
    return nil
}

func (hw *HTTPWriter) Flush() error {
    hw.Flusher.Flush()
    return nil
}
