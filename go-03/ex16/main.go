/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 14:15:54 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"ft"
"piscine"
)
func main() {
piscine.PrintNbrBase(125, "0123456789")
ft.PrintRune('\n')
piscine.PrintNbrBase(-125, "01")
ft.PrintRune('\n')
piscine.PrintNbrBase(125, "0123456789ABCDEF")
ft.PrintRune('\n')
piscine.PrintNbrBase(-125, "choumi")
ft.PrintRune('\n')
piscine.PrintNbrBase(125, "aa")
ft.PrintRune('\n')
}