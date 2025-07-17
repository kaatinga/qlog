# qlog

a very basic implementation of a logger

---

## Package Description

`qlog` provides a minimalistic logging interface for Go applications. The `Logger` interface (defined in `interface.go`) exposes `Debugf` and `Errorf` methods, supporting printf-style formatting with variadic arguments (`a ...any`). The package includes a default implementation that writes log messages to standard output or standard error, depending on the log level. Log messages are formatted with the log level and the formatted message.

The package can be used as a default logger for helper packages.

### Usage Example

```go
logger := qlog.New()
logger.Debugf("starting app on port %d", 8080)
logger.Errorf("failed to connect: %v", err)
```
