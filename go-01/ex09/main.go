/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 15:01:16 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"piscine"
)
func main() {
fmt.Println(piscine.BasicAtoi2("12345"))
fmt.Println(piscine.BasicAtoi2("0000000012345"))
fmt.Println(piscine.BasicAtoi2("012 345"))
fmt.Println(piscine.BasicAtoi2("Hello World!"))
}