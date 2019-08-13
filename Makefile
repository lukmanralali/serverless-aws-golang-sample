.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/sample_test_function sample_test_function/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/movie_list movie_list/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_movie get_movie/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/insert_movie insert_movie/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete_movie delete_movie/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update_movie update_movie/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose