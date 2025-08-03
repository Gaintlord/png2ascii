package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var asciiCharaArray = []string{
	"$",
	"@",
	"B",
	"%",
	"8",
	"&",
	"W",
	"M",
	"#",
	"*",
	"o",
	"a",
	"h",
	"k",
	"b",
	"d",
	"p",
	"q",
	"w",
	"m",
	"Z",
	"O",
	"0",
	"Q",
	"L",
	"C",
	"J",
	"U",
	"Y",
	"X",
	"z",
	"c",
	"v",
	"u",
	"n",
	"x",
	"r",
	"j",
	"f",
	"t",
	"/",
	"|",
	"(",
	")",
	"1",
	"{",
	"}",
	"[",
	"]",
	"?",
	"-",
	"_",
	"+",
	"~",
	"<",
	">",
	"i",
	"!",
	"l",
	"I",
	";",
	":",
	",",
	`"`,
	"^",
	"'",
	".",
}

func CreateAsciiWithLoop(wid int, height int, img image.Image) {

	var greyscalevalue float64
	var ratio float64
	for i := 0; i < wid; i++ {
		for j := 0; j < height; j++ {
			red, green, blue, _ := img.At(j, i).RGBA()
			ratio = ((0.3 * float64(red/257)) + (0.59 * float64(green/257)) + (0.11 * float64(blue/257)))
			greyscalevalue = ratio
			// fmt.Printf(" #%v# ", int(greyscalevalue/3.835))
			fmt.Print(asciiCharaArray[66-int(greyscalevalue/3.835)])
		}
		fmt.Print("\n")
	}

}

func main() {

	path := flag.String("path", "gaintlord.jpg", "print an ascii art in terminal")
	flag.Parse()
	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("error loading the image")
	}
	println(file)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("something went wrong:-", err)
	}
	// println(file)
	bound := img.Bounds()
	width := bound.Max.X
	height := bound.Max.Y
	fmt.Printf("width :%v \n, height:%v\n", width, height)

	CreateAsciiWithLoop(width, height, img)

}
