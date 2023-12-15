/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   makerange.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 14:41:13 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func MakeRange(min, max int) []int {
    if min >= max {
        return nil
    }
    result := make([]int, max-min)
    
    for i := min; i < max; i++ {
        result[i-min] = i
    }
    return result
}