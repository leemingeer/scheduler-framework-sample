
BUILDENVVAR=CGO_ENABLED=0
.PHONY: build
build:
	$(BUILDENVVAR) GOARCH=arm64 go build -ldflags '-w' -o sample-scheduler-framework main.go

.PHONY: image
image:
	docker build -t leemingeer/sample-scheduler:v1.0.2 .
	docker push leemingeer/sample-scheduler:v1.0.2

.PHONY: test
test:
	kubectl apply -f deploy/sample-scheduler.yaml
	kubectl apply -f deploy/test-scheduler.yaml