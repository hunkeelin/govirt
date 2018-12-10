package controller
import(
    "io/ioutil"
    "errors"
    "github.com/hunkeelin/klinutils"
    "fmt"
    "github.com/hunkeelin/mtls/klinreq"
    "github.com/hunkeelin/govirt/govirtlib"
)
func (c *Conn) setimage(dest,image,hostname string) error {
    vm := govirtlib.CreateVmForm {
        Image: image,
        Hostname: hostname,
    }
    p := &govirtlib.PostPayload {
        Action: "setimage",
        VmForm: vm,
    }
    i := &klinreq.ReqInfo {
        Dest: dest,
        Dport: klinutils.Stringtoport("storagehost"),
        Method: "POST",
        Payload: p,
        TrustBytes: c.Tb,
        CertBytes: c.Cb,
        KeyBytes: c.Kb,
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
        return errors.New("Failed, check logs on the storage server")
    }
    return nil
}
func (c *Conn) delimage(dest,hostname string) error {
    p := &govirtlib.PostPayload {
        Action: "host",
        Target: hostname,
    }
    i := &klinreq.ReqInfo {
        Dest: dest,
        Dport: klinutils.Stringtoport("storagehost"),
        Method: "DELETE",
        Payload: p,
        TrustBytes: c.Tb,
        CertBytes: c.Cb,
        KeyBytes: c.Kb,
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
        return errors.New("Failed, check logs on the storage server")
    }
    return nil
}
func (c *Conn)storagedup(dest string,d map[string]int) error {
    p := &govirtlib.PostPayload {
        Action: "dup",
        DuplicateInfo: d,
    }
    i := &klinreq.ReqInfo {
        Dest: dest,
        Dport: klinutils.Stringtoport("storagehost"),
        Method: "POST",
        Payload: p,
        TrustBytes: c.Tb,
        CertBytes: c.Cb,
        KeyBytes: c.Kb,
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
        return errors.New("Failed, check logs on the storage server")
    }
    return nil
}
