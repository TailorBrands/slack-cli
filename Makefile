BINARY=slack-cli
SRC_DIR=$(shell pwd)

build:
	docker run --rm \
	  -e LDFLAGS='-extldflags "-static"' \
	  -e COMPRESS_BINARY=true \
	  -v ${SRC_DIR}:/src \
	  centurylink/golang-builder
	docker build -t ${BINARY} .

install:
	go install

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
