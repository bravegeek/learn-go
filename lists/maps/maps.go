package maps

import "fmt"

func main() {
	// websites:=[]string{"https://google.com","https://aws.com"}

	websites := map[string]string{ //like a dictionary
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
