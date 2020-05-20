package common

//type PkgHeader struct {
//	HeaderFlag [2]byte
//	DataLength uint32
//}
//type SocketUtil struct {
//	Coon net.Conn
//}
//
//const (
//	HeaderLength = 6
//)
//const (
//	SOCKET_ERROR_CLIENT_CLOSED = "Client has been closed"
//	SOCKET_ERROR_SERVER_CLOSED = "Server has been closed"
//	SOCKET_ERROR_TIMEOUT       = "Timeout"
//)
//
//func (fd *SocketUtil) WritePkg(data []byte) (int, error) {
//	if fd == nil {
//		return -1, errors.New("error")
//	}
//	if len(data) == 0 {
//		return 0, nil
//	}
//	buff := bytes.NewBuffer([]byte{})
//	binary.Write(buff, binary.BigEndian, []byte{0xaa, 0xbb}) //添加协议头
//	binary.Write(buff, binary.BigEndian, uint32(len(data)))  //添加长度
//	binary.Write(buff, binary.BigEndian, data)               //数据部分
//	allBytes := buff.Bytes()
//	return fd.writeNByte(allBytes)
//}
//
//func (fd *SocketUtil) writeNByte(data []byte) (int, error) {
//	n, err := fd.Coon.Write(data)
//	if err != nil {
//		return -1, err
//	} else {
//		return n, nil
//	}
//}
//
//func (fd *SocketUtil) ReadPkg() ([]byte, error) {
//	if fd == nil || fd.Coon == nil {
//		return nil, errors.New("error")
//	}
//	head, err := fd.readHead() //先读取数据头
//	if err != nil {
//		return nil, err
//	}
//	//数据头和约定不一样，报错
//	if head.HeaderFlag != [2]byte{0xaa, 0xbb} {
//		return nil, errors.New("Head package error")
//	}
//	//读取指定长度的数据
//	datas, err := fd.readNByte(head.DataLength)
//	if err != nil {
//		return nil, err
//	}
//	return datas, nil
//}
//
//func (fd *SocketUtil) readHead() (*PkgHeader, error) {
//	data, err := fd.readNByte(HeaderLength)
//	if err != nil {
//		return nil, err
//	}
//	h := PkgHeader{}
//	buff := bytes.NewBuffer(data)
//	binary.Read(buff, binary.BigEndian, &h.HeaderFlag) //读取0xaa 0xbb连个字节
//	binary.Read(buff, binary.BigEndian, &h.DataLength) //读取四个字节的长度
//	return &h, nil
//}
//
//func (fd *SocketUtil) readNByte(n uint32) ([]byte, error) {
//	data := make([]byte, n)
//	for x := 0; x < int(n); {
//		length, err := fd.Coon.Read(data[x:]) //有数据则读，没有则阻塞
//		if length == 0 {
//			return nil, errors.New(SOCKET_ERROR_CLIENT_CLOSED)
//		}
//		if err != nil {
//			return nil, err
//		}
//		x += length
//	}
//	return data, nil
//}
