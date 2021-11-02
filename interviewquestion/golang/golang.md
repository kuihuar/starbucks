## channel
当未为channel分配内存时，channel就是nil channel
nil channel会阻塞对该channel的所有读、写
+ 写入nil channel
+ channel panic
  
  关闭一个已关闭的 channel 会引发 panic
  关闭一个 nil 值 channel 会引发 panic
  向一个已关闭的 channel 发送数据。

+ conetext cancel()原理

关闭当前done channel
递归调用children,并关闭


## slice
函数传切片，只要没有扩容或者缩容也没有修改指定索引的值，原来的切片不变。
