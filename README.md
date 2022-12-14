# Effective Go

> A repository designed to practice Effective Go heuristics.

## Brief

### Introduction

### Formatting

> Formatting issues are the most contentious but the least consequential.

```bash
gofmt path/file.go
```


* Indentation
We use tabs for indentation and gofmt emits them by default. Use spaces only if you must.
* Line length
Go has no line length limit. Don't worry about overflowing a punched card. If a line feels too long, wrap it and indent with an extra tab.
* Parentheses
Go needs fewer parentheses than C and Java: control structures (if, for, switch) do not have parentheses in their syntax. Also, the operator precedence hierarchy is shorter and clearer, so
``` 
x<<8 + y<<16
```
means what the spacing implies, unlike in the other languages.

### Commentary
> Comments that appear before top-level declarations, with no intervening newlines.

```go
package bytes

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	
}
```

### Names

#### Package names
* Lowercase
* Single word
* The package name is the base name of its source directory

### Getters
> If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner.

### Interfaces
> By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

