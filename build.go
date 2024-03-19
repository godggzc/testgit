package testgit

import "fmt"

type myStruct struct {
	Cat string `json:"cat"`
	Dog string `json:"dog"`
}

func (m *myStruct) fmtAll() string {
	return fmt.Sprintf("Cat: %s, Dog: %s", m.Cat, m.Dog)
}
