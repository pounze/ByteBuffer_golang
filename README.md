# ByteBuffer_golang
Golang ByteBuffer Library

This library helps you to work with bytes packets and create buffer array very easily in golang.
Here you can put, get bytes with its data types.

There are several methods to do it. Followign are the examples of ByteBufferPacket

To import package you can do the following

import (
	"fmt"
	"ByteBuffer"
)

Now we will run the method one by one

To Initialize the ByteBuffer packet we do the following.

var byteBuffer = ByteBuffer.Buffer{}

To put short value that is uint16 in golang

byteBuffer.PutShort(20)

To put int value that is uint32 in golang

byteBuffer.PutInt(20)

To put long value that is uint64 in golang

byteBuffer.PutLong(20)

# To put float value that is float32 in golang

byteBuffer.PutFloat(20.0)

# To put double value that is float64 in golang

byteBuffer.PutDouble(20.0)

To put String value

var sampleData = []byte("Hello World")
byteBuffer.Put(sampleData)

To convert ByteBuffer to byte array

var res = byteBuffer.Array()

To convert byte array to ByteBuffer

var byteBufferNew = ByteBuffer.Buffer{}
byteBufferNew.Wrap(res)

To fetch short value from byte buffer

var shortData = byteBufferNew.GetShort()

fmt.Println(shortData)

To fetch int value from byte buffer

var intData = byteBufferNew.GetInt()

fmt.Println(intData)

To fetch double value from byte buffer

var longData = byteBufferNew.GetLong()

fmt.Println(longData)

To fetch float value from byte buffer

var floatData = byteBufferNew.GetFloat()

fmt.Println(floatData)

To fetch double value from byte buffer

var doubleData = byteBufferNew.GetDouble()

fmt.Println(doubleData)

To fetch bytes value from byte buffer

var data = byteBufferNew.Get(len(sampleData))

fmt.Println(data)

To get size value from byte buffer

fmt.Println(byteBufferNew.Size())

To flip byte buffer for littleEndian

byteBuffer.Flip()

fmt.Println(byteBuffer.Array())

To clear byte buffer

byteBuffer.Clear()

fmt.Println(byteBuffer.Array())

To slice a bytebuffer

err := byteBuffer.Slice(2,2)

if err != nil{
  fmt.Println(err)
}

To convert byte array to string
	
fmt.Println(byteBuffer.Bytes2Str(data))

To convert string to byte array

fmt.Println(byteBuffer.Str2Bytes(byteBuffer.Bytes2Str(data)))

There are more method like Bytes2Short, Short2Bytes, Byte2Int, Int2Bytes, Bytes2Long, Long2Bytes, Float2Bytes, Bytes2Float, Double2Bytes, Bytes2Double

More method will be introduced very soon. Happy coding
