setup:

	@echo "Running setup" \

	brew update
	brew install pyenv
	pyenv install ${cat .python-version}
	brew install go


test-py:

	@echo "Running Python tests" \

	python -m pytest

test-go:

	@echo "Running Go tests" \
	
	go test ./tests/...

test:

	@echo "Running all tests" \

	make test-py && make test-go
