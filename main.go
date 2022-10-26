package main

import (
	"fmt"
	"sort"
)

func main() {
	// Load CSV
	csvName := "uni-pairs202210241405.csv"
	records, err := loadCSV(csvName)
	if err != nil {
		panic(err)
	}

	// sort records
	sort.Sort(records)

	// list by records
	sb, count := records[0].StartBlock, 0
	for _, rec := range records {
		if rec.StartBlock == sb {
			count++
		} else {
			fmt.Printf("sb: %v\tcount: %v\n", sb, count)
			sb = rec.StartBlock
			count = 1
		}
	}
}
