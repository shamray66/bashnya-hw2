package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scan(&num)
    if (num < 12307) {
        for (num < 12307) {
            if !(num % 13 == 0 && num % 9 == 0) {
                if (num < 0) {
                    num *= -1
                } else if (num % 7 == 0) {
                    num *= 39
                } else if (num % 9 == 0) {
                    num *= 13
                    num +=1
                    continue
                } else {
                    num += 2
                    num *= 3
                }
                /*
                Для ревьюера башни: я честно не особо смог представить структуру 
                бинарного дерева (Именно в моменте с блоком 2), поэтому 
                сделал так, как сделал. Следующая инструкция относится именно ко 2 блоку.
                */
                num += 1
            } else {
                fmt.Println("service error")
                return
            }
        }
        fmt.Println("Number after all manips: ", num)
    } else {
        fmt.Println("Number >= 12307\nWe dont want work with this number.")
    }
}