all:
	make ui && make core && make package

ui: 
	mkdir -p ui && \
	cd app/ && npm build && cp -r build/* ../ui

core:
	go build -o core main.go

package:
	tar -czvf app.tar.gz ui/ core