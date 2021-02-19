package main

import "github.com/yuin/gopher-lua"
import "github.com/cjoudrey/gluahttp"
import "net/http"

func main() {
    L := lua.NewState()
    defer L.Close()

    L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)

    if err := L.DoString(`

        local http = require("http")

        response, error_message = http.request("GET", "http://bing.com", {
            query="page=1",
            timeout="30s",
            headers={
                Accept="*/*"
            }
        })
        print(response.body)

    `); err != nil {
        panic(err)
    }
}
