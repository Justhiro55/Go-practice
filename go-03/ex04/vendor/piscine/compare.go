/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   compare.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:13:28 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Compare(a, b string) int {
    lenA := len(a)
    lenB := len(b)    
    minLen := lenA
    if lenB < lenA {
        minLen = lenB
    }
    for i := 0; i < minLen; i++ {
        c1 := a[i]
        c2 := b[i]
        
        if c1 > c2 {
            return 1
        } else if c2 > c1 {
			return -1
		}
    }
	if lenA > lenB {
		return 1
	} else if lenB > lenA {
		return -1
	}
	return 0
}