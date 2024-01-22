package main
import (
  "log"
  "fmt"
  "time"
  "net/http"
  "io/ioutil"
  "strings"
  "github.com/corpix/uarand"
  "github.com/PuerkitoBio/goquery"
  )
const FilePath="/sdcard/project/telegram-core/allwallet.txt"
const bsc="https://bscscan.com/address/"
type DbgidRequest interface{
  Fetch() string
}
type do struct{
  checker string
}
func(c *do) Fetch() string{
  ua:=&ua{}
  req,err:=http.NewRequest("GET",bsc+c.checker,nil)
  if err !=nil{
    log.Fatal(err)
    return ""
  }
  req.Header.Set("Accept","text/html")
  req.Header.Set("User-Agent",ua.getUa())
  client:=&http.Client{}
  res,err:=client.Do(req)
  if err !=nil{
    log.Fatal(err)
    return ""
  }
  defer res.Body.Close()
  if res.StatusCode !=200{
    log.Printf("Error: server error, server returning status code -> %d",res.StatusCode)
    return ""
  }
  body,err:=ioutil.ReadAll(res.Body)
  if err !=nil{
    panic("Failed to get response body")
    return ""
  }
  bodyStr:=string(body)
  doc,err:=goquery.NewDocumentFromReader(strings.NewReader(bodyStr))
  if err !=nil{
    log.Fatal(err)
    return ""
  }
  bep20:=doc.Find("html#html>body#body>main#content>section#ContentPlaceHolder1_divSummary>div:nth-of-type(2)>div>div>div>div#ContentPlaceHolder1_divTokenHolding>div#ContentPlaceHolder1_tokenbalance>div>div#availableBalance>div:nth-of-type(2)>ul>li>div>span").Text()
  if strings.Contains(bep20,"(0)")==true{
    log.Printf("Skip this wallet is not have token yet -> %s",bep20)
    return ""
  }
  log.Printf("[info] Total Token: %s",bep20)
  doc.Find("html#html>body#body>main#content>section#ContentPlaceHolder1_divSummary>div:nth-of-type(2)>div>div>div>div#ContentPlaceHolder1_divTokenHolding>div#ContentPlaceHolder1_tokenbalance>div>div#availableBalance>div>ul>li").Each(func(i int, s *goquery.Selection){
    token:=s.Find("span").Text()
    contract:=s.Find("a").AttrOr("href","")
    contract=strings.Replace(contract,"?a="+c.checker,"",-1)
    contract=strings.Replace(contract,"/token/","",-1)
    if i==0{
      fmt.Println("")
    }else{
    log.Printf("[%d] %s (%s)",i,token,contract)
    }
  })
  return ""
}
type HttpUtil interface{
  getUa() string
}
type ua struct {}
func(u *ua) getUa() string {
  return uarand.GetRandom()
}
type Address interface{
  getWallet() []string
}
type wallet struct{
  fileWallet string
}
func (w *wallet) getWallet() []string{
  data,err:=ioutil.ReadFile(w.fileWallet)
  if err !=nil{
    log.Fatal(err)
    return []string{}
  }
  walletStr:=string(data)
  walletSplit:=strings.Split(walletStr,"\n")
  return walletSplit
}
func main(){
  w:=&wallet{fileWallet:FilePath}
  for index,data:=range w.getWallet(){
    fmt.Printf("[info] Wallet ke %d/%d -> %s\n",index,len(w.getWallet()),data)
    client:=&do{checker:data}
    time.Sleep(2 * time.Second)
    client.Fetch()
  }
}