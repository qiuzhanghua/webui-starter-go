package main

import (
	"fmt"
	"github.com/webui-dev/go-webui/v2"
	"os"
)

func main() {
	myWindow := webui.NewWindow()
	err := myWindow.Show("<html><script src=\"webui.js\"></script> Goto <a href='https://cn.bing.com/'>bing</a> </html>")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	// or
	// myWindow.Navigate("https://cn.bing.com/")
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
