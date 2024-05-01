package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"modulosn/Analyzer"
	"modulosn/DiskManagement"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type responseList struct {
	Status int64    `json:"Status"`
	List   []string `json:"List"`
}

type responseString struct {
	Status int64  `json:"Status"`
	Value  string `json:"Value"`
}

type loginValues struct {
	User     string `json:"User"`
	Password string `json:"Password"`
}

func postMethod(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var newLoginValues loginValues
	if err := json.Unmarshal(reqBody, &newLoginValues); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	var response interface{}

	// Verificar el comando recibido
	// Lógica para procesar el comando "lista"
	fmt.Println(newLoginValues.User)
	response = Analyzer.Analyze(newLoginValues.User)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response JSON", http.StatusInternalServerError)
		return
	}
}

func getMethod(w http.ResponseWriter, r *http.Request) {
	var newResponseString responseString
	newResponseString.Status = 200
	newResponseString.Value = "Hello World"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newResponseString)
}

func postMethodDisk(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Data")
	}
	var newLoginValues loginValues
	json.Unmarshal(reqBody, &newLoginValues)

	fmt.Println(newLoginValues.User)
	fmt.Println(newLoginValues.Password)

	var newResponseList responseList
	newResponseList.Status = 200

	ruta := "./MIA"
	nombres, err := obtenerNombresYExtensiones(ruta)
	if err != nil {
		fmt.Println("Error al obtener los nombres de los archivos:", err)
		return
	}
	newResponseList.List = nombres

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newResponseList)
}

func postMethoPartition(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Data")
	}
	var newLoginValues loginValues
	json.Unmarshal(reqBody, &newLoginValues)

	fmt.Println(newLoginValues.User)
	fmt.Println(newLoginValues.Password)

	var newResponseList responseList
	newResponseList.Status = 200

	cadenaMinuscula := strings.ToLower(newLoginValues.User)
	// Quitar la extensión ".disk"
	cadenaSinExtension := strings.TrimSuffix(cadenaMinuscula, ".disk")
	partition := DiskManagement.ListarParticionesMontadas(cadenaSinExtension)
	newResponseList.List = partition

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newResponseList)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API :D")
}

func obtenerNombresYExtensiones(ruta string) ([]string, error) {
	var nombres []string

	// Leer el contenido del directorio
	archivos, err := ioutil.ReadDir(ruta)
	if err != nil {
		return nil, err
	}

	// Iterar sobre cada archivo
	for _, archivo := range archivos {
		// Obtener el nombre y la extensión del archivo
		nombre := archivo.Name()
		extension := filepath.Ext(nombre)

		// Verificar si el nombre ya incluye una extensión
		if extension != "" {
			// Si ya incluye una extensión, agregar el nombre sin modificar
			nombres = append(nombres, nombre)
		} else {
			// Si no incluye una extensión, combinar el nombre y la extensión y agregarlo a la lista
			nombreCompleto := nombre + extension
			nombres = append(nombres, nombreCompleto)
		}
	}

	return nombres, nil
}

func main() {
	fmt.Println("Server started on port 4000")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoute)
	router.HandleFunc("/comand", postMethod).Methods("POST") //comandos

	router.HandleFunc("/tasks", postMethodDisk).Methods("POST") //DISCOS
	router.HandleFunc("/tasks", getMethod).Methods("GET")

	router.HandleFunc("/partition", postMethoPartition).Methods("POST") //PARTITION

	// router.HandleFunc("/login", postMethoPartition).Methods("POST") //login

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":4000", handler))
}
