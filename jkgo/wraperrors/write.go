package wraperrors

import (
	"github.com/pkg/errors"
	"io"
)

func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	return errors.Wrap(err, "write failed")
}
