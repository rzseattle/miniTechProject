package main

import (
	"fmt"
	"net/http"
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

	cmd := exec.Command("export YOLO_VERBOSE=False", "python3.11", "test.py")
	cmd.Dir = "../"
	cmdOutput, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "Command Output:\n%s\n", string(cmdOutput))
}
