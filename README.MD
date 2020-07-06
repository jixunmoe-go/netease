# 网易云音乐 API

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go Version](https://img.shields.io/github/go-mod/go-version/jixunmoe-go/netease)
[![GoDoc](https://godoc.org/github.com/jixunmoe-go/netease?status.svg)](https://godoc.org/github.com/jixunmoe-go/netease)

一个很简单的 POC package。

## 使用

直接引入即可 (应该)；如果需要手动安装，执行 `go get -u github.com/jixunmoe-go/netease`。

```
package main

import (
	"fmt"
	"github.com/jixunmoe-go/netease"
)

func main() {
	api := netease.New()
	resp, err := api.Song(37460604)
	if err == nil && resp != nil {
		song := resp.Songs[0]
		fmt.Printf("Song: %s (album: %s), singer: %s\n", song.Name, song.Album.Name, song.Artists[0].Name)
	} else {
		fmt.Printf("Could not fetch song info: %s", err)
	}
}
```

[Go Playground](https://play.golang.org/p/HZxVi1G01OD)

## 致谢

本项目参考了以下仓库里的代码与内容：

* [FakerGit/go-tools](https://github.com/FakerGit/go-tools)
* [nondanee/UnblockNeteaseMusic](https://github.com/nondanee/UnblockNeteaseMusic)
* [picone/MusicUnionSearch](https://github.com/picone/MusicUnionSearch)
* [sqaiyan/netmusic-node](https://github.com/sqaiyan/netmusic-node)
* [surmon-china/simple-netease-cloud-music](https://github.com/surmon-china/simple-netease-cloud-music)
