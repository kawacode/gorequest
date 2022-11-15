package gorequest

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	http "github.com/kawacode/fhttp"
	gostruct "github.com/kawacode/gostruct"
	gotools "github.com/kawacode/gotools"
)

// It makes an HTTP request and returns the response
// bot is the bot where datas of response or request are stored/saved
// dstbot it is the bot from the Request is being used by HttpRequest
func HttpRequest(bot *gostruct.BotData) http.Response {
	URL, err := url.Parse(bot.HttpRequest.Request.URL)
	if err != nil {
		log.Panic(err)
	}
	if URL.String() == "" {
		log.Panic("please provide a URL parameter at bot.HttpRequest.Request.URL")
	}
	if bot.HttpRequest.Request.Method == "" {
		log.Panic("please provide a method parameter at bot.HttpRequest.Request.Method")
	}
	if !gotools.IsInt(bot.HttpRequest.Request.Protocol) {
		bot.HttpRequest.Request.Protocol = "2"
	}
	req, err := http.NewRequest(strings.ToUpper(bot.HttpRequest.Request.Method), URL.String(), bytes.NewBuffer([]byte(bot.HttpRequest.Request.Payload)))
	if err != nil {
		log.Panic(err)
	}
	req.Close = true
	req.Header = gotools.MapStringToMapStringSlice(bot.HttpRequest.Request.Headers, bot)
	client, err := CreateClient(bot)
	if err != nil {
		log.Panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	location, err := res.Location()
	if err == nil {
		bot.HttpRequest.Response.Location = *location
	}
	if bot.HttpRequest.Request.ReadResponseCookies {
		cookies := make(map[string]string)
		for _, cookie := range res.Cookies() {
			cookies[cookie.Name] = cookie.Value
		}
		bot.HttpRequest.Response.Cookies = cookies
	}
	bot.HttpRequest.Response.Status = res.Status
	bot.HttpRequest.Response.StatusCode = res.StatusCode
	if bot.HttpRequest.Request.ReadResponseHeaders {
		bot.HttpRequest.Response.Headers = gotools.MapStringSliceToMapString(res.Header)
	}
	bot.HttpRequest.Response.Protocol = res.Proto
	bot.HttpRequest.Response.ContentLength = res.ContentLength
	if bot.HttpRequest.Request.ReadResponseBody {
		resp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Panic(err)
		}
		if !bot.HttpRequest.Request.ReadResponseHeaders {
			source, err := gotools.DecompressGzip(string(resp))
			if err != nil {
				bot.HttpRequest.Response.Source = string(resp)
			}
			bot.HttpRequest.Response.Source = source
		} else {
			bot.HttpRequest.Response.Source = string(resp)
		}
	}
	bot.HttpRequest.Response.ProtoMajor = res.ProtoMajor
	bot.HttpRequest.Response.ProtoMinor = res.ProtoMinor
	bot.HttpRequest.Response.WasUncompressed = res.Uncompressed
	client.CloseIdleConnections()
	return *res
}
