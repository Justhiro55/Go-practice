/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoibase.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 15:50:46 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

// import "fmt"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	charSet := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || charSet[char] == true {
			return false
		}
		charSet[char] = true
	}

	return true
}

func isValidDigit(char rune, base string) bool {
	for _, validChar := range base {
		if char == validChar {
			return true
		}
	}
	return false
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, char := range s {
		if !isValidDigit(char, base) {
			return 0
		}

		value := 0
		for i, validChar := range base {
			if char == validChar {
				value = i
				break
			}
		}

		result = result*baseLen + value
	}

	return result
}