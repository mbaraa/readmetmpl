,PHONY: tailwindcss build

BINARY_NAME=readmetmpl

tailwindcss:
	cd tailwindcss && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css -m && \
	cd ..

run_tailwindcss_watcher:
	cd tailwindcss && \
	npx tailwindcss -i ../resources/css/style.css -o ../resources/css/tailwind.css --watch  \

build:
	go mod tidy && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

# install inotify-tools
dev:
	export `xargs < .env` && \
	while true; do \
	  go build -o ${BINARY_NAME}; \
	  ./${BINARY_NAME} & \
	  PID=$$!; \
	  echo "PID=$$PID"; \
	  inotifywait -r -e modify ./**/*; \
	  kill $$PID; \
	done

clean:
	go clean
