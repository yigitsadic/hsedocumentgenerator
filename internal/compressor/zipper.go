package compressor

import "io/ioutil"

type Zipper struct {
}

func (z Zipper) WriteAsZip(fileName string, bytes []byte) error {
	return ioutil.WriteFile(fileName, bytes, 0644)
}
