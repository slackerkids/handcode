package service

type TextBuffer interface {
	MoveCursor(i int)
	InsertAtCursor(r rune)
	DeleteAtCursor()
	LookBehind(length int) []byte
	LookAhead(length int) []byte
}
