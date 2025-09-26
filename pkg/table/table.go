package table

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorReset  = "\033[0m"
)

func Tables(mirror io.Writer, length int, Title string, Header table.Row) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(mirror)
	t.SetAllowedRowLength(length)
	t.SetTitle(Title)
	t.AppendHeader(Header)
	return t
}

func ColorData(row table.Row) table.Row {
	colorRows := make(table.Row, len(row))
	colors := []string{colorRed, colorGreen, colorBlue, colorPurple}

	for i, v := range row {
		colorIndex := i % len(colors)
		colorRows[i] = colors[colorIndex] + fmt.Sprintf("%v", v) + colorReset
	}
	return colorRows
}
