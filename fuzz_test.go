package archive

import (
	"bytes"
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

func FuzzUntar(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		ff := fuzz.NewConsumer(data)
		tarBytes, err := ff.TarBytes()
		if err != nil {
			return
		}
		options := &TarOptions{}
		err = ff.GenerateStruct(options)
		if err != nil {
			return
		}
		tmpDir := t.TempDir()
		_ = Untar(bytes.NewReader(tarBytes), tmpDir, options)
	})
}

func FuzzApplyLayer(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tmpDir := t.TempDir()
		_, _ = ApplyLayer(tmpDir, bytes.NewReader(data))
	})
}
