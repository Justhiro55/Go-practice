/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 14:26:48 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("File name missing")
        return
    }
    filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    buffer := make([]byte, 10240)
    for {
        n, err := file.Read(buffer)
        if n == 0{
            break
        }
        if err != nil{
            break
        }
        fmt.Println(string(buffer[:n]))
    fmt.Print(string(buffer))
    }
}