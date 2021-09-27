package main

// go mod init srv
// go mod tidy
// go build -o bin/srv -v .

// import (
// "fmt"
// // "io"
// "os"
// "log"
// "net/http"
// // "net/url"
// "time"
// // "github.com/yhat/wsutil"
// )

// func goPort(p string) string {
// pt:= os.Getenv("PORT")
// if pt == "" {
// pt = p
// }
// return pt
// }

import (
        "fmt"
        "net/http"
        "os"
)

func GoPort(p string) string {
        pt:= os.Getenv("PORT")
        if pt == "" {
                pt = p
        }
        return pt
}

func hello(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
        for name, headers := range req.Header {
                for _, h := range headers {
                fmt.Fprintf(w, "%v: %v\n", name, h)
                }
        }
}

func main() {
        http.HandleFunc("/", hello)
        http.HandleFunc("/headers", headers)
        // http.ListenAndServe(":22222", nil)
        http.ListenAndServe(":"+GoPort("22222"),nil)
}

// func Handler(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w,"hello")
        // dialer := &websocket.Dialer{
        //         ReadBufferSize:128,
        //         WriteBufferSize:128,
        //         Proxy:http.ProxyFromEnvironment,
        //         HandshakeTimeout:8*time.Second,
        // }
        // fmt.Println("1...",r.URL)
        // fmt.Println("2...",r.URL.Hostname())
        // fmt.Println("1...",r.URL.String())       
        // if r.URL.String()=="/" {
        //         fmt.Fprintf(w,"hello")
        // }
        // else {
                // fmt.Println("2...",r.URL.String())
                // if r.URL.String()=="/ws/" {
        //                 // fmt.Println("3...",r.URL.String())
        //                 url := &url.URL{Scheme: "ws://", Host: "127.0.0.1:3333"}
	//                 p := wsutil.NewSingleHostReverseProxy(url)
        //                 r.URL.Host = url.Host
        //                 r.URL.Scheme = url.Scheme
        //                 r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        //                 r.Host = url.Host
        //                 p.ServeHTTP(w, r)
        //         }
        // }
// }

// func main() {
//         srv := &http.Server{
// 		Addr:":"+goPort("22222"),
// 		WriteTimeout:5*time.Second,
// 		ReadTimeout:5*time.Second,
// 	}
//         http.HandleFunc("/", Handler)
//         log.Fatal(srv.ListenAndServe())
// }