
package main

import (
    "fmt"
    "net/http"
    "io"
//    "io/ioutil"
    "os"
)

func main() {

    resp, err := http.Get("http://127.0.0.1:8080/api")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()
//    body, err := ioutil.ReadAll(resp.Body)
    
//    if err != nil {
//        fmt.Println(err)
//        return
//    }

    io.Copy(os.Stdout, resp.Body)
}
