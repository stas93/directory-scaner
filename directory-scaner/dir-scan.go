package directory_scaner

import (
	"io/ioutil"
	"os"
	"sync"
)

func s(d string) (result int64, need []string) {
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return
	}
	for _, f := range files {
		if f.IsDir() {
			need = append(need, d+"/"+f.Name())
		} else {
			result += f.Size()
		}
	}
	return
}
func Scan(d string) int {
	result, need := s(d)
	for len(need) > 0 {
		resTMP, needTMP := s(need[0])
		need = need[1:]
		result += resTMP
		need = append(need, needTMP...)
	}
	return int(result) / 1000000
}

/////////////////////////////////////////////////////////////////////////
func s22(d string, result chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println()
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return
	}
	wg2 := sync.WaitGroup{}
	for _, f := range files {
		if f.IsDir() {
			wg2.Add(1)
			go s22(d+"/"+f.Name(), result, &wg2)
		} else {
			result <- f.Size()
		}
	}
	wg2.Wait()
}
func s2(d string, result chan int64) {
	defer close(result)
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			go s22(d+"/"+f.Name(), result, &wg)
		} else {
			result <- f.Size()
		}
	}
	wg.Wait()
}
func Scan2(d string) int {
	result := make(chan int64)
	go s2(d, result)
	var resTMP int64
	for r := range result {
		resTMP += r
	}
	return int(resTMP) / 1000000
}

////////////////////////////
func Scan3(d string) int {
	var resTMP int64
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	workList := make(chan string)
	wg.Add(1)
	go func() {
		workList <- d
	}()
	for i := 0; i < 1; i++ {
		go func() {
			for work := range workList {
				files, err := ioutil.ReadDir(work)
				if err == nil {
					for _, f := range files {
						if f.IsDir() {
							wg.Add(1)
							go func(work string, f os.FileInfo) {
								workList <- work + "/" + f.Name()
							}(work, f)
						} else {
							mu.Lock()
							resTMP += f.Size()
							mu.Unlock()
						}
					}
				}
				wg.Done()
			}
		}()
	}
	wg.Wait()
	go func() {
		close(workList)
	}()
	return int(resTMP) / 1000000
}
