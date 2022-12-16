package helper

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func base_dir() string {

	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return mydir
}

func InputReader(filename string)string{
	var path string = base_dir()
	path = filepath.Join(path, filename)
	
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}