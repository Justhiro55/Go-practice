/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   split.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/15 14:27:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/15 17:42:20 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Split(s, sep string) []string {
	result := []string{}
	tmp := ""

	if sep == "" {
		result = append(result, s)
		return result
	}
	for i := 0; i < len(s); i++ {
		if len(sep) == 1 && s[0] == sep[0] && i == 0 {
			continue
		}
		if s[i] != sep[0]{
			tmp += string(s[i])
		} else {
			if s[i] == sep[0] {
				match := true
				for j := 0; j < len(sep); j++ {
					if i+j >= len(s) || s[i+j] != sep[j]{
						match = false
						break
					}
				}
				if match && tmp != ""{
					result = append(result, tmp)
					tmp = ""
					i += len(sep) - 1
				} else {
					tmp += string(s[i])
				}
			}
		}
	}
	if tmp != "" {
		result = append(result, tmp)
	}
	return result
}