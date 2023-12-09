/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   SortIntegerTable.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/09 15:18:07 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 15:29:31 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine
import (
// "fmt"
)
func SortIntegerTable(table []int) {
	n := len(table)
	for i:= 0; i < n - 1; i++{
		for j:= i + 1; j < n; j++{
			if table[i] > table[j]{
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}