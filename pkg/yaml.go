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

func CheckYamlFormat(data string,obj interface{}) error {
	if err := yaml.Unmarshal([]byte(data), &obj); err != nil {
        return err 
    }
	return nil 
}