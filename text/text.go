package text

import (
	"bufio"
	"fmt"
	"os"
)

//ReadText reads the .txt file
//and adds it back into the array
//one line at a time
func ReadText(fpath string) (texts []string) {
	file, error := os.Open(fpath)
	if error != nil {
		fmt.Println("error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	return
}
