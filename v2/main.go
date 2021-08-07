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

type MyFileInfo struct {
	FilePath string
	name     string
	size     int64
}

func (f *MyFileInfo) Contains(list []MyFileInfo) bool {
	for _, val := range list {
		if val.size == f.size && val.name == f.name {
			return true
		}
	}

	return false
}

func FindDuplicates(dirPath string) {
	ch := make(chan MyFileInfo, 100)

	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	all, err := GetListOfFiles(ch, absDirPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	close(ch)

	//for v := <-ch {
	//
	//}

	var uniques []MyFileInfo
	var doubles []MyFileInfo
	for _, fi := range all {
		// fmt.Println(fi.FilePath)
		if !fi.Contains(uniques) {
			uniques = append(uniques, fi)
		} else {
			doubles = append(doubles, fi)
		}
	}

	fmt.Println("Doubles:")
	for _, double := range doubles {
		fmt.Println(double)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Unique:")
	for _, unique := range uniques {
		fmt.Println(unique)
	}
}

func GetListOfFiles(ch chan MyFileInfo, dirPath string) ([]MyFileInfo, error) {
	var files []MyFileInfo
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, MyFileInfo{
				path,
				info.Name(),
				info.Size(),
			})
		}
		return nil
	})
	return files, err
}

func main() {
	FindDuplicates("./files")
}
