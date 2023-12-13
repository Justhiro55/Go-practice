/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 15:44:21 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"piscine"
)
func main() {
fmt.Println(piscine.AtoiBase("125", "0123456789"))
fmt.Println(piscine.AtoiBase("1111101", "01"))
fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
fmt.Println(piscine.AtoiBase("uoi", "choumi"))
fmt.Println(piscine.AtoiBase("-bbbbbab", "-ab"))
}