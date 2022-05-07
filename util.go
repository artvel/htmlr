package htmlr

import (
	"io/ioutil"
	"os"
)

func fileToBytes(fpath string) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	bts, er := ioutil.ReadAll(f)
	_ = f.Close()
	return bts, er
}
