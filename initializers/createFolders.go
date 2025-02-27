package initializers

import (
	"fmt"
	"os"

	"github.com/go-fast-cdn/util"
)

func CreateFolders() {
	uploadsFolder := fmt.Sprintf("%v/uploads", util.ExPath)
	os.Mkdir(uploadsFolder, 0755)
	os.Mkdir(fmt.Sprintf("%v/docs", uploadsFolder), 0755)
	os.Mkdir(fmt.Sprintf("%v/images", uploadsFolder), 0755)
}
