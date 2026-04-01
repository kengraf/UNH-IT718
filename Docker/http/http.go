package main

import (
    "net"
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello from: " + GetLocalIP() + "\n" )
}

func wait(w http.ResponseWriter, req *http.Request) {

    var t = time.Duration(rand.Intn(10)) * time.Second
    time.Sleep( t )
    fmt.Fprintf(w, GetLocalIP() + " delayed for " + t.String() + " seconds \n" )
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func dowork(w http.ResponseWriter, req *http.Request) {

    var sum = 0
    for i := 1; i < 20000; i++ {
        for j := 1; j < i; j++ {
            sum += i * rand.Intn(j)
        }
    }
    fmt.Fprintf(w, "done %d\n", sum )
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/wait", wait)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/dowork", dowork)

    http.ListenAndServe(":8090", nil)
}
