# Go ABC

## 1. hello-world

在如下目录建立

```
+--gopath
|  +--bin
|  +--pkg
|  +--src
|  |  +--goabc
|  |  |  +--hello.go
+-...
```

- bin目录下放的是可执行文件，pkg目录下放的是.a或.so的库文件，src目录下放置源码

在`hello.go`中输入如下内容：

```go
package main

func main(){
    println("hello world")
}
```

在`goabc`文件夹下执行`go build`，即可得到名称为hello的可执行文件。执行`go install`，将可知行文件安装到`gopath/bin`目录下。

- 只有package name为main的包才可被编译为可执行文件，该包下的`main`函数为主函数。

## 2. import & package

### import

```
+--gopath
|  +--bin
|  +--pkg
|  +--src
|  |  +--goabc
|  |  |  +--hello.go
|  |  +--goabc-link
|  |  |  +--importme.go
+-...
```

在`importme.go`中输入如下内容：

```go
package goabc_link

func Echo() string {
	return "you have imported me"
}
```

在`hello.go`中加入如下内容

```go
package main

import (
	"goabc-link" //该包相对于gopath/src的相对路径，gopath/src/goabc-link
)

func main(){
	println(goabc_link.Echo())
}

```

-  go会在goroot/vendor/gopath下找这个包是否存在，排在前面的优先级高。

将hello.go改写为：

```go
package main

import (
	"goabc-link" //该包相对于gopath/src的相对路径，gopath/src/goabc-link
)

func main(){
     println("hello world")
}
```

再次编译，将引起编译错误

- 没有使用到的包出现在import中会引编译报错，解决办法

  1. 删掉无用的package，可使用工具`go fmt`或`go import`，自动格式化go工程或文件

  2. 将无用包改为匿名包，如下所示：

     ```go
     import (
     	_ "goabc-link" 
     )
     ```

### package

- 在同一目录下的文件隶属于同一个package，如下图中的hello.go和samepackage.go都属于package main。

  ```
  +--gopath
  |  +--bin
  |  +--pkg
  |  +--src
  |  |  +--goabc
  |  |  |  +--hello.go
  |  |  |  +--samepackage.go
  +-...
  ```

- golang中大写开头的变量，方法，函数是包外可见的，反之只有包内可见。

  ```go
  package goabc_link
  
  func echo() string {
  	return "you have imported me"
  }
  
  ```

  将goabc_link中的`Echo`方法改为`echo`，再次编译hello.go会引起编译报错。

## 3. slice & map

### map

- 简单使用

  ```go
  func testMap(){
      Roster := make(map[string]string,0) //初始化
      Roster["bob"]="01"    //插入数据
      delete(Roster, "bob") //删除数据
      Roster["bob"]="02"    //修改数据
  }
  ```

- 初始化参数对性能的影响

  map的初始化使用的是make方法申请内存，make方法第一个参数是申请内存的类型，第二个参数是所要申请的容量。

  对于map/slice而言，当使用的容量超过申请的容量时，该容量是可以自动扩容的，但是自动扩容的机制会带来一定的性能损耗。因此，在已知容量为n的情况下，应使用`make(map[{key_type}]{value_type},n)`在初始化时为map申请足够内存，避免扩容开销。

  ```go
  //未指定容量
  func Mapinit(){
  	Roster := make(map[string]string,0)
  	for i:=0;i<100;i++{
  		Roster["bob"+strconv.Itoa(i)] = strconv.Itoa(i)
  	}
  }
  func BenchmarkMapinit(t *testing.B) {
  	for i := 0; i < t.N; i++ {
  		Mapinit()
  	}
  }
  
  //指定容量
  func Mapinit2(){
  	Roster := make(map[string]string,100)
  	for i:=0;i<100;i++{
  		Roster["bob"+strconv.Itoa(i)] = strconv.Itoa(i)
  	}
  }
  func BenchmarkMapinit2(t *testing.B) {
  	for i := 0; i < t.N; i++ {
  		Mapinit2()
  	}
  }
  ```

  压测结果：

  ```
  BenchmarkMapinit-8    	  100000	     16067 ns/op
  BenchmarkMapinit2-8   	  200000	      9674 ns/op
  ```

### slice

- 简单使用

  ```go
  func testSlice() {
  	array := make([]int, 0)  //初始化
  	array = append(array, 0) //增加元素
  	printSlice(array)
  	array2 := []int{1, 2, 3}
  	array = append(array, array2...) //合并切片
  	printSlice(array)
  	array3 := array[1:3] //截取切片
  	printSlice(array3)
  	array = append(array[:1], array[2:]...) //删除元素
  	printSlice(array)
  }
  func printSlice(array []int) {
  	for i := range array {
  		print(array[i], ",")
  	}
  	println()
  }
  ```

删除元素是截取切片和合并切片的组合操作：

![slice-delete](D:\doc\Time-to-go\media\slice-delete.PNG)

需要注意的时，当使用`make([]string,n)`，n不为0时，`array[0],...,array[n]`均为`nil`。

![slice](D:\doc\Time-to-go\media\slice.PNG)



## 5. 条件和循环

### 条件

```go
func testIf(){
    a = Random
}
```

### 循环

```go
for i:=0i<100;i++{
    println(i)
}

//遍历切片
slicex := make([]int,100)
for index := range slicex{
    slicex[index] = index
}
for _,value := range slicex{
    println(v)
}

//遍历map
mapx := map[string]int{"a": 1, "b": 2, "c": 3}
for k,v := range mapx {
    println(k,v)
}
```

## 6.函数

- 函数的返回值可以定义多个

```go
func testFunc(args []string)(int, string){
    return 1,"1"
}
```

- 函数是golang的第一成员，可以作为参数，成员，变量....

```go
func main(){
    f := testFunc //将函数赋值给本地变量
    doFunc(f)//函数作为入参
}

func doFunc(f func([]string)(int, string)) string {
	_,v := f([]string{})
	return v
}
```

## 7. 下节预告

part2：从goroutine到运行时调度

part3：从interface到设计模式

part4：go tools





