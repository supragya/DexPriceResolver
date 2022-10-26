package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Load CSV
	csvName := "uni-pairs202210241405.csv"
	records, err := loadCSV(csvName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", records)
}

type Record struct {
	Pair       common.Address
	Token0     common.Address
	Token1     common.Address
	StartBlock int64
	Exchange   common.Address
}

func loadCSV(filePath string) ([]Record, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	resRecords := []Record{}

	for _, rec := range records[1:] {
		startBlock, err := strconv.Atoi(strings.Split(rec[4], "-")[0])
		if err != nil {
			panic(err)
		}
		record := Record{
			Pair:       common.HexToAddress(strings.Split(rec[0], "-")[0]),
			Token0:     common.HexToAddress(strings.Split(rec[1], "-")[0]),
			Token1:     common.HexToAddress(strings.Split(rec[2], "-")[0]),
			StartBlock: int64(startBlock),
			Exchange:   common.HexToAddress(strings.Split(rec[6], "-")[0]),
		}
		resRecords = append(resRecords, record)
	}

	return resRecords, nil
}
