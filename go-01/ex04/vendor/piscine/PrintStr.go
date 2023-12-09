/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   PrintStr.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 14:03:38 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 13:36:32 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"fmt"
)

func PrintStr(a string) {
    for _, char := range a {
        fmt.Print(string(char))
    }
}
