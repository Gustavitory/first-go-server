package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){//action con template
	// fmt.Fprintf(w,"Hola mundo")//por que Fprintf? => porque escribe en un formato especifico, el primer parametro es el writer
	template,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
		fmt.Fprintf(w,"Page not found.")
	}else{
		template.Execute(w,nil)
	}
}

func readParameters(w http.ResponseWriter, r *http.Request){
	namesList, exist := r.URL.Query()["name"]
	fmt.Printf("%v\n", namesList)
	if exist{
		io.WriteString(w,"Hola, "+namesList[0])
	}else {io.WriteString(w,"Hola, desconocido")}
}

func main(){
	http.HandleFunc("/",index) // endpoint
	http.HandleFunc("/param",readParameters)
	fmt.Println("El servidor esta escuchando en el puerto 3001")
	http.ListenAndServe(":3001",nil) // inicia el server en el puerto 3001
}