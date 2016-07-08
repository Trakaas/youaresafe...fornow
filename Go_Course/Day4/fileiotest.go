package main

import (
  "fmt"
  "io"
  "os"
)

func readFile(name string) ([]byte, error) {
  buf := make([]byte,4096)
  var data []byte
  fid, err := os.Open(name)
  if err != nil {
    fmt.Println("Problem opening file.")
    return nil, err
  }

  for {
    num,err := fid.Read(buf)
    if err != nil {
      if err == io.EOF {
        return data, nil
      }
      return nil, err
    }

    data = append(data, buf[:num]...)
  }
  return nil, nil
}

func writeFile (name string, data []byte) error {
  f, err := os.Create(name)
  if err != nil {
    return err
  }

  _, err = f.Write([]byte("hello"))
  if err != nil {
    return err
  }
  return nil
}



func main() {
  err := writeFile("example.text",make([]byte,8))
  if err != nil {
    fmt.Println("Error")
    return
  }
  fmt.Println()
}
