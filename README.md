# lockie
A dead simple locking mechanism

## Usage
```go
package main 

import "fmt"

type DB struct {
	mux *Lockie
	m   map[string]string
}

func (db *DB) Get(k string) (v int) {
	db.mux.Lock()
	v = t.m[k]
	db.mux.Unlock()
	return v
}

func (db *DB) Put(k string, v int) {
	db.mux.Lock()
	db.m[k] = v
	db.mux.Unlock()
}

func main(){
	db := DB{
		mux: NewLockie(),
		m:   make(map[string]string),
	}
	
	db.Put("name", "John Doe")
	fmt.Println(db.Get("name"))
}
```

## Benchmarks
```markdown
# Command
go test --bench=. -cpu=1,2,4 

# Results
BenchmarkMuxR           50000000		37.1 ns/op		0 B/op		0 allocs/op
BenchmarkMuxR-2         20000000		112 ns/op		0 B/op		0 allocs/op
BenchmarkMuxR-4         10000000		133 ns/op		0 B/op		0 allocs/op

BenchmarkMuxW           30000000		58.2 ns/op		0 B/op		0 allocs/op
BenchmarkMuxW-2         10000000		197 ns/op		0 B/op		0 allocs/op
BenchmarkMuxW-4         10000000		173 ns/op		0 B/op		0 allocs/op

BenchmarkMuxRW          20000000		94.8 ns/op		0 B/op		0 allocs/op
BenchmarkMuxRW-2         5000000		266 ns/op		0 B/op		0 allocs/op
BenchmarkMuxRW-4         5000000		320 ns/op		0 B/op		0 allocs/op



BenchmarkRWMuxR         30000000		58.3 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxR-2       10000000		191 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxR-4       10000000		186 ns/op		0 B/op		0 allocs/op

BenchmarkRWMuxW         20000000		79.0 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxW-2        5000000		318 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxW-4       10000000		208 ns/op		0 B/op		0 allocs/op

BenchmarkRWMuxRW        10000000		137 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxRW-2       3000000		513 ns/op		0 B/op		0 allocs/op
BenchmarkRWMuxRW-4       5000000		400 ns/op		0 B/op		0 allocs/op



BenchmarkLockieR        50000000		25.5 ns/op		0 B/op		0 allocs/op
BenchmarkLockieR-2      50000000		35.8 ns/op		0 B/op		0 allocs/op
BenchmarkLockieR-4      30000000		51.0 ns/op		0 B/op		0 allocs/op

BenchmarkLockieW        30000000		46.4 ns/op		0 B/op		0 allocs/op
BenchmarkLockieW-2      20000000		72.0 ns/op		0 B/op		0 allocs/op
BenchmarkLockieW-4      20000000		114 ns/op		0 B/op		0 allocs/op

BenchmarkLockieRW       20000000		68.8 ns/op		0 B/op		0 allocs/op
BenchmarkLockieRW-2     20000000		111 ns/op		0 B/op		0 allocs/op
BenchmarkLockieRW-4     10000000		176 ns/op		0 B/op		0 allocs/op
```