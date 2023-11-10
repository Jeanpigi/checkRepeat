package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Print("Ingresa la ruta de la carpeta donde se encuentran los archivos: ")
	var folderPath string
	_, err := fmt.Scan(&folderPath)
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	// Verifica si la ruta es válida
	_, err = os.Stat(folderPath)
	if err != nil {
		fmt.Println("Error: La ruta de la carpeta no es válida.")
		return
	}

	// Crear un array para almacenar nombres de los archivos
	musicMap := make(map[string][]string)

	// Recorre la carpeta en busca de archivos
	dirEntries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al escanear la carpeta:", err)
		return
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			// Obtén el nombre del archivo
			songName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

			// Agrega la canción al mapa
			musicMap[songName] = append(musicMap[songName], filepath.Join(folderPath, entry.Name()))
		}
	}

	// Encuentra canciones con nombres duplicados
	duplicatesFound := false
	for name, paths := range musicMap {
		if len(paths) > 1 {
			fmt.Printf("Nombre del archivo duplicado: %s\n", name)
			fmt.Println("Archivos:")
			for _, path := range paths {
				fmt.Println("  ", path)
			}
			duplicatesFound = true
		}
	}

	if !duplicatesFound {
		fmt.Println("No se encontraron archivos duplicados.")
	}
}
