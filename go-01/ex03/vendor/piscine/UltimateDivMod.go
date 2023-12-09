/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   UltimateDivMod.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 13:57:45 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/10/03 14:03:07 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	// "fmt"
)

func UltimateDivMod(a *int, b *int) {
	var  tmp_a int = *a
	var  tmp_b int = *b
	
	*a = tmp_a / tmp_b
	*b = tmp_a % tmp_b
}
