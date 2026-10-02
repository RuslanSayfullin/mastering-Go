package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// вызов os.OpenFile() создает файл журнала для записи,
	// если он еще не существует, или же открывает его для записи
	// путем добавления новых данных в конце (ос.O_APPEND)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

}
