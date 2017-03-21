package main

import "fmt"
import "strings"

func main() {



str := "001001001001001001001001001001001001001001001001001" +
"00000000000000000000000000000000000000000011111111111110101010101010101" +
"01010101010101010101010101101101101101101101101101101101101101101100101" +
"0101010101010101010101010101010101010101010101010101010101"

subStr1 := "001"  //Teller 2 for mange elever.

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr1) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Helse har %d elever\n", strings.Count(str, subStr1))

}

subStr2 := "000"

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr2) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Humanoira har %d elever\n", strings.Count(str, subStr2))

}

subStr3 := "111"

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr3) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Kunst har %d elever\n", strings.Count(str, subStr3))

/*  //Telte alt for mange. Klarte ikke å skille ut.
}

subStr4 := "10"

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr4) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Teknologi har %d elever\n", strings.Count(str, subStr4))
*/

}

subStr5 := "110" //Telte 1 for mye.

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr5) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Lærer har %d elever\n", strings.Count(str, subStr5))

}

/* //Funket ikke.
subStr6 := "01"

for i := 0; i < len(str); i++ {

if strings.Contains(str, subStr6) {
//	H := strings.Count(str, subStr1)
}
		fmt.Printf("Økonomi har %d elever\n", strings.Count(str, subStr6))

}
*/
}