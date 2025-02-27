package database

import (
	"github.com/go-fast-cdn/models"
)

func AddDoc(fileName string, fileHashBuffer []byte) (string, bool) {
	var doc models.Doc
	doc.FileName = fileName
	doc.Checksum = fileHashBuffer

	var entries []models.Doc

	DB.First(&entries, "checksum = $1", string(fileHashBuffer))
	if len(entries) == 0 {
		DB.Create(&doc)
		return fileName, false
	} else {
		return entries[0].FileName, true
	}
}


func DeleteDoc(fileName string) (string, bool) {
	var doc models.Doc

	result := DB.Where("file_name = ?", fileName).First(&doc)

	if result.Error == nil {
		DB.Delete(&doc)
		return fileName, true
	} else {
		return "", false
	}
}
