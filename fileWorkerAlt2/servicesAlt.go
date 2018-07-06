package fileWorkerAlt2

import (
	"strings"
	"fmt"
	"github.com/spf13/afero"
)

func FileReadAlt(path string, fs afero.Fs) error {
	path = strings.TrimRight(path, "/") + "/" // независим от заверщающего слега
	files, err := afero.ReadDir(fs, path)
	if err != nil {
		return fmt.Errorf("cannot read from file, %v", err)
	}

	for _, f := range files {
		deleteFileName := path + f.Name()
		_, err := afero.ReadFile(fs, path+f.Name())
		if err != nil {
			return err
		}
		err = fs.Remove(deleteFileName) // после вывода удаляем файл
	}
	return nil
}
