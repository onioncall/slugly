package slugly

import (
    "io"
    "net/http"
)

// Helper function to create a SlugConfig with console writer
func NewConsoleSlugConfig(speed PrintSpeed, writer io.Writer) *SlugConfig {
    return &SlugConfig{
        Speed:  speed,
        Writer: &ConsoleWriter{Writer: writer},
    }
}

// Helper function to create a SlugConfig with HTTP writer
func NewHTTPSlugConfig(speed PrintSpeed, w http.ResponseWriter) (*SlugConfig, error) {
    httpWriter, err := NewHTTPWriter(w)
    if err != nil {
        return nil, err
    }
    
    return &SlugConfig{
        Speed:  speed,
        Writer: httpWriter,
    }, nil
}
