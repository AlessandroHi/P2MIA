package DiskManagement

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"modulosn/Prints"
	"modulosn/Structs"
	"modulosn/Utilities"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// ================== MKDISK ============================
func Mkdisk(size int, fit string, unit string) {
	Prints.Prints = append(Prints.Prints, "\n======Start MKDISK======")
	Prints.Prints = append(Prints.Prints, "Size:"+strconv.Itoa(size))
	Prints.Prints = append(Prints.Prints, "Fit:"+fit)
	Prints.Prints = append(Prints.Prints, "Unit:"+unit)

	// validate fit equals to b/w/f
	if fit != "bf" && fit != "wf" && fit != "ff" {
		Prints.Prints = append(Prints.Prints, "Error: Fit must be b, w or f")
		return
	}

	// validate size > 0
	if size <= 0 {
		Prints.Prints = append(Prints.Prints, "Error: Size must be greater than 0")
		return
	}

	// validate unit equals to k/m
	if unit != "k" && unit != "m" {
		Prints.Prints = append(Prints.Prints, "Error: Unit must be k or m")
		return
	}

	num := countFilesInFolder()
	letter := letterDisk(num)
	pathdisk := "./MIA/" + strings.ToUpper(letter) + ".disk"

	// Create file
	err := Utilities.CreateFile(pathdisk)
	if err != nil {
		fmt.Println("Error: ", err)

	}

	// Set the size in bytes
	if unit == "k" {
		size = size * 1024
	} else {
		size = size * 1024 * 1024
	}

	// Open bin file
	file, err := Utilities.OpenFile(pathdisk)
	if err != nil {
		return
	}

	// Write 0 binary data to the file

	// create array of byte(0)
	arreglo := make([]byte, 1024)
	// create array of byte(0)
	for i := 0; i <= size/1024; i++ {
		err := Utilities.WriteObject(file, arreglo, int64(i*1024))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// Create a new instance of MRB
	var newMRB Structs.MBR
	newMRB.MbrSize = int32(size) //MBRSize

	// Formatear la fecha --- creationDate
	fechaActual := time.Now()
	formato := "02-01-2006"
	fechaFormateada := fechaActual.Format(formato)
	copy(newMRB.CreationDate[:], fechaFormateada)

	newMRB.Signature = 10 // random ID TOCA VER...................
	copy(newMRB.Fit[:], fit)

	// Write object in bin file
	if err := Utilities.WriteObject(file, newMRB, 0); err != nil {
		return
	}

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)

	// Close bin file
	defer file.Close()

	Prints.Prints = append(Prints.Prints, "======End MKDISK======")
}

func countFilesInFolder() int { // VER CUANTOS ARCHIVOS HAY EN EL DIRECTORIO
	// Lee el contenido de la carpeta
	files, err := ioutil.ReadDir("./MIA")
	if err != nil {
		return 0
	}

	// Contador para los archivos
	count := 0

	// Itera sobre los archivos y cuenta solo los archivos regulares (no directorios)
	for _, file := range files {
		if file.Mode().IsRegular() {
			count++
		}
	}

	// Retorna el número de archivos y ningún error
	return count
}
func letterDisk(number int) string {
	// Convertir el número en un índice válido para el alfabeto (0-25)
	letter := 'a' + byte(number)

	return string(letter)
}

//============================================================
//==================== RMDISK ================================

func Rmdisk(driveletter string) {
	Prints.Prints = append(Prints.Prints, "\n======START RMDISK======")
	// Directorio donde buscar
	directorio := "./MIA"

	// Nombre del archivo a buscar
	nombreArchivo := strings.ToUpper(driveletter) + ".disk"

	// Realizar la búsqueda
	rutaArchivo, err := buscarArchivoPorNombre(directorio, nombreArchivo)
	if err != nil {
		fmt.Println(err) // Imprimir el error
	} else {
		// Se elimina el archivo
		fmt.Println("Confirm deletion: Y/N")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var confir string
		confir = scanner.Text()
		if strings.ToUpper(confir) == "Y" {
			Utilities.DeleteFile(rutaArchivo)
		} else {
			Prints.Prints = append(Prints.Prints, "It was not eliminated\n")
			fmt.Println("It was not eliminated\n")
		}
	}

	fmt.Println("======End RMDISK======")

}

func buscarArchivoPorNombre(dir string, nombre string) (string, error) {
	// Leer el contenido del directorio
	archivos, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	// Iterar sobre los archivos y buscar coincidencias por nombre
	for _, archivo := range archivos {
		if !archivo.IsDir() && archivo.Name() == nombre {
			rutaAbsoluta, err := filepath.Abs(filepath.Join(dir, archivo.Name()))
			if err != nil {
				return "", err
			}
			return rutaAbsoluta, nil
		}
	}

	return "", fmt.Errorf("The disk to be deleted was not found")
}

//============================================================
//==================== FDISK ================================
func Fdisk(size int, driveletter string, name string, unit string, type_ string, fit string) {

	Prints.Prints = append(Prints.Prints, "\n======Start FDISK======")
	Prints.Prints = append(Prints.Prints, "Driveletter:"+driveletter)
	Prints.Prints = append(Prints.Prints, "Name: "+name)
	Prints.Prints = append(Prints.Prints, "Unit: "+unit)
	Prints.Prints = append(Prints.Prints, "Type: "+type_)
	Prints.Prints = append(Prints.Prints, "Size: "+strconv.Itoa(size))
	Prints.Prints = append(Prints.Prints, "Fit: "+fit)

	// validate fit equals to b/w/f
	if fit != "bf" && fit != "wf" && fit != "ff" {
		fmt.Println("Error: Fit must be b, w or f")
		Prints.Prints = append(Prints.Prints, "Error: Fit must be b, w or f")

		return
	}

	// validate size > 0
	if size <= 0 {
		fmt.Println("Error: Size must be greater than 0")
		Prints.Prints = append(Prints.Prints, "Error: Size must be greater than 0")
		return
	}

	// validate unit equals to b/k/m
	if unit != "b" && unit != "k" && unit != "m" {
		fmt.Println("Error: Unit must be b, k or m")
		Prints.Prints = append(Prints.Prints, "Error: Unit must be b, k or m")
		return
	}

	// validate type equals to p/e/l
	if type_ != "p" && type_ != "e" && type_ != "l" {
		fmt.Println("Error: Type must be p, e or l")
		Prints.Prints = append(Prints.Prints, "Error: Type must be p, e or l")
		return
	}

	// Set the size in bytes
	if unit == "k" {
		size = size * 1024
	} else if unit == "M" {
		size = size * 1024 * 1024
	}

	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(driveletter) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}
	defer file.Close() // Close the file when done

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)
	Prints.Prints = append(Prints.Prints, "------------------------")

	var count = 0
	var gap = int32(0)
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			count++
			gap = TempMBR.Partitions[i].Start + TempMBR.Partitions[i].Size
		}
	}

	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size == 0 {
			TempMBR.Partitions[i].Size = int32(size)

			if count == 0 {
				TempMBR.Partitions[i].Start = int32(binary.Size(TempMBR))
			} else {
				TempMBR.Partitions[i].Start = gap
				gap += TempMBR.Partitions[i].Size
			}

			copy(TempMBR.Partitions[i].Name[:], name)
			copy(TempMBR.Partitions[i].Fit[:], fit)
			copy(TempMBR.Partitions[i].Status[:], "0")
			copy(TempMBR.Partitions[i].Type[:], type_)
			TempMBR.Partitions[i].Correlative = int32(count + 1)
			break
		}
	}

	// Overwrite the MBR
	if err := Utilities.WriteObject(file, TempMBR, 0); err != nil {
		return
	}

	// Reopen the file for reading the updated MBR
	file, err = Utilities.OpenFile(filepath)
	if err != nil {
		return
	}
	defer file.Close() // Close the file when done

	var TempMBR2 Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR2, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR2)
	Prints.Prints = append(Prints.Prints, "======End FDISK======")
}

func DeletePartition(driveLetter string, partitionName string) {
	fmt.Println("\n====== Start DeletePartition FULL ======")
	fmt.Println("Drive letter:", driveLetter)
	fmt.Println("Partition name:", partitionName)

	// Confirmación de eliminación
	fmt.Print("¿Está seguro que desea eliminar la partición ", partitionName, "? (s/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "s" {
		fmt.Println("Eliminación de la partición cancelada")
		return
	}

	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(driveLetter) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)

	fmt.Println("-------------")

	// Find the partition by name and reset its data
	for i := 0; i < 4; i++ {
		if strings.Contains(string(TempMBR.Partitions[i].Name[:]), partitionName) {
			// Reset partition data
			TempMBR.Partitions[i] = Structs.Partition{}
			// Overwrite the MBR
			if err := Utilities.WriteObject(file, TempMBR, 0); err != nil {
				return
			}
			fmt.Println("Partition", partitionName, "deleted successfully")
			// Close bin file
			var TempMBR2 Structs.MBR
			// Read object from bin file
			if err := Utilities.ReadObject(file, &TempMBR2, 0); err != nil {
				file.Close() // Close file in case of error
				return
			}

			// Print object
			Structs.PrintMBR(TempMBR2)
			file.Close()
			fmt.Println("====== End DeletePartition full ======")
			return
		}
	}
	fmt.Println("Partition", partitionName, "not found")

	// Close bin file
	file.Close()

	fmt.Println("====== End DeletePartition FULL======")
}

//===========================================================

func Mount(driveletter string, name string) {
	Prints.Prints = append(Prints.Prints, "\n======Start MOUNT======")
	Prints.Prints = append(Prints.Prints, "Driveletter:"+driveletter)
	Prints.Prints = append(Prints.Prints, "Name:"+name)

	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(driveletter) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)
	Prints.Prints = append(Prints.Prints, "----------------")

	var index int = -1
	var count = 0
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			count++
			if strings.Contains(string(TempMBR.Partitions[i].Name[:]), name) {
				index = i
				break
			}
		}
	}

	if index != -1 {
		// Check if partition is already mounted
		if TempMBR.Partitions[index].Status[0] == '1' {
			Prints.Prints = append(Prints.Prints, "Partition is already mounted")
			return
		}
		Prints.Prints = append(Prints.Prints, "Partition found")

		Structs.PrintPartition(TempMBR.Partitions[index])
	} else {
		Prints.Prints = append(Prints.Prints, "Partition not found")
		return
	}

	// carnet: 201902888
	// id = DriveLetter + Correlative + 88

	id := strings.ToUpper(driveletter) + strconv.Itoa(count) + "88"

	copy(TempMBR.Partitions[index].Status[:], "1")
	copy(TempMBR.Partitions[index].Id[:], id)

	// Overwrite the MBR
	if err := Utilities.WriteObject(file, TempMBR, 0); err != nil {
		return
	}

	var TempMBR2 Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR2, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR2)

	// Close bin file
	defer file.Close()
	Prints.Prints = append(Prints.Prints, "======End MOUNT======")

}

func ListarParticionesMontadas(driveletter string) []string {
	var partitionList []string

	Prints.Prints = append(Prints.Prints, "\n======Listar Particiones Montadas======")
	Prints.Prints = append(Prints.Prints, "Driveletter:"+driveletter)

	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(driveletter) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return nil
	}

	Prints.Prints = append(Prints.Prints, "----------------")

	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		// Check if the partition is marked as mounted
		if TempMBR.Partitions[i].Status[0] == '1' {
			// Attempt to open the partition
			partitionPath := "./MIA/" + strings.ToUpper(driveletter) + ".disk"
			partitionFile, err := Utilities.OpenFile(partitionPath)
			if err == nil {
				// If no error, the partition is mounted
				partitionList = append(partitionList, string(TempMBR.Partitions[i].Id[:]))
				Structs.PrintPartition(TempMBR.Partitions[i])
				partitionFile.Close() // Close the partition file
			}
		}
	}

	return partitionList
}

//============================================================
//==================== UNMOUNT ===============================

func Unmount(id string) {
	fmt.Println("\n======Start UNMOUNT======")
	fmt.Println("id:", id)

	driveletter := string(id[0])

	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(driveletter) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR)

	fmt.Println("-------------")

	var index int = -1
	var count = 0
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			count++
			if strings.Contains(string(TempMBR.Partitions[i].Id[:]), id) {
				index = i
				break
			}
		}
	}

	if index != -1 {
		fmt.Println("Partition found")
		Structs.PrintPartition(TempMBR.Partitions[index])
	} else {
		fmt.Println("Partition not found")
		return
	}

	var nullId = []byte{0x00, 0x00, 0x00, 0x00}
	copy(TempMBR.Partitions[index].Status[:], "0")
	copy(TempMBR.Partitions[index].Id[:], nullId)

	fmt.Println((string(TempMBR.Partitions[index].Id[:])))

	// Overwrite the MBR
	if err := Utilities.WriteObject(file, TempMBR, 0); err != nil {
		return
	}

	var TempMBR2 Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR2, 0); err != nil {
		return
	}

	// Print object
	Structs.PrintMBR(TempMBR2)

	// Close bin file
	defer file.Close()

	fmt.Println("======End UNMOUNT======")
}

func Rep(name string, path string, id string) {

	Prints.Prints = append(Prints.Prints, "\n======Start REP======")
	Prints.Prints = append(Prints.Prints, "Name: "+name)
	Prints.Prints = append(Prints.Prints, "path: "+path)
	Prints.Prints = append(Prints.Prints, "id: "+id)
	// Open bin file
	filepath := "./MIA/" + strings.ToUpper(string(id[0])) + ".disk"
	file, err := Utilities.OpenFile(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	var TempMBR Structs.MBR
	// Read object from bin file
	if err := Utilities.ReadObject(file, &TempMBR, 0); err != nil {
		return
	}

	var totalSize int32

	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			totalSize += TempMBR.Partitions[i].Size
		}
	}

	var freezise int32
	freezise = TempMBR.MbrSize - totalSize

	Prints.Prints = append(Prints.Prints, "Total Partition Size: "+strconv.FormatInt(int64(int(freezise)), 10))

	Prints.Prints = append(Prints.Prints, "======End REP======")
}

func reporDisk(free string) string {
	graph := `
	digraph D {
	    subgraph cluster_0 {
	        bgcolor="#3731f4"
	        node [style="rounded" style=filled];
	       
	        node_A [shape=record label="MBR|Libre\n` + free + `|{PRIMARIA|{EBR|}}|PArticion"];
	    }
	}`

	return graph
}
