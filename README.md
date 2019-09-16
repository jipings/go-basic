# go-basic

## GO中的指针

1)  基本数据类型，变量存的就是值，也叫做值类型
2)  获取变量的地址，用&， 比如： var num int, 获取num的地址: &num
3)  指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
4)  获取指针类型所指向的值，使用：*，比如: var ptr *int, 使用*ptr获取ptr指向的值

* 指针的使用细节

1)  值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int，float32对应的指针类型就是 *float32
2)  值类型包括： 基本数据类型 int 系列， float 系列，bool，string、数组和结构体 struct

* 值类型和引用类型的说明
1)  值类型: 基本数据类型 int 系列， float系列，bool，string、数组和结构体
2)  引用类型: 指针、slice切片、map、管道chan、interface 等都是引用类型


## 原码、反码、补码


## 位运算符和移位运算符

## 函数的递归调用

## init 函数

## 函数 defer
1)  当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中, 然后继续执行函数下一个语句
2)  当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行(注：遵守栈 先入后出的机制)，
3)  在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。

## 内置函数
Golang 设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为 Go 的内置函 数。文档：https://studygolang.com/pkgdoc -> builtin

1) len：用来求长度，比如 string、array、slice、map、channel
2) new：用来分配内存，主要用来分配值类型，比如 int、float32, struct...返回的是指针
3) make: 用来分配内存，主要用来分配引用类型，比如 channel、map、slice

## 错误处理
Go 中引入的处理方式为：defer, panic, recover

* recover() 
    使用 defer+recover 来处理错误
* panic( error )
    打印出error并终止程序    
* 错误处理的好处
    进行错误处理后，程序不会轻易挂掉，如果加入预警代码，就可以让程序更加的健壮。

## goroutine
1) 进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的一个单位

2) 线程是进程的一个执行实例，是程序执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的基本单元

3) 一个进程可以创建和销毁多个线程，同一个程序中的多个线程可以并发执行

4) 一个程序至少有一个进程，一个进程至少有一个线程

## go 协程和 go 主线程

### go协程的特点

1) 有独立的栈空间
2) 共享程序堆空间
3) 调度由用户控制
4) 协程是轻量级的线程

## forloop



## 第三方包

### [fake-useragent](https://github.com/EDDYCJY/fake-useragent) 
模拟浏览器头（User-Agent）

> go get github.com/EDDYCJY/fake-useragent