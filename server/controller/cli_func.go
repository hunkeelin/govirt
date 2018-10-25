package controller

import ()

func (c *Conn) Migratehost(ori, dest, vm string) error {
	err := c.migrate(ori, dest, vm)
	if err != nil {
		return err
	}
	xml, err := c.Getxml(vm, ori)
	if err != nil {
		return err
	}
	err = c.Statevm("unDefine", vm, ori)
	if err != nil {
		return err
	}
	err = c.Define(xml, dest)
	if err != nil {
		return err
	}
	return nil
}
