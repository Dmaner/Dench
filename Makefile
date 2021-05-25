.PHONY:default clean import
default:
	go run main.go
import:
	sh importdb/ArangodbDataImport.sh dman test mydb
clean:
	rm -rf log/* data/*
