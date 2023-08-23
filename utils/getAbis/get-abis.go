package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
* main is used as a gateway for get-abis.go util. It is used to write and delete proper ABI files from given paths.
* It is used to generate go files from ABIs using abigen.
 */
func main() {
	args := os.Args
	if len(args) > 2 {

		switch args[1] {
		case "--get":
			for i := 2; i < len(args); i++ {
				getABI(args[i])
			}
		case "--rm":
			for i := 2; i < len(args); i++ {
				removeABI(args[i])
			}
		default:
			fmt.Println("get-abis util: invalid command")
		}

	} else {
		fmt.Println("get-abis util: no command or args")
	}
}

// getABI is used to get ABI from given path and write proper ABI file in ./contracts/
func getABI(path string) {

	// Get and read source file from given path

	sourceFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("get-abis util: error parsing abi %v: %v", path, err.Error())
		return
	}
	defer sourceFile.Close()

	b, _ := io.ReadAll(sourceFile)

	// Unmarshal given file and get bytes from the "abi" block

	var res map[string]interface{}
	if err = json.Unmarshal(b, &res); err != nil {
		fmt.Printf("get-abis util: error unmarshal abi %v: %v", path, err.Error())
		return
	}

	abiB, err := json.Marshal(res["abi"])
	if err != nil {
		fmt.Printf("get-abis util: error marshal abi %v: %v", path, err.Error())
	}

	// Get file name and create new temp file

	pathSplit := strings.Split(path, "/")
	fileName := pathSplit[len(pathSplit)-1]

	writeFile, err := os.Create(fmt.Sprintf("./contracts/%v", fileName))
	if err != nil {
		fmt.Printf("get-abis util: error create file %v: %v", fileName, err.Error())
		return
	}
	defer writeFile.Close()

	// Write file

	_, err = writeFile.Write(abiB)
	if err != nil {
		fmt.Printf("get-abis util: error writing file %v: %v", fileName, err.Error())
	}
}

// removeABI is used to remove file by given path
func removeABI(path string) {
	if err := os.Remove(path); err != nil {
		fmt.Printf("get-abis util: error removing %v: %v", path, err.Error())
	}
}
