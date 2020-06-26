package directory_scaner

import (
	"io/ioutil"
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
func s2(d string, result chan int64) {
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return
	}
	var rrrr int64
	for _, f := range files {
		if f.IsDir() {
			//dd <- d + "/" + f.Name()
		} else {
			rrrr += f.Size()
		}
	}
	result <- rrrr
}
func Scan2(d string) int {
	result := make(chan int64)
	need := make(chan string)
	go func() {
		need <- d
	}()
	//wg := sync.WaitGroup{}wg.Add(1)wg.Wait()
	go s2(<-need, result)
	var resTMP int64
	for r := range result {
		resTMP += r
		close(result)
	}
	return int(resTMP) / 1000000
}
