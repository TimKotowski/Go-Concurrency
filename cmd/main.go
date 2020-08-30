	package main


	import (
		"fmt"
		"net/http"
		"time"
	)

	func main() {
		links := []string {
			"https://www.youtube.com/",
			"https://www.twitch.tv/",
			"https://golang.org/",
			"http://timkotowski.com/",
		}

		c := make(chan string)

		for _, link := range links {
			go checkLink(link, c)
		}

		for  l := range c {
			go func (link string) {
				time.Slee=pop(5 * time.Second)
				checkLink(link, c)
			}(l)
		}
	}

		func checkLink(link string, c chan string) {

				_, err := http.Get(link)
				if err != nil {
					fmt.Println(link, "might be down!")
					c <- link
					return
				}
				fmt.Println(link, "is up!")
				c <- link
		}


