package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ErrPanic(err error, snapshot interface{}) {
	println("ERROR:")
	println(snapshot)
	panic(err)
}

type Exporter struct {
	file *os.File
}

func (e *Exporter) WriteToFile(content string) {
	data := fmt.Sprintf("%s\n", content)
	_, err := e.file.WriteString(data)
	if err != nil {
		ErrPanic(err, content)
	}
}

type BookInfoCollector struct {
	index int
	url   string
}

func (c *BookInfoCollector) refineAuthor(raw string) string {
	out := strings.Replace(raw, "Tác giả : ", "", 1)
	out = strings.Replace(out, " - ", ", ", -1)
	return out
}

func (c *BookInfoCollector) ExportData() string {
	doc, err := goquery.NewDocument(c.url)
	if err != nil {
		ErrPanic(err, c.url)
	}

	var title, publisher, description, authors, isbn, image string
	var create_at, updated_at, publish_date time.Time
	create_at, updated_at = time.Now(), time.Now()
	publish_date = time.Now() // TODO no publish date

	bookInfo := doc.Find(".thong_tin_ebook")
	title = bookInfo.Find("h1").Text()
	image, _ = bookInfo.Find("img").Attr("src")
	authors = c.refineAuthor(bookInfo.Find("h5").First().Text())

	description = doc.Find(".gioi_thieu_sach").Text()

	return strconv.Itoa(c.index) + "\t" +
		title + "\t" +
		fmt.Sprintf(create_at.Format("2006-01-02 15:04:05")) + "\t" +
		fmt.Sprintf(updated_at.Format("2006-01-02 15:04:05")) + "\t" +
		publisher + "\t" +
		fmt.Sprintf(publish_date.Format("2006-01-02 15:04:05")) + "\t" +
		description + "\t" +
		authors + "\t" +
		isbn + "\t" +
		image
}

type Link struct {
	url string
	id  int
}

func cloneBookWorker(id int, linksQueue chan Link, wg *sync.WaitGroup) {
	outFilePath := "./chan_" + strconv.Itoa(id) + "_out.tsv"
	file, err := os.Create(outFilePath)
	if err != nil {
		ErrPanic(err, outFilePath)
	}
	exporter := Exporter{file: file}

	defer wg.Done()
	for {
		link, ok := <-linksQueue
		if !ok {
			return
		}
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		println("Clone book:", link.id, link.url)
		col := BookInfoCollector{index: link.id, url: link.url}
		exporter.WriteToFile(col.ExportData())
	}
}

func linkCollector(wg *sync.WaitGroup, linksQueue chan Link) {

	seedUrl := "https://sachvui.com/the-loai/tat-ca.html"

	for i := 1; i < 3; i++ {
		println("Collect link:", seedUrl+"/"+strconv.Itoa(i))
		doc, _ := goquery.NewDocument(seedUrl + "/" + strconv.Itoa(i))
		books := doc.Find(".panel-body").Find(".ebook")
		books.Each(func(j int, s *goquery.Selection) {
			url, _ := s.Find("a").Attr("href")
			println("Push to queue", (i*books.Length())+j, url)
			linksQueue <- Link{id: (i * books.Length()) + j, url: url}
		})
	}

	close(linksQueue)
}

const WORKERS = 5

func main() {
	var wg sync.WaitGroup
	wg.Add(WORKERS)

	linksQueue := make(chan Link, WORKERS)
	for i := 0; i < WORKERS; i++ {
		go cloneBookWorker(i, linksQueue, &wg)
	}

	go linkCollector(&wg, linksQueue)
	wg.Wait()
}
