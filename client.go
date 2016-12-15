package gopackt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://www.packtpub.com"
)

type Client struct {
	httpClient  *http.Client
	loginValues url.Values
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	c := Client{
		httpClient: &http.Client{
			Jar: jar,
		},
	}

	c.init()

	return &c
}

func (c *Client) init() {
	c.loginValues = url.Values{}

	resp, _ := c.httpClient.Get(baseURL)
	htmlByte, _ := ioutil.ReadAll(resp.Body)

	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(htmlByte))
	loginForm := doc.Find("#packt-user-login-form")
	loginForm.Find("input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")

		c.loginValues.Add(name, value)
	})

	fmt.Println("Packt Client initialized")
}

func (c *Client) Login(email string, password string) (username string, err bool) {
	username = email

	c.loginValues.Add("email", email)
	c.loginValues.Add("password", password)

	resp, _ := c.httpClient.Post(baseURL, "application/x-www-form-urlencoded", strings.NewReader(c.loginValues.Encode()))

	htmlByte, _ := ioutil.ReadAll(resp.Body)
	// ioutil.WriteFile("logged2.html", htmlByte, 0644)

	// check if login was successfull
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(htmlByte))
	loginForm := doc.Find("#packt-user-login-form")
	loginForm.Find("input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if name == "email" || name == "password" {
			inputClass, _ := s.Attr("class")
			if strings.Contains(inputClass, "error") {
				err = true
			}
		}
	})

	return username, err
}

func (c *Client) ListBooks() {
	fmt.Println("Getting books")

	quit := make(chan bool)
	go FancyLoad(quit)

	resp, _ := c.httpClient.Get("https://www.packtpub.com/account/my-ebooks")
	htmlByte, _ := ioutil.ReadAll(resp.Body)
	//ioutil.WriteFile("listBooks.html", htmlByte, 0644)

	// Quit goroutine
	quit <- true

	// check if login was successfull
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(htmlByte))
	bookList := doc.Find("#product-account-list")
	bookList.Find("div .product-line").Each(func(i int, s *goquery.Selection) {
		if title, _ := s.Attr("title"); title != "" {
			fmt.Println("", strconv.Itoa(i+1)+")", title)
		}
	})
}
