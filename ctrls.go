package ctrls

import "os"
import "encoding/json"
import "os/user"
import "runtime"
import "fmt"

func parselocation(location string) string {

	user, err := user.Current()
	if err != nil {
		fmt.Println("ctrls: ", err)
		return ""
	}
	
	if len(location) == 0 {
		fmt.Println("ctrls: missing location")
		return ""
	}

	if location[0] == '~' {
		if runtime.GOOS == "linux" { 
			location = "."+location[2:]
		} else {
			location = location[2:]
		}
		location = user.HomeDir+location
	}
	
	return location
}

func Load(value interface{}, location string) {
	
	location = parselocation(location)
	
	file, err := os.Open(location)
	if err != nil {
		fmt.Println("ctrls: ", err)
		return
	}

	dec := json.NewDecoder(file)
	
	err = dec.Decode(value)
	if err != nil {
		fmt.Println("ctrls: ", err)
		return
	}
}

func Save(value interface{}, location string) {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println("ctrls: ", err)
		return
	}
	
	location = parselocation(location)
	
	file, err := os.Create(location)
	if err != nil {
		fmt.Println("ctrls: ", err)
		return
	}
	
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("ctrls: ", err)
		return
	}
}
