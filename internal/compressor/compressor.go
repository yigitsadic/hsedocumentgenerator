package compressor

type ZipWriter interface {
	WriteAsZip([]byte) error
}
