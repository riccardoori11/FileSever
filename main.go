package main

import(
	"FS/greetings/fileserver"
	"os"
	"fmt"

)

func main(){


	
	if len(os.Args) < 2{

		fmt.Println("Wrong usage")
		return
	}
	
	path := os.Args[1]
	if !fileserver.Verify_path(path){
		
		fmt.Println("The file was not found")
		return

	}
	
	fileserver.FileSever(path)

}
