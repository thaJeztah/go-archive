package chrootarchive

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/moby/go-archive"
	"github.com/moby/go-archive/compression"
)

// applyLayerHandler parses a diff in the standard layer format from `layer`, and
// applies it to the directory `dest`. Returns the size in bytes of the
// contents of the layer.
func applyLayerHandler(dest string, layer io.Reader, options *archive.TarOptions, decompress bool) (size int64, err error) {
	if decompress {
		decompressed, err := compression.DecompressStream(layer)
		if err != nil {
			return 0, err
		}
		defer decompressed.Close()

		layer = decompressed
	}

	// Ensure it is a Windows-style volume path
	dest = addLongPathPrefix(filepath.Clean(dest))
	s, err := archive.UnpackLayer(dest, layer, nil)
	if err != nil {
		return 0, fmt.Errorf("ApplyLayer %s failed UnpackLayer to %s: %w", layer, dest, err)
	}

	return s, nil
}
