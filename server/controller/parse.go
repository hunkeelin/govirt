package controller

import (
	"bytes"
	"errors"
	"io/ioutil"
)

type ClusterInfo struct {
	ClusterName string
	Godhcp      string
	Govirt      []string
	Storage     string
}

func Parse(f string) (map[string]ClusterInfo, error) {
	m := make(map[string]ClusterInfo)
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return m, err
	}
	blines := bytes.Split(b, []byte("\n"))
	var current string
	var split [][]byte
	var tmp ClusterInfo
	for _, lines := range blines {
		clines := bytes.Replace(lines, []byte(" "), []byte(""), -1)
		switch {
		case bytes.HasPrefix(clines, []byte("cluster")):
			split = bytes.Split(clines, []byte(":"))
			if len(split) != 2 {
				return m, errors.New("Syntax error on cluster")
			}
			current = string(split[1])
			tmp = m[current]
			tmp.ClusterName = current
			m[current] = tmp
		case bytes.HasPrefix(clines, []byte("-dhcpd")):
			split = bytes.Split(clines, []byte(":"))
			if len(split) != 2 {
				return m, errors.New("Syntax error on dhcpd")
			}
			tmp = m[current]
			tmp.Godhcp = string(split[1])
			m[current] = tmp
		case bytes.HasPrefix(clines, []byte("-govirt")):
			split = bytes.Split(clines, []byte(":"))
			if len(split) != 2 {
				return m, errors.New("Syntax error on govirt")
			}
			govirthosts := bytes.Split(split[1], []byte(","))
			var govirts []string
			for _, hosts := range govirthosts {
				govirts = append(govirts, string(hosts))
			}
			tmp = m[current]
			tmp.Govirt = govirts
			m[current] = tmp
		case bytes.HasPrefix(clines, []byte("-storage")):
			split = bytes.Split(clines, []byte(":"))
			if len(split) != 2 {
				return m, errors.New("Syntax error on storage")
			}
			tmp = m[current]
			tmp.Storage = string(split[1])
			m[current] = tmp
		default:
			continue
		}
	}
	return m, nil
}
