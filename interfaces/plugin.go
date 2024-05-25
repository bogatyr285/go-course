package main

import (
	"fmt"
	"strings"
)

// Plugin defines the interface that all plugins must implement
type Plugin interface {
	Execute(data interface{}) interface{}
}

// register a map to hold the plugin instances
var pluginRegistry = make(map[string]Plugin)

// RegisterPlugin adds a plugin to the registry map with a given name
func RegisterPlugin(name string, plugin Plugin) {
	pluginRegistry[name] = plugin
}

// GetPlugin retrieves a plugin from the registry by name
func GetPlugin(name string) Plugin {
	return pluginRegistry[name]
}

type UppercasePlugin struct{}

func (p *UppercasePlugin) Execute(data interface{}) interface{} {
	str, ok := data.(string)
	if !ok {
		return "Invalid data"
	}
	return strings.ToUpper(str)
}

// ReversePlugin reverses a string
type ReversePlugin struct{}

func (p *ReversePlugin) Execute(data interface{}) interface{} {
	str, ok := data.(string)
	if !ok {
		return "Invalid data"
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {

	RegisterPlugin("uppercase", &UppercasePlugin{})
	RegisterPlugin("reverse", &ReversePlugin{})

	data := "hello, world!"

	upperPlugin := GetPlugin("uppercase")
	if upperPlugin != nil {
		fmt.Println("Uppercase Plugin Output:", upperPlugin.Execute(data))
	} else {
		fmt.Println("Uppercase Plugin not found")
	}

	reversePlugin := GetPlugin("reverse")
	if reversePlugin != nil {
		fmt.Println("Reverse Plugin Output:", reversePlugin.Execute(data))
	} else {
		fmt.Println("Reverse Plugin not found")
	}
}
