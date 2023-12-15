/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   appendrange.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 14:32:43 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func AppendRange(min, max int) []int {
    if min >= max {
        return nil
    }

    result := []int{}
    for i := min; i < max; i++ {
        result = append(result, i)
    }
    
    return result
}