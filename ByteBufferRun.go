package main

import (
	"fmt"
	"ByteBuffer"
)

func main(){
	
	var byteBuffer = ByteBuffer.Buffer{}

	byteBuffer.PutShort(20)

	byteBuffer.PutInt(20)

	byteBuffer.PutLong(20)

	byteBuffer.PutFloat(20.0)

	byteBuffer.PutDouble(20.0)

	var sampleData = []byte("Sudeep Dasgupta")

	byteBuffer.Put(sampleData)

	var res = byteBuffer.Array()
	
	fmt.Println(res)

	var byteBuffer1 = ByteBuffer.Buffer{}

	byteBuffer1.Wrap(res)

	var shortData = byteBuffer1.GetShort()

	fmt.Println(shortData)

	var intData = byteBuffer1.GetInt()

	fmt.Println(intData)

	var longData = byteBuffer1.GetLong()

	fmt.Println(longData)

	var floatData = byteBuffer1.GetFloat()

	fmt.Println(floatData)

	var doubleData = byteBuffer1.GetDouble()

	fmt.Println(doubleData)

	var data = byteBuffer1.Get(len(sampleData))

	fmt.Println(data)

	fmt.Println(byteBuffer1.Size())

	byteBuffer.Flip()

	fmt.Println(byteBuffer.Array())

	byteBuffer.Clear()

	fmt.Println(byteBuffer.Array())

	err := byteBuffer.Slice(2,2)

	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println(byteBuffer.Bytes2Str(data))

	fmt.Println(byteBuffer.Str2Bytes(byteBuffer.Bytes2Str(data)))

	//var sampl1Data = byteBuffer.Get(len(sampleData))

	//fmt.Println(sampl1Data)

}