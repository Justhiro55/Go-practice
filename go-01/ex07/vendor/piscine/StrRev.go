/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   StrRev.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 15:07:48 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 14:20:10 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func StrRev(s string) string {
    runes := []rune(s)
    n := len(runes)
    reversed := make([]rune, n)
    for i := 0; i < n; i++ {
        reversed[i] = runes[n-1-i]
    }
    return string(reversed)
}