package wb

import (
	"fmt"

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
		log.Infof("File Creato \n")
	}

	createReq(c, filename, value)

}

func createFile(filename string) {

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
