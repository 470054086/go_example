package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

type pkgHeader struct {
	HeaderData   [2]byte
	HeaderLength uint32
}

var (
	HeaderData          = [2]byte{0xaa, 0xbb}
	HeaderLength uint32 = 6
)

type socketUtil struct {
	conn net.Conn
}

func NewSocketUtil(c net.Conn) *socketUtil {
	return &socketUtil{conn: c}
}

// 写入流数据
func (s *socketUtil) PkgWrite(data []byte) (int, error) {
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
func (s *socketUtil) PkgReader() ([]byte, error) {
	//先读入头部 并且判断 是不是一个流
	header, err := s.readerHeader()
	if err != nil {
		return nil, err
	}
	if header.HeaderData != HeaderData {
		return nil, errors.New("package reader inivad")
	}
	// 读取存储的数字长度
	return s.readNByte(header.HeaderLength)
}

func (s *socketUtil) readerHeader() (*pkgHeader, error) {
	nByte, err := s.readNByte(HeaderLength)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nByte)
	var p pkgHeader
	binary.Read(buffer, binary.BigEndian, &p.HeaderData)
	binary.Read(buffer, binary.BigEndian, &p.HeaderLength)
	return &p, nil
}
func (s *socketUtil) readNByte(n uint32) ([]byte, error) {
	data := make([]byte, n)
	//开始进行读取
	for x := 0; x < int(n); {
		//读取数据到read
		read, err := s.conn.Read(data[x:])
		if read == 0 {
			return nil, errors.New("read package error")
		}
		if err != nil {
			return nil, err
		}
		// 每次读取的流之后 加入到下次流读取的位置
		x += read
	}
	return data, nil
}
