package fileWorkerAlt

import (
	"testing"
	"github.com/spf13/afero"
)

const testSqlDir = "test/"
const wrongDir = "no/data/here/"

//func TestFileWorkerAlt(t *testing.T) {
//
//	appFS := afero.NewMemMapFs()
//	// create test files and directories
//	appFS.MkdirAll(testSqlDir, 0755)
//	afero.WriteFile(appFS, testSqlDir+"good.data", []byte("Readable"), 0644)
//	afero.WriteFile(appFS, testSqlDir+"bad.data", []byte("Invalid"), 0000)
//
//	err := FileReadAlt(wrongDir, appFS)
//	if err == nil {
//		t.Error("No error from wrong dir")
//	}
//	FileReadAlt(testSqlDir+"/", appFS)
//}

func BenchmarkFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appFS := afero.NewMemMapFs()
		// create test files and directories
		appFS.MkdirAll(testSqlDir, 0755)
		afero.WriteFile(appFS, testSqlDir+"good.data", []byte("Readable"), 0644)

		FileReadAlt(testSqlDir+"/", appFS)
	}
}
