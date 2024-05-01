package UtilitiesInodes

import (
	"encoding/binary"
	"fmt"
	"modulosn/Prints"
	"modulosn/Structs"
	"modulosn/Utilities"
	"os"
	"strings"
)

// login -user=root -pass=123 -id=A119
func InitSearch(path string, file *os.File, tempSuperblock Structs.Superblock) int32 {
	Prints.Prints = append(Prints.Prints, "======Start INITSEARCH======")
	Prints.Prints = append(Prints.Prints, "path: "+path)
	fmt.Println("======Start INITSEARCH======")
	fmt.Println("path:", path)
	// path = "/ruta/nueva"

	// split the path by /
	TempStepsPath := strings.Split(path, "/")
	StepsPath := TempStepsPath[1:]

	fmt.Println("StepsPath:", StepsPath, "len(StepsPath):", len(StepsPath))
	for _, step := range StepsPath {
		fmt.Println("step:", step)
		Prints.Prints = append(Prints.Prints, "step: "+step)
	}

	var Inode0 Structs.Inode
	// Read object from bin file
	if err := Utilities.ReadObject(file, &Inode0, int64(tempSuperblock.S_inode_start)); err != nil {
		return -1
	}
	Prints.Prints = append(Prints.Prints, "======End INITSEARCH======")
	fmt.Println("======End INITSEARCH======")

	return SarchInodeByPath(StepsPath, Inode0, file, tempSuperblock)
}

func pop(s *[]string) string {
	lastIndex := len(*s) - 1
	last := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return last
}

// login -user=root -pass=123 -id=A119
func SarchInodeByPath(StepsPath []string, Inode Structs.Inode, file *os.File, tempSuperblock Structs.Superblock) int32 {
	Prints.Prints = append(Prints.Prints, "======Start SARCHINODEBYPATH======")
	fmt.Println("======Start SARCHINODEBYPATH======")
	index := int32(0)
	SearchedName := strings.Replace(pop(&StepsPath), " ", "", -1)
	Prints.Prints = append(Prints.Prints, "========== SearchedName: "+SearchedName)
	fmt.Println("========== SearchedName:", SearchedName)

	// Iterate over i_blocks from Inode
	for _, block := range Inode.I_block {
		if block != -1 {
			if index < 13 {
				//CASO DIRECTO

				var crrFolderBlock Structs.Folderblock
				// Read object from bin file
				if err := Utilities.ReadObject(file, &crrFolderBlock, int64(tempSuperblock.S_block_start+block*int32(binary.Size(Structs.Folderblock{})))); err != nil {
					return -1
				}

				for _, folder := range crrFolderBlock.B_content {
					// fmt.Println("Folder found======")
					Prints.Prints = append(Prints.Prints, "Folder === Name:", string(folder.B_name[:]))
					fmt.Println("Folder === Name:", string(folder.B_name[:]), "B_inodo", folder.B_inodo)

					if strings.Contains(string(folder.B_name[:]), SearchedName) {

						fmt.Println("len(StepsPath)", len(StepsPath), "StepsPath", StepsPath)
						if len(StepsPath) == 0 {
							Prints.Prints = append(Prints.Prints, "Folder found======")
							fmt.Println("Folder found======")
							return folder.B_inodo
						} else {

							Prints.Prints = append(Prints.Prints, "NextInode======")
							fmt.Println("NextInode======")
							var NextInode Structs.Inode
							// Read object from bin file
							if err := Utilities.ReadObject(file, &NextInode, int64(tempSuperblock.S_inode_start+folder.B_inodo*int32(binary.Size(Structs.Inode{})))); err != nil {
								return -1
							}
							return SarchInodeByPath(StepsPath, NextInode, file, tempSuperblock)
						}
					}
				}

			} else {
				//CASO INDIRECTO
			}
		}
		index++
	}
	Prints.Prints = append(Prints.Prints, "======End SARCHINODEBYPATH======")
	fmt.Println("======End SARCHINODEBYPATH======")
	return 0
}

func GetInodeFileData(Inode Structs.Inode, file *os.File, tempSuperblock Structs.Superblock) string {
	Prints.Prints = append(Prints.Prints, "======Start GETINODEFILEDATA======")
	fmt.Println("======Start GETINODEFILEDATA======")
	index := int32(0)
	// define content as a string
	var content string

	// Iterate over i_blocks from Inode
	for _, block := range Inode.I_block {
		if block != -1 {
			if index < 13 {
				//CASO DIRECTO

				var crrFileBlock Structs.Fileblock
				// Read object from bin file
				if err := Utilities.ReadObject(file, &crrFileBlock, int64(tempSuperblock.S_block_start+block*int32(binary.Size(Structs.Fileblock{})))); err != nil {
					return ""
				}

				content += string(crrFileBlock.B_content[:])

			} else {
				//CASO INDIRECTO
			}
		}
		index++
	}
	Prints.Prints = append(Prints.Prints, "======End GETINODEFILEDATA======")
	fmt.Println("======End GETINODEFILEDATA======")
	return content
}
