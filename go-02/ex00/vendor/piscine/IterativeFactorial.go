/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   IterativeFactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:58:32 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/10 15:58:40 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IterativeFactorial(nb int) int {
	result := 1
	var tmp		int64 = 1
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for nb > 1 {
		tmp = tmp * int64(nb)
		if tmp > 2147483647 || tmp < -2147483648{
			return 0
		}
		result = result * nb
		nb = nb - 1
	}
	return result
}