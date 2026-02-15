package files

import (
	"fmt"
	"os"
	"strings"

	"jsonStorage/pkg/consts"
)

func ReadFile(fileName string) ([]byte, error) {
	if strings.HasSuffix(fileName, ".json") {
		return nil, fmt.Errorf(consts.ERR_JSON_FILE)
	}
	
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func WriteFile(content []byte, name string) error{
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	
	return nil
}

func CreateFile(name string) error {
	_, err := os.Create(name)
	if err != nil {
		return err
	}
	return nil
}

func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}