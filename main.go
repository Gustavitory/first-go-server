package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){//action
	// fmt.Fprintf(w,"Hola mundo")//por que Fprintf? => porque escribe en un formato especifico, el primer parametro es el writer
	template,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
		fmt.Fprintf(w,"Page not found.")
	}else{
		template.Execute(w,nil)
	}
}

func main(){
	http.HandleFunc("/",index); // endpoint
	fmt.Println("El servidor esta escuchando en el puerto 3001")
	http.ListenAndServe(":3001",nil) // inicia el server en el puerto 3001
}