package task8

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
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

func GetUniqueAndDoubles(done chan struct{}, chOut chan MyFileInfo) ([]MyFileInfo, []MyFileInfo) {
	var uniques []MyFileInfo
	var doubles []MyFileInfo

	chanClosed := false

	for {
		select {
		case <-done:
			if !chanClosed {
				close(chOut)
				chanClosed = true
			}
		case fi, ok := <-chOut:
			if !ok {
				if chanClosed {
					return uniques, doubles
				}
			} else {
				// fmt.Println(fi, ok, "<-- reading ...")
				if !fi.Contains(uniques) {
					uniques = append(uniques, fi)
				} else {
					doubles = append(doubles, fi)
				}
			}

		default:
		}
	}
}

func FindDuplicates(dirPath string) {
	chOut := make(chan MyFileInfo)

	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	done := make(chan struct{})
	go GetListOfFiles(done, chOut, absDirPath)

	uniques, doubles := GetUniqueAndDoubles(done, chOut)

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

func GetListOfFiles(chDone chan struct{}, chOut chan MyFileInfo, dirPath string) {
	defer close(chDone)

	var dones []chan struct{}

	fileInfo, _ := ioutil.ReadDir(dirPath)
	for _, info := range fileInfo {
		fullPath := dirPath + "/" + info.Name()
		if !info.IsDir() {
			chOut <- MyFileInfo{
				fullPath,
				info.Name(),
				info.Size(),
			}
		} else {
			done := make(chan struct{})
			dones = append(dones, done)
			go func(done chan struct{}, chOut chan MyFileInfo, fullPath string) {
				GetListOfFiles(done, chOut, fullPath)
			}(done, chOut, fullPath)
		}
	}

	for _, ch := range dones {
		select {
		case <-ch:
		}
	}
}
