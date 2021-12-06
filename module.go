package core

type Module struct {
	plugins map[string]func(map[string]interface{}, *Scenario) error
}

func (module *Module) GetPluginModule(name string) (func(map[string]interface{}, *Scenario) error, error) {
	v, found := module.plugins[name]
	if found {
		return v, nil
	}
	return nil, nil
}

func (module *Module) AddPluginModule(name string, v func(map[string]interface{}, *Scenario) error) {
	module.plugins[name] = v
}
