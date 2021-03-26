package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
)

func main() {
	response,err := httplib.Get("https://cn.investing.com/equities/apple-computer-inc-historical-data").Response()
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range response.Cookies() {
		fmt.Println(cookie.Name,cookie.Value)
	}



	req := httplib.Post("https://cn.investing.com/instruments/HistoricalDataAjax")
	req.Header("curr_id","100673")
	req.Header("smlID","1437916")
	req.Header("header","600519历史数据")
	req.Header("st_date","2021/02/11")
	req.Header("end_date","2021/03/11")
	req.Header("interval_sec","Daily")
	req.Header("sort_col","date")
	req.Header("sort_ord","DESC")
	req.Header("action","historical_data")
	for _, cookie := range response.Cookies() {
		_ = cookie
		//req.SetCookie(cookie)
	}
	str,err := req.String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
