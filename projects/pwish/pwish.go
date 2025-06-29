package main

import (
	"fmt"
	"embed"
	"github.com/fatih/color"
)
//go:embed pwish.txt
var pwishContent embed.FS

func printPwishAvatar(){
	pwishAvatar,err := pwishContent.ReadFile("pwish.txt")
	if err != nil{
		fmt.Println("Avatar not found",err)
	}
	fmt.Println(string(pwishAvatar))
}

func main(){
	PurpleOnBlack := color.New(color.FgMagenta,color.BgBlack)
	BlackOnPurple := color.New(color.FgBlack,color.BgMagenta)
	WhiteOnPurple := color.New(color.FgWhite,color.BgMagenta)
	PurpleOnWhite := color.New(color.FgMagenta,color.BgWhite)
	printPwishAvatar()
	PurpleOnBlack.Println("\t\t\t\t                    Pwish-Go @v1.0.0                           ")
	BlackOnPurple.Println("\t\t\t\tAd: Pwish                                                      ")
	PurpleOnBlack.Println("\t\t\t\tProgramlama dilleri: Javascript                                ")
	WhiteOnPurple.Println("\t\t\t\tKullandığı DE: KDE/XFCE4                                       ")
	PurpleOnWhite.Println("\t\t\t\tOS: Gentoo                                                     ")
	PurpleOnBlack.Println("\t\t\t\t                    Pwish-Go @v1.0.0                           ")
	WhiteOnPurple.Println("\t\t\t\t                      Made by isa                              ")

	
}
