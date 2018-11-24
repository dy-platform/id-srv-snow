# This is how we want to name the binary output
OUTPUT=id-srv-snowflake

# These are the values we want to pass for Version and BuildTime
GITTAG=`git describe --tags`
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
# LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"
LDFLAGS=-ldflags "-X main.BuildTime=${BUILD_TIME}"

exclude_dirs := Makefile

dirs := $(shell ls -l | grep '^d' |awk '{print $$9}')


.PHONY:all clean proto release
all:clean proto release

clean:
	rm -f ${OUTPUT}

proto:
	make -C idl

release:
	go build ${LDFLAGS} -o ${OUTPUT} main.go


