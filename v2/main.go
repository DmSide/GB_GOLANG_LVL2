package main

import (
	"fmt"
	"os"
	// _ "github.com/gorilla/websocket"
	// _ "github.com/valyala/fasthttp"
	_ "os"
	"path/filepath"
	_ "path/filepath"
)

//В качестве завершающего задания нужно выполнить программу поиска дубликатов файлов. Дубликаты файлов - это файлы, которые совпадают по имени файла и по его размеру. Нужно написать консольную программу, которая проверяет наличие дублирующихся файлов.
//Программа должна работать на локальном компьютере и получать на вход путь до директории. Программа должна вывести в стандартный поток вывода список дублирующихся файлов, которые находятся как в директории, так и в поддиректориях директории,  переданной через аргумент командной строки. Данная функция должна работать эффективно при помощи распараллеливания программы
//Программа должна принимать дополнительный ключ - возможность удаления обнаруженных дубликатов файлов после поиска. Дополнительно нужно придумать, как обезопасить пользователей от случайного удаления файлов. В качестве ключей желательно придерживаться общепринятых практик по использованию командных опций.
//Критерии приемки программы:
//Программа компилируется
//Программа выполняет функциональность, описанную выше.
//Программа покрыта тестами
//Программа содержит документацию и примеры использования
//Программа обладает флагом “-h/--help” для краткого объяснения функциональности
//Программа должна уведомлять пользователя об ошибках, возникающих во время выполнения
//Дополнительно можете выполнить следующие задания:
//Написать программу которая по случайному принципу генерирует копии уже имеющихся файлов, относительно указанной директории
//Сравнить производительность программы в однопоточном и многопоточном режимах

func FindDuplicates(dirPath string) {
	ch := make(chan int, 100)

	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	GetListOfFiles(ch, absDirPath)

	//for v := <-ch {
	//
	//}

}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func GetListOfFiles(ch chan, dirPath string) {
	FilePathWalkDir(dirPath)
}

func main() {
	FindDuplicates("../files")
}
