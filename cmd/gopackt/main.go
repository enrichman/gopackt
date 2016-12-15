package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/enrichman/gopackt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	pClient := gopackt.NewClient()

mainloop:
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		switch text {
		case "":
			continue
		case "exit":
			break mainloop
		case "login":
			email, pwd := credentials()

			done := false
			doneChan := make(chan bool)

			go gopackt.FancyLoad(&done, doneChan)
			user, err := pClient.Login(email, pwd)
			done = true
			<-doneChan

			fmt.Println("User", user, "Error", err)

		case "ls":
			pClient.ListBooks()
		default:
			fmt.Println(text)
		}
	}

	fmt.Println("Bye!")
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Println("Error:", err)
	}
	password := string(bytePassword)
	fmt.Println()

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

/*
func main() {
	fmt.Println("Hello")

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	// "email":"enrico.candino@gmail.com","password":"3nr1c0Pakt"

	baseURL := "https://www.packtpub.com"
	resp, _ := client.Get(baseURL)
	htmlByte, _ := ioutil.ReadAll(resp.Body)

	ioutil.WriteFile("hello.html", htmlByte, 0644)
	// fmt.Println(string(htmlByte))

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlByte))
	if err != nil {
		log.Fatal(err)
	}

	form := url.Values{}

	loginForm := doc.Find("#packt-user-login-form")
	loginForm.Find("input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")

		fmt.Printf("input %d: %s - %s\n", i, name, value)

		form.Add(name, value)
	})

	fmt.Printf("form", form)

	form.Add("email", "enrico.candino@gmail.com")
	form.Add("password", "3nr1c0Pakt")

	fmt.Printf("form", form)

	resp, _ = client.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	htmlByte, _ = ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("logged.html", htmlByte, 0644)

}
*/
