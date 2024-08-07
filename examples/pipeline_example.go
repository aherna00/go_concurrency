package examples

import (
	"net/http"
	"sync"
)

type House struct {
	Name          string
	Address       string
	Owner         string
	Value         int
	Neighborvalue int
}

func HandleHouses(r *http.Request) {

	var authTokens string

	////////////stage1/////////////

	var stage1 sync.WaitGroup

	var house House

	stage1.Add(1)
	go func() {

		defer stage1.Done()
		authTokens = fetchAuthTokens()
	}()

	stage1.Add(1)
	go func() {

		defer stage1.Done()
		house = unmarshalHouse(r)
	}()

	stage1.Wait()

	////////////////////////////////

	////////////stage2/////////////

	var stage2 sync.WaitGroup

	var houseValue int

	var neighborValue int

	stage2.Add(1)
	go func() {

		defer stage2.Done()

		houseValue = fetchHouseValue(house, authTokens)
	}()

	stage2.Add(1)
	go func() {

		defer stage2.Done()

		neighborValue = fetchNeighborValue(house, authTokens)

	}()

	stage2.Wait()
	////////////////////////////////

	////////////stage3/////////////

	var stage3 sync.WaitGroup

	house.Value = houseValue
	house.Neighborvalue = neighborValue

	stage3.Add(1) //should always increment wait group before starting go routine

	go func() {

		defer stage3.Done()

		err := writeHouseToDb(house)
		if err != nil {
			panic(err)
		}
	}()

	stage3.Add(1)

	go func() {

		defer stage3.Done()

		err := postToAPI(house, authTokens)
		if err != nil {
			panic(err)
		}
	}()

	stage3.Wait()

	return

	////////////////////////////////

}
func fetchAuthTokens() string {
	return ""
}

func unmarshalHouse(r *http.Request) House {
	return House{}
}

func fetchHouseValue(house House, a string) int {
	return 1
}

func fetchNeighborValue(house House, a string) int {
	return 2
}

func writeHouseToDb(house House) error {

	return nil
}

func postToAPI(house House, authTokens string) error {
	return nil
}
