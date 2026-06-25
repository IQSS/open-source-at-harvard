download:
	docker build --target download -t osah-download .
	docker run --rm -v "$(PWD):/data" osah-download

parse:
	docker build --target parse -t osah-parse .
	docker run --rm -v "$(PWD):/data" osah-parse
