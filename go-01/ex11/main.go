/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42Tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/09 15:17:49 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
    "fmt"
    "piscine"
)

func main() {
    s := []int{5, 4, 3, 2, 1, 0}
    piscine.SortIntegerTable(s)
    fmt.Println(s)
}