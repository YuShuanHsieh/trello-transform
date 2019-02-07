all:
	make ui && make core && make package

ui: 
	mkdir -p ui && \
	cd app/ && npm run build && cp -r build/* ../ui

core:
	go build -o core main.go

package:
	tar -czvf app.tgz ui/ core