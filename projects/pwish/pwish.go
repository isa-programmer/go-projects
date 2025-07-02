package main
// MIT LICENSE
// İşsiz gibi bununla uğraştım .d
// Muhtemelen kimse okumayacak
import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
	"embed"
)


func customBox(String string){
	var style = lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#ff00ff")).
		MarginLeft(15).
		Foreground(lipgloss.Color("#f8f8ff")).
		Background(lipgloss.Color("#8c018c"))

	fmt.Println(style.Render(String))
}

//go:embed pwish.txt
var pwishContent embed.FS

func ReadPwish() (string,error) {
	content, err := pwishContent.ReadFile("pwish.txt")
	if err != nil{
		return "",err
	}

	return string(content), nil
}

func main(){
	avatar,err := ReadPwish()
	if err != nil{
		fmt.Println(err)
		return
	}
    customBox(avatar)
}
