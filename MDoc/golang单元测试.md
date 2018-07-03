### golang单元测试


#### 使用总结

```
go test xxx_test.go
go test -v xxx_test.go  # -v是显示出详细的测试结果

```


#### 单元测试

测试源码文件是名称以 `“_test.go”` 为后缀的

测试函数一般是以“Test”为名称前缀并有一个类型为“testing.T”的参数声明的函数

函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。

执行：

```
go test envir_test.go
```


测试整个utils包时：

```
# -v是显示出详细的测试结果, -cover 显示出执行的测试用例的测试覆盖率。

go test -v 

go test -cover=true 
```

执行单个测试用例：

```
#./src/utils为包utils的路径
go test -v -cover=true ./src/utils -run TestSuccessStringInSlice
```

#### 基准测试

基准测试的函数必须以Benchmark开头
基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
基准测试函数不能有返回值

```
go test -bench=. -benchtime=3s -run=none
```


