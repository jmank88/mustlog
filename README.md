# MustLog [![GoDoc](https://godoc.org/github.com/jmank88/mustlog?status.svg)](https://godoc.org/github.com/jmank88/mustlog)
An error-less go-kit/log extension.

## MustLogger

MustLogger extends log.Logger with Must, an error-less version of Log.

```go
mustLogger := NewMustLogger(&errLogger{}, func(err error, keyvals ...interface{}) {
    fmt.Println("failed to log:", err)
})
mustLogger.Log("k", "v")
mustLogger.Log(poison, "poison")

mustLogger.Must("k", "v")
mustLogger.Must(poison, "poison")
```

Output:

```text
k v
k v
failed to log: poisoned
```

## MustContext

MustContext extends log.Context to implement MustLogger.

```go
mustContext := NewMustContext(&errLogger{}, func(err error, keyvals ...interface{}) {
    fmt.Println("failed to log:", err)
})

mustContext.Must("k", "v")
mustContext.Must(poison, "poison")

mustContext = mustContext.With("withK", "withV")
mustContext.Must("k", "v")
mustContext.Must(poison, "poison")

mustContext = mustContext.WithPrefix("prefixK", "prefixV")
mustContext.Must("k", "v")
mustContext.Must(poison, "poison")
```

Output:

```text
k v
failed to log: poisoned
withK withV k v
failed to log: poisoned
prefixK prefixV withK withV k v
failed to log: poisoned
```