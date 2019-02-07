all:
	make ui && make core && make package

ui: 
	mkdir -p tmp/web && \
	cd app/ && npm run build && cp -r build/* ../tmp/web

core:
	go build -o core main.go && \
	cp core appspec.yml tmp/ && \
	mkdir -p tmp/scripts && cp -r scripts/* tmp/scripts

package:
	mkdir -p release/ && cd tmp/ && \
	tar -czvf app.tgz --exclude=./*.tgz . && \
	mv app.tgz ../release/ && \
	cd ../ && rm -r tmp/
	