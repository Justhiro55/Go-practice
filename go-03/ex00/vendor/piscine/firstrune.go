/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   firstrune.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:58:32 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 11:56:54 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func FirstRune(s string) rune {	
	slice := []rune(s)
	return slice[0]
}