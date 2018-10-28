package controller

import ()

func (c *Conn) post(w http.ResponseWriter, r *http.Request) error {
	var p govirtlib.PostPayload
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("unable to read response body post")
		return err
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("unable to unmarshal json post")
		return err
	}
	switch strings.ToLower(p.Action) {
	case "host":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err := start(p.Domain, c.L)
		if err != nil {
			return err
		}
	case "env":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err = shutdown(p.Domain, c.L)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
