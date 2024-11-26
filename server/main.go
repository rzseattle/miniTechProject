package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
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
	http.HandleFunc("/example", example)

	http.ListenAndServe(":8090", nil)
}

func example(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	label := req.URL.Query().Get("label")

	f_path := path.Join("../train", label)
	files, err := os.ReadDir(f_path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading directory: %s", err), http.StatusInternalServerError)
		return
	}

	if len(files) == 0 {
		http.Error(w, "No files found in directory", http.StatusNotFound)
		return
	}

	fileContent, err := os.ReadFile(path.Join("../train", label, files[0].Name()))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(fileContent)
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
