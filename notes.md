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