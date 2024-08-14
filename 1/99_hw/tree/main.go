package main

import (
  "fmt"
  "io"
  "os"
  "strings"
  "strconv"
)

const (
  be = "\t"
  bl = "│\t"
  ms = "├───"
  ls = "└───"
)

func isLastDir(list []os.DirEntry, name string, printFiles bool) bool {
  if printFiles && name == list[len(list) - 1].Name() { return true }
  for i := len(list) - 1; i >= 0; i-- {
    if list[i].IsDir() {
      if name == list[i].Name() && i != 0{ return !printFiles }
      break
    }
  } 
  return false
}

func dirTree(out io.Writer, path string, printFiles bool) error {
  var str0 string
  spath := strings.Split(path, "/")
  if len(spath) > 1 {
    var tpath string
    for i := 0; i < len(spath) - 1; i++ {
      tpath += spath[i]
      files, err := os.ReadDir(tpath)
      if err != nil { return err }
      if isLastDir(files, spath[i + 1], printFiles) {
        str0 += be
      } else { str0 += bl }
      tpath += "/"
    }
  }
  files, err := os.ReadDir(path)
  if err != nil { return err }
  if len(files) == 0 { return nil }
  for _, file := range files {
    str := str0
    if !isLastDir(files, file.Name(), printFiles) {
      str += ms
    } else { 
      str += ls 
    }
    if file.IsDir() {
      str += file.Name()
  	  fmt.Fprintln(out, str)
      npath := path + "/" + file.Name()
      dirTree(out, npath, printFiles)
    } else if printFiles {
      finfo, err := file.Info()
      if err != nil { return err }
      var size string
      if finfo.Size() == 0 {
        size = "empty"
      } else {
        size = strconv.Itoa(int(finfo.Size())) + "b"
      }
      str += file.Name() + " (" + size + ")"
      fmt.Fprintln(out, str)
    }
  }
  return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
