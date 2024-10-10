all: 
	mkdir -p output
	go build -o output/calendar cmd/main.go
	cp scripts/gridview.sh output/gridview.sh
	cp -r icon output/icon

run: all
	./output/calendar

clean:
	rm -rf output