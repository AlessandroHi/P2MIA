package Analyzer

import (
	"bufio"
	"flag"
	"fmt"
	"modulosn/DiskManagement"
	"modulosn/FileManager"
	"modulosn/FileSystem"
	"modulosn/Prints"
	"modulosn/User"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

func getCommandAndParams(input string) (string, string) {
	parts := strings.Fields(input)
	if len(parts) > 0 {
		command := strings.ToLower(parts[0])
		params := strings.Join(parts[1:], " ")
		return command, params
	}
	return "", input
}

func Analyze(input string) []string {

	command, params := getCommandAndParams(input)

	AnalyzeCommnad(command, params)

	return Prints.Prints

}

func AnalyzeCommnad(command string, params string) {
	if strings.Contains(command, "mkdisk") {
		fn_mkdisk(params)
	} else if strings.Contains(command, "rmdisk") {
		fn_rmdisk(params)
	} else if strings.Contains(command, "unmount") {
		fn_unmount(params)
	} else if strings.Contains(command, "fdisk") {
		fn_fdisk(params)
	} else if strings.Contains(command, "mount") {
		fn_mount(params)
	} else if strings.Contains(command, "mkfs") {
		fn_mkfs(params)
	} else if strings.Contains(command, "login") {
		fn_login(params)
	} else if strings.Contains(command, "logout") {
		fn_logout()
	} else if strings.Contains(command, "mkusr") {
		fn_mkusr(params)
	} else if strings.Contains(command, "pause") {
		fn_pause()
	} else if strings.Contains(command, "execute") {
		fn_execute(params)
	} else {
		Prints.Prints = append(Prints.Prints, "Error: Command not found")
	}
}

// ================== MKDISK ============================
func fn_mkdisk(params string) {
	// Define flags
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	fit := fs.String("fit", "ff", "Ajuste")
	unit := fs.String("unit", "m", "Unidad")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit":
			fs.Set(flagName, flagValue)
		default:
			Prints.Prints = append(Prints.Prints, "Error: Flag not found")
			fmt.Println("Error: Flag not found")
			return
		}
	}

	// Call the function
	DiskManagement.Mkdisk(*size, *fit, *unit)

}

// ===================================================
// ================== RMDISK =========================
func fn_rmdisk(params string) {
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "letter")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	DiskManagement.Rmdisk(*driveletter)

}

// ===================================================
//================== FDISK =========================
func fn_fdisk(input string) {
	// Define flags
	fs := flag.NewFlagSet("fdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")
	unit := fs.String("unit", "m", "Unidad")
	type_ := fs.String("type", "p", "Tipo")
	fit := fs.String("fit", "wf", "Ajuste")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit", "driveletter", "name", "type", "delete", "add":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}
	// Call the function
	DiskManagement.Fdisk(*size, *driveletter, *name, *unit, *type_, *fit)
}

// ===================================================
//================== MOUNT =========================
func fn_mount(input string) {
	// Define flags
	fs := flag.NewFlagSet("mount", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter", "name":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Mount(*driveletter, *name)
}

// ===================================================
//================== UNMOUNT =========================
func fn_unmount(input string) {
	// Define flags
	fs := flag.NewFlagSet("unmount", flag.ExitOnError)
	id := fs.String("id", "", "letter")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Unmount(strings.ToUpper(*id))
}

func fn_mkusr(input string) {
	// Define flags
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	user := fs.String("user", "", "Usuario")
	pass := fs.String("pass", "", "Contraseña")
	grp := fs.String("grp", "", "grupo")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "user", "pass", "grp":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	FileManager.Mkusr(*user, *pass, *grp)

}

func fn_logout() {
	User.Logout()
}

func fn_login(input string) {
	// Define flags
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	user := fs.String("user", "", "Usuario")
	pass := fs.String("pass", "", "Contraseña")
	id := fs.String("id", "", "Id")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "user", "pass", "id":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	User.Login(*user, *pass, *id)

}

func fn_mkfs(input string) {
	// Define flags
	fs := flag.NewFlagSet("mkfs", flag.ExitOnError)
	id := fs.String("id", "", "Id")
	type_ := fs.String("type", "", "Tipo")
	fs_ := fs.String("fs", "2fs", "Fs")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id", "type", "fs":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	FileSystem.Mkfs(*id, *type_, *fs_)

}

func fn_pause() {
	// Define flags
	fs := flag.NewFlagSet("pause", flag.ExitOnError)

	// Parse the flags
	fs.Parse(os.Args[1:])
	fmt.Println("Presione ENTER para continuar...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	fmt.Println("Continuando la ejecución...")

}

func fn_execute(input string) {
	fs := flag.NewFlagSet("execute", flag.ExitOnError)
	pathScrti := fs.String("path", "", "path")

	// Parse the flags
	fs.Parse(strings.Fields(input))

	// Abrir el archivo para lectura
	file, err := os.Open(*pathScrti)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Iterar sobre cada línea del archivo
	for scanner.Scan() {
		linea := scanner.Text()

		// Si la línea está vacía, continuar con la siguiente línea
		if len(strings.TrimSpace(linea)) == 0 {
			continue
		}

		// Si la línea es un comentario, ignorarla y continuar con la siguiente línea
		if strings.HasPrefix(strings.TrimSpace(linea), "#") {
			continue
		}

		// Si la línea contiene un comentario (#), eliminar solo el comentario y continuar con la ejecución
		if idx := strings.Index(linea, "#"); idx != -1 {
			linea = linea[:idx]
		}

		// Si la línea no está vacía ni es un comentario, ejecutarla y mostrar el resultado
		command, params := getCommandAndParams(linea)
		fmt.Println("\nEjecutando:", linea)
		fmt.Println("Command: ", command, "Params: ", params)

		AnalyzeCommnad(command, params)
	}

	// Verificar si hubo errores durante el escaneo
	if err := scanner.Err(); err != nil {
		fmt.Println("Hubo un error:", err)
	}
}
