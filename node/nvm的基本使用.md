**我们在一些项目中需要切换node版本，需要一个node版本管理工具，于是nvm出现啦**
**name如果使用nvm 管理node版本呢？** 

### 1.列出已经安装的node版本

```bash
nvm ls    
复制代码
```

### 2.列出所有可以安装的node版本号

```
nvm ls-remote
复制代码
```

### 3.安装node

```arduino
// 安装指定版本的node
nvm install v10.4.0
// 安装最新稳定版 node
nvm install stable
复制代码
```

### 4.切换node版本

```perl
nvm use v12.13.0    
复制代码
```

### 5.查看当前node版本

```sql
nvm current
复制代码
```

### 6.指定默认的node版本

```csharp
nvm alias default v12.13.0    
```



