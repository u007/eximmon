package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveSubFileFolder(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		log("deleting: %+v", filepath.Join(dir, name))
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func log(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}
