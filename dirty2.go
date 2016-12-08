package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func zFoo(arg string) int {
	return 1
	return 1
}

func complexFunc5(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	b, err = ioutil.ReadAll(r)
	r.Read([]byte{})
	if err != nil {
		return err
	}
	a := 0
	c := make(chan string, 0)
	for a == 1 {
		select {
		case <-c:
		default:
			switch a {
			case 1:
				fmt.Println("hoge")
			case 0:
				fmt.Printf("hoge\n")
			default:
				os.Stdout.Write(b)
			case 2:
				if a == 4 {
					if len(b) != 0 {
						panic("panic!")
					}
				}
				return nil
			}

		}
	}

	return nil
}

func ComplexFunc3(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	b, err = ioutil.ReadAll(r)
	r.Read([]byte{})
	if err != nil {
		return err
	}
	a := 0
	c := make(chan string, 0)
	for a == 1 {
		select {
		case <-c:
		default:
			switch a {
			case 1:
				fmt.Println("hoge")
			case 0:
				fmt.Printf("hoge\n")
			default:
				os.Stdout.Write(b)
			case 2:
				if a == 4 {
					if len(b) != 0 {
						panic("panic!")
					}
				}
				return nil
			}

		}
	}
	return nil
}

type BadnameUrl struct {
	foo string
}

type BrokenStruct2 struct {
	hoge string `aabbcc`
}

// dirty MarshalJSON
func (b *BrokenStruct2) MarshalJSON() int {
	return 0
}
func (s *BrokenStruct2) Hoge() {
	return
}
