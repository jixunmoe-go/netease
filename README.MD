# 网易云音乐 API

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![License: MIT](https://img.shields.io/github/go-mod/go-version/JixunMoe/netease-api-go)

一个很简单的 POC package。

## 使用

直接引入即可 (应该)；如果需要手动安装，执行 `go get -u github.com/JixunMoe/netease-api-go`。

```
package main

import (
	"fmt"
	NetEaseAPI "github.com/JixunMoe/netease-api-go/pkg/api"
)

func main() {
	api := NetEaseAPI.New()
	resp, err := api.Song(37460604)
	if err == nil && resp != nil {
		song := resp.Songs[0]
		fmt.Printf("Song: %s (album: %s), singer: %s\n", song.Name, song.Album.Name, song.Artists[0].Name)
	} else {
		fmt.Printf("Could not fetch song info: %s", err)
	}
}
```

## 致谢

本项目参考了以下仓库里的代码与内容：

* [surmon-china/simple-netease-cloud-music](https://github.com/surmon-china/simple-netease-cloud-music)
* [FakerGit/go-tools](https://github.com/FakerGit/go-tools)
