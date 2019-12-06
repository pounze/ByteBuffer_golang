package ByteBuffer

import (
	_"fmt"
	"errors"
	"bytes"
	"encoding/binary"
	"math"
)

type Buffer struct{
	packetBuffer bytes.Buffer
}

func (obj *Buffer) Wrap(data []byte) {
	obj.packetBuffer.Write(data)
}

func (obj *Buffer) PutShort(value int) {

	var buff = make([]byte, 2)
	
	binary.BigEndian.PutUint16(buff, uint16(value))

	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetShort() []byte{
	
	var tempBuff = obj.packetBuffer.Bytes()

	var shortValue = tempBuff[:2]

	var restValue = tempBuff[2:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return shortValue
}

func (obj *Buffer) PutInt(value int) {

	var buff = make([]byte, 4)

	binary.BigEndian.PutUint32(buff, uint32(value))

	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetInt() []byte{
	
	var tempBuff = obj.packetBuffer.Bytes()

	var intValue = tempBuff[:4]

	var restValue = tempBuff[4:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return intValue
}


func (obj *Buffer) PutLong(value int) {

	var buff = make([]byte, 8)

	binary.BigEndian.PutUint64(buff, uint64(value))

	obj.packetBuffer.Write(buff)

}

func (obj *Buffer) GetLong() []byte{
	
	var tempBuff = obj.packetBuffer.Bytes()

	var longValue = tempBuff[:8]

	var restValue = tempBuff[8:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return longValue
}

func (obj *Buffer) PutFloat(value float32) {

	var bits = math.Float32bits(value)

	var buff = make([]byte, 4)
	
	binary.BigEndian.PutUint32(buff, bits)
	
	obj.packetBuffer.Write(buff)
}

func (obj *Buffer) GetFloat() []byte{
	
	var tempBuff = obj.packetBuffer.Bytes()

	var floatValue = tempBuff[:4]

	var restValue = tempBuff[4:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return floatValue
}

func (obj *Buffer) PutDouble(value float64) {

	var bits = math.Float64bits(value)

	var buff = make([]byte, 8)
	
	binary.BigEndian.PutUint64(buff, bits)
	
	obj.packetBuffer.Write(buff)
}

func (obj *Buffer) GetDouble() []byte{
	
	var tempBuff = obj.packetBuffer.Bytes()

	var doubleValue = tempBuff[:8]

	var restValue = tempBuff[8:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return doubleValue
}

func (obj *Buffer) Put(value []byte) {

	obj.packetBuffer.Write(value)
}

func (obj *Buffer) Get(size int) []byte{

	var tempBuff = obj.packetBuffer.Bytes()

	var value = tempBuff[:size]

	var restValue = tempBuff[size:]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(restValue)

	obj.packetBuffer = byteBuffer

	return value
}

func (obj *Buffer) Array() []byte{
	return obj.packetBuffer.Bytes()
}

func (obj *Buffer) Size() int{
	return len(obj.packetBuffer.Bytes())
}

func (obj *Buffer) Flip(){

	var bytesArr = obj.packetBuffer.Bytes()

	for i, j := 0, len(bytesArr)-1; i < j; i, j = i+1, j-1 {
		bytesArr[i], bytesArr[j] = bytesArr[j], bytesArr[i]
	}

	var byteBuffer bytes.Buffer

	byteBuffer.Write(bytesArr) 

	obj.packetBuffer = byteBuffer;
}

func (obj *Buffer) Clear(){
	var byteBuffer bytes.Buffer
	obj.packetBuffer = byteBuffer;
}

func (obj *Buffer) Slice(start int, end int) error{

	var bytesArr = obj.packetBuffer.Bytes()

	if len(bytesArr) < (start + end){
		return errors.New("Buffer does not contain that much of limit")
	}

	bytesArr = bytesArr[start:end]

	var byteBuffer bytes.Buffer

	byteBuffer.Write(bytesArr) 

	obj.packetBuffer = byteBuffer;

	return nil

}