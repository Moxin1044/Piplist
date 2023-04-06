# 关于
这是一个使用Go语言编写的一个IP列表工具，可以生成一个IP的字典，并且返回为一个txt文件，主要供由`Moxin1044`编写的Pscan（ https://github.com/Moxin1044/Pscan ）等相关的P-tools使用。

# 使用说明
## 1.Windows：
```
Piplist.exe -IP <ip Interval> -O <out file>
```
## Linux 或 Mac：
```
chmod +x Piplist
Piplist -IP <ip Interval> -O <out file>
```
**两者在使用方法上其实并异样，只是Windows系统中是有exe后缀的！**

## 例子
我需要生成192.168.1.1-192.168.1.255的IP段：
```
Piplist -IP 129.168.1.1/24
```
- **如果你是Windows，请加上.exe**
- 如果不指定-O，则默认保存到此目录下的ip.txt
- 如果指定-O，如下：
```
Piplist -IP 129.168.1.1/24 -O 123.txt
```
当然你也可以使用另一种表达方式：
```
Piplist -IP 129.168.1.1-192.168.1.255
```
- 我们还支持跨越扫描：
```
Piplist -IP 129.168.1.1-192.168.2.255
```
- **如果你使用的是上面的表达方式，请注意：**后面的IP不要小于前面的IP。


# 其他
## 开源协议

本项目遵循MIT协议，项目被允许修改和共享，且允许商业使用，但需要保留LICENSE和相关版权。
如果你有更好的建议，倒不如提交一份issues：https://github.com/Moxin1044/Piplist/issues