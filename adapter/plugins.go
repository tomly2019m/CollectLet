package adapter

import "sync"

type Plugin interface {
	Process(data interface{}) interface{}
}

type DefaultPlugin struct {
	Plugin
}

func (plugin *DefaultPlugin) Process(data interface{}) interface{} {
	return data
}

type PluginRegistry struct {
	plugins map[string]Plugin
	mu      sync.RWMutex
}

var registry = &PluginRegistry{plugins: make(map[string]Plugin)}

func (r *PluginRegistry) Register(name string, plugin Plugin) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.plugins[name] = plugin
}

// Get 获取插件，如果没有实现插件则返回默认插件
func (r *PluginRegistry) Get(name string) Plugin {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if plugin, exists := r.plugins[name]; exists {
		return plugin
	}
	return &DefaultPlugin{}
}

// Middleware 中间件
type Middleware func(data interface{}) interface{}

// PluginManager 插件管理器，将所有插件串联起来，当启用插件时，数据依次通过所有的插件，并传递给下一个插件或下一层
type PluginManager struct {
	middlewares []Middleware
	mu          sync.Mutex
}

// Use 将某个插件添加到插件列表中，从而启用插件
func (pm *PluginManager) Use(middleware Middleware) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.middlewares = append(pm.middlewares, middleware)
}

// Process 数据依次通过所有插件
func (pm *PluginManager) Process(data interface{}) interface{} {
	for _, middleware := range pm.middlewares {
		data = middleware(data)
	}
	return data
}
