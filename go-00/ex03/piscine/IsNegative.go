/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   IsNegative.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/02 13:07:16 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/10/02 19:56:53 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "fmt"

func IsNegative(nb int){
	if nb >= 0{
		fmt.Printf("F\n")
	}else {
		fmt.Printf("T\n")
	}	
}
