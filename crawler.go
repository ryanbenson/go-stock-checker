package main

import (
  "io"
  "errors"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

type Crawler struct{}

// getPage will get the page contents from the URL we want
// @params: url (string)
// @return: io.ReadCloser
// @return: error
func (c Crawler) getPage(url string) (io.ReadCloser, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if res.StatusCode != 200 {
    return nil, errors.New("Page not available")
  }
  return res.Body, err
}

// getHtmlDoc gets the DOM from the html contents
// @params: body (io.ReadCloser)
// @return: *goquery.Document
// @return: error
func (c Crawler) getHTMLDoc(body io.ReadCloser) (*goquery.Document, error) {
  doc, err := goquery.NewDocumentFromReader(body)
  if err != nil {
    return nil, err
  }
  return doc, nil
}

// getNodeText gets the text from the HTML Node
// @params: doc (*goquery.Document)
// @params: ref (string)
// @return string
func (c Crawler) getNodeText(doc *goquery.Document, ref string) string {
  return doc.Find(ref).Text()
}
