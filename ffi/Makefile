all: install-dependence build-shard-lib benchmark

.PHONY: install-dependence
install-dependence:
	bundle install --path=./vendor/bundle

.PHONY: build-shard-lib
build-shard-lib:
	go build -buildmode=c-shared -o fib.so fib.go

.PHONY: benchmark
benchmark:
	echo "---go run fib.go---" && \
		time go run fib.go && \
		echo "---bundle exec ruby fib.rb---" && \
		time bundle exec ruby fib.rb && \
		echo "---bundle exec ruby go-fib.rb---" && \
		time bundle exec ruby go-fib.rb
