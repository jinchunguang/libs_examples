/**
 * @Author : jinchunguang
 * @Date : 19-10-31 上午11:28
 * @Project : sty
 */
package main

import (
    "bytes"
    "log"
    "os/exec"
)

func main() {

    // out, _ := UseOutput("/bin/sh", "-c", "ls -l .")
    // log.Println(string(out))


    // out, _ := FillStd("/bin/sh", "-c", "ls -l .")
    // log.Println(string(out))

    out, _ := FillStd("/bin/sh", "-c", "ls -l .")
    log.Println(string(out))
}
func FillStd(name string, arg ...string) ([]byte, error) {
    cmd := exec.Command(name, arg...)
    var out = new(bytes.Buffer)
    var outerr = new(bytes.Buffer)

    cmd.Stdout = out
    cmd.Stderr = outerr

    err := cmd.Run()
    if err != nil {
        return nil, err
    }

    log.Println("xx",outerr)
    return out.Bytes(), nil
}

func UseOutput(name string, arg ...string) ([]byte, error) {
    return exec.Command(name, arg...).Output()
}
func UsePipe(name string, arg ...string) ([]byte, error) {
    cmd := exec.Command(name, arg...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return nil, err
    }

    if err = cmd.Start(); err != nil {
        return nil, err
    }

    var out = make([]byte, 0, 1024)
    for {
        tmp := make([]byte, 128)
        n, err := stdout.Read(tmp)
        out = append(out, tmp[:n]...)
        if err != nil {
            break
        }
    }

    if err = cmd.Wait(); err != nil {
        return nil, err
    }

    return out, nil
}
