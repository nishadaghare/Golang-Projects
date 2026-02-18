package services

import (
	"encoding/csv"
	"os"
	"yourapp/models"
)

func LoadRegions(filePath string) map[string]*models.Region {
	file, _ := os.Open(filePath)
	reader := csv.NewReader(file)

	rows, _ := reader.ReadAll()

	regionMap := make(map[string]*models.Region)

	getOrCreate := func(code string) *models.Region {
		if r, ok := regionMap[code]; ok {
			return r
		}
		r := &models.Region{Code: code}
		regionMap[code] = r
		return r
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		city := row[0]
		state := row[1]
		country := row[2]

		c := getOrCreate(country)
		s := getOrCreate(state)
		ci := getOrCreate(city)

		s.Parent = c
		c.Children = append(c.Children, s)

		ci.Parent = s
		s.Children = append(s.Children, ci)
	}

	return regionMap
}
