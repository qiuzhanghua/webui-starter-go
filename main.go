package main

import (
	"fmt"
	"github.com/webui-dev/go-webui/v2"
	"os"
)

func main() {
	myWindow := webui.NewWindow()
	err := myWindow.Show("<html><script src=\"webui.js\"></script> Hello World from Go! </html>")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	webui.Wait()
}

//package main
//
///*
//#include <stdio.h>
//int SayHello() {
// puts("Hello World");
//    return 0;
//}
//*/
//import "C"
//import "fmt"
//
//func main() {
//	//ret := C.SayHello()
//	//fmt.Println(ret)
//	fmt.Println("Hello World")
//}
