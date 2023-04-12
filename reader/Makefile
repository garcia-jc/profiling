LENS_FILTER=github.com/garcia-jc
BUILD_FLAGS=-gcflags="-m -m" -o=reader

build-original:
	@go build $(BUILD_FLAGS) ./functional/... 2> inline-original.log
	@open inline-original.log

build-pgo:
	@go build -pgo=functional/profile-regular.out $(BUILD_FLAGS) ./functional/... 2> inline-pgo.log
	@open inline-pgo.log

clean:
	@rm inline-*.* reader