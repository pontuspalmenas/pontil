package files

import (
	"fmt"
	"os"
	"testing"
)

func Test_ForAllFilesInDir(t *testing.T) {
	err := ForAllFilesInDir("/tmp", func(file *os.File) error {
		defer file.Close()
		return nil
	})
	if err != nil {
		t.Errorf("ForAllFilesInDir = %v, want nil", err)
	}
}

func Test_ForAllFilesInDir_WithForcedError(t *testing.T) {
	err := ForAllFilesInDir("/tmp", func(file *os.File) error {
		defer file.Close()
		return fmt.Errorf("forced error")
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_ForAllFilesInDir_WithUnexpectedError(t *testing.T) {
	err := ForAllFilesInDir("/this-directory-does-not-exist", func(file *os.File) error {
		return nil
	})
	if err == nil {
		t.Errorf("expected error")
	}
}
