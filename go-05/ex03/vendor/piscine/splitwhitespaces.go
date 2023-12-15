/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   splitwhitespaces.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 15:04:11 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func SplitWhiteSpaces(s string) []string {
    result := []string{}
    tmp := ""
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            tmp += string(s[i])
        } else {
            if tmp != "" {
                result = append(result, tmp)
            }
            tmp = ""
        }
    }

    if tmp != "" {
        result = append(result, tmp)
    }

    return result
}