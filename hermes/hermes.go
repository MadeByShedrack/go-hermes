package hermes

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type student struct {
	studentID int
	name      string
	gpa       float64
	subjects  []string
}

var myStudent []student
var studentMutex = &sync.Mutex{}

func getStudents(w http.ResponseWriter, myRequest *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func addStudents(w http.ResponseWriter, studentRequest *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent student
	err := json.NewDecoder(studentRequest.Body).Decode(&newStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	studentMutex.Lock()
	myStudent = append(myStudent, newStudent)
	studentMutex.Unlock()
	json.NewEncoder(w).Encode(newStudent)
}

func Hermes() {
	http.HandleFunc("/students", addStudents)
	http.HandleFunc("/students/add", getStudents)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
