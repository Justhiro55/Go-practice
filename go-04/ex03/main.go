/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/14 18:44:18 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"os"
)

func sortParams(array []string) {
	// arrayLen := len(array) - 1
	var tmp string = ""
	// fmt.Println(arrayLen)
	for i := 1; i <len(array) - 1; i++ {
		for j := i + 1; j < len(array) - 1; j++ {
			if array[i] > array[j] {
				tmp = array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
		fmt.Println(array[i])
	}
}

func main() {
	sortParams(os.Args)
}