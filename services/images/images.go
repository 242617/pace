package images

import (
	"io"
	"log"
	"os"

	"github.com/242617/pace/config"
	"github.com/242617/pace/utils"
)

var g = utils.NewGenerator()

func Upload(reader io.Reader) (string, error) {

	name := g.Generate(32)
	filename := config.ImagesFilePath + "/" + name

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("err", err)
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		log.Println("err", err)
		return "", err
	}

	return name, nil

}
