package core

type Module struct {
	plugins map[string]func(map[string]interface{}, *Scenario) error
}

func (module *Module) GetPluginModule(name string) (func(map[string]interface{}, *Scenario) error, bool) {
	v, found := module.plugins[name]
	return v, found
}

func (module *Module) AddPluginModule(name string, v func(map[string]interface{}, *Scenario) error) {
	if module.plugins == nil {
		// Init map
		module.plugins = make(map[string]func(map[string]interface{}, *Scenario) error)
	}
	module.plugins[name] = v
}
