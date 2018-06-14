package html

import (
	"testing"
	"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"bufio"
)

func TestHTML(t *testing.T) {
	file, err := os.Open("article.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text:=s.Text()
		fmt.Println(text)
		s.ReplaceWithHtml(text)

	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		s.Remove()

	})

	text,err:=doc.Find("body").Html()
	if err != nil {
		fmt.Println(doc.Text())
	}
	write(text)

}

func write(s string) {
	file, err := os.OpenFile("new.html", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(s)
	writer.Flush()

}

/*func TestString(t *testing.T){
	s:="1111111"
	buffer:=bytes.NewBufferString(s)
	buffer.Read()

}*/
