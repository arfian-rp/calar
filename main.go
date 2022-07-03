package main

//CREATED BY ARFIAN

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const filename = "/calarinit"

func main() {
	args := os.Args
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args[1] {
	case "info":
		fmt.Println(`
		CALAR v1
website			: https://arfian-id.web.app
instagram		: https://instagram.com/arfian_rp_
github			: https://github.com/arfian-rp/calar`)
	case "init":
		CreateInit(pwd)
	default:
		fmt.Println(RunScript(pwd, args))
	}

}
func CreateInit(pwd string) {
	if _, err := os.Stat("/path/to/whatever"); err != nil {
		f, err := os.Create(pwd + filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err2 := f.WriteString(`{
"start":"echo run start"
}`)
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("done")
	}

}
func RunScript(pwd string, args []string) string {
	jsonData, err := ioutil.ReadFile(pwd + filename)
	if err != nil {
		panic(err)
	}
	var setup map[string]string
	json.Unmarshal(jsonData, &setup)
	com := strings.Fields(setup[args[1]])
	out, err2 := exec.Command(com[0], com[1:]...).Output()
	if err2 != nil {
		panic(err2)
	}
	return string(out)
}
