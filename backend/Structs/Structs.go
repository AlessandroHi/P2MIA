package Structs

import (
	"fmt"
	"modulosn/Prints"
)

//  =================== ESTRUCTURAS PARA DISCO ===========================
//  =================== MBR - INICO DEL DISCO ============================

type MBR struct {
	MbrSize      int32
	CreationDate [10]byte
	Signature    int32
	Fit          [1]byte
	Partitions   [4]Partition
}

func PrintMBR(data MBR) {
	fmt.Println(fmt.Sprintf("CreationDate: %s, fit: %s, size: %d", string(data.CreationDate[:]), string(data.Fit[:]), data.MbrSize))
	Prints.Prints = append(Prints.Prints, fmt.Sprintf("CreationDate: %s, fit: %s, size: %d", string(data.CreationDate[:]), string(data.Fit[:]), data.MbrSize))
	for i := 0; i < 4; i++ {
		PrintPartition(data.Partitions[i])
	}
}

//  ====================== PARTITION ============================

type Partition struct {
	Status      [1]byte
	Type        [1]byte //P - E
	Fit         [1]byte // B - F - W
	Start       int32
	Size        int32
	Name        [16]byte
	Correlative int32
	Id          [4]byte
}

func PrintPartition(data Partition) {
	Prints.Prints = append(Prints.Prints, fmt.Sprintf("Name: %s, type: %s, start: %d, size: %d, status: %s, id: %s", string(data.Name[:]), string(data.Type[:]), data.Start, data.Size, string(data.Status[:]), string(data.Id[:])))
}

// ========================= EBR - PARA LOGICA ================================

type EBR struct {
	Mount  [1]byte
	Fit    [1]byte // B - F - W
	Start  int32
	Part_s int32
	Next   int32
	Name   [16]byte
}

//  =============== ESTRUCUTURAS PARA CARPETAS Y ARCHIVOS =============
//  ================== SUPER BLOQUE ===================================

type Superblock struct {
	S_filesystem_type   int32
	S_inodes_count      int32
	S_blocks_count      int32
	S_free_blocks_count int32
	S_free_inodes_count int32
	S_mtime             [17]byte
	S_umtime            [17]byte
	S_mnt_count         int32
	S_magic             int32
	S_inode_size        int32
	S_block_size        int32
	S_fist_ino          int32
	S_first_blo         int32
	S_bm_inode_start    int32
	S_bm_block_start    int32
	S_inode_start       int32
	S_block_start       int32
}

//  ======================= INODOS =============================

type Inode struct {
	I_uid   int32
	I_gid   int32
	I_size  int32
	I_atime [17]byte
	I_ctime [17]byte
	I_mtime [17]byte
	I_block [15]int32 // 13 simple -  14 doble - 15 triple
	I_type  [1]byte   // 1 archivo - 0 carpeta
	I_perm  [3]byte
}

//  =================== BLOQUE DE CARPETAS ================================
type Folderblock struct {
	B_content [4]Content
}

type Content struct {
	B_name  [12]byte
	B_inodo int32
}

//  =================== BLOQUE DE ARCHIVOS ================================
type Fileblock struct {
	B_content [64]byte
}

//  =============================================================

type Pointerblock struct {
	B_pointers [16]int32
}

//  ============== JOURNALING ==================================

type Journaling struct {
	Size      int32
	Ultimo    int32
	Contenido [50]Content_J
}

type Content_J struct {
	Operation [10]byte
	Path      [100]byte
	Content   [100]byte
	Date      [17]byte
}
