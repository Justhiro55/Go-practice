/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 11:58:44 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"ft"
"piscine"
)
func main() {
ft.PrintRune(piscine.NRune("Hello!", 3))
ft.PrintRune(piscine.NRune("Salut!", 2))
ft.PrintRune(piscine.NRune("Bye!", -1))
ft.PrintRune(piscine.NRune("Bye!", 5))
ft.PrintRune(piscine.NRune("Ola!", 4))
ft.PrintRune('\n')
}