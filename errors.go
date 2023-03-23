package tiktokengrpc

import (
	"fmt"
)

// InvalidEncodingNameErr is returned if the given TiktokenEncodingName is not recognized by the library.
type InvalidEncodingNameErr TiktokenEncodingName

func (e InvalidEncodingNameErr) Error() string {
	return fmt.Sprintf("%s is not a valid encoding name", string(e))
}
