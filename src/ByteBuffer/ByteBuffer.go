package ByteBuffer

import (
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

func (obj *Buffer) Bytes2Str(data []byte) string{

	return string(data)

}

func (obj *Buffer) Str2Bytes(data string) []byte{

	return []byte(data)

}

func (obj *Buffer) Bytes2Short(data []byte) uint16{

	return binary.BigEndian.Uint16(data)
}

func (obj *Buffer) Bytes2Int(data []byte) uint32{

	return binary.BigEndian.Uint32(data)

}

func (obj *Buffer) Bytes2Long(data []byte) uint64{

	return binary.BigEndian.Uint64(data)

}

func (obj *Buffer) Short2Bytes(data uint16) []byte{

	bs := make([]byte, 2)

	binary.BigEndian.PutUint16(bs, data)

	return bs
}

func (obj *Buffer) Int2Bytes(data uint32) []byte{

	bs := make([]byte, 4)

	binary.BigEndian.PutUint32(bs, data)

	return bs

}

func (obj *Buffer) Long2Bytes(data uint64) []byte{

	bs := make([]byte, 8)

	binary.BigEndian.PutUint64(bs, data)

	return bs

}

func Bytes2Float(bytes []byte) float32 {
    bits := binary.BigEndian.Uint32(bytes)
    float := math.Float32frombits(bits)
    return float
}

func Float2Bytes(float float32) []byte {
    bits := math.Float32bits(float)
    bytes := make([]byte, 4)
    binary.BigEndian.PutUint32(bytes, bits)
    return bytes
}

func Bytes2Double(bytes []byte) float64 {
    bits := binary.BigEndian.Uint64(bytes)
    float := math.Float64frombits(bits)
    return float
}

func Double2Bytes(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.BigEndian.PutUint64(bytes, bits)
    return bytes
}