// Copyright (c) 2018 Mark K Mueller <github.com/mkmueller>

// Package warn implements methods to utilize multiple warnings
//
package warn


type Warn interface {
	// Append another warning message
	Append(string)

	// Warnings will return a slice of warning messages
	Warnings() []string
}

type warnStr struct {
	s []string
}

// New creates a new warning message
func New(text string) Warn {
	return &warnStr{[]string{text}}
}

func (w *warnStr) Append (s string) {
	w.s = append(w.s, s)
}

func (w *warnStr) Warnings() []string {
	return w.s
}
