package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os/exec"
	"os"
)

func FetchGithubProfile(username string) (map[string]interface{},error){ 
	var data map[string]interface{}
	resp,err := http.Get("https://api.github.com/users/" + username)
	if err != nil{
		return nil, err
	}
	
	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil{
		return nil, err
	}
	return data, nil
}

func main(){
	var avatar_url string
	var username string
	if len(os.Args) < 2 {
		fmt.Print("\x1b[38;5;2m Usage: ./githubfetch <username> \x1b[0m")
	}
	username = os.Args[1]
	data,err := FetchGithubProfile(username)
	if err != nil{
		fmt.Println(err)
		return
	}
	if data["bio"] == nil{
		data["bio"] = "N/A"
	}
	
	avatar_url = data["avatar_url"].(string)
	cmd := exec.Command("kitty","+kitten","icat","--align=left","--place=20x20@0x1",avatar_url)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
	        fmt.Println("Error:", err)
	}
	fmt.Printf("\n")
	fmt.Printf("\x1b[38;5;4m \t\t\t Username:\x1b[0m %s \n",data["name"])
	fmt.Printf("\x1b[38;5;3m \t\t\t Repos:\x1b[0m %1.f \n",data["public_repos"])
	fmt.Printf("\x1b[38;5;2m \t\t\t Bio:\x1b[0m %s \n",data["bio"])
	fmt.Printf("\x1b[38;5;1m \t\t\t Location:\x1b[0m %s \n",data["location"])
	fmt.Printf("\x1b[38;5;12m \t\t\t Followers:\x1b[0m %1.f \n",data["followers"])
	fmt.Printf("\x1b[38;5;13m \t\t\t Following:\x1b[0m %1.f \n",data["following"])
	fmt.Printf("\n\n\n\n")
}
