package wb

import "fmt"

type Request struct {
	file string
	data string
	resp chan bool
}

func createReq(req chan Request, filename string, value interface{}) {

	strValue := fmt.Sprintf("%v", value)

	resp := make(chan bool)

	req <- Request{filename, strValue, resp}
	<-resp

}
