/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   PrintComb.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/02 17:12:45 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/10/02 19:12:49 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "fmt"

func PrintComb(){
	for i := 0; i < 10; i++ {
		for j := i+1; j < 10; j++ {
			for k := j+1; k < 10; k++ {
				fmt.Printf("%d%d%d", i, j, k)
				if i != 7 {
					fmt.Printf(", ")
				}
			}
		}
	}
}