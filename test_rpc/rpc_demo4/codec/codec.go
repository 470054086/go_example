package codec

import "github.com/vmihailenco/msgpack"
type SerializeType byte

//  这里到时候 可以移出去
const (
	MessagePack SerializeType = iota
)
var codecs = map[SerializeType]Codec{
	MessagePack: &MessagePackCodec{},
}
type Codec interface {
	// 解压缩的形式
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

// 获取当前的codec
func GetCodec(t SerializeType) Codec {
	return codecs[t]
}


// 使用package进行序列化
type MessagePackCodec struct {
}

func (m MessagePackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (m MessagePackCodec) Decode(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b,v)
}



