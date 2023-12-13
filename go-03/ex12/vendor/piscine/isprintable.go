/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isprintable.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:51:31 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IsPrintable(s string) bool {
    for _, c := range s {
        if c < 32 || 126 < c {
            return false
        }
    }
    return true
}
