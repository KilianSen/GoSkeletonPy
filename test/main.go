package main

import (
	"GoSkeletonPy"
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
