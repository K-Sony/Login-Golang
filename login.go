package main


import (
  "net/http"
  "fmt"
  "os"
  "io"
  "encoding/csv"
)

var err error
	
func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")
	file, err := os.Open("Files/user.csv")
  if err != nil {
  http.Error(res, "Server error", 500)
  fmt.Println("Error:", err)
    return
  }
  defer file.Close()
  reader := csv.NewReader(file)
  lineCount := 0
  for {
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			res.Write([]byte("Invaid User or password"))
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		if( lineCount != 0){
			if(username == record[0] && password == record[1]){
				res.Write([]byte("Hello  " + username))
				break
			}
			
		}
		
		lineCount += 1
	}

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}