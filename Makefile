run: main.out
	./$^

main.out: ./main.go
	go build -v -i -o $@ $^
