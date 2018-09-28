language: go

os:
	- linux

before_install:
	- make debs

after_success:
	- codecov

go:
	- 1.8.5
	- 1.9.2
	- tip

matrix:
	allow_failures:
		- go: tip

script:
	- go test -cpu=2 -race -v ./...
	- go test -cpu=2 -coverprofile=coverage.txt -covermode=atomic ./...

notifications:
	email: false