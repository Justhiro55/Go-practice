/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   recursivepower.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/12 19:29:22 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func RecursivePower(nb int, power int) int {
	var num int = nb

	if power == 0{
		return 1
	} else if nb == 0 || nb < 0 {
		return 0
	} else if power == 1 {
		return nb
	}
	num = num * RecursivePower(nb, power - 1)
	return num
}
