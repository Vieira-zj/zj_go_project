package demos

import (
	"fmt"
	"io"
	"strings"
)

// demo 01, map
func testMapGetEmpty() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println("item at 2 =>", m[2])
	fmt.Printf("first char: %c\n", m[2][0])
	fmt.Printf("item length: %d\n", len(m[2]))

	if len(m) > 0 && len(m[3]) > 0 {
		fmt.Println("item at 3 =>", m[3])
	}
}

// demo 02-01, custom reader
type alphaReader1 struct {
	src string
	cur int
}

func newAlphaReader1(src string) *alphaReader1 {
	return &alphaReader1{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader1) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	bound := 0
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}

	buf := make([]byte, bound)
	for n := 0; n < bound; n++ {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		a.cur++
	}
	copy(p, buf)
	return bound, nil
}

func testAlphaReader1() {
	reader := newAlphaReader1("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		// fmt.Printf("%d\n", n)
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

// demo 02-02, custom reader
type alphaReader2 struct {
	reader io.Reader
}

func newAlphaReader2(reader io.Reader) *alphaReader2 {
	return &alphaReader2{reader: reader}
}

func (a *alphaReader2) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}

func testAlphaReader2() {
	reader := newAlphaReader2(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(p[:n]))
				break
			}
			panic(err.Error())
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

// demo 02-03, custom writer

// MainDemo03 : main
func MainDemo03() {
	// testMapGetEmpty()

	// testAlphaReader1()
	// testAlphaReader2()

	fmt.Println("demo 03 done.")
}
