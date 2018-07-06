package fileWorkerAlt2

import (
	"testing"
	"github.com/spf13/afero"
)

const testSqlDir = "test/"
const wrongDir = "no/data/here/"


func BenchmarkFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appFS := afero.NewMemMapFs()
		// create test files and directories
		appFS.MkdirAll(testSqlDir, 0755)
		afero.WriteFile(appFS, testSqlDir+"good.data", []byte("Readable"), 0644)

		FileReadAlt(testSqlDir+"/", appFS)
	}
}
