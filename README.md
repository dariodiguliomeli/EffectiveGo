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