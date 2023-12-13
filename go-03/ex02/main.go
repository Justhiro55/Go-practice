/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 12:02:54 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"ft"
"piscine"
)
func main() {
ft.PrintRune(piscine.LastRune("Hello!"))
ft.PrintRune(piscine.LastRune("Salut!"))
ft.PrintRune(piscine.LastRune("Ola!"))
ft.PrintRune('\n')
}