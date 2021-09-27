package main

import (
	"go-ip-mmdb/internal/model"
	"log"
	"net"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/inserter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

var (
	InCityDbPath  string = os.Getenv("GEOIP2_CITY_INPUT")
	InIspDbPath   string = os.Getenv("GEOIP2_ISP_INPUT")
	OutCityDbPath string = os.Getenv("GEOIP2_CITY_OUTPUT")
	OutIspDbPath  string = os.Getenv("GEOIP2_ISP_OUTPUT")
	InputTsvPath  string = os.Getenv("INPUT_TSV")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println("Begin")
	// Reads and processes TSV File
	if InputTsvPath == "" {
		InputTsvPath = "../input.tsv"
	}

	records, err := model.ReadTsv(InputTsvPath)
	if err != nil {
		log.Fatal(err)
	}
	if len(records) == 0 {
		log.Fatal("Empty input file")
	}

	log.Println("Load GeoIP2 database")
	IspWriter, err := mmdbwriter.Load(InIspDbPath, mmdbwriter.Options{})
	if err != nil {
		log.Fatal(err)
	}
	CityWriter, err := mmdbwriter.Load(InCityDbPath, mmdbwriter.Options{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Create new records")
	for _, record := range records {
		_, ipNet, err := net.ParseCIDR(record.Network)
		if err != nil {
			log.Printf("[Error] %s | %s", record.Network, err)
		}

		// Writes to ISP Database
		log.Println("\t" + record.Network + ": ISP record")
		ispData := mmdbtype.Map{
			"isp":          mmdbtype.String(record.IspName),
			"organization": mmdbtype.String(record.IspName),
		}
		if err := IspWriter.InsertFunc(ipNet, inserter.TopLevelMergeWith(ispData)); err != nil {
			log.Printf("[Error] %s | %s", record.Network, err)
		}

		log.Println("\t" + record.Network + ": City record")
		cityData := mmdbtype.Map{
			"city": mmdbtype.Map{
				"geoname_id": mmdbtype.Uint16(record.CityGeoNameID),
				"names": mmdbtype.Map{
					"en": mmdbtype.String(record.CityName),
				},
			},
		}
		if err := CityWriter.InsertFunc(ipNet, inserter.TopLevelMergeWith(cityData)); err != nil {
			log.Printf("[Error] %s | %s", record.Network, err)
		}
	}

	// Write the newly enriched DB to the filesystem.
	log.Println("Write to " + OutIspDbPath)
	ispFh, err := os.Create(OutIspDbPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = IspWriter.WriteTo(ispFh)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Write to " + OutIspDbPath)
	cityFh, err := os.Create(OutCityDbPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = CityWriter.WriteTo(cityFh)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
