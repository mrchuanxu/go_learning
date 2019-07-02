package go_html_regex

import (
	// "bufio"
	// "fmt"
	"net/http"
	"testing"
	"io/ioutil"
	"regexp"
	"os"
)

func TestHtml(t *testing.T){
	var str string = "http://www.love-fitness.com.cn/2019/06/14/%e5%b0%8f%e4%bc%99%e5%81%a5%e8%ba%ab2%e4%b8%aa%e6%9c%88%e5%90%8e%e5%90%91%e5%a5%b3%e5%8f%8b%e7%82%ab%e8%80%80%e8%85%b9%e8%82%8c%ef%bc%8c%e7%bb%93%e6%9e%9c%e5%b0%b4%e5%b0%ac%e4%ba%86%ef%bc%81/"
	resp,err := http.Get(str)
	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	body,errbody := ioutil.ReadAll(resp.Body)
	if errbody!=nil{
		t.Log("wrong body type")
		panic(errbody)
	}

	databody := string(body)

	// fmt.Println(databody)
	
	regx,_ := regexp.Compile(`<article[\d\D]*<\/article>`)
	// fmt.Printf(regx.FindString(databody))

	file1,err2 := os.Create("test.html")
	if err2 != nil{
		panic(err2)
	}
	
	defer file1.Close()

	file1.WriteString(regx.FindString(databody))

}