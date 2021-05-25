default:
	./test-wc.sh
mutli:
	go run wc.go master localhost:7777 pg-*.txt
clean:
	rm -rf mrtmp.* diff.out *.log
.PHONY: clean mutli