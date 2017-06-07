1. net/http/pprofをインポート

```
import _ "net/http/pprof"
```

↑これを追記するだけで良い。

2. webアプリを起動

```
$ go run main.go // => localhost:8080でサーバー起動
```

3. 負荷ツールを使ったりして高い負荷がかかった状態を作る

4. go tool pprofを起動

```
// メモリ
go tool pprof http://localhost:8080/debug/pprof/heap

// CPU
go tool pprof http://localhost:8080/debug/pprof/goroutine
```

5. 対話型インタプリタが起動したら、「tree」「list」「top」などで負荷の原因を調べる

```
(pprof) list github.com/oyakata/std-go/prof/memeater.

Total: 40MB
ROUTINE ======================== github.com/oyakata/std-go/prof/memeater.Eat in /home/echizen/_go/src/github.com/oyakata/std-go/prof/memeater/memeater.go
         0       32MB (flat, cum) 80.00% of Total
         .          .     10:    data = strings.Repeat("0", 8*1024*1024)
         .          .     11:)
         .          .     12:
         .          .     13:func Eat() int {
         .          .     14:    buf.WriteString(data)
         .       32MB     15:    return buf.Len()
         .          .     16:}
ROUTINE ======================== github.com/oyakata/std-go/prof/memeater.init in /home/echizen/_go/src/github.com/oyakata/std-go/prof/memeater/memeater.go
         0        8MB (flat, cum) 20.00% of Total
         .          .      5:    "strings"
         .          .      6:)
         .          .      7:
         .          .      8:var (
         .          .      9:    buf  = bytes.NewBufferString("")
         .        8MB     10:    data = strings.Repeat("0", 8*1024*1024)
         .          .     11:)
         .          .     12:
         .          .     13:func Eat() int {
         .          .     14:    buf.WriteString(data)
         .          .     15:    return buf.Len()
```
