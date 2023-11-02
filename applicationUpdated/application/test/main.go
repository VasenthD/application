package main

import "fmt"

// "bufio"

// "io"
// "os"

func main() {
	// localFile, err := os.Open("/home/vasenth/Documents/umamaheswari/download.jpeg")
	// if err != nil {
	// 	fmt.Println("error opening local file: ", err)
	// 	return
	// }

	// stat, err := localFile.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// bs := make([]byte, stat.Size())
	// _, err = bufio.NewReader(localFile).Read(bs)
	// if err != nil && err != io.EOF {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("size: ",stat.Size)
	// fmt.Println("byte array: ", bs)
	// folpath := "/sft/testint"
	// bytefol := []rune(folpath)
	// // if folpath[len(folpath)-1]
	// fmt.Println(bytefol[len(bytefol)-1])
	// // fmt.Println("/sftp"+"/"+"new.jpg")

	str := "sftp/testint" +"/"+ "new.txt"
	fmt.Println(str)

}
