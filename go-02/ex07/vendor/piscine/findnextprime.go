/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   findnextprime.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/12 21:50:28 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func is_prime(nb int) bool {
	for i:= 2; i < nb / 2 + 1; i++{
		if nb % i == 0{
			return false
		}
	}
	return true
}

func find_next_prime(nb int) int {
	if nb % 2 == 0{
		nb++
	}
	for {
		if is_prime(nb) == true {
			return nb
		}
		nb += 2
	}
	return -1
}

func FindNextPrime(nb int) int {
	if nb == 1 || nb < 0{
		return 2
	} else if ( nb <= 3){
		return nb
	}
	if is_prime(nb) == true{
		return nb
	}
	return find_next_prime(nb)
}