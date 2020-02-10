package main

import (
	"errors"
	"fmt"
	"strings"
)

func GetDefaultGroup() string  {
	return "imageName"
}

func Group(groupBy string, podsDataset PodsData) (PodsData, error)  {
	switch groupBy {
		case "none":
			return podsDataset, nil
		case "imageName":
			podsDataset.Data = groupByImageName(podsDataset.Data)
			return podsDataset, nil
	}

	return podsDataset, errors.New(fmt.Sprintf("Something went very wrong, give group-by [%s] not exist in [%s]", groupBy, strings.Join(ListAvailableGroupers(), ",")))

}


func ListAvailableGroupers() []string {
	return []string{"none", "imageName"}
}


func groupByImageName(data [][]string) [][]string  {
	for rowId, row := range data {
		err, idOfAfterNextElement := findAfterNextElement(rowId, row[1], data)
		if err != nil {
			continue
		}

		temp := data[rowId+1]
		data[rowId+1] = data[idOfAfterNextElement]
		data[idOfAfterNextElement] = temp
	}
	return data
}

func findAfterNextElement(rowIdx int, searchableString string, dataset [][]string) (error, int) {
	for rowIdx += 2; rowIdx < len(dataset); rowIdx += 1 {
		if searchableString == dataset[rowIdx][1] {
			return nil, rowIdx
		}
	}
	return errors.New(fmt.Sprintf("Could not find id [%d] and string [%s] in given Data set", rowIdx, searchableString)), 0
}
