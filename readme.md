# cjungo demo

cjungo 的使用示例。

## 加入 cjungo 的开发

使用 git 克隆 [cjungo 项目](https://github.com/cjungo/cjungo)，和 [demo 项目](https://github.com/cjungo/demo)
使用 go work 指定本地 cjungo 和 demo 的源码，得到如下目录结构

```bash
go work init
go work use ./cjungo
go work use ./demo
```

```
- cjungo/
    - go.mod
    - ...
- demo/
    - go.mod
    - ...
- go.work
- go.work.sum
```