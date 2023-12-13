/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isupper.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:46:08 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func IsLower(s string) bool {
    for _, c := range s {
        if !('a' <= c && c <= 'z') {
            return false
        }
    }
    return true
}