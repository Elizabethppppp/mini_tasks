package _interface

type Reader interface{ Read(p []byte) (int, error) }
type Writer interface{ Write(p []byte) (int, error) }
type ReadWriter interface {
	Reader
	Writer
}

//мои копии

type CopyReader interface {
	MyRead(p []byte) (int, error)
}

type CopyWriter interface {
	MyWrite(p []byte) (int, error)
}

type CopyReadWriter interface {
	CopyReader
	CopyWriter
}
