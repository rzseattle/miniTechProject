package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/test", test)
	http.HandleFunc("/file", file)
	http.HandleFunc("/added", added)
	fmt.Println("Server listening on port: 8090")
	http.ListenAndServe(":8090", nil)

}

func added(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	if req.Method == http.MethodOptions {
		return
	}

	file := req.URL.Query().Get("file")

	f_path := path.Join("front/public/fruits", file)

	fmt.Println("make", "test_external", "FILE="+f_path)

	cmd := exec.Command("make", "test_external", "FILE="+f_path)
	cmd.Env = os.Environ()
	cmd.Dir = "../"
	cmdOutput, err := cmd.Output()
	if err != nil {
		retJsonString, _ := json.Marshal(struct{ Error string }{Error: err.Error()})
		fmt.Fprintf(w, "%s", retJsonString)
		fmt.Print(err)
		return
	}

	ret := string(cmdOutput)

	parts := strings.Split(ret, "-->")
	found := strings.Trim(parts[len(parts)-1], " \n")
	fmt.Println("------------")
	fmt.Println(found)
	fmt.Println("------------")

	price := findPrice(found)

	retJson := struct {
		Fruit string  `json:"fruit"`
		Price float64 `json:"price"`
	}{Fruit: found, Price: price}
	retJsonString, _ := json.Marshal(retJson)

	fmt.Printf("cmdRet: %s\n", cmdOutput)

	//fmt.Fprintf(w, "%s", string(cmdOutput))
	//fmt.Fprintf(w, "%s", ret)
	fmt.Fprintf(w, "%s", retJsonString)

}

func file(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	filePath := req.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "Missing 'file' parameter", http.StatusBadRequest)
		return
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(fileContent)
}

func test(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	cmd := exec.Command("make", "test_repo")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "YOLO_VERBOSE=False")
	cmd.Dir = "../"
	cmdOutput, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	ret := string(cmdOutput)
	var cmdRet CmdRet
	err = json.Unmarshal([]byte(ret), &cmdRet)
	if err != nil {
		fmt.Fprintf(w, "Error unmarshalling JSON: %s\n", err)
		return
	}

	fmt.Printf("cmdRet: %v\n", cmdRet)

	//fmt.Fprintf(w, "%s", string(cmdOutput))
	fmt.Fprintf(w, "%s", ret)
}

type CmdRet []struct {
	File  string `json:"file"`
	Probs []struct {
		Confidence float64 `json:"confidence"`
		Label      string  `json:"label"`
	} `json:"probs"`
}

func findPrice(fruit string) float64 {
	pricesFile := "../assets/prices.csv"

	file, err := os.Open(pricesFile)
	if err != nil {
		fmt.Printf("Error opening prices file: %s\n", err)
		return 0.0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}
		if parts[0] == fruit {
			price, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Printf("Error parsing price: %s\n", err)
				return 0.0
			}
			return price
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading prices file: %s\n", err)
	}

	return -1.0
}
