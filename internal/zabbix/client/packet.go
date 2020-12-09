package client

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
)

// Packet.
type Packet struct {
	Request string
	IsCompress bool
	Data    []byte
	CData   []byte
}

// DataLen Packet method, return 8 bytes with packet length in little endian order.
func (p *Packet) DataLen() []byte {
	dataLen := make([]byte, 8)
	// JSONData, _ := json.Marshal(p)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(p.Data)))
	return dataLen
}

func (p *Packet) DataLenCompress() (dataLen1 []byte, dataLen2 []byte) {
	dataLen1 = make([]byte, 4)
	dataLen2 = make([]byte, 4)
	// JSONData, _ := json.Marshal(p)
	binary.LittleEndian.PutUint32(dataLen1, uint32(len(p.Data)))
	binary.LittleEndian.PutUint32(dataLen2, uint32(len(p.CData)))
	return
}


func (p *Packet) Compress() (dataLen1 []byte, dataLen2 []byte) {
	var buf bytes.Buffer
	dataLen1 =  make([]byte, 4)
	dataLen2 =  make([]byte, 4)
	binary.LittleEndian.PutUint32(dataLen1, uint32(len(p.Data)))
	w := zlib.NewWriter(&buf)
	w.Write(p.Data)
	w.Close()
	p.CData = buf.Bytes()
	binary.LittleEndian.PutUint32(dataLen2, uint32(len(p.CData)))
	return
}

func (p *Packet) Decompress() bool {
	buf := bytes.NewReader(p.CData)
	r, err := zlib.NewReader(buf)
	if err != nil {
		return false
	}
	p.Data, err = ioutil.ReadAll(r)
	if err != nil {
		return false
	}
	return true
}

func AnyByteCompress(data []byte) (cdata []byte, dataLen1 []byte, dataLen2 []byte) {
	var buf bytes.Buffer
	dataLen1 =  make([]byte, 4)
	dataLen2 =  make([]byte, 4)
	binary.LittleEndian.PutUint32(dataLen1, uint32(len(data)))
	w := zlib.NewWriter(&buf)
	w.Write(data)
	w.Close()
	cdata = buf.Bytes()
	binary.LittleEndian.PutUint32(dataLen2, uint32(len(cdata)))
	return
}
