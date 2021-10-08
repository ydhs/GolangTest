package main

import ("fmt"
        "os"
        "os/exec"
        )

func main() {
// начало. Проверка на количество аргументов
if len(os.Args) == 1 {
  fmt.Println("Добавьте аргумент\n\n")
  fmt.Println("Версия программы 1.0")
  os.Exit(0)
  } else if len(os.Args) > 2 {
    fmt.Println("Нужен только один аргумент")
    os.Exit(0)
  }
// конец. Проверка на количество аргументов



switch os.Args[1] {
case "NBAZDB":
  status, err := exec.Command("bash", "-c", "/usr/openv/db/bin/nbdb_ping -dbn NBAZDB").Output()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(status)
case "NBDB":
  status, err := exec.Command("bash", "-c", "/usr/openv/db/bin/nbdb_ping -dbn NBDB").Output()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(status)
case "BMRDB":
  status, err := exec.Command("bash", "-c", "/usr/openv/db/bin/nbdb_ping -dbn BMRDB").Output()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(status)
}
}
