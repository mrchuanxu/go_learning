package go_html_regex

import (
	"net/http"
	"testing"
	"io/ioutil"
	"regexp"
	"os"
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
	var str string = "http://www.love-fitness.com.cn/2019/06/14/%e5%b0%8f%e4%bc%99%e5%81%a5%e8%ba%ab2%e4%b8%aa%e6%9c%88%e5%90%8e%e5%90%91%e5%a5%b3%e5%8f%8b%e7%82%ab%e8%80%80%e8%85%b9%e8%82%8c%ef%bc%8c%e7%bb%93%e6%9e%9c%e5%b0%b4%e5%b0%ac%e4%ba%86%ef%bc%81/"
	// resp,err := http.Get(str)
	// if err != nil{
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// body,errbody := ioutil.ReadAll(resp.Body)
	// if errbody!=nil{
	// 	t.Log("wrong body type")
	// 	panic(errbody)
	// }

	// databody := string(body)
	str = GetHtmlText(str)
	str = GetTextComplie(str,`<article[\d\D]*<\/article>`)
	imgStr := GetTextComplie(str,`<img class[\d\D]*\/>`)
	t.Log(imgStr)
	// strH1Get := GetTextComplie(str,`<h1 class=\"entry-title\">[\d\D]*<\/h1>`)
	//strhead:= GetArticleHeadName(strH1Get,"\\<[\\S\\s]+?\\>")
	// fmt.Println(databody)
	// t.Log(strhead)
	//strClearHtml := GetArticleHeadName(str,"\\<[\\S\\s]+?\\>")
	//isMake,err := MakeDir(strhead)
	// isMake,err = MakeTxt(strhead,strClearHtml)
	


	// if err!=nil&&isMake == false{
	// 	panic(nil)
	// }
	// regx,_ := regexp.Compile(`<article[\d\D]*<\/article>`)
	// // fmt.Printf(regx.FindString(databody))
	// regx1,_ := regexp.Compile(`<h1 class=\"entry-title\">[\d\D]*<\/h1>`)
	
	// strBody := regx.FindString(databody)
	// t.Log(regx1.FindString(strBody))
	
	// re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	// t.Log(re.ReplaceAllString(regx1.FindString(strBody), ""))

	// err = os.Mkdir(re.ReplaceAllString(regx1.FindString(strBody), ""),os.ModePerm)
	// file1,err2 := os.Create("test.txt")

	
	// if err2 != nil{
	// 	panic(err2)
	// }
	
	// defer file1.Close()
	// strBody = re.ReplaceAllString(strBody," ")
	// file1.WriteString(strBody)

}