package main

/*
#include <stdio.h>
#include <malloc.h>
#include <string.h>

char** getStringArray(){
	char** list = malloc(sizeof(char*));
	list[0] = "xing";
	list[1] = "jack";
	list[2] = "john";
	return list;
}

void showStringArray(char** list){
	int num = strlen(*list)-1;
	int i = 0;
	for(i=0;i<num;i++){
		printf("%s\n",list[i]);
	}
}

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//char** 转化成 []string
	list := C.getStringArray()
	var str []string
	var pbuf []*C.char
	header := (*reflect.SliceHeader)(unsafe.Pointer(&pbuf))
	header.Cap = 3
	header.Len = 3
	header.Data = uintptr(unsafe.Pointer(list))
	for _, i := range pbuf {
		str = append(str, C.GoString(i))
	}
	fmt.Println(str)

	//[]string 转化成 char**
	box := []string{"xing", "jack", "john"}
	var buf []*C.char
	for i := range box {
		buf = append(buf, (*C.char)(unsafe.Pointer(C.CString(box[i]))))
	}
	box2 := (**C.char)(unsafe.Pointer(&buf[0]))
	C.showStringArray(box2)

}
