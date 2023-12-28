/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/28 13:09:49 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] != "-c" {
        fmt.Println("error")
        return
    }
    for i := 2; i < len(os.Args); i++ {
        filename := os.Args[i]
        file, err := os.Open(filename)
        if err != nil {
            continue
        }
        fmt.Println("==>", os.Args[i], "<==")
        ztail(file)
        defer file.Close()
        if i + 1 < len(os.Args) {
            fmt.Println("")
        }
    }
}

func ztail(file *os.File) {
    var lastLine string    
    buffer := make([]byte, 10240)
    
    n, err := file.Read(buffer)
    if err != nil || n == 0 {
        fmt.Println("")
        return
    }
    for i := n - 1; i > 0; i-- {
        if buffer[i] == '\n' {
            fmt.Println(lastLine)
            return
        }
        lastLine = string(buffer[i]) + lastLine
    }
    fmt.Println(lastLine)
}