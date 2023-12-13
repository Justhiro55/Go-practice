/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isupper.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:44:59 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IsUpper(s string) bool {
    for _, c := range s {
        if !('A' <= c && c <= 'Z') {
            return false
        }
    }
    return true
}