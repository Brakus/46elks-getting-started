package main

import "net/url"
import "io/ioutil"
import "fmt"
import "net/http"
import "bytes"
import "strconv"

func main() {

    fmt.Println("I will now try to allocate a new number!")

    data := url.Values{
        "country": {"se"},
        //"sms_url": {"http://www.example.com/sms"}, //Optional
        //"mms_url": {"http://www.example.com/mms"}, //Optional
        //"voice_start": {"http://www.example.com/calls"}, //Optional
        //"capabilities": {"sms", "voice", "mms" }, //Optional
    }

    req, err := http.NewRequest("POST", "https://api.46elks.com/a1/Numbers", bytes.NewBufferString(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    req.SetBasicAuth("<API-username>", "<API-password>")

    client := &http.Client{}
    resp, err := client.Do(req)

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Oh dear!!!")
    }

    fmt.Println(string(body))

}