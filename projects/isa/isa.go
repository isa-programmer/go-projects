package main

import (
	"fmt"
	"embed"
	"github.com/fatih/color"
)
//go:embed isa.txt
var isaContent embed.FS

func printIsaAvatar(){
	isaAvatar,err := isaContent.ReadFile("isa.txt")
	if err != nil{
		fmt.Println("Avatar not found",err)
	}
	fmt.Println(string(isaAvatar))
}

func main(){
	BlueOnBlack := color.New(color.FgBlue,color.BgBlack)
	BlackOnBlue := color.New(color.FgBlack,color.BgBlue)
	WhiteOnBlue := color.New(color.FgWhite,color.BgBlue)
	BlueOnWhite := color.New(color.FgBlue,color.BgWhite)
	printIsaAvatar()
	BlueOnBlack.Println("\t\t\t\t                    İsa-Go @v1.0.1                             ")
	BlackOnBlue.Println("\t\t\t\tAd: İsa                                                        ")
	BlueOnBlack.Println("\t\t\t\tProgramlama dilleri: Go,Bash,Python                            ")
	WhiteOnBlue.Println("\t\t\t\tKullandığı DE: GNOME                                           ")
	BlueOnWhite.Println("\t\t\t\tOS: Debian                                                     ")
	BlueOnBlack.Println("\t\t\t\tProjeler: goscan, githubfetch,yt-api-wrapper                   ")
	BlackOnBlue.Println("\t\t\t\tGithub: https://github.com/isa-programmer                      ")
	BlueOnBlack.Println("\t\t\t\t                    İsa-Go @v1.0.1                             ")
	WhiteOnBlue.Println("\t\t\t\t               Made by isa-programmer for İsa                  ")

	
}
