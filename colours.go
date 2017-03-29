package main

import(
	"fmt"
	"image"
    //"sync"
    "os"
    "time"
    //"math"
	//"log"
	_ "image/png"
	_ "image/jpeg"
)

type Pixel struct {
    R int
    G int
    B int
}

type SavedColor struct{
	next *SavedColor
	col Pixel
}

var pixel_number int
var scanned bool

func open_image(name string)(image.Image,int,int){
	infile, err := os.Open(name)
    if err != nil {
        fmt.Println("err nr.1")
        panic(err)
    }
    defer infile.Close()
    src, _, err := image.DecodeConfig(infile)
    if err != nil {
        fmt.Println("err nr.2")
        panic(err)
    }
    fmt.Println(src.Width, src.Height)
    width := src.Width
    height := src.Height
    infile2, err := os.Open(name)
    if err != nil {
        fmt.Println("err nr.3")
        panic(err)
    }
    defer infile2.Close()
    src2, _, err := image.Decode(infile2)
    if err != nil {
        fmt.Println("err nr.4")
        panic(err)
    }
    return src2,width,height;
}

func scan(w, h, offset_x_p, offset_x_m, offset_y_p, offset_y_m int, col [256][256][256]int, src2 image.Image)*SavedColor{
	var ptr *SavedColor
    for y := offset_y_p; y < (h-offset_y_m); y++ {
        for x := offset_x_p; x < (w-offset_x_m); x++ {
            r, g, b, _ := src2.At(x, y).RGBA()
            intr,intg,intb := (int)(r>>8),(int)(g>>8),(int)(b>>8)
            if(col[intr][intg][intb]==0){
	            ptr = insert(ptr,intr,intg,intb)
	            pixel_number++
            }
            col[intr][intg][intb]++;
            //fmt.Println(intr,intg,intb,col[intr][intg][intb])
        }
    }
    top3(col,ptr)
    return ptr
}

func insert(ptr *SavedColor,r,g,b int)*SavedColor {
	newPointer := new(SavedColor)
	newPointer.col.R = r
	newPointer.col.G = g
	newPointer.col.B = b
	newPointer.next = ptr
	return newPointer
}

func top3(col [256][256][256]int, ptr *SavedColor) {
	var f,s,t Pixel
	var first,second,third int
	if(ptr != nil){
		//inizializzo tutto al primo elemento
		f,s,t = ptr.col, ptr.col, ptr.col
		//fmt.Println("color times : ", col[39][40][34])
		first,second,third = col[f.R][f.G][f.B],0,0
    	for(ptr.next != nil){
    		//fmt.Println(ptr.col.R,ptr.col.G,ptr.col.B)
    		ptr = ptr.next
    		num := col[ptr.col.R][ptr.col.G][ptr.col.B]
    		if(num > first){
    			f = ptr.col
    			first = num
			}else if(num > second){
				s = ptr.col
    			second = num
			}else if(num > third){
				t = ptr.col
    			third = num
			}
    	}
    }
    fmt.Println("First  : ",f.R,f.G,f.B,first)
    fmt.Println("Second : ",s.R,s.G,s.B,second)
    fmt.Println("Third  : ",t.R,t.G,t.B,third)
}

func main() {
    start := time.Now()
    //apro l'immagine
    image,width,height := open_image(os.Args[1])
    //ometto la trasparenza per problemi di memoria impliciti di go
	var col [256][256][256] int
    _ = scan(width,height,0,0,0,0,col,image)
    fmt.Println("colours count: ",pixel_number)
    elapsed := time.Since(start)
    fmt.Println("elapsed time: ",elapsed)
}