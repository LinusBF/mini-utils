package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	fmt.Println("IMPORTANT!!! This is intended to be ran with no files other than MainActivity.kt in the android src")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Whats the new BundleId?")
	oldId, _ := reader.ReadString('\n')
	oldId = strings.Replace(oldId, "\n", "", -1)

	oldPath := strings.Join(strings.Split(oldId, "."), "/")
	oldPath = "./android/app/src/main/kotlin" + oldPath + "/MainActivity.kt"

	fmt.Println("Whats the new BundleId?")
	bundleId, _ := reader.ReadString('\n')
	bundleId = strings.Replace(bundleId, "\n", "", -1)

	bundlePath := strings.Join(strings.Split(bundleId, "."), "/")
	bundlePath = "./android/app/src/main/kotlin" + bundlePath

	err := os.MkdirAll(bundlePath, fs.ModeDir)
	if err != nil {
		fmt.Printf("Failed to create new dir %s", bundlePath)
		return
	}

	err = os.Rename(oldPath, bundlePath+"/MainActivity.kt")
	if err != nil {
		fmt.Printf("Failed to move to new dir %s", bundlePath)
		return
	}
}
