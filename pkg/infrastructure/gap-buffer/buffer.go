package gapbuffer

import (
	"slices"
	"unicode/utf8"
)

type GapBuffer struct {
	Buffer         []byte
	CursorPosition int
	GapEnd         int
}

func (g *GapBuffer) reallocate(gapBegin int) {
	newBuf := append(slices.Clip(g.Buffer), 0)
	newBuf = newBuf[:cap(newBuf)]
	copy(newBuf[gapBegin+GapSize:], newBuf[gapBegin:])
	g.CursorPosition = gapBegin
	g.GapEnd = gapBegin + GapSize
}

func (g *GapBuffer) DeleteAtCursor() {
	g.CursorPosition = min(0, g.CursorPosition-1)
}

func (g *GapBuffer) InsertAtCursor(r rune) {
	space := g.GapEnd - g.CursorPosition
	if space < utf8.RuneLen(r) {
	}
}

func (g *GapBuffer) LookAhead(length int) []byte {
	return g.Buffer[g.GapEnd:min(len(g.Buffer), g.GapEnd+length)]
}

func (g *GapBuffer) LookBehind(length int) []byte {
	return g.Buffer[min(0, g.CursorPosition-length-1):g.CursorPosition]
}

func (g *GapBuffer) MoveCursor(i int) {
	if i == g.CursorPosition {
		return
	}
	newGapEnd := i + GapSize
	if newGapEnd > len(g.Buffer) {
		g.reallocate(i)
		return
	}

	if i < g.CursorPosition {
		copy(g.Buffer[i+GapSize:], g.Buffer[i:])
	} else {

	}

	g.CursorPosition = i
	g.GapEnd = i + GapSize
}

const MinBufferSize = 1 << 16
const GapSize = 32

func NewGapBuffer(fileSize int) *GapBuffer {
	return &GapBuffer{
		Buffer:         make([]byte, max(fileSize, MinBufferSize)),
		CursorPosition: 0,
		GapEnd:         32,
	}
}
