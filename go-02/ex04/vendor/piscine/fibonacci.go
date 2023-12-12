/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fibonacci.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/12/10 15:57:43 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/12 19:36:34 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func Fibonacci(index int) int {
	if index == 0{
		return 0
	} else if index == 1 || index == 2 {
		return 1
	} 
	if index < 0 {
		return -1
	} else {
		return Fibonacci(index - 1) + Fibonacci(index - 2)
	}
}
