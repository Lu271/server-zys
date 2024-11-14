DEFAULT := ./cmd
OUTPUT := ./bin/server-zys

default: build

build: clean
	go build -gcflags "-N -l" -o ${OUTPUT} ${DEFAULT}

clean:
	rm -f ${OUTPUT}