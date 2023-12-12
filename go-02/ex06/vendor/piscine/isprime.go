/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isprime.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/12 21:51:02 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IsPrime(nb int) bool {
	if nb == 1 || nb < 0{
		return false
	} else if ( nb <= 3){
		return true
	}
	for i:= 2; i < nb / 2 + 1; i++{
		if nb % i == 0{
			return false
		}
	}
	return true
}