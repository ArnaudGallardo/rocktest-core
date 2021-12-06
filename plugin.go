package core

import (
	"errors"
	"fmt"
	"plugin"
)

func (module *Module) Plugin(params map[string]interface{}, scenario *Scenario) error {

	paramsEx, err := scenario.ExpandMap(params)
	if err != nil {
		return err
	}

	val, err := scenario.GetString(paramsEx, "value", nil)
	if err != nil {
		return err
	}

	fmt.Println("toto")
	p, err := plugin.Open(val)
	if err != nil {
		panic(err)
	}
	fmt.Println("test")
	name, err := p.Lookup("Name")
	if err != nil {
		return errors.New("Plugin is missing the Name variable.")
	}
	nameStr := *name.(*string)
	fmt.Println(*name.(*string))
	f, err := p.Lookup(nameStr)
	if err != nil {
		panic(err)
	}
	module.AddPluginModule(nameStr, f.(func(map[string]interface{}, *Scenario) error))
	// f.(func(map[string]interface {}) error)(params) // prints "Hello, number 7"

	return nil
}
