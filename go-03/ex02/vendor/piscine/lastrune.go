/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   lastrune.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 12:06:08 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func LastRune(s string) rune {
	slice := []rune(s)
	if len(slice) == 0 {
		return 0
	}
	return slice[len(slice) - 1]
}
