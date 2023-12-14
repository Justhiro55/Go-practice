/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hhagiwar <hhagiwar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/10/03 12:48:04 by hhagiwar          #+#    #+#             */
/*   Updated: 2023/12/14 16:09:21 by hhagiwar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main
import (
"fmt"
"os"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	programName := os.Args[0]

	lastSlashIndex := -1
	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			lastSlashIndex = i
			break
		}
	}

	if lastSlashIndex != -1 {
		programName = programName[lastSlashIndex+1:]
	}

	fmt.Print(programName)
}