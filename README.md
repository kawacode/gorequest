# Gorequest

This Project is made for spoofing Ja3 and Akamai fingerprint, it is also very simple to use. I made it noob-friendly as possiable.

# 🚀 Features

```markdown
- High-performance
- Custom Header ordering via https://github.com/kawacode/fhttp
- Custom Pseudo header ordering via https://github.com/kawacode/fhttp
- Custom Frame Setting ordering via https://github.com/kawacode/fhttp
- Able to spoof akamai fingerprint using frame & priority settings
- All Proxy types supported (SOCKS4/SOCKS4a, HTTP/s, SOCKS5, Proxyless)
- Ja3 Customization
- Supported many HelloClients
- HTTP/2 and HTTP1 Support
- HTTP2 Setting customization
- HTTP1 Setting customization
```

# ⬇️ How to install it?

```
go get github.com/kawacode/gorequest
go get github.com/kawacode/gostruct
go get github.com/kawacode/gotools
```

# 🪧 How to use it?

### HTTP request using gorequest without proxy or Ja3 customization

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
)

func main() {
	// Ja3 is by default HelloChrome_105
	var bot gostruct.BotData                                    // Creating struct
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // URL
	bot.HttpRequest.Request.Method = "GET"                      // All method that are supported by fhttp is here supported too
	bot.HttpRequest.Request.Protocol = "2"                      // 2.0 and 1.1 supported
	bot.HttpRequest.Request.ReadResponseBody = true                                  // Used to read the body from the server.
	bot.HttpRequest.Request.ReadResponseCookies = true                                  // Used to read the cookies from the server.
	bot.HttpRequest.Request.ReadResponseHeaders = true                                  // Used to read the headers from the server.
	gorequest.HttpRequest(&bot)                                  // Request
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.request
}

```

### HTTP request using gorequest with custom and ordered headers

```go
package main

import (
  gorequest "github.com/kawacode/gorequest"
  gotools "github.com/kawacode/gotools"
  gostruct "github.com/kawacode/gostruct"
)

func main() {
	var bot gostruct.BotData                                    // Creating a variable called `bot` and setting it to the type `gorequest.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                      // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                      // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponseBody = true                                  // Used to read the body from the server.
	bot.HttpRequest.Request.ReadResponseCookies = true                                  // Used to read the cookies from the server.
	bot.HttpRequest.Request.ReadResponseHeaders = true                                  // Used to read the headers from the server.
	bot.HttpRequest.Request.ClientSpec = "<YourFingerprintbyteOrJa3>" // Your Fingerprint byte or Ja3 it is recommended to use Fingerprint byte
	bot.HttpRequest.Request.Client = gotools.GetHelloClient("HelloCustom") // The Client
	bot.HttpRequest.Request.Headers = map[string]string{        // Used to add headers to the request.
		"Content-Type": "application/json",
		"TestToken":    "123492190391239102301293",
		"user-agent":   "Go-http-client/2.0",
	}
	gorequest.HttpRequest(&bot)               // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.

```

### HTTP request using gorequest in debugging mode

```go
package main

import (
  gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
  gotools "github.com/kawacode/gotools"

)

func main() {
	var bot gostruct.BotData                                    // Creating a variable called `bot` and setting it to the type `gorequest.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                      // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                      // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponseBody = true                                  // Used to read the body from the server.
	bot.HttpRequest.Request.ReadResponseCookies = true                                  // Used to read the cookies from the server.
	bot.HttpRequest.Request.ReadResponseHeaders = true                                  // Used to read the headers from the server.
	bot.HttpRequest.Request.ClientSpec = "<YourFingerprintbyteOrJa3>" // Your Fingerprint byte or Ja3 it is recommended to use Fingerprint byte
	bot.HttpRequest.Request.Client = gotools.GetHelloClient("HelloCustom") // The Client
	bot.HttpRequest.Request.InsecureSkipVerify = true // Turn on debugging mode.
	bot.HttpRequest.Request.Headers = map[string]string{        // Used to add headers to the request.
		"Content-Type": "application/json",
		"TestToken":    "123492190391239102301293",
		"user-agent":   "Go-http-client/2.0",
	}
	bot.HttpRequest.Request.HeaderOrderKey = []string{ // Used to order the headers in the request.
		"content-type",
		"user-agent",
		"testtoken",
	}
	gorequest.HttpRequest(&bot)               // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}

```

### HTTP Request using gorequest with proxy

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
)

func main() {
	var bot gostruct.BotData                                            // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all"         // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                              // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                              // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                         // Used to read the response from the server.
	bot.HttpRequest.Request.Proxy = "type://127.0.0.1:1080:user:password" // Setting the proxy of the request. Supported types: ( "proxyless","http", "https", "socks4", "socks4a", "socks5" ), Without Auth format: type://example.com:port
	bot.HttpRequest.Request.UseProxy = true                               // Used to tell the program to use the proxy.
	gorequest.HttpRequest(&bot)                                          // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}

```

### HTTP Request using gorequest with custom Ja3

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
  gotools "github.com/kawacode/gotools"
)

func main() {
	var bot gostruct.BotData                                                         // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all"                      // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                                           // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                                           // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                                      // Used to read the response from the server.
	bot.HttpRequest.Request.ClientSpec = "<YourFingerprintbyteOrJa3>" // Your Fingerprint byte or Ja3 it is recommended to use Fingerprint byte
	bot.HttpRequest.Request.Client = gotools.GetHelloClient("HelloCustom") // The Client
	gorequest.HttpRequest(&bot)                                                      // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}

```

### HTTP Request using gorequest with HelloClient

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
  gotools "github.com/kawacode/gotools"
)

func main() {
	var bot gostruct.BotData                                                         // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all"                      // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                                           // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                                           // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                                      // Used to read the response from the server.
	bot.HttpRequest.Request.Client = gotools.GetHelloClient("HelloChrome_Auto") // The Client
	gorequest.HttpRequest(&bot)                                                      // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}

```

### HTTP Request using gorequest with MaxRedirects or disable redirects

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
)

func main() {
	var bot gostruct.BotData                                    // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "GET"                      // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                      // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                 // Used to read the response from the server.
	bot.HttpRequest.Request.MaxRedirects = "10"                 // Used to set the max number of redirects allowed from server. By default 10 redirects. to disable use "false"
	gorequest.HttpRequest(&bot)                                  // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}

```

### HTTP Request using gorequest with Payload

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
  gostruct "github.com/kawacode/gostruct"
)

func main() {
	var bot gostruct.BotData                                    // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "POST"                     // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                      // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                 // Used to read the response from the server.
	bot.HttpRequest.Request.Payload = "user=joe&pass=ok"        // Setting the payload of the request.
	gorequest.HttpRequest(&bot)                                  // A function that is used to send a request to the server.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}
```

### HTTP Request using gorequest with Frame and Priority ordering and customization

```go
package main

import (
	gorequest "github.com/kawacode/gorequest"
        gostruct "github.com/kawacode/gostruct"
	json "encoding/json"
)

func main() {
	var bot gostruct.BotData                                    // Creating a variable called `bot` and setting it to the type `gostruct.BotData`.
	bot.HttpRequest.Request.URL = "https://tls.peet.ws/api/all" // Setting the URL of the request.
	bot.HttpRequest.Request.Method = "POST"                     // Used to set the method of the request.
	bot.HttpRequest.Request.Protocol = "2"                      // Used to set the protocol version of the request.
	bot.HttpRequest.Request.ReadResponse = true                 // Used to read the response from the server.
	gorequest.HttpRequest(&bot)                                  // A function that is used to send a request to the server.
	bot.HttpRequest.Request.Payload = "user=joe&pass=ok"        // Setting the payload of the request.
	for k, v := range bot.HttpRequest.Request.Headers { // Printing the Headers from the server.
		println(k, ":	", v)
	}
	for k, v := range bot.HttpRequest.Request.Cookies { // Printing the Cookies from the server.
		println(k, ":	", v)
	}
	println(bot.HttpRequest.Response.Source) // Printing the source from the server.
}
```

### Supported HelloClient's:

```go
HelloCustom
HelloChrome_58
HelloChrome_62
HelloChrome_70
HelloChrome_72
HelloChrome_83
HelloChrome_87
HelloChrome_96
HelloChrome_100
HelloChrome_103
HelloChrome_104
HelloChrome_105
HelloChrome_106
HelloChrome_107
HelloChrome_Auto
HelloFirefox_55
HelloFirefox_56
HelloFirefox_63
HelloFirefox_65
HelloFirefox_102
HelloFirefox_104
HelloFirefox_105
HelloFirefox_106
HelloFirefox_Auto
HelloAndroid_11_OkHttp
HelloIOS_11_1
HelloIOS_12_1
HelloIOS_13
HelloIOS_14
HelloIOS_15_5
HelloIOS_15_6
HelloIOS_16_0
HelloIOS_Auto
HelloSafari_15_3
HelloSafari_15_5
HelloSafari_16_0
HelloSafari_15_6_1
HelloSafari_Auto
HelloIPad_15_6
HelloIPad_Auto
Hellocsslang
HelloOpera_89
HelloOpera_90
HelloOpera_Auto
HelloRandomized
HelloRandomizedALPN
HelloRandomizedNoALPN

```

## LICENSE

### GPL3 LICENSE SYNOPSIS

**_TL;DR_\*** Here's what the GPL3 license entails:

```markdown
1. Anyone can copy, modify and distribute this software.
2. You have to include the license and copyright notice with each and every distribution.
3. You can use this software privately.
4. You can use this software for commercial purposes.
5. Source code MUST be made available when the software is distributed.
6. Any modifications of this code base MUST be distributed with the same license, GPLv3.
7. This software is provided without warranty.
8. The software author or license can not be held liable for any damages inflicted by the software.
```

More information on about the [LICENSE can be found here](http://choosealicense.com/licenses/gpl-3.0/)
