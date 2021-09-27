package model

import (
	"encoding/csv"
	"os"
	"strconv"
)

type DataRow struct {
	Network        string
	CityName       string
	CityGeoNameID  uint
	CountryName    string
	CountryIsoCode string
	IspName        string
}

// ReadTsv reads input TSV file, whose headers are: Network, CityName, CityGeoNameID, CountryIsoCode, IspName
func ReadTsv(filepath string) ([]DataRow, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	result := []DataRow{}
	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = '\t'
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range csvLines {
		cityGeoNameId, _ := strconv.ParseUint(line[2], 10, 64)
		d := DataRow{
			Network:        line[0],
			CityName:       line[1],
			CityGeoNameID:  uint(cityGeoNameId),
			CountryName:    CountryName[line[3]], // readmore at country.go
			CountryIsoCode: line[3],
			IspName:        line[4],
		}
		result = append(result, d)
	}
	return result, nil
}
