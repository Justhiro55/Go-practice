/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicjoin.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:57:32 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Join(strs []string, sep string) string {
    result := ""
    for i, str := range strs {
        result += str
        if i < len(strs)-1 {
            result += sep
        }
    }
    return result
}