package util

import (
	"os"
	"testing"
)

func TestPathExistsAndRemoveDir(t *testing.T) {
	dirPath := "testdir"
	if err := os.Mkdir(dirPath, 0755); err != nil {
		t.Fatal("Make dir failed")
	}
	t.Log("Make dir success")

	if exists, err := PathExists(dirPath); err != nil {
		t.Fatal("Check if path exists failed")
	} else {
		if exists {
			t.Log("Path exists")
		} else {
			t.Log("Path dose not exists")
		}
	}

	if err := RemoveDir(dirPath); err != nil {
		t.Fatal("Remove dir error")
	}
	t.Log("Remove dir success")
}
