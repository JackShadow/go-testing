package fileWorker

import (
	"testing"
	"os"
)

const testSqlDir = "test/"
const wrongDir = "no/data/here/"

func TestMain(m *testing.M) {
	os.Mkdir(testSqlDir, 0777)
	file, _ := os.Create(testSqlDir + "1good.data")
	file.WriteString("Readable")
	file.Close()
	file, _ = os.Create(testSqlDir + "2bad.data")
	file.Chmod(0000)
	file.WriteString("Invalid data")
	v := m.Run()
	file.Close()

	os.Remove(testSqlDir)
	os.Exit(v)
}

func TestFileWorker(t *testing.T) {
	err := FileRead(wrongDir)
	if err == nil {
		t.Error("No error from wrong dir")
	}
	FileRead(testSqlDir + "/")
}

func BenchmarkFileRead(b *testing.B) {
	os.Mkdir(testSqlDir, 0777)
	for i := 0; i < b.N; i++ {
		file, _ := os.Create(testSqlDir + "1good.data")
		file.WriteString("Readable")
		file.Close()
		FileRead(testSqlDir + "/")
	}
}
