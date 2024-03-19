package testgit

import "fmt"

type MyStruct struct {
	Cat string `json:"cat"`
	Dog string `json:"dog"`
}

func (m *MyStruct) fmtAll() string {
	return fmt.Sprintf("Cat: %s, Dog: %s", m.Cat, m.Dog)
}
