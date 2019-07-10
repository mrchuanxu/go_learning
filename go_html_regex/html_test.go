package go_html_regex

import (
	"net/http"
	"testing"
	"io/ioutil"
	"regexp"
	"os"
	"github.com/opesun/goquery"
	"io"
	"bytes"
)
// 判断错误
func CheckErr(err error){
	if err != nil{
		panic(err) // 直接终止程序
	}
}

func GetHtmlText(str string)string{
	resp,err := http.Get(str)
	CheckErr(err)
	defer resp.Body.Close()

	body,errbody := ioutil.ReadAll(resp.Body)
	CheckErr(errbody)
	// 格式化string
	str = string(body)
	return str
}

// 将string与compile绑定 返回compile后的结果
func GetTextComplie(str string,compstr string)string{
	regx,_ := regexp.Compile(compstr)
	strBody := regx.FindString(str)
	return strBody
}

func GetArticleHeadName(str string,compstr string)string{
	regx,_ := regexp.Compile(compstr)
	strBody := regx.ReplaceAllString(str, "")
	return strBody
}

// 获取图片的正则书写 将图片下载到本地
func GetSrcImg(str string,compstr string){

}

// 新建文件夹 先判断文件夹是否存在
func pathExists(path string)(bool,error){
	_,err := os.Stat("./"+path)
	if err != nil{
		return true,err
	}
	return false,err
}
// 新建文件夹
func MakeDir(str string)(bool,error){
	isExist,err:=pathExists(str)
	if err==nil{
		panic("Can't create Dir")
	}
	err = os.Mkdir(str,os.ModePerm)
	return isExist,err
}
// 将string写到 新建txt文件 str写入的是带标签的
func MakeTxt(strPath string,str string)(bool,error){
	file1,err := os.Create("./"+strPath+"/"+strPath+".txt")
	CheckErr(err)
	defer file1.Close()
	file1.WriteString(str)
	if err != nil{
		return false,err
	}
	return true,err
}


func TestHtml(t *testing.T){
	var str string = "http://www.love-fitness.com.cn/2019/07/05/%e7%9c%8b%e5%88%b0%e5%8f%91%e8%83%96%e5%90%8e%e7%9a%84%e5%90%b4%e4%ba%a6%e5%87%a1%ef%bc%8c%e5%90%93%e5%be%97%e6%88%91%e6%89%94%e6%8e%89%e4%ba%86%e6%89%8b%e4%b8%ad%e7%9a%84%e9%b8%a1%e8%85%bf/"
	str = GetHtmlText(str)
	str = GetTextComplie(str,`<article[\d\D]*<\/article>`)
	// imgStr := GetTextComplie(str,`<img class[\d\D]*\/>\s`)
	// t.Log(imgStr)
	xUrl, err := goquery.ParseUrl("http://www.love-fitness.com.cn/2019/07/05/%e7%9c%8b%e5%88%b0%e5%8f%91%e8%83%96%e5%90%8e%e7%9a%84%e5%90%b4%e4%ba%a6%e5%87%a1%ef%bc%8c%e5%90%93%e5%be%97%e6%88%91%e6%89%94%e6%8e%89%e4%ba%86%e6%89%8b%e4%b8%ad%e7%9a%84%e9%b8%a1%e8%85%bf/")
	CheckErr(err)
	imgUrls := xUrl.Find("img").Attrs("src")
	reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	for i:=0;i<len(imgUrls);i++{
		name := reg.FindStringSubmatch(imgUrls[i])[0]
		resp, _ := http.Get(imgUrls[i])
		body, _ := ioutil.ReadAll(resp.Body)
		out, _ := os.Create(name)
		io.Copy(out, bytes.NewReader(body))
	}
	// need to get the http img
	// strH1Get := GetTextComplie(str,`<h1 class=\"entry-title\">[\d\D]*<\/h1>`)
	//strhead:= GetArticleHeadName(strH1Get,"\\<[\\S\\s]+?\\>")
	// fmt.Println(databody)
	// t.Log(strhead)
	//strClearHtml := GetArticleHeadName(str,"\\<[\\S\\s]+?\\>")
	//isMake,err := MakeDir(strhead)
	// isMake,err = MakeTxt(strhead,strClearHtml)
}