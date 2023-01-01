package makecmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var FileExtension = ".go"

func FormatFilename(name string) string {
	var mySli []string

	for _, val := range strings.Split(name, "-") {
		mySli = append(mySli, strings.Title(val))
	}

	return strings.Join(mySli, "")
}

func WriteFile(path string, content string, args []string) {
	pathSet := path + "s"

	_, err := os.Stat(pathSet)

	if os.IsNotExist(err) {
		errDir := os.Mkdir(pathSet, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	fileName := strings.Join(args, " ")

	f, err := os.Create(filepath.Join(pathSet, filepath.Base(fileName+FileExtension)))
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

	fmt.Printf("%v %v created successfully\n", fileName, path)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
