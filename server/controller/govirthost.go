package controller
import(
    "github.com/hunkeelin/mtls/klinreq"
    "github.com/hunkeelin/govirt/govirtlib"
    "io/ioutil"
    "fmt"
    "errors"
    "github.com/hunkeelin/klinutils"
)
func createvm(xml []byte, dest string) error {
    p := govirtlib.PostPayload {
        Action: "define",
        Xml: xml,
    }
    i := &klinreq.ReqInfo{
        Dest:    dest,
        Dport:   klinutils.Stringtoport("govirthost"),
        Method:  "POST",
        Payload: p,
        Http:    true,
    }
    resp, err := klinreq.SendPayload(i)
    if err != nil {
        return err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    resp.Body.Close()
    if resp.StatusCode != 200 {
        fmt.Println(string(body))
        return errors.New("Failed, check logs on the godhcp server")
    }
    return nil
}
