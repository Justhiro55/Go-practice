/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   BasicAtoi.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 17:34:12 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 14:59:46 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func BasicAtoi(s string) int {
    var result int = 0
	var tmp		int64 = 0
	for _, char := range s {
        if char < '0' || char > '9' {
            return 0
        }
		tmp = int64(int64(result*10) + int64(char-'0'))
		if tmp > 2147483647 {
			return 0
		}
        result = result*10 + int(char-'0')
    }
    return result
}