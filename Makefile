# Keep test at the top so that it is default when `make` is called.
# This is used by Travis CI.
coverage.txt:
	cd ./http/mux; make coverage.txt; cd ../../.
	mv ./http/mux/coverage.txt .
view-cover: clean coverage.txt
	go tool cover -html=coverage.txt
test: build
	cd ./http/mux; make test; cd ../../.
build:
	cd ./http/mux; make build; cd ../../.
inspect: build
	cd ./http/mux; make inspect; cd ../../.
pre-commit: clean coverage.txt inspect
	cd ./http/mux; make pre-commit; cd ../../.
clean:
	cd ./http/mux; make clean; cd ../../.
	rm -f coverage.txt