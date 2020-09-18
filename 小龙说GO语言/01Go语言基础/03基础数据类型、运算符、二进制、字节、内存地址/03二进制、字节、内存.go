package main

import (
	"fmt"
	"unsafe"
)

func main ()  {
	//计算机是4GB，计算机内存地址是以字节为单位
	//字节又是以二进制为单位，实际上就是以0和1为单位
	x := 1
	y := 2
	z := 3
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(&z)
	fmt.Println(unsafe.Sizeof(x))

	//这就是计算机的内存地址
	//说明我电脑时64位，并且我的电脑上的int类型是8个字节
	//0xc000016090
	//0xc000016098
	//0xc0000160a0

	//0x0000000000
	//0x0000000001
	//0x0000000002
	//0x0000000003
	//0x0000000004
	//0x0000000005
	//............
	//0xc000016060 1
	//0xc000016061 //二进制 0000000
	//0xc000016062 //二进制 0000000
	//0xc000016063 //二进制 0000000
	//0xc000016064 //二进制 0000000
	//0xc000016065 //二进制 0000000
	//0xc000016066 //二进制 0000000
	//0xc000016067 //二进制 0000001
	//0xc000016068 2
	//0xc000016070 3
}	
// package main

// import (

// 	"fmt"

// 	"unsafe"

// )



// func main ()  {

// 	//计算机是4GB，计算机内存地址是以字节为单位
//         //计算机地址的表示是16进制的
//         //起始地址是：0x0000000000
//         //终止地址是：0xFFFFFFFFFF

// 	//字节又是以二进制为单位，实际上就是以0和1为单位(bit)
//         //一个字节是8个bit，

// 	x := 1

// 	y := 2

// 	z := 3

// 	fmt.Println(&x)

// 	fmt.Println(&y)

// 	fmt.Println(&z)

// 	fmt.Println(unsafe.Sizeof(x)) //看看我的电脑是几个字节



// 	//这就是计算机的内存地址

// 	//说明我电脑是64位，并且我的电脑上的int类型是8个字节

// 	//0xc000016090

// 	//0xc000016098

// 	//0xc0000160a0



// 	//0x0000000000  //计算机起始内存地址，每次加一个就是一个字节

// 	//0x0000000001

// 	//0x0000000002

// 	//0x0000000003

// 	//0x0000000004

// 	//0x0000000005

// 	//............

// 	//0xc000016060 //0000000 //这个就是指针，就是地址块的起始地址，你的x实际上计算机根本不知道x，只会将x映射成地址：0xc000016060，这是一个字节，这个字节的结构是8个bit，结构是（二进制表示）：00000000，但是1又没有这么大，所以填充0

// 	//0xc000016061 //二进制 0000000

// 	//0xc000016062 //二进制 0000000

// 	//0xc000016063 //二进制 0000000

// 	//0xc000016064 //二进制 0000000

// 	//0xc000016065 //二进制 0000000

// 	//0xc000016066 //二进制 0000000

// 	//0xc000016067 //二进制 0000001

// 	//0xc000016068 2 //同上

// 	//0xc000016070 3 //同上

// }
//     //上面的例子是GO语言的，其实java语言、js、python也一样，可以这样理解，你可以用上面的例子尝试去理解java中的二进制、字节、内存地址、指针

//     //java中面向对象，说白了就是对象的名字映射成了一个内存地址，只要计算机知道了这个内存地址，那么就可以找到这个内存地址中存储的值，当然内存地址肯定存在“套娃”的问题，就是这个内存地址存了另外一个对象（实际上就是存了另外一个对象的内存地址），然后另外一个对象的内存地址又存了其它对象的（实际上就是其它对象的内存地址），有了内存地址，就可以为所欲为了。当然计算机的内存地址远不止这些知识，但是我觉得只要你会了我说的这些，其实就差不多了。
//     //总结一下，计算机其实只认识0和1，是根据内存地址找值，计算机内存地址是以字节为单位，每个字节又以bit为单位，大小是8个bit，其它什么变量的名字，什么=号，什么其它的东西，计算机统统不认识，只是因为每个语言有编译器和虚拟机，翻译这些人能看懂的字符成二进制。

