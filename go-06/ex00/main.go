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

const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"

func printStr(s string) {
    for _, r := range s {
        fmt.Print(string(r))
    }
    fmt.Println()
}

func isEven(nbr int) bool {
    return nbr%2 == 0
}

func main() {
    lengthOfArg := len(os.Args[1:])
    if isEven(lengthOfArg) {
        printStr(EvenMsg)
    } else {
        printStr(OddMsg)
    }
}