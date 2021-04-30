# panic 的错误日志输出到了哪里

## 1. 输出到了哪里

`panic()` 函数的具体实现是调用了 `gopanic() ` 函数，其部分实现如图所示：

![](https://raw.githubusercontent.com/xmge/image/main/gonote/panic.png)

可以看到 `gopanic()` 调用内置函数 `print()` 进行日志的输出，通过看 `print()` 方法，其把日志的输出到了标准错误中。

```go
// The print built-in function formats its arguments in an
// implementation-specific way and writes the result to standard error.
// Print is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func print(args ...Type)
```

## 2. 如何让日志输出到自己的日志文件中呢

```go
var stdErrFileHandler *os.File

func RewriteStderrFile() error {
      if runtime.GOOS == "windows" {
            return nil
      }

    file, err := os.OpenFile(stdErrFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
            fmt.Println(err)
        return err
    }
    stdErrFileHandler = file //把文件句柄保存到全局变量，避免被GC回收

    if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
        fmt.Println(err)
        return err
    }
    // 内存回收前关闭文件描述符
      runtime.SetFinalizer(stdErrFileHandler, func(fd *os.File) {
            fd.Close()
      })

    return nil
}

```

参考文档：https://zhuanlan.zhihu.com/p/245369778