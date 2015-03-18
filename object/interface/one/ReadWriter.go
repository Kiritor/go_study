package one

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Writer(buf []byte) (n int, err error)
}
