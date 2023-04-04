BENCHTIME?=60s

all: clean bench bench-pgo bench-compare

clean:
	echo "===> cleaning up generated files"
	rm -f bench-* functional-*.test inline-*.txt profile-*.out

bench:
	echo "===> start benchmark"
	go test -cpuprofile=profile-regular.out -cpu=1,2,4,6 -benchtime=$(BENCHTIME) -bench . | tee bench-original.txt
	mv functional.test functional-regular.test

bench-pgo:
	echo "===> start benchmark using PGO"
	go test -pgo=profile-regular.out -cpuprofile=profile-pgo.out -benchtime=$(BENCHTIME) -cpu=1,2,4,6 -bench=. | tee bench-pgo.txt
	mv functional.test functional-pgo.test

bench-compare: dev-tools
	echo "===> comparing benchmarks"
	benchstat bench-original.txt bench-pgo.txt

pprof:
	go tool pprof -http=:9001 functional-regular.test profile-regular.out;

pprof-pgo:
	go tool pprof -http=:9002 functional-pgo.test profile-pgo.out;
	
dev-tools:
	go install golang.org/x/perf/cmd/benchstat@latest