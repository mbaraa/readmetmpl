.PHONY: build

BINARY_NAME=readmetmpl

build:
	cd tailwindcss && \
	npm i && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css -m && \
	cd .. && \
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}


# install inotify-tools
dev:
	cd tailwindcss && \
	npm i && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css --watch & \
	while true; do \
	 go build -o ${BINARY_NAME}; \
	 ./${BINARY_NAME} & \
	 PID=$$!; \
	 echo "PID=$$PID"; \
	 inotifywait -r -e modify ./**; \
	 kill $$PID; \
	done

clean:
	go clean
