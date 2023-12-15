/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   makerange.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 14:52:00 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func ConcatParams(args []string) string {
    if args == nil {
        return "\n"
    }
    var str string
    for i := 0; i < len(args); i++ {
        str = str + args[i]
        if i+1 < len(args) {
            str = str + "\n"
        }
    }
    return str
}