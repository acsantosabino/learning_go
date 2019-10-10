package page

import (
	"io/ioutil"
    "os"
)

type Page struct {
    Title string
    Body  []byte
}

var Filepath = "../data/"

func CheckPagesPath() error{
    _, err := os.Stat(Filepath)
    if os.IsNotExist(err){
        err = os.MkdirAll(Filepath, 0777)
    }
    return err
}

func AllPages() (pages []string, err error){
    fp, err := os.Open(Filepath)
    if err == nil {
        pages, err = fp.Readdirnames(0)
    }
    return pages, err
}

func (p *Page) Save() error {
    filename := Filepath + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
    filename := Filepath + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}