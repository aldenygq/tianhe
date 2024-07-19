package pkg
import (
	"gopkg.in/yaml.v2"
)
func ToYAML(obj interface{}) (string, error) {
    out, err := yaml.Marshal(obj)
    if err != nil {
        return "", err
    }
    return string(out), nil
}