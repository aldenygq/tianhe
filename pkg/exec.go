package pkg
import (
	"os/exec"
	"bytes"
)

func RunCmd(script string) (string,error) {
    cmd := exec.Command("/bin/bash", script)
    var stdin, stdout, stderr bytes.Buffer
    cmd.Stdin = &stdin
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
       return "",err 
    }
    outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
    return outStr,nil 
}