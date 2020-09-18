package files

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)


func TestCopyFile(t *testing.T) {
	file, err := ioutil.TempFile("", "source.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	initialData := "file content"
	_, err = file.WriteString(initialData)
	assert.NoError(t, err)

	dir, err := ioutil.TempDir("", "dest.*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dir)

	destFile := path.Join(dir, "dest.txt")
	assert.NoError(t, CopyFile(file.Name(), destFile))

	data, err := ioutil.ReadFile(destFile)
	assert.NoError(t, err)

	assert.Equal(t, string(data), initialData)
}
