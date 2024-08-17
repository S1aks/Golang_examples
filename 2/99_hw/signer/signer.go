package main
import (
  "fmt"
)

func SingleHash(in, out chan interface{}) {
  data := <-in
  crc32 := DataSignerCrc32(data.(string))
  crc32md5 := DataSignerCrc32(DataSignerMd5(data.(string)))
  result := crc32 + "~" + crc32md5
  out <- result
}

func MultiHash(in, out chan interface{}) {

}

func CombineResults(in, out chan interface{})  {
  
}

func ExecutePipeline(jobs... job) {
  fmt.Println()
}
