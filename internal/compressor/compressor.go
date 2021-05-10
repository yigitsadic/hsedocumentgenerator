package compressor

type ZipWriter interface {
	WriteAsZip(string, []byte) error
}
