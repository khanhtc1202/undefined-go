OUT:=tmp

build:
	GOOS=linux GOARCH=amd64 go build -o $(OUT) $(FILE).go

asm:
	GOOS=linux GOARCH=amd64 go tool compile -S $(FILE).go

dump:
	GOOS=linux GOARCH=amd64 go build -o $(OUT) $(FILE).go
	go tool objdump -s main.main $(OUT) 

clean:
	rm -rf $(OUT)
	rm -rf *.o
