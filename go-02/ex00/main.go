/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/10 16:03:54 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"piscine"
)
func main() {
arg := 4
fmt.Println(piscine.IterativeFactorial(arg))
}