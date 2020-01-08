package wb

import "fmt"

type Work struct {
	idx     int
	wok     chan Request
	pending int
}

func (w *Work) doWork(done chan *Work) {
	for {
		req := <-w.wok
		execute(req.resp, req.file, req.data)
		done <- w
	}
}

func execute(crsp chan bool, filename, value string) {
	fref := m[filename]

	for _, sheet := range fref.file.Sheets {
		if sheet.Name == "Sheet1" {
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.Value = value
			err := fref.file.Save(fmt.Sprintf("./%s.xlsx", filename))
			if err != nil {
				fmt.Println(err.Error())
				crsp <- false
			}
			crsp <- true
		}
	}
}
