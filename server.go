package lrpc

import "github.com/Ekreke/simple-rpc-implentation/codec"

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

// | Option{MagicNumber: xxx, CodecType: xxx} | Header{ServiceMethod ...} | Body interface{} |
// the message may like this | Option | Header1 | Body1 | Header2 | Body2 | ...
var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}
