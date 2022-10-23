package main

import (
	"io"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"os/exec"
)

// run action shell script located at ./src/notification.sh
func sendNotification() {
	exec.Command("/bin/sh", "./src/notification.sh").Output()
}

func main() {
		PORT := 8080
		http.HandleFunc("/", request)

		log.Println("Started Server on port", PORT)
		err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

// actions to run when a user connects
func request(w http.ResponseWriter, r *http.Request) {
	programName := os.Args[1]
	io.WriteString(w, readFile(programName))
	log.Println("Received Request at", programName)
	sendNotification()
}

func readFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "Error 404, File not Found!"
    }
    
	return string(data)
}
