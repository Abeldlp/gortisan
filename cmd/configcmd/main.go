package configcmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteFile(path string, content string) {
	_, err := os.Stat("config")

	if os.IsNotExist(err) {
		errDir := os.Mkdir("config", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(filepath.Join("config", filepath.Base("database.go")))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Printf("%v config created successfully\n", path)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
