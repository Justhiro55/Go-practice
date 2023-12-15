/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   convertbase.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 17:00:45 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func convertToDecimal(nbr, baseFrom string) int {
	decimalValue := 0
	baseLen := len(baseFrom)

	for _, digit := range nbr {
		decimalValue = decimalValue * baseLen + indexOf(baseFrom, digit)
	}

	return decimalValue
}

func convertFromDecimal(value int, baseTo string) string {
	base := len(baseTo)
	result := ""

	for value > 0 {
		tmp := value % base
		result = string(baseTo[tmp]) + result
		value = value / base
	}

	return result
}

func indexOf(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
    value := convertToDecimal(nbr, baseFrom)

    result := convertFromDecimal(value, baseTo)
    return result
}
