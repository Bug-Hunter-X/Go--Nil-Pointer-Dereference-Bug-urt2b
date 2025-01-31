# Go Nil Pointer Dereference Bug

This repository demonstrates a common error in Go: dereferencing a nil pointer. This leads to a runtime panic, crashing the program.  The `bug.go` file shows the erroneous code, while `bugSolution.go` provides a corrected version.

**Problem:**

The `bug.go` file contains a pointer variable (`y`) that is initially `nil`.  The code attempts to dereference this `nil` pointer, causing a runtime panic.

**Solution:**

The `bugSolution.go` file demonstrates the correct way to handle pointers.  It checks for `nil` before dereferencing, preventing the panic.