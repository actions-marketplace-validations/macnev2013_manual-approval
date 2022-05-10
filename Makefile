IMAGE_REPO=ghcr.io/macnev2013/manual-approval

.PHONY: build
build:
	@if [ -z "$$VERSION" ]; then \
		echo "VERSION is required"; \
		exit 1; \
	fi
	docker build -t $(IMAGE_REPO):$$VERSION .
	docker tag $(IMAGE_REPO):$$VERSION $(IMAGE_REPO):latest

.PHONY: push
push:
	@if [ -z "$$VERSION" ]; then \
		echo "VERSION is required"; \
		exit 1; \
	fi
	docker push $(IMAGE_REPO):$$VERSION
	docker push $(IMAGE_REPO):latest

.PHONY: test
test:
	go test -v .
