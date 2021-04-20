package main

import (
	"crypto/rand"
	"flag"
	"log"
	"math/big"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Git struct {
	StartTime       time.Time
	EndTime         time.Time
	NumberOfCommits int
	sync.Mutex
}

// func (g *Git) MakeGitCommits() {
// 	td := g.StartTime.Hour() - g.EndTime.Hour()
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < g.NumberOfCommits; i++ {
// 		t := g.StartTime.Add(time.Hour * time.Duration(td*(i+1)))
// 		wg.Add(1)
// 		g.commit(t, &wg)
// 	}
// 	wg.Wait()
// }

func (g *Git) AutoGitCommits() {
	// get time one year back
	t := time.Now().Add(-time.Duration(24*365) * time.Hour)
	// now randomly add days to this time.
	log.Println("number of commits", g.NumberOfCommits)
	for i := 0; i < g.NumberOfCommits; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(365))
		if err != nil {
			log.Println("error getting random integer ", err)
		}
		tt := t.Add(time.Duration(nBig.Int64()*24) * time.Hour)
		g.commit(tt)
	}

}

func (g *Git) commit(t time.Time) {
	// text := fmt.Sprintf("%v-%v-%v", t.Hour(), t.Minute(), t.Second())
	name := "abcd.txt"
	text := t.String()
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("error opening file", err)
		return
	}
	f.WriteString(text)
	f.Write([]byte("\n"))
	f.Close()
	add := exec.Command("git", "add", ".")
	if err := add.Run(); err != nil {
		log.Println("error adding files to tracking git", err)
	}
	args := []string{"commit", "-am", text, "--date", t.String()}
	cmd := exec.Command("git", args...)
	if err := cmd.Run(); err != nil {
		log.Println("errorrr", err.Error())
	}
}

func main() {
	var nc int
	flag.IntVar(&nc, "nc", 50, "number of commits")
	flag.Parse()
	log.Println(nc)
	// two dates between which we need to make commits, and the number of commits in that range
	g := Git{
		NumberOfCommits: nc,
	}
	g.AutoGitCommits()

}
