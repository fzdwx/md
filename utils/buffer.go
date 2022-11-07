package utils

import (
	"strings"
)

type StrBuffer struct {
	builder strings.Builder
}

// NewStrBuffer the constructor, only use val[0]
func NewStrBuffer(val ...string) *StrBuffer {
	b := &StrBuffer{builder: strings.Builder{}}
	if len(val) > 0 {
		b.Write(val[0])
	}
	return b
}

// NewLine append '\n' to buffer
func (b *StrBuffer) NewLine() *StrBuffer {
	b.builder.WriteRune('\n')
	return b
}

func (b *StrBuffer) String() string {
	return b.builder.String()
}

func (b *StrBuffer) Write(str string) *StrBuffer {
	b.builder.WriteString(str)

	return b
}
