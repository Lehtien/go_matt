package main

import (
      "net/http"
      "github.com/labstack/echo"
)
import (
   "fmt"
   "io/ioutil"
   "path/filepath"
   "os"
   "encoding/base64"
)


func main () {
   e := echo.New()
   e.GET("/", func(c echo.Context) error {
      paths := readDir("./image")
      images := getImages(paths)
      return c.JSON(http.StatusOK, images)
   })
   e.Logger.Fatal(e.Start(":8082"))

}

func readDir(dir string) []string {
   files, err := ioutil.ReadDir(dir)
   if err != nil {
      panic(err)
   }

   var paths []string
   for _, file := range files {
      paths = append(paths, filepath.Join(dir, file.Name()))
   }
   return paths
}

func getImages(paths []string) []string {
   var dataURLs []string
   for _, value := range paths {
      file, _ := os.Open(value)
      defer file.Close()

      body, _ := ioutil.ReadAll(file)
      str := base64.StdEncoding.EncodeToString(body)

      dataURLs = append(dataURLs, fmt.Sprintf("data:image/png;base64,%s", str))
   }
   return dataURLs
}