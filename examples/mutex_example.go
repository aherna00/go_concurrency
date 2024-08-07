package examples

import (
	"fmt"
	"sync"
)

type VeryImportantResource struct {
	VeryImportantThings []string
	mutex               *sync.RWMutex
	AddChan             chan string
	DeleteChan          chan string
	UpdateChan          chan string
	ReadChan            chan string
}

func NewVeryImportantResource() *VeryImportantResource {
	return &VeryImportantResource{
		AddChan:    make(chan string),
		DeleteChan: make(chan string),
		UpdateChan: make(chan string),
		ReadChan:   make(chan string),
	}
}

func (v *VeryImportantResource) Run() {

	for {
		select {
		case add := <-v.AddChan:
			v.mutex.Lock()
			v.VeryImportantThings = append(v.VeryImportantThings, add)
			v.mutex.Unlock()
		case d := <-v.DeleteChan:
			v.mutex.Lock()
			deleteFromSlice(d)
			v.mutex.Unlock()
		case update := <-v.UpdateChan:
			v.mutex.Lock()
			v.VeryImportantThings = append(v.VeryImportantThings, update)
			v.mutex.Unlock()
		case <-v.ReadChan:
			v.mutex.RLock()
			fmt.Println(v.VeryImportantThings)
			v.mutex.RUnlock()
		}
	}

}

func deleteFromSlice(delete string) {

}
