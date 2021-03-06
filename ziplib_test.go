package ziplib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestZipAndUnzip(t *testing.T) {
	tmp, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.RemoveAll(tmp)

	// zip
	src := filepath.Join(tmp, "src")
	dst := filepath.Join(tmp, "dst.zip")
	if err = os.MkdirAll(src, os.ModePerm); err != nil {
		t.Error(err.Error())
	}
	if err = ioutil.WriteFile(filepath.Join(src, "a.txt"), []byte("test"), os.ModePerm); err != nil {
		t.Error(err.Error())
	}
	if err = ioutil.WriteFile(filepath.Join(src, "b.txt"), []byte("test"), os.ModePerm); err != nil {
		t.Error(err.Error())
	}
	if err = os.MkdirAll(filepath.Join(src, "c"), os.ModePerm); err != nil {
		t.Error(err.Error())
	}
	if err = ioutil.WriteFile(filepath.Join(src, "c", "c.txt"), []byte("test"), os.ModePerm); err != nil {
		t.Error(err.Error())
	}

	if err = ZipFolder(src, dst); err != nil {
		t.Error(err.Error())
	}

	if _, err = os.Stat(dst); err != nil {
		t.Error(err.Error())
	}

	// unzip
	srcNew := filepath.Join(tmp, "srcnew")
	if err = os.MkdirAll(srcNew, os.ModePerm); err != nil {
		t.Error(err.Error())
	}

	if err = Unzip(dst, srcNew); err != nil {
		t.Error(err.Error())
	}

	if _, err = os.Stat(filepath.Join(srcNew, "a.txt")); err != nil {
		t.Error(err.Error())
	}
	if _, err = os.Stat(filepath.Join(srcNew, "b.txt")); err != nil {
		t.Error(err.Error())
	}
	if _, err = os.Stat(filepath.Join(srcNew, "c", "c.txt")); err != nil {
		t.Error(err.Error())
	}
}

func TestCompress(t *testing.T) {
	tmp, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.RemoveAll(tmp)

	if err = ioutil.WriteFile(filepath.Join(tmp, "a.txt"), []byte("test"), os.ModePerm); err != nil {
		t.Error(err.Error())
	}

	data, err := ioutil.ReadFile(filepath.Join(tmp, "a.txt"))
	if err != nil {
		t.Error(err)
	}

	data, err = Compress(data)
	if err != nil {
		t.Error(err)
	}
	if err = ioutil.WriteFile(filepath.Join(tmp, "a.zip"), data, os.ModePerm); err != nil {
		t.Error(err)
	}

	data, err = ioutil.ReadFile(filepath.Join(tmp, "a.zip"))
	if err != nil {
		t.Error(err)
	}

	data, err = Decompress(data)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "test" {
		t.Error(fmt.Errorf("content error"))
	}
}
