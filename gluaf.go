package main

import "github.com/yuin/gopher-lua"
import "github.com/cjoudrey/gluahttp"
import "net/http"
import "os"

func main() {
    if len(os.Args) !=2 {
        panic("Unexpected args\n")
        os.Exit(-1)
    }
    L := lua.NewState()
    defer L.Close()

    L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)

    if err := L.DoFile(os.Args[1]); err != nil {
        panic(err)
    }
}
