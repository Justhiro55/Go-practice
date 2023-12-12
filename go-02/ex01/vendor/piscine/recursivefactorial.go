/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   recursivefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/10 16:20:10 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	tmp := int64(nb) * int64(RecursiveFactorial(nb - 1))
	if tmp > 9223372036854775807 || tmp < -9223372036854775808 {
		return 0
	}
	return int(tmp)
}
