/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicjoin.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 13:58:58 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func BasicJoin(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}