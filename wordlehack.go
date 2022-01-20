package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

/*
posc = []
for _,guess in scores:
    count = 0
    for target in la:
        if guess == target:
            continue
        res = matchres(target,guess)
        for w in la:
            hit = True
            for rei in res:
                m = rei.match(w)
                if not m:
                    hit = False
                    break
            if hit:
                count += 1
    print(guess,count)
    posc.append((count,guess))
posc.sort()
*/

func maybefail(err error, msg string, args ...interface{}) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

type StringInt struct {
	S string
	I int
}

func guessThread(la []string, guesses chan string, out chan StringInt, wg *sync.WaitGroup) {
	defer wg.Done()
	var hasb [5]byte
	var notb [5]byte
	var posb [5]string
	for guess := range guesses {
		gb := []byte(guess)
		count := 0
		for _, target := range la {
			if target == guess {
				continue
			}
			has := hasb[:0]
			not := notb[:0]
			for gci, gc := range gb {
				if gc == target[gci] {
					posb[gci] = string(gb[gci : gci+1])
				} else if strings.IndexByte(target, gc) != -1 {
					has = append(has, gc)
					posb[gci] = fmt.Sprintf("[^%c]", gc)
				} else {
					posb[gci] = "."
					not = append(not, gc)
				}
			}
			res := strings.Join(posb[:], "")
			re := regexp.MustCompile(res)

			for _, w := range la {
				if !re.MatchString(w) {
					continue
				}
				hit := true
				for _, c := range has {
					if strings.IndexByte(w, c) == -1 {
						hit = false
						break
					}
				}
				if !hit {
					continue
				}
				for _, c := range not {
					if strings.IndexByte(w, c) != -1 {
						hit = false
						break
					}
				}
				if !hit {
					continue
				}
				count++
			}
		}
		out <- StringInt{guess, count}
		//fmt.Printf("%s\t%d\n", guess, count)
	}
}

func submitter(la []string, guesses chan string) {
	for _, guess := range la {
		guesses <- guess
	}
	close(guesses)
}

func main() {
	fin, err := os.Open("La.json")
	maybefail(err, "La.json: %v", err)
	dec := json.NewDecoder(fin)
	var la []string
	err = dec.Decode(&la)
	maybefail(err, "La.json []string: %v", err)
	// var hasb [5]byte
	// var notb [5]byte
	// var posb [5]string
	wg := sync.WaitGroup{}
	guesses := make(chan string, 20)
	results := make(chan StringInt, 20)
	for i := 0; i < runtime.NumCPU(); i++ {
		go guessThread(la, guesses, results, &wg)
		wg.Add(1)
	}
	go submitter(la, guesses)
	/*
		gb := []byte(guess)
		count := 0
		for ti, target := range la {
			if ti == gi {
				continue
			}
			has := hasb[:0]
			not := notb[:0]
			for gci, gc := range gb {
				if gc == target[gci] {
					posb[gci] = string(gb[gci : gci+1])
				} else if strings.IndexByte(target, gc) != -1 {
					has = append(has, gc)
					posb[gci] = fmt.Sprintf("[^%c]", gc)
				} else {
					posb[gci] = "."
					not = append(not, gc)
				}
			}
			res := strings.Join(posb[:], "")
			re := regexp.MustCompile(res)

			for _, w := range la {
				if !re.MatchString(w) {
					continue
				}
				hit := true
				for _, c := range has {
					if strings.IndexByte(w, c) == -1 {
						hit = false
						break
					}
				}
				if !hit {
					continue
				}
				for _, c := range not {
					if strings.IndexByte(w, c) != -1 {
						hit = false
						break
					}
				}
				if !hit {
					continue
				}
				count++
			}
		}
	*/
	for xr := range results {
		guess := xr.S
		count := xr.I
		fmt.Printf("%s\t%d\n", guess, count)
	}
}
