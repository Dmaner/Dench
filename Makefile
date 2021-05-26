.PHONY:default clean import
default:
	go run main.go
import:
	sh database/ArangodbDataImport.sh dman test mydb
clean:
	rm -rf log/* data/*
