# go-lib-dev

go-lib-dev用来开发go依赖库，
每个git子模块都是一个独立的go依赖库。
**请不要在代码中直接依赖本库！请按需依赖子库！**

- [go-dispatcher](https://github.com/xmh19936688/go-dispatcher)用于将需要处理的数据分发到协程执行，可控制最大并发量
- [go-requester](https://github.com/xmh19936688/go-requester)用于发送http请求，避免写重复代码
- [go-number](https://github.com/xmh19936688/go-number)用于在API中兼容字符串和数值类型的数字
- [go-seeker](https://github.com/xmh19936688/go-seeker)用于逐行处理字符串
