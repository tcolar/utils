// History: May 02 14 tcolar Creation

package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Recursively copy the contents of directory src into the "to" folder
func RecursiveCopy(src, to string) error {
	_, err := os.Stat(src)
	if err != nil {
		return err
	}
	files, _ := ioutil.ReadDir(src)
	for _, f := range files {
		nm := f.Name()
		if f.IsDir() {
			folder := path.Join(to, nm)
			err := CreateFolder(folder, f.Mode())
			if err != nil {
				return err
			}
			err = RecursiveCopy(path.Join(src, nm), folder)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(path.Join(src, nm), path.Join(to, nm))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CreateFolder creates a folder
// Does nothing if the folder already exists
func CreateFolder(folder string, mode os.FileMode) error {
	f, err := os.Stat(folder)
	if err == nil {
		if f.IsDir() {
			// Alreday exists, that's fine, leave it alone
			return nil
		} else {
			return fmt.Errorf("Can't create folder %s because their is a file with the same name.", folder)
		}
	}
	// Ok trying to create it now
	err = os.MkdirAll(folder, mode)
	return err
}

// CopyFile copies the contents of the src file as the file denoted by dst
// The dest file will be created or overwritten.
// The src File mode will be applied to dst.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		CloseFile(in)
	}()
	fz, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		CloseFile(out)
	}()
	// Actual copy of bytes
	sz, err := io.Copy(out, in)
	if err != nil {
		return err
	}
	if sz != fz.Size() {
		return fmt.Errorf("Failed to copy the whole file ?? %s size: %d, %s size : %d,",
			src, sz, dst, fz.Size())
	}
	// Trying to preserve file mode
	if err = out.Chmod(fz.Mode()); err != nil {
		return err
	}
	err = out.Sync()
	return err
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		// I'm not clear what can be done if close fails
		// Ether ignore or panic ?
		panic(err)
	}
}
