/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   iterativepower.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/12 19:24:54 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IterativePower(nb int, power int) int {
	var tmp int = nb
	if power == 0 || (nb < 0 && power == 0){
		return 1
	} else if (nb == 0 && power != 0) || power < 0 {
		return 0
	} else if power == 1{
		return nb
	}
	for i := 0; i < power; i++ {
		tmp = tmp * nb
		i++
	}
	return tmp
}
