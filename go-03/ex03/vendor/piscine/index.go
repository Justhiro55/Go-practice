/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   index.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:02:52 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Index(s string, toFind string) int {
	for i:=0; i <= len(s) - len(toFind); i++ {
		if toFind == s[i: i+len(toFind)]{
			return i
		}
	}
	return -1
}
