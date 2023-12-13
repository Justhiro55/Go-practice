/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   nrune.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 12:02:04 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func NRune(s string, n int) rune {
	if n < 0 || n > len([]rune(s)) {
		return 0
	}
	slice := []rune(s)
	return slice[n - 1]
}
