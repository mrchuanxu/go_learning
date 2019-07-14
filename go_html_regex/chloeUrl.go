package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"os"
	"github.com/opesun/goquery"
	"io"
	"bytes"
	"sync"
	"time"
	"fmt"
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
func GetSrcImg(str string,outDir string){
	xUrl, err := goquery.ParseUrl(str)
	CheckErr(err)
	imgUrls := xUrl.Find("img").Attrs("src") // 这里使用gojquery库，感觉这个程序很多地方都可以使用这个库
	reg, _ := regexp.Compile(`(\w|\d|_)*.(jpg|png|gif)`)
	numRoutine := len(imgUrls)
	var wg sync.WaitGroup
	for i:=0;i<numRoutine;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			strName := string(i) //TODO: know routine best tomorrow!
			var name string
			if imgName:=reg.FindStringSubmatch(imgUrls[i]);imgName!= nil{
				name = imgName[0]
				resp, _ := http.Get(imgUrls[i])
				body, _ := ioutil.ReadAll(resp.Body)
				out, _ := os.Create("./"+outDir+"/"+strName+name)// 写入
				io.Copy(out, bytes.NewReader(body))
			}else{
				fmt.Println("无法下载该image")
			}
		}(i)
	}
	wg.Wait()
	time.Sleep(3*time.Second)
}

// 新建文件夹 先判断文件夹是否存在
func pathExists(path string)(bool,error){
	_,err := os.Stat("./"+path)
	if err != nil{
		return true,err
	}
	fmt.Println("对不起，该网址已经被爬取，请更换地址！")
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


func main(){
	if len(os.Args)==2{
		fmt.Println("谢谢，已经正在进行爬取")
	}else{
		panic("请输入一个需要爬取网址")
	}
	var strUrl string = os.Args[1]
	str := GetHtmlText(strUrl)
	str = GetTextComplie(str,`<article[\d\D]*<\/article>`)
	// need to get the http img
	strH1Get := GetTextComplie(str,`<h1 class=\"entry-title\">[\d\D]*<\/h1>`)
	strhead:= GetArticleHeadName(strH1Get,"\\<[\\S\\s]+?\\>")
	strClearHtml := GetArticleHeadName(str,"\\<[\\S\\s]+?\\>")
	isMake,err := MakeDir(strhead)
	isMake,err = MakeTxt(strhead,strClearHtml)
	GetSrcImg(strUrl,strhead)
	if err != nil{
		fmt.Println(isMake)
		panic(err)
	}
}