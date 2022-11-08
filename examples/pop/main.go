package main

import (
	"fmt"
	"strings"
)

func composite(main, overlay string, xoffset, yoffset int) string {
	doc := strings.Builder{}
	m := strings.Split(main, "\n")
	o := strings.Split(overlay, "\n")

	for i, row := range m {

		for j, char := range row {
			if j < xoffset || i < yoffset || i >= len(o)+yoffset || j >= len(o[i-yoffset])+xoffset {

				doc.WriteRune(char)
				continue
			}

			doc.WriteByte(o[i-yoffset][j-xoffset])

		}
		doc.WriteString("\n")

	}

	return doc.String()
}

const (
	A = `_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________
_______________________________________________________________`

	B = `OOOOOOOOOO
OOOOOOOOOO
OOOOOOOOOO
OOOOOOOOOO`
)

func main() {
	fmt.Println(composite(A, B, 10, 3))
}
