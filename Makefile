generate-mock:
	go install github.com/vektra/mockery/v2@v2.53.3
	mockery --dir src/core/ --all --output src/core/_mocks
	mockery --dir src/repository/ --all --output src/core/_mocks
