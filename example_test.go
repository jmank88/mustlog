package mustlog

import (
	"errors"
	"fmt"
)

func ExampleMustLogger() {
	mustLogger := NewMustLogger(&errLogger{}, func(err error, keyvals ...interface{}) {
		fmt.Println("failed to log:", err)
	})
	mustLogger.Log("k", "v")
	mustLogger.Log(poison, "poison")

	mustLogger.Must("k", "v")
	mustLogger.Must(poison, "poison")

	// Output:
	// k v
	// k v
	// failed to log: poisoned
}

func ExampleMustContext() {
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

	// Output:
	// k v
	// failed to log: poisoned
	// withK withV k v
	// failed to log: poisoned
	// prefixK prefixV withK withV k v
	// failed to log: poisoned
}

var poison interface{} = nil

type errLogger struct{}

func (*errLogger) Log(keyvals ...interface{}) error {
	for _, v := range keyvals {
		if v == poison {
			return errors.New("poisoned")
		}
	}
	fmt.Println(keyvals...)
	return nil
}
