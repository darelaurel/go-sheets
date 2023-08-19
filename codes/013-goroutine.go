package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

/**
* WaitGroup is used to wait for all the goroutines launched here to finish.
**/

/*
**Goroutine execute code de facon asynchrone
Go routine is a non-blocking function which starts executing on it’s own once called.
GO routine can be understood as a light weight thread
All the program contains at least one Go routine, it is the Main Go Routine. One thing that is to be known is that, all the other GO routines are actually running inside this main GO routine. So if this main GO routine is terminated then all the other GO routines also gets terminated.
*/

var wg sync.WaitGroup

func fetchUrl(url string) {
	defer wg.Done()
	start_time := time.Now()
	resp, err := http.Get(url)
	end_time := time.Since(start_time)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Print("%Site %s length %d Processing time %s \n", url, len(body), end_time)
}

func main() {
/**
**Dans ce bloc main, fetchUrl asynchrone, une fois lancé le programme n'attend pas sa fin 
** pour continuer donc le waitGroup informe le nombre d'element en pending a terminer
**avant la fin de lexecution du programme
*/

	urls := []string{"right-com.com", "opensi.co"}
	wg.Add(2)

	for _, x := range urls {
		go fetchUrl(x)
	}
	wg.Wait()
}
