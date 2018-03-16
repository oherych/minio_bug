This is sandbox created for reproducing possible minion bug. Attached ssh key is generated special for this example and dont using in any servers.

For reproducing this bug please start main.go and open "http://localhost:8000". This page will make request to "/file" 1000 times

```
{"time":"2018-03-16T20:40:46.811144+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.053516+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.053661+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.0553+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.058796+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.058829+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:49.991704+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"http: read on closed response body"}
{"time":"2018-03-16T20:40:49.992093+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"http: read on closed response body"}
{"time":"2018-03-16T20:40:47.061494+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.478363+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:47.478403+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
{"time":"2018-03-16T20:40:50.032778+02:00","level":"ERROR","prefix":"-","file":"main.go","line":"144","message":"invalid padding size"}
panic: runtime error: slice bounds out of range

goroutine 1100 [running]:
bytes.(*Buffer).ReadFrom(0xc420125030, 0x14cefc0, 0xc420327480, 0x2660140, 0xc420125030, 0x1)
	/usr/local/Cellar/go/1.9.1/libexec/src/bytes/buffer.go:205 +0x2c9
io.copyBuffer(0x14ceb00, 0xc420125030, 0x14cefc0, 0xc420327480, 0x0, 0x0, 0x0, 0x130a660, 0x1700d01, 0xc420327480)
	/usr/local/Cellar/go/1.9.1/libexec/src/io/io.go:386 +0x2bb
io.Copy(0x14ceb00, 0xc420125030, 0x14cefc0, 0xc420327480, 0x10, 0x12dca80, 0x1)
	/usr/local/Cellar/go/1.9.1/libexec/src/io/io.go:362 +0x68
io.CopyN(0x14ceb00, 0xc420125030, 0x14cf4c0, 0xc44429cc40, 0x10, 0x10, 0x0, 0x0)
	/usr/local/Cellar/go/1.9.1/libexec/src/io/io.go:338 +0x8b
github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go/pkg/encrypt.(*CBCSecureMaterials).Read(0xc4200fc0c0, 0xc42010c000, 0x8000, 0x8000, 0x0, 0x14ceec0, 0xc4204b1310)
	/Users/oherych/src/github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go/pkg/encrypt/cbc.go:224 +0x138
io.ReadAtLeast(0x26600f0, 0xc4200fc0c0, 0xc42010c000, 0x8000, 0x8000, 0x8000, 0x131f240, 0x12a2c00, 0x26600f0)
	/usr/local/Cellar/go/1.9.1/libexec/src/io/io.go:309 +0x86
io.ReadFull(0x26600f0, 0xc4200fc0c0, 0xc42010c000, 0x8000, 0x8000, 0x0, 0x0, 0x0)
	/usr/local/Cellar/go/1.9.1/libexec/src/io/io.go:327 +0x58
github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go.Client.getObjectWithContext.func1(0xc420380ea0, 0xc420380f00, 0xc420381260, 0xc42039a100, 0xc4203fe080, 0xc420140140, 0xc42039a110, 0xc420140000, 0x14d4000, 0xc420016090, ...)
	/Users/oherych/src/github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go/api-get-object.go:198 +0xd80
created by github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go.Client.getObjectWithContext
	/Users/oherych/src/github.com/oherych/mimio_bug/vendor/github.com/minio/minio-go/api-get-object.go:71 +0x2c7
```