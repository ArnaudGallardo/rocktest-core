package core

import (
	"errors"
	"plugin"
	"strings"

	log "github.com/sirupsen/logrus"
)

type PluginInterface struct {
	Name    string
	Version string
	Author  string
	Methods []string
}

func (module *Module) Plugin(params map[string]interface{}, scenario *Scenario) error {

	paramsEx, err := scenario.ExpandMap(params)
	if err != nil {
		return err
	}

	val, err := scenario.GetString(paramsEx, "value", nil)
	if err != nil {
		return err
	}

	p, err := plugin.Open(val)
	if err != nil {
		panic(err)
	}
	metadata, err := p.Lookup("PluginMetadata")
	if err != nil {
		return errors.New("Plugin is missing the Name variable.")
	}
	pluginMetadata := *metadata.(*PluginInterface)

	// Display plugin infos
	log.Debugf("Loading plugin %s - version %s - author: %s\n", pluginMetadata.Name, pluginMetadata.Version, pluginMetadata.Author)

	// Load methods
	for _, v := range pluginMetadata.Methods {
		f, err := p.Lookup(v)
		if err != nil {
			return err
		}
		name := strings.ReplaceAll(v, ".", "_")
		name = strings.ToLower(name)
		name = strings.Title(name)
		module.AddPluginModule(name, f.(func(map[string]interface{}, *Scenario) error))
	}

	return nil
}
