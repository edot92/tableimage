package tableimage

import (
	"image"
)

const (
	rowSpace         = 30
	tablePadding     = 20
	letterPerPx      = 10
	separatorPadding = 10
)

type tableImage struct {
	width           int
	height          int
	backgroundColor string
	img             *image.RGBA
}

func Start(ths []string, trs [][]string) {

	var totalRowNo = len(trs) + 1
	var totalColumnNo = len(ths)
	var columnSpace = getMaxAmountOfLetters(ths, trs) * letterPerPx
	var H = totalRowNo*rowSpace + rowSpace
	var W = totalColumnNo * int(columnSpace)

	ti := createContext(W, H, "#FFFFFF")

	for colNo, th := range ths {
		ti.addString(colNo*columnSpace+tablePadding, 1*rowSpace, th, "#000000")
	}

	//draw the double line for th's
	ti.addLine(1, 1*rowSpace+separatorPadding, W, 1*rowSpace+separatorPadding, "#000000")
	ti.addLine(1, 1*rowSpace+separatorPadding+2, W, 1*rowSpace+separatorPadding+2, "#000000")

	for rowNo, tds := range trs {
		//start with the second row since the firsone is the th
		fRowNo := rowNo + 2
		for colNo, td := range tds {

			if colNo == 0 {
				ti.addString(columnSpace-columnSpace+tablePadding, fRowNo*rowSpace, td, "#000")
			}
			ti.addString(colNo*columnSpace+tablePadding, fRowNo*rowSpace, td, "#000")
		}
		ti.addLine(1, fRowNo*rowSpace+separatorPadding, W, fRowNo*rowSpace+separatorPadding, "#000")
	}

	ti.save("png", "./out.png")
}
