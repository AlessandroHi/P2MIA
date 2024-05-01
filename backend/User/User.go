package User

import (
	//  "os"
	"encoding/binary"
	"fmt"
	"modulosn/Global"
	"modulosn/Prints"
	"modulosn/Structs"
	"modulosn/Utilities"
	"modulosn/UtilitiesInodes"
	"strings"
)

// login -user=root -pass=123 -id=A119
func Login(user string, pass string, id string) {
	Prints.Prints = append(Prints.Prints, "======Start LOGIN======")
	Prints.Prints = append(Prints.Prints, "User: "+user)
	Prints.Prints = append(Prints.Prints, "Pass: "+pass)
	Prints.Prints = append(Prints.Prints, "Id: "+id)

	if Global.Usuario.Status {
		fmt.Println("User already logged in")
		Prints.Prints = append(Prints.Prints, "User already logged in")
		return
	}

	var login bool = false
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
	Prints.Prints = append(Prints.Prints, "---------------------")
	fmt.Println("-------------")

	var index int = -1
	// Iterate over the partitions
	for i := 0; i < 4; i++ {
		if TempMBR.Partitions[i].Size != 0 {
			if strings.Contains(string(TempMBR.Partitions[i].Id[:]), id) {
				Prints.Prints = append(Prints.Prints, "Partition found")
				fmt.Println("Partition found")
				if strings.Contains(string(TempMBR.Partitions[i].Status[:]), "1") {
					Prints.Prints = append(Prints.Prints, "Partition is mounted")
					fmt.Println("Partition is mounted")
					index = i
				} else {
					Prints.Prints = append(Prints.Prints, "Partition is not mounted")
					fmt.Println("Partition is not mounted")
					return
				}
				break
			}
		}
	}

	if index != -1 {
		Structs.PrintPartition(TempMBR.Partitions[index])
	} else {
		Prints.Prints = append(Prints.Prints, "Partition not found")
		fmt.Println("Partition not found")
		return
	}

	var tempSuperblock Structs.Superblock
	// Read object from bin file
	if err := Utilities.ReadObject(file, &tempSuperblock, int64(TempMBR.Partitions[index].Start)); err != nil {
		return
	}

	// initSearch /users.txt -> regresa no Inodo
	// initSearch -> 1
	indexInode := UtilitiesInodes.InitSearch("/users.txt", file, tempSuperblock)

	// indexInode := int32(1)

	var crrInode Structs.Inode
	// Read object from bin file
	if err := Utilities.ReadObject(file, &crrInode, int64(tempSuperblock.S_inode_start+indexInode*int32(binary.Size(Structs.Inode{})))); err != nil {
		return
	}

	// read file data
	data := UtilitiesInodes.GetInodeFileData(crrInode, file, tempSuperblock)
	Prints.Prints = append(Prints.Prints, "Fileblock------------")
	fmt.Println("Fileblock------------")
	// Dividir la cadena en líneas
	lines := strings.Split(data, "\n")

	// login -user=root -pass=123 -id=A119

	// Iterar a través de las líneas
	for _, line := range lines {
		// Imprimir cada línea
		// fmt.Println(line)
		words := strings.Split(line, ",")

		if len(words) == 5 {
			if (strings.Contains(words[3], user)) && (strings.Contains(words[4], pass)) {
				login = true
				break
			}
		}
	}

	// Print object
	fmt.Println("Inode", crrInode.I_block)

	// Close bin file
	defer file.Close()

	if login {
		Prints.Prints = append(Prints.Prints, "User logged in")
		fmt.Println("User logged in")
		Global.Usuario.ID = id
		Global.Usuario.Status = true
	}
	Prints.Prints = append(Prints.Prints, "======End LOGIN======")
	fmt.Println("======End LOGIN======")
}

func Logout() {

	Prints.Prints = append(Prints.Prints, "======Start LOGOUT======")
	fmt.Println("======Start LOGOUT======")
	if Global.Usuario.Status {
		Global.Usuario.ID = ""
		Global.Usuario.Status = false
		Prints.Prints = append(Prints.Prints, "User logged out")
		fmt.Println("User logged out")
	} else {
		Prints.Prints = append(Prints.Prints, "No user logged in")
		fmt.Println("No user logged in")
	}
	Prints.Prints = append(Prints.Prints, "======End LOGOUT======")
	fmt.Println("======End LOGOUT======")
}
