/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printnbrbase.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 14:42:22 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "fmt"

func check_arg(base string) bool {
    charCount := make(map[rune]int)
    if len(base) < 2{
        return false
    }
    for _, char := range base {
        charCount[char]++
        if(charCount[char] > 1) {
            return false
        }
    }
    return true
}

func PrintNbrBase(nbr int, base string) {
    baseLen := len(base)
    var result string

    if !check_arg(base) {
        fmt.Print("NV")
        return
    }
    if nbr < 0 {
        fmt.Print("-")
        nbr = -nbr
    } else if nbr == 0 {
        fmt.Print("0")
        return
    }
    for nbr != 0 {
        index := nbr % baseLen
        result = string(base[index]) + result
        nbr = nbr / baseLen
    }
    fmt.Print(result)
}
