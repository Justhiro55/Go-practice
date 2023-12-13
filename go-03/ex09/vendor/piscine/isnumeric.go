/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isnumeric.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:42:02 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IsNumeric(s string) bool {
    for _, c := range s {
        if !('0' <= c && c <= '9') {
            return false
        }
    }
    return true
}