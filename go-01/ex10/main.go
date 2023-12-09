/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 15:03:43 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"piscine"
)
func main() {
fmt.Println(piscine.Atoi("12345"))
fmt.Println(piscine.Atoi("0000000012345"))
fmt.Println(piscine.Atoi("012 345"))
fmt.Println(piscine.Atoi("Hello World!"))
fmt.Println(piscine.Atoi("+1234"))
fmt.Println(piscine.Atoi("-1234"))
fmt.Println(piscine.Atoi("++1234"))
fmt.Println(piscine.Atoi("--1234"))
}