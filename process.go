package wb

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/tealeg/xlsx"
)

var m map[string]fileRef
var c chan Request
var nWorker int

type fileRef struct {
	id   string
	file *xlsx.File
}

func Init() {
	m = make(map[string]fileRef)
	c = make(chan Request, 0)
	nWorker = 2

	go initBalancer().balance(c)
}

func Process(filename string, value interface{}) {

	_, ok := m[filename]
	if !ok {
		createFile(filename)
	}

	createReq(c, filename, value)

}

func fileExists(filename string) bool {
	info, err := os.Stat(fmt.Sprintf("%s.xlsx", filename))

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func createFile(filename string) {

	if fileExists(filename) {
		xlFile, err := xlsx.OpenFile(fmt.Sprintf("%s.xlsx", filename))
		if err != nil {
			log.Errorf("Error opening file xlsx %v", err)
		}
		wref := fileRef{id: filename, file: xlFile}
		m[filename] = wref
		return
	}

	var file *xlsx.File
	var err error

	file = xlsx.NewFile()

	wref := fileRef{id: filename, file: file}
	m[filename] = wref

	_, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = file.Save(fmt.Sprintf("%s.xlsx", filename))
	if err != nil {
		fmt.Printf(err.Error())
	}

}
