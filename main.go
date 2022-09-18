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
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
)

func ejemplon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API of EDD, hopefully you enjoy it! :)")
}

func cargaArchivos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")
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
	carga_BD(db, "pais.csv", "pais")
	carga_BD(db, "cliente.csv", "cliente")
	carga_BD(db, "categoria.csv", "categoria")
	carga_BD(db, "producto.csv", "producto")
	carga_BD(db, "vendedor.csv", "vendedor")
	carga_BD(db, "orden.csv", "temp_orden")
}

func consulta_1(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 1 :3")
	sql := structs.Peticion1()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons1
	for result.Next() {
		var Cons1 structs.Cons1
		result.Scan(&Cons1.ID_Cliente, &Cons1.NombreCompleto, &Cons1.Pais, &Cons1.Monto_Total)
		envio = append(envio, Cons1)
		fmt.Printf("%v \n", Cons1)
	}

	enviar := structs.EnvCons1{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)

}

func consulta_2(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 1 :3")
	sql := structs.Peticion2()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons2
	for result.Next() {
		var Cons2 structs.Cons2
		result.Scan(&Cons2.Id_Producto, &Cons2.Nombre_Producto, &Cons2.Categoria, &Cons2.Cantidad, &Cons2.Precio)
		envio = append(envio, Cons2)
		fmt.Printf("%v \n", Cons2)
	}

	enviar := structs.EnvCons2{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_3(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 3 :3")
	sql := structs.Peticion3()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons3
	for result.Next() {
		var Cons3 structs.Cons3
		result.Scan(&Cons3.Id_vendedor, &Cons3.Nombre, &Cons3.Vendido)
		envio = append(envio, Cons3)
		fmt.Printf("%v \n", Cons3)
	}

	enviar := structs.EnvCons3{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_4(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 4 :3")
	sql := structs.Peticion4()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons4
	for result.Next() {
		var Cons4 structs.Cons4
		result.Scan(&Cons4.Nombre, &Cons4.Precio)
		envio = append(envio, Cons4)
		fmt.Printf("%v \n", Cons4)
	}

	enviar := structs.EnvCons4{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_5(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 5 :3")
	sql := structs.Peticion5()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons5
	for result.Next() {
		var Cons5 structs.Cons5
		result.Scan(&Cons5.Id_pais, &Cons5.Pais, &Cons5.Monto)
		envio = append(envio, Cons5)
		fmt.Printf("%v \n", Cons5)
	}

	enviar := structs.EnvCons5{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_6(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 6 :3")
	sql := structs.Peticion6()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons6
	for result.Next() {
		var Cons6 structs.Cons6
		result.Scan(&Cons6.Nombre, &Cons6.Cantidad)
		envio = append(envio, Cons6)
		fmt.Printf("%v \n", Cons6)
	}

	enviar := structs.EnvCons6{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_7(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 7 :3")
	sql := structs.Peticion7()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons7
	for result.Next() {
		var Cons7 structs.Cons7
		result.Scan(&Cons7.Pais, &Cons7.Categoria, &Cons7.Cantidad)
		envio = append(envio, Cons7)
		fmt.Printf("%v \n", Cons7)
	}

	enviar := structs.EnvCons7{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_8(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 8 :3")
	sql := structs.Peticion8()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons8
	for result.Next() {
		var Cons8 structs.Cons8
		result.Scan(&Cons8.Mes, &Cons8.Total)
		envio = append(envio, Cons8)
		fmt.Printf("%v \n", Cons8)
	}

	enviar := structs.EnvCons8{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_9(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 9 :3")
	sql := structs.Peticion9()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons9
	for result.Next() {
		var Cons9 structs.Cons9
		result.Scan(&Cons9.Mes, &Cons9.Total)
		envio = append(envio, Cons9)
		fmt.Printf("%v \n", Cons9)
	}

	enviar := structs.EnvCons9{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
}
func consulta_10(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "root:Ak22/03/2022@tcp(127.0.0.1:3306)/bases1_p1")

	fmt.Println("Consulta 10 :3")
	sql := structs.Peticion10()
	result, _ := db.Query(sql)
	defer result.Close()
	var envio []structs.Cons10
	for result.Next() {
		var Cons10 structs.Cons10
		result.Scan(&Cons10.Id_Producto, &Cons10.Nombre_Producto, &Cons10.Total)
		envio = append(envio, Cons10)
		fmt.Printf("%v \n", Cons10)
	}

	enviar := structs.EnvCons10{
		Lista: envio,
	}

	json.NewEncoder(w).Encode(enviar)
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
	if sheet == "Orden" {
		for i, fila := range rows {
			if i != 0 {
				for j, celda := range fila {
					if j == 2 {
						rows[i][j] = parseDate(celda)
					}
				}
			}
		}
	}

	fa, err := os.Create("C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/" + csv_name)
	defer fa.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(fa)
	err = w.WriteAll(rows)

	if err != nil {
		log.Fatal(err)
	}
}

func carga_BD(db *sql.DB, tipo string, tabla string) {
	sql := "LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/" + tipo + "' INTO TABLE bases1_p1." + tabla + " FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'	IGNORE 1 LINES "
	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
}

func parseDate(fecha string) string {
	joinstring := "/"
	temsplit := strings.Split(fecha, "-")
	return temsplit[2] + joinstring + temsplit[0] + joinstring + temsplit[1]
}

func main() {
	fmt.Println("Empiece a realizar consultas")
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
