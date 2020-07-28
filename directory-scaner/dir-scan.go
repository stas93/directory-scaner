package directory_scaner

import (
	"io/ioutil"
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
	result := make(chan int64)
	workList := make(chan string, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		workList <- d
	}()
	for i := 0; i < 2; i++ {
		go func() {
			for work := range workList {
				go func(work string) {
					files, err := ioutil.ReadDir(work)
					if err != nil {
						panic("err")
					}
					for _, f := range files {
						if f.IsDir() {
							wg.Add(1)
							workList <- d + "/" + f.Name()
						} else {
							result <- f.Size()
						}
					}
				}(work)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(workList)
	}()
	var resTMP int64
	for r := range result {
		resTMP += r
	}
	return int(resTMP) / 1000000
}
