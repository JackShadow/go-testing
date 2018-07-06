package fileWorker

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
)

func FileRead(path string) error {
	path = strings.TrimRight(path, "/") + "/" // независим от заверщающего слега
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
		err = os.Remove(deleteFileName) // после вывода удаляем файл
	}
	return nil
}
