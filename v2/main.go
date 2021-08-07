package main

import (
	"fmt"
	"io/ioutil"
	"log"
	// _ "github.com/gorilla/websocket"
	// _ "github.com/valyala/fasthttp"
	_ "os"
	"path/filepath"
	_ "path/filepath"
	"sync"
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
	var ch []chan MyFileInfo

	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	go GetListOfFiles(ch, absDirPath)

	var uniques []MyFileInfo
	var doubles []MyFileInfo

	for {
		fi, ok := <-Merge(ch...)
		if ok == false {
			fmt.Println(fi, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			if !fi.Contains(uniques) {
				uniques = append(uniques, fi)
			} else {
				doubles = append(doubles, fi)
			}
		}
	}

	//for fi := range ch {
	//	// fmt.Println(fi.FilePath)
	//	if !fi.Contains(uniques) {
	//		uniques = append(uniques, fi)
	//	} else {
	//		doubles = append(doubles, fi)
	//	}
	//}

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

// Merge Use fan-in pattern
func Merge(cs ...chan MyFileInfo) chan MyFileInfo {
	var wg sync.WaitGroup

	out := make(chan MyFileInfo)

	// Запускаем send goroutine
	// для каждого входящего канала в cs.
	// send копирует значения из c в out
	// до тех пор пока c не закрыт, затем вызываем wg.Done.
	send := func(c <-chan MyFileInfo) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Запускаем goroutine чтобы закрыть out
	// когда все send goroutine выполнены
	// Это должно начаться после вызова wg.Add.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func GetListOfFiles(chOut []chan MyFileInfo, dirPath string) {
	chIn := make(chan MyFileInfo)
	chOut = append(chOut, chIn)

	go func(chIn chan MyFileInfo, chOut []chan MyFileInfo, dirPath string) {
		fs, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range fs {
			fullPath := dirPath + "/" + val.Name()
			if val.IsDir() {
				// go GetListOfFiles(chOut, fullPath)
			} else {
				chIn <- MyFileInfo{
					fullPath,
					val.Name(),
					val.Size(),
				}
			}
		}

		close(chIn)
	}(chIn, chOut, dirPath)

	//err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
	//	if !info.IsDir() {
	//
	//		files = append(files, MyFileInfo{
	//			path,
	//			info.Name(),
	//			info.Size(),
	//		})
	//	}
	//	return nil
	//})
	//return files, err
}

func main() {
	FindDuplicates("./files")
}
