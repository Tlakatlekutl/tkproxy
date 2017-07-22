package parser

import (
	"io"
	"github.com/PuerkitoBio/goquery"
	"github.com/Tlakatlekutl/tkproxy/log"
	"strings"
)

func ChangeSourceUrl(rbody io.Reader, prefixPath, hostPath string) (out string, err error) {
	doc, err := goquery.NewDocumentFromReader(rbody)


	doc.Find("*").Each(func(i int, s *goquery.Selection){
		if url, ok := s.Attr("href"); ok {
			log.Debug("%d) href=%s\n", i, url)
			if strings.HasPrefix(url, "//") {
				s.SetAttr("href", prefixPath + hostPath + strings.TrimPrefix(url, "//"))
			} else {
				s.SetAttr("href", prefixPath + hostPath + url)
			}
		}
	})
	out, err = doc.Html()
	return
}


