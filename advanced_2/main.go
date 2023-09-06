package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("%v (%d €)", p.Name, p.Price)
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Bacon ipsum dolor amet porchetta short ribs short loin, spare ribs t-bone kielbasa bresaola ")
	fmt.Fprint(w, "tail ribeye pastrami flank doner. Turducken shankle kevin, landjaeger rump bresaola \n")
	// don't forget flush
	w.Flush()

	number_input := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36"
	input_scanner := bufio.NewScanner(strings.NewReader(number_input))

	//custom split by comma function
	splitByComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := strings.IndexRune(string(data), ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return
	}

	input_scanner.Split(splitByComma)
	buf := make([]byte, 2)
	// scan 2 bytes at the time
	input_scanner.Buffer(buf, bufio.MaxScanTokenSize)
	for input_scanner.Scan() {
		fmt.Printf("%s ", input_scanner.Text())
	}

	s := fmt.Sprint(Product{
		Name:  "Quicky Pants",
		Price: 100,
	})

	println(s)
}
