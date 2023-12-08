/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   PrintComb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/02 17:12:45 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/10/02 19:25:15 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "fmt"

func PrintComb2(){
	for i := 0; i < 99; i++ {
		for j := i+1; j < 100; j++{
			fmt.Printf("%d%d %d%d", 0 + i/10, i%10, j/10, j%10)
			if i != 98{
				fmt.Printf(", ")
			}
		}
	}
}