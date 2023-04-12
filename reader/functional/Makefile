BENCHTIME?=60s
BENCHFLAGS=-benchtime=30s -cpu=2,4,8 # -benchtime=$(BENCHTIME)

all: clean bench bench-pgo bench-compare

clean:
	echo "===> cleaning up generated files"
	rm -f bench-* functional-*.test inline-*.txt profile-*.out

bench:
	echo "===> start benchmark"
	go test -cpuprofile=profile-regular.out $(BENCHFLAGS) -bench . | tee bench-original.txt
	mv functional.test functional-regular.test

bench-pgo:
	echo "===> start benchmark using PGO"
	go test -pgo=profile-regular.out -cpuprofile=profile-pgo.out $(BENCHFLAGS) -bench=. | tee bench-pgo.txt
	mv functional.test functional-pgo.test

bench-compare: dev-tools
	echo "===> comparing benchmarks"
	benchstat bench-original.txt bench-pgo.txt

pprof: bench
	go tool pprof -http=:9001 functional-regular.test profile-regular.out;

pprof-pgo: bench-pgo
	go tool pprof -http=:9002 functional-pgo.test profile-pgo.out
	
dev-tools:
	go install golang.org/x/perf/cmd/benchstat@latest