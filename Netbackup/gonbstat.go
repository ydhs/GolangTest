package main

import ("fmt"
        "os"
        "os/exec"
        "strings"
        )

func Stat_DB(arg string)  {
  statusByte, err:= exec.Command("bash", "-c", "/usr/openv/db/bin/nbdb_ping -dbn " + arg).Output()

  if err != nil {
    fmt.Println(err)
    return
  }

  statusString := string(statusByte)

  if strings.Contains (statusString, "Database ["+ arg +"] is alive and well") {
    fmt.Println("OK")
    fmt.Println("Database ["+ arg +"] is alive and well")
    } else {
      fmt.Println("Статус не норма: "+ statusString)
    }
}

func main() {
help0 := "Версия программы 1.0\n\n Добавьте один аргумент: NBAZDB, NBDB или BMRDB\n\n"

// начало. Проверка на количество аргументов и на сами аргументы
if len(os.Args) == 1 {
    fmt.Println(help0)

  os.Exit(0)
  } else if len(os.Args) > 2 {
    fmt.Println("Нужен только один аргумент")
    os.Exit(0)
  } else if (os.Args[1] != "NBDB" && os.Args[1] != "NBAZDB" && os.Args[1] != "BMRDB") {
    fmt.Println(help0)
    os.Exit(0)
  }
// конец. Проверка на количество аргументов  и на сами аргументы


Stat_DB(os.Args[1])
}
