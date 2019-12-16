package main

import(
	"os"
	"fmt"
	"io"
)

var buf = make(chan []byte,1024)

func readFile(buf chan<-[]byte, f *os.File ){
	fileData := make([]byte,1024)
	for {
		n , err := f.Read(fileData)
		if err != nil && err == io.EOF {
			fmt.Printf("文件读完，n=%d\n",n)
			close(buf)
			return
		}
		buf <- fileData
		//fileWrite.Write(buf[:n])
	}
}

func writeFile(buf <-chan []byte,f *os.File ){
	for fileData := range buf {
		f.Write(fileData)
	}
}
func main(){
	fileRead, err := os.Open("CustomerReportRightsControlController.php")
	if err != nil {
		fmt.Println("Open err:",err)
	}
	defer fileRead.Close()

	//创建写文件

	fileWrite ,err := os.Create("test.php")
	if err != nil {
		fmt.Println("Open err:",err)
	}
	defer fileWrite.Close()

	// fileData := make([]byte,1024)

	// for {
	// 	n , err := fileRead.Read(fileData)
	// 	if err != nil && err == io.EOF {
	// 		fmt.Printf("文件读完，n=%d\n",n)
	// 		return
	// 	}
	// 	fileWrite.Write(buf[:n])
	// }
	go readFile(buf,fileRead)
	go writeFile(buf,fileWrite)
	
}