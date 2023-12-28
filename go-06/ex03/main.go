/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/28 11:54:50 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 1 {
        copyStdin(os.Stdin)
        return
    }
    filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil{
        fmt.Println("ERROR: open ", filename, ": no such file or directory")
        return
    }
    copyFile(file)
    defer file.Close()
}

func copyStdin(src *os.File) {
    buffer := make([]byte, 10240)
    
    for {
        n, err := src.Read(buffer)
        if err != nil {
            break
        }
        if n == 0 {
            break
        }
        fmt.Print(string(buffer[:n]))
    }
}

func copyFile(src *os.File) {
    buffer := make([]byte, 10240)
    
    for {
        n, err := src.Read(buffer)
        if err != nil {
            break
        }
        if n == 0 {
            break
        }
        fmt.Println(string(buffer[:n]))
    }
}
