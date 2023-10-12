/*1.数组中元素数量固定，不能增减；而切片则是动态的，可以进行增减，数量不固定
2.在进行拷贝时，数组进行整体拷贝，而切片则是被拷贝内存地址
3.切片的返回值有三个，分别是cap（容量），len（长度），array（数组地址）
而数组则只有len和array这两个值*/

/*
切片创建
1.常规方法，var 变量名 []+切片数据类型 以及先声明后赋值等等通用方法

2.使用make进行创建 即 var 变量名：=make([]+切片数据类型，len（长度），cap（容量）)

3.先创建一个指针 利用指针创建切片
I.var 变量名 *[]+切片数据类型
II.var 变量名 =new([]切片数据类型)
区别是I创建的切片初始值为nil，而II创建的切片的初始值为对应数据类型的初始值

map创建
1.常规方法，不过多赘述

2.使用make进行创建
变量名:=make(map[键类型][值类型])

3.使用new创建指针的同时初始化一个map（基本同上，不过多赘述）

 */

