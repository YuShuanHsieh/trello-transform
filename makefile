all:
	make ui && make core && make package

ui: 
	mkdir -p release/web && \
	cd app/ && npm run build && cp -r build/* ../release/web

core:
	go build -o core main.go && \
	cp core release/

package:
	cd release/ && tar -czvf app.tgz web/ core