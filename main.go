package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func find(out io.Writer, format string, filePath string, printFiles bool) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if !printFiles {
		var lastDir int
		for idx, f := range files {
			if f.IsDir() {
				lastDir = idx
			}
		}
		for idx, f := range files {
			if f.IsDir() {
				if idx == lastDir {
					fmt.Fprintln(out, format+"└"+"───"+f.Name())
					find(out, format+"\t", filePath+f.Name()+"/", printFiles)
				} else {
					fmt.Fprintln(out, format+"├"+"───"+f.Name())
					find(out, format+"│"+"\t", filePath+f.Name()+"/", printFiles)
				}
			}
		}
	} else {
		for idx, f := range files {
			if f.IsDir() {
				if idx == len(files)-1 {
					fmt.Fprintln(out, format+"└"+"───"+f.Name())
					find(out, format+"\t", filePath+f.Name()+"/", printFiles)
				} else {
					fmt.Fprintln(out, format+"├"+"───"+f.Name())
					find(out, format+"│"+"\t", filePath+f.Name()+"/", printFiles)
				}
			} else {
				var temp string
				if f.Size() == 0 {
					temp = "empty"
				} else {
					temp = strconv.Itoa(int(f.Size()))
					temp += "b"
				}
				if idx == len(files)-1 {
					fmt.Fprintln(out, format+"└"+"───"+f.Name()+" ("+temp+")")
				} else {
					fmt.Fprintln(out, format+"├"+"───"+f.Name()+" ("+temp+")")
				}
			}
		}
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	path += "/"
	find(out, "", path, printFiles)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

}
