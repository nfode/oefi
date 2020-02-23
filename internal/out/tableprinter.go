package out

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

const SEPARATOR = "-----"

type TablePrinter struct {
	Cols   int
	Header []string
	lines  []string
}

func (printer *TablePrinter) AddLine(values map[string]string) error {
	builder := strings.Builder{}
	for i, header := range printer.Header {
		if value, ok := values[header]; ok {
			if i == 0 {
				builder.WriteString(fmt.Sprintf("\n %s", value))
			} else {
				builder.WriteString(fmt.Sprintf("\t%s", value))
			}
		} else {
			return errors.New(fmt.Sprintf("No entry for header: %s", header))
		}
	}
	builder.WriteString("\t")
	printer.lines = append(printer.lines, builder.String())
	return nil
}

func (printer *TablePrinter) headerLine() string {
	header := strings.Builder{}
	sep := strings.Builder{}
	for i, s := range printer.Header {
		if i == 0 {
			header.WriteString(s)
			sep.WriteString(SEPARATOR)
		} else {
			header.WriteString(fmt.Sprintf("\t%s", s))
			sep.WriteString(fmt.Sprintf("\t%s", SEPARATOR))
		}
	}
	header.WriteString("\t")
	sep.WriteString("\t")
	header.WriteString(fmt.Sprintf("\n %s", sep.String()))
	return header.String()
}

func (printer *TablePrinter) Print() {
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	fmt.Fprint(w, printer.headerLine())
	for _, line := range printer.lines {
		fmt.Fprint(w, line)
	}
	w.Flush()
}
