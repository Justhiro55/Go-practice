/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/13 11:56:45 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"ft"
"piscine"
)
func main() {
ft.PrintRune(piscine.FirstRune("Hello!"))
ft.PrintRune(piscine.FirstRune("Salut!"))
ft.PrintRune(piscine.FirstRune("Ola!"))
ft.PrintRune('\n')
}