# Makefile

lint:
	npm run eslint .

test:
	npm test

list:
	npm list -g --depth=0

.PHONY: test