package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"time"
)

type pkgHeader struct {
	HeaderData   [2]byte
	HeaderLength uint32
}

var (
	HeaderData          = [2]byte{0xaa, 0xbb}
	HeaderLength uint32 = 6
)

type SocketUtil struct {
	conn   net.Conn
	reader []byte
}

//todo 这个返回读取多少长度的 还需要优化
func (s *SocketUtil) Read(b []byte) (n int, err error) {
	b, n, err = s.pkgReader()
	s.reader = b // 保存读取的数据
	return n, err
}

func (s *SocketUtil) Write(b []byte) (n int, err error) {
	return s.pkgWrite(b)
}

func (s *SocketUtil) Close() error {
	return nil
}

func (s *SocketUtil) LocalAddr() net.Addr {
	return nil
}

func (s *SocketUtil) RemoteAddr() net.Addr {
	return nil
}

func (s *SocketUtil) SetDeadline(t time.Time) error {
	return nil
}

func (s *SocketUtil) SetReadDeadline(t time.Time) error {
	return nil
}

func (s *SocketUtil) SetWriteDeadline(t time.Time) error {
	return nil
}

func NewSocketUtil(c net.Conn) *SocketUtil {
	return &SocketUtil{conn: c}
}

/**
获取到写入的数据
*/
func (s *SocketUtil) GetBytes() []byte {
	return s.reader
}

// 写入流数据
func (s *SocketUtil) pkgWrite(data []byte) (int, error) {
	// 写入数据 先写入头部数据
	buffer := bytes.NewBuffer([]byte{})
	// 二进制的方式写入
	binary.Write(buffer, binary.BigEndian, HeaderData)
	binary.Write(buffer, binary.BigEndian, uint32(len(data)))
	binary.Write(buffer, binary.BigEndian, data)
	//获取所有的内容 写入到socket之中
	allBytes := buffer.Bytes()
	return s.conn.Write(allBytes)
}

// 读入流数据
func (s *SocketUtil) pkgReader() ([]byte, int, error) {
	//先读入头部 并且判断 是不是一个流
	header, n, err := s.readerHeader()
	if err != nil {
		return nil, n, err
	}
	if header.HeaderData != HeaderData {
		return nil, 0, errors.New("package reader inivad")
	}
	// 读取存储的数字长度
	return s.readNByte(header.HeaderLength)
}

func (s *SocketUtil) readerHeader() (*pkgHeader, int, error) {
	nByte, n, err := s.readNByte(HeaderLength)
	if err != nil {
		return nil, n, err
	}
	buffer := bytes.NewBuffer(nByte)
	var p pkgHeader
	binary.Read(buffer, binary.BigEndian, &p.HeaderData)
	binary.Read(buffer, binary.BigEndian, &p.HeaderLength)
	return &p, n, nil
}
func (s *SocketUtil) readNByte(n uint32) ([]byte, int, error) {
	data := make([]byte, n)
	//开始进行读取
	var x = 0
	for x = 0; x < int(n); {
		//读取数据到read
		read, err := s.conn.Read(data[x:])
		if read == 0 {
			return nil, 0, errors.New("read package error")
		}
		if err != nil {
			return nil, 0, err
		}
		// 每次读取的流之后 加入到下次流读取的位置
		x += read
	}
	return data, x, nil
}
