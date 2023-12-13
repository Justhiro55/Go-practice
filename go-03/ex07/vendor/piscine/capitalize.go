/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   capitalize.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:37:55 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Capitalize(s string) string {
    result := ""
    flag := 1
    for _, c := range s {
        if flag == 1 {
            if 'a' <= c && c <= 'z' {
                result += string(c - 'a' + 'A')
                flag = 0
            } else {
                if !(('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')) {
                    flag = 1
                } else {
                    flag = 0
                }
                result += string(c)
            } 
        } else {
            if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9'){
                result += string(c)
            } else if ('A' <= c && c <= 'Z'){
                result += string(c - 'A' + 'a')
            } else {
                result += string(c)
                flag = 1
            }
        }
    }
    return result
}