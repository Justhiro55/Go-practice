/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/28 10:53:27 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
    "fmt"
    // "os"
)

type point struct {
    x int;
    y int;
}

func setPoint(ptr *point) {
    ptr.x = 42
    ptr.y = 21
}

func main() {
points := &point{}
setPoint(points)
fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}