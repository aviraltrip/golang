memory management
garbage collection happens automatically
out of score or nil 

new() - allocate memory but no init
get memory addr
zeroed storage

make() - allocate mem and init, get mem addr, non-zeroed storage

go list -m -versions github.com/gorilla/mux
go mod tidy
go mod -edit -module (version)
go run -mod=vendor main.go  (pahle vendor folder me dekhega fir maal uthaega udhar se)
go mod vendor - copies all req 3rd party dependencies into vendor/ folder & records their versions in vendor/modules.txt

mod ke ops are expensive btw

gorilla mux : URL router & request dispatcher (matches incoming HTTP requests to specific handler functions)

go get -u github.com/gorilla/mux

Concurrency: handling multiple tasks at the same time by switching between them (using goroutines)

Parallelism: running multiple tasks simultaneously on different CPU cores.

goroutines - do not communicate by sharing memory, instead share memory by communicating 