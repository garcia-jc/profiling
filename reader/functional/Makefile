create-profile:
	go test -cpuprofile=profile-regular.out -benchtime=180s -bench .

pprof:
	go tool pprof -http=:9001 functional.test profile-regular.out;

pprof-pgo:
	mv profile.out profile-pgo.out;
	go tool pprof -http=:9002 functional.test profile-pgo.out;

inline:
	go build -gcflags="-m" . 2> inline-original.txt

inline-pgo:
	go build -pgo=profile-regular.out -gcflags="-m" . 2> inline-pgo.txt

bench:
	go test -bench=. -benchtime=180s > bench-original.txt

bench-pgo:
	go test -pgo=profile-regular.out -bench=. -benchtime=180s > bench-pgo.txt

bench-compare:
	benchstat bench-original.txt bench-pgo.txt