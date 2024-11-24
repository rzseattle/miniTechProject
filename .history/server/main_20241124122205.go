package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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

	http.ListenAndServe(":8090", nil)
}

func test(w http.ResponseWriter, req *http.Request) {

	cmd := exec.Command("python3.11", "test.py")
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

	fmt.Fprintf(w, "%s", string(cmdOutput))
}

type CmdRet []struct {
	File  string `json:"file"`
	Probs []struct {
		Confidence float64 `json:"confidence"`
		Label      string  `json:"label"`
	} `json:"probs"`
}