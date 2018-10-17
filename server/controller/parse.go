package controller
import(
    "bytes"
    "io/ioutil"
)
type clusterInfo struct {
    ClusterName string
    Godhcp  string
    Govirt []string
    Storage string
}
func parse(f string) (map[string]clusterInfo, error){
    m := make(map[string]clusterInfo)
    b,err := ioutil.ReadAll(f)
    if err != nil {
        return m,err
    }
 
}
