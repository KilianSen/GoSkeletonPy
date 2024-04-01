# GoSkeletonPy
GoSkeletonPy is a project that allows you to generate a skeletonized file from a Python file and restore it.
This is done by encrypting the Python file using AES and storing it in a file, then replacing all the function bodies in the original file with `...`. 
The original file can be restored from the skeleton file.

# Example

### Original File
```python
# test.py
def add(a, b):
    return a + b
    
valueA = 5
    
def sub(a, b):
    return a - b
```

### Skeletonized File
```python
# test.py
def add(a, b):
    ...
    
valueA = ...
    
def sub(a, b):
    ...
```

## Usage
Here's a basic usage example:

```go
package main

import (
	GoSkeletonPy "GoSkeletonPy/src"
	"time"
)

func main() {
	// Generate a skeleton file from a python file
	println("Generating skeleton file...")
	skeleton, err := GoSkeletonPy.FileToSkeleton("./test/test.py", "", "")
	if err != nil {
		panic(err)
	}
	println(skeleton)

	// sleep for 5 seconds
	time.Sleep(5 * time.Second)

	file, err := GoSkeletonPy.RestoreSkeletonFile("./test/test.skeletonpy", "", "")
	if err != nil {
		panic(err)
	}

	println(file)

}