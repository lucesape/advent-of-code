setup:

	@echo "Running setup" \

	brew update
	brew install pyenv
	pyenv install ${cat .python-version}
	brew install go


test-py:

	@echo "Running Python tests" \

	python -m unittest advent_of_code/**/*_test.py

test-go:

	@echo "Running Go tests" \
	
	go test ./...

test:

	@echo "Running all tests" \

	make test-py && make test-go
