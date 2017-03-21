package runecount

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	//"strconv"
)

const Items = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,!?`

func FindAll(filename string) {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	var biggest int = 0
	var andre int = 0
	var tredje int = 0
	var fjerde int = 0
	var femte int = 0

	var bokstav byte
	var andreb byte
	var tredjeb byte
	var fjerdeb byte
	var femteb byte

	for i := 0; i < len(Items); i++ {

		if bytes.Contains(byteSlice, []byte{Items[i]}) {
			//fmt.Println("Det er", , " i denne filen.)
			fmt.Printf("Det er %d %q i denne filen\n", bytes.Count(byteSlice, []byte{Items[i]}), Items[i])
			v, _ := bytes.Count(byteSlice, []byte{Items[i]}), Items[i]
			if v > femte {
				femte = v
				femteb = Items[i]
				if fjerde < femte {
					fjerde = femte
					fjerdeb = femteb
					femte = 0

					if tredje < fjerde {
						tredje = fjerde
						tredjeb = fjerdeb
						fjerde = 0
						if andre < tredje {
							andre = tredje
							andreb = fjerdeb
							tredje = 0
							if biggest < andre {
								biggest = andre
								bokstav = andreb
								andre = 0

							}

						}

					}

				}

			}

		} else {
			//fmt.Println("Det er ingen a i denne filen.\")
			fmt.Printf("Der er ingen %q i denne filen.\n", Items[i])
		}
	}
	fmt.Println("\n")
	fmt.Printf("Det er flest av %q i denne filen og det er %d stk\n", bokstav, biggest)
	fmt.Printf("Det er nest flest av %q i denne filen og det er %d stk\n", andreb, andre)
	fmt.Printf("Det er tredje flest av %q i denne filen og det er %d stk\n", tredjeb, tredje)
	fmt.Printf("Det er fjerde flest av %q i denne filen og det er %d stk\n", fjerdeb, fjerde)
	fmt.Printf("Det er femte flest av %q i denne filen og det er %d stk\n", femteb, femte)

}
