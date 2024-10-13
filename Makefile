all: 
	mkdir -p output
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o output/calendar_arm64 cmd/main.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o output/calendar_amd64 cmd/main.go
	makefat output/calendar output/calendar_*
	rm -rf output/calendar_*
	cp scripts/gridview.sh output/gridview.sh
	cp scripts/note.sh output/note.sh
	cp scripts/cat.sh output/cat.sh
	cp -r icon output/icon

run: all
	./output/calendar

clean:
	rm -rf output