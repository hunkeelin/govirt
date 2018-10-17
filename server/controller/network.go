package controller
import(
    "github.com/hunkeelin/govirt/govirtlib"
    "errors"
    "fmt"
    "io/ioutil"
    "github.com/hunkeelin/mtls/klinreq"
    "github.com/hunkeelin/klinutils"
)
func editnetwork(n govirtlib.Network) error {
    p := govirtlib.PostPayload {
        Target: "network",
        Netinfo: n,
    }
    i := &klinreq.ReqInfo {
        Dest: "sf01-lab-1.squaretrade.com",
        Dport: klinutils.Stringtoport("godhcp"),
        Method: "PATCH",
        Payload: p,
        Http: true,
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
func edithost(n govirtlib.CreateVmForm) error {
    p := govirtlib.PostPayload {
        Target: "host",
        VmForm: n,
    }
    i := &klinreq.ReqInfo {
        Dest: "sf01-lab-1.squaretrade.com",
        Dport: klinutils.Stringtoport("godhcp"),
        Method: "POST",
        Payload: p,
        Http: true,
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
