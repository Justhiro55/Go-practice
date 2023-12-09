/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Atoi.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/09 15:03:27 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 15:12:39 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine
import (
	// "fmt"
)
func Atoi(s string) int {
    var result int = 0
	var tmp		int64 = 0
	var minus	int = 1
	var count 	int = 0

	for _, char := range s {
		if char == '+' || char == '-'{
			if count != 0{
				return 0
			} else{
				if(char == '-'){
					minus = -1
				}
			}
		} else if char < '0' || char > '9' {
            return 0
        } else {
			tmp = int64(int64(result*10) + int64(char-'0'))
			if tmp > 2147483647 {
				return 0
			}
			result = result*10 + int(char-'0')
		}
		count++
    }
    return result*minus
}
