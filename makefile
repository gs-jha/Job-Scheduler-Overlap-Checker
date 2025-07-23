.PHONY: build run clean test

build:
	docker build -t job-scheduler-overlap-checker .

test:
	docker build --target builder --tag job-scheduler-overlap-checker-build .
	docker run --rm job-scheduler-overlap-checker-build go test ./...

coverage:
	docker build --target builder --tag job-scheduler-overlap-checker-build .
	docker run --rm -v $(PWD):/app job-scheduler-overlap-checker-build /bin/sh -c "go test -coverprofile=coverage.out ./..."
	docker run --rm -v $(PWD):/app job-scheduler-overlap-checker-build go tool cover -html=coverage.out -o coverage.html

run:
	docker run -p 8080:8080 job-scheduler-overlap-checker

clean:
	docker rmi job-scheduler-overlap-checker
