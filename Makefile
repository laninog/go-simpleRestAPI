BINARY=go_simple_rest_api

build: clean
	go build -o ${BINARY}

install: clean
	go build -o ${BINARY}

run: build
	./${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
