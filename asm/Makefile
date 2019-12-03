asm:
	GOOS=linux GOARCH=amd64 go tool compile -S $(FILE).go

dump:
	GOOS=linux GOARCH=amd64 go build -o tmp $(FILE).go
	go tool objdump -s main.main tmp

clean:
	rm -rf tmp
	rm -rf *.o
