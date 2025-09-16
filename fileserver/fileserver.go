package fileserver

import (
	"log"
	"net/http"
	"os"
)

func Verify_path(path string) bool{

	
	
	_,err := os.Stat(path)
	
	if os.IsNotExist(err){
		
		return false


		
	}else{

		return true
	}
}


func FileSever(path string){

	fs := http.FileServer(http.Dir(path))
	
	http.Handle("/", fs)
	
	port := ":8443"
	
	err := 	http.ListenAndServeTLS(port, "server.crt","server.key",nil)
	
	if err != nil{
		
		log.Fatal("ListenAndServe", err)

	}
	
}
