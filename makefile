.PHONY: build run clean test

build:
	docker build -t job-scheduler-overlap-checker .

test:
	docker run --rm job-scheduler-overlap-checker go test ./...

run:
	docker run -p 8080:8080 job-scheduler-overlap-checker

clean:
	docker rmi job-scheduler-overlap-checker
