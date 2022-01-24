package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const gitPath = ""

func main() {
	http.HandleFunc("/", pullRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func pullRequest(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = gitPath
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "err: %s", err.Error())
		return
	}
	fmt.Fprint(w, "pull successfully!")
}
