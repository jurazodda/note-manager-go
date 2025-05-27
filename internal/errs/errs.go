package errs

import "errors"

var (
	ErrNoteIDRequired      = errors.New("err note id required")
	ErrNoteTitleRequired   = errors.New("err note title required")
	ErrNoteContentRequired = errors.New("err note content required")
	ErrNoteNotFound        = errors.New("err note not found")
)
