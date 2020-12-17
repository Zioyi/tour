package csv_generator

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvGenerator struct {
}

func (c *CsvGenerator) Open(fileName string) [][]string {
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
		os.Exit(0)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
		os.Exit(0)
	}

	return records
}

func (c *CsvGenerator) Generate(records [][]string, fileName string, overWrite bool) {
	if !overWrite && exists(fileName) {
		fmt.Println(1123)
		log.Fatalf("the file is existed, can not overwrite")
		os.Exit(0)
	}

	fs, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("can not create the file, err is %+v", err)
		os.Exit(0)
	}
	defer fs.Close()

	w := csv.NewWriter(fs)
	err = w.WriteAll(records)
	if err != nil {
		log.Fatalf("can not writeall, err is %+v", err)
		os.Exit(0)
	}
}

// 判断文件或文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
