package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/structs"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
)

func ejemplon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API of EDD, hopefully you enjoy it! :)")
}

func cargaArchivos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/prueba")
	fmt.Fprintf(w, "Empezando Carga de Archivo \n")
	var newPath structs.Path
	reqPath, err := ioutil.ReadAll(r.Body)

	if err != nil {
		error := structs.Error{Mensaje: "Ha ocurrido un problema :c"}
		json.NewEncoder(w).Encode(error)
	}
	json.Unmarshal(reqPath, &newPath)
	fmt.Println(newPath.Path)
	mensaje := structs.Mensajito{Mensaje: newPath.Path}
	excelization(mensaje.Mensaje)
	carga_BD(db, "orden.csv", "tabla_temp")
	carga_BD(db, "pais.csv", "pais")
}

func consulta_1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Consulta 1 :3")
}

func consulta_2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Consulta 2 :3")
}
func consulta_3(w http.ResponseWriter, r *http.Request) {

}
func consulta_4(w http.ResponseWriter, r *http.Request) {

}
func consulta_5(w http.ResponseWriter, r *http.Request) {

}
func consulta_6(w http.ResponseWriter, r *http.Request) {

}
func consulta_7(w http.ResponseWriter, r *http.Request) {

}
func consulta_8(w http.ResponseWriter, r *http.Request) {

}
func consulta_9(w http.ResponseWriter, r *http.Request) {

}
func consulta_10(w http.ResponseWriter, r *http.Request) {

}

func excelization(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	createCSV(f, "Orden", "orden.csv")
	createCSV(f, "País", "pais.csv")
	createCSV(f, "Cliente", "cliente.csv")
	createCSV(f, "Categoría", "categoria.csv")
	createCSV(f, "Producto", "producto.csv")
	createCSV(f, "Vendedor", "vendedor.csv")

}

func createCSV(file *excelize.File, sheet string, csv_name string) {
	rows, err := file.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}

	fa, err := os.Create("C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/" + csv_name)
	defer fa.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(fa)
	err = w.WriteAll(rows) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func carga_BD(db *sql.DB, tipo string, tabla string) {
	sql := "LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/" + tipo + "' INTO TABLE prueba." + tabla + " FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'	IGNORE 1 LINES "
	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Holis")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ejemplon).Methods("GET")
	router.HandleFunc("/carga", cargaArchivos).Methods("POST")
	router.HandleFunc("/consulta_1", consulta_1).Methods("GET")
	router.HandleFunc("/consulta_2", consulta_2).Methods("GET")
	router.HandleFunc("/consulta_3", consulta_3).Methods("GET")
	router.HandleFunc("/consulta_4", consulta_4).Methods("GET")
	router.HandleFunc("/consulta_5", consulta_5).Methods("GET")
	router.HandleFunc("/consulta_6", consulta_6).Methods("GET")
	router.HandleFunc("/consulta_7", consulta_7).Methods("GET")
	router.HandleFunc("/consulta_8", consulta_8).Methods("GET")
	router.HandleFunc("/consulta_9", consulta_9).Methods("GET")
	router.HandleFunc("/consulta_10", consulta_10).Methods("GET")

	header := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(header, methods, origins)(router)))
	fmt.Println("Adios")

}
