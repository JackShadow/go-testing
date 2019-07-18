package fileWorker

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func FileRead(path string) error {
	path = strings.TrimRight(path, "/") + "/" // so not to depend on the closing slash
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("cannot read from file, %v", err)
	}

	for _, f := range files {
		deleteFileName := path + f.Name()
		_, err := ioutil.ReadFile(path + f.Name())
		if err != nil {
			return err
		}
		err = os.Remove(deleteFileName) // clearing test files
	}
	return nil
}
