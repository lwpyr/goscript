// Code generated from goscript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // goscript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 62, 640,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 6, 2, 58, 10, 2, 13, 2, 14, 2, 59, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 69, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 77,
	10, 4, 12, 4, 14, 4, 80, 11, 4, 3, 4, 3, 4, 5, 4, 84, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 93, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 102, 10, 4, 12, 4, 14, 4, 105, 11, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 7, 4, 112, 10, 4, 12, 4, 14, 4, 115, 11, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 128,
	10, 4, 12, 4, 14, 4, 131, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 142, 10, 4, 12, 4, 14, 4, 145, 11, 4, 3, 4, 3, 4,
	5, 4, 149, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 158,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 167, 10, 4, 12,
	4, 14, 4, 170, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 177, 10, 4, 12,
	4, 14, 4, 180, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 193, 10, 4, 12, 4, 14, 4, 196, 11, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 201, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	222, 10, 9, 12, 9, 14, 9, 225, 11, 9, 3, 9, 5, 9, 228, 10, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 243, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 252, 10, 11, 12, 11, 14, 11, 255, 11, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 263, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 273, 10, 13, 5, 13, 275, 10, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 6, 13, 287,
	10, 13, 13, 13, 14, 13, 288, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 331, 10, 13, 12, 13, 14, 13, 334, 11, 13, 3, 13,
	3, 13, 5, 13, 338, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 7, 15, 345,
	10, 15, 12, 15, 14, 15, 348, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 7, 16, 356, 10, 16, 12, 16, 14, 16, 359, 11, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 7, 16, 367, 10, 16, 12, 16, 14, 16, 370, 11, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 376, 10, 16, 12, 16, 14, 16, 379, 11,
	16, 5, 16, 381, 10, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 388,
	10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 408,
	10, 18, 12, 18, 14, 18, 411, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 7, 18, 421, 10, 18, 12, 18, 14, 18, 424, 11, 18, 3,
	18, 3, 18, 3, 18, 7, 18, 429, 10, 18, 12, 18, 14, 18, 432, 11, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 452, 10, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 479, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 7, 21, 499, 10, 21, 12, 21, 14, 21, 502, 11, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 510, 10, 23, 12, 23, 14,
	23, 513, 11, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 528, 10, 23, 12, 23, 14, 23,
	531, 11, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 7, 23, 544, 10, 23, 12, 23, 14, 23, 547, 11, 23, 3, 23,
	3, 23, 3, 23, 5, 23, 552, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 559, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 566, 10, 25,
	12, 25, 14, 25, 569, 11, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 7, 25, 589, 10, 25, 12, 25, 14, 25, 592, 11, 25, 3, 25, 3, 25,
	5, 25, 596, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 7, 26, 608, 10, 26, 12, 26, 14, 26, 611, 11, 26, 3, 26,
	3, 26, 5, 26, 615, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 628, 10, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 638, 10, 28, 3, 28, 2, 4, 34,
	40, 29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 2, 8, 3, 2, 40, 41, 4, 2, 39, 39,
	51, 51, 3, 2, 52, 56, 3, 2, 35, 37, 3, 2, 38, 39, 3, 2, 42, 48, 2, 708,
	2, 57, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 200, 3, 2, 2, 2, 8, 202, 3, 2,
	2, 2, 10, 204, 3, 2, 2, 2, 12, 206, 3, 2, 2, 2, 14, 208, 3, 2, 2, 2, 16,
	227, 3, 2, 2, 2, 18, 242, 3, 2, 2, 2, 20, 244, 3, 2, 2, 2, 22, 262, 3,
	2, 2, 2, 24, 337, 3, 2, 2, 2, 26, 339, 3, 2, 2, 2, 28, 341, 3, 2, 2, 2,
	30, 380, 3, 2, 2, 2, 32, 382, 3, 2, 2, 2, 34, 387, 3, 2, 2, 2, 36, 433,
	3, 2, 2, 2, 38, 451, 3, 2, 2, 2, 40, 478, 3, 2, 2, 2, 42, 503, 3, 2, 2,
	2, 44, 551, 3, 2, 2, 2, 46, 558, 3, 2, 2, 2, 48, 595, 3, 2, 2, 2, 50, 614,
	3, 2, 2, 2, 52, 627, 3, 2, 2, 2, 54, 637, 3, 2, 2, 2, 56, 58, 5, 4, 3,
	2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60,
	3, 2, 2, 2, 60, 3, 3, 2, 2, 2, 61, 69, 5, 6, 4, 2, 62, 63, 5, 16, 9, 2,
	63, 64, 7, 3, 2, 2, 64, 69, 3, 2, 2, 2, 65, 69, 5, 52, 27, 2, 66, 69, 5,
	20, 11, 2, 67, 69, 5, 22, 12, 2, 68, 61, 3, 2, 2, 2, 68, 62, 3, 2, 2, 2,
	68, 65, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 5, 3, 2,
	2, 2, 70, 71, 7, 30, 2, 2, 71, 72, 7, 60, 2, 2, 72, 73, 7, 4, 2, 2, 73,
	78, 5, 8, 5, 2, 74, 75, 7, 5, 2, 2, 75, 77, 5, 8, 5, 2, 76, 74, 3, 2, 2,
	2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81,
	3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 83, 7, 6, 2, 2, 82, 84, 5, 12, 7, 2,
	83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 5,
	28, 15, 2, 86, 201, 3, 2, 2, 2, 87, 88, 7, 30, 2, 2, 88, 89, 7, 60, 2,
	2, 89, 90, 7, 4, 2, 2, 90, 92, 7, 6, 2, 2, 91, 93, 5, 12, 7, 2, 92, 91,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 201, 5, 28, 15,
	2, 95, 96, 7, 30, 2, 2, 96, 97, 7, 60, 2, 2, 97, 98, 7, 4, 2, 2, 98, 103,
	5, 8, 5, 2, 99, 100, 7, 5, 2, 2, 100, 102, 5, 8, 5, 2, 101, 99, 3, 2, 2,
	2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	106, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 107, 7, 6, 2, 2, 107, 108,
	7, 4, 2, 2, 108, 113, 5, 12, 7, 2, 109, 110, 7, 5, 2, 2, 110, 112, 5, 12,
	7, 2, 111, 109, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2,
	113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116,
	117, 7, 6, 2, 2, 117, 118, 5, 28, 15, 2, 118, 201, 3, 2, 2, 2, 119, 120,
	7, 30, 2, 2, 120, 121, 7, 60, 2, 2, 121, 122, 7, 4, 2, 2, 122, 123, 7,
	6, 2, 2, 123, 124, 7, 4, 2, 2, 124, 129, 5, 12, 7, 2, 125, 126, 7, 5, 2,
	2, 126, 128, 5, 12, 7, 2, 127, 125, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 132, 133, 7, 6, 2, 2, 133, 134, 5, 28, 15, 2, 134, 201, 3,
	2, 2, 2, 135, 136, 7, 30, 2, 2, 136, 137, 7, 60, 2, 2, 137, 138, 7, 4,
	2, 2, 138, 143, 5, 8, 5, 2, 139, 140, 7, 5, 2, 2, 140, 142, 5, 8, 5, 2,
	141, 139, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143,
	144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 148,
	7, 6, 2, 2, 147, 149, 5, 10, 6, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 5, 28, 15, 2, 151, 201, 3, 2, 2,
	2, 152, 153, 7, 30, 2, 2, 153, 154, 7, 60, 2, 2, 154, 155, 7, 4, 2, 2,
	155, 157, 7, 6, 2, 2, 156, 158, 5, 18, 10, 2, 157, 156, 3, 2, 2, 2, 157,
	158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 201, 5, 28, 15, 2, 160, 161,
	7, 30, 2, 2, 161, 162, 7, 60, 2, 2, 162, 163, 7, 4, 2, 2, 163, 168, 5,
	8, 5, 2, 164, 165, 7, 5, 2, 2, 165, 167, 5, 8, 5, 2, 166, 164, 3, 2, 2,
	2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169,
	171, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171, 172, 7, 6, 2, 2, 172, 173,
	7, 4, 2, 2, 173, 178, 5, 10, 6, 2, 174, 175, 7, 5, 2, 2, 175, 177, 5, 10,
	6, 2, 176, 174, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181,
	182, 7, 6, 2, 2, 182, 183, 5, 28, 15, 2, 183, 201, 3, 2, 2, 2, 184, 185,
	7, 30, 2, 2, 185, 186, 7, 60, 2, 2, 186, 187, 7, 4, 2, 2, 187, 188, 7,
	6, 2, 2, 188, 189, 7, 4, 2, 2, 189, 194, 5, 10, 6, 2, 190, 191, 7, 5, 2,
	2, 191, 193, 5, 10, 6, 2, 192, 190, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2, 194,
	192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196, 194,
	3, 2, 2, 2, 197, 198, 7, 6, 2, 2, 198, 199, 5, 28, 15, 2, 199, 201, 3,
	2, 2, 2, 200, 70, 3, 2, 2, 2, 200, 87, 3, 2, 2, 2, 200, 95, 3, 2, 2, 2,
	200, 119, 3, 2, 2, 2, 200, 135, 3, 2, 2, 2, 200, 152, 3, 2, 2, 2, 200,
	160, 3, 2, 2, 2, 200, 184, 3, 2, 2, 2, 201, 7, 3, 2, 2, 2, 202, 203, 5,
	14, 8, 2, 203, 9, 3, 2, 2, 2, 204, 205, 5, 14, 8, 2, 205, 11, 3, 2, 2,
	2, 206, 207, 5, 18, 10, 2, 207, 13, 3, 2, 2, 2, 208, 209, 7, 60, 2, 2,
	209, 210, 5, 18, 10, 2, 210, 15, 3, 2, 2, 2, 211, 212, 7, 31, 2, 2, 212,
	213, 7, 60, 2, 2, 213, 228, 5, 18, 10, 2, 214, 215, 7, 31, 2, 2, 215, 216,
	7, 60, 2, 2, 216, 223, 7, 7, 2, 2, 217, 218, 7, 60, 2, 2, 218, 219, 5,
	18, 10, 2, 219, 220, 7, 5, 2, 2, 220, 222, 3, 2, 2, 2, 221, 217, 3, 2,
	2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2,
	224, 226, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 228, 7, 8, 2, 2, 227,
	211, 3, 2, 2, 2, 227, 214, 3, 2, 2, 2, 228, 17, 3, 2, 2, 2, 229, 243, 7,
	60, 2, 2, 230, 231, 7, 9, 2, 2, 231, 232, 7, 47, 2, 2, 232, 233, 7, 60,
	2, 2, 233, 234, 7, 5, 2, 2, 234, 235, 5, 18, 10, 2, 235, 236, 7, 44, 2,
	2, 236, 243, 3, 2, 2, 2, 237, 238, 7, 10, 2, 2, 238, 239, 7, 47, 2, 2,
	239, 240, 5, 18, 10, 2, 240, 241, 7, 44, 2, 2, 241, 243, 3, 2, 2, 2, 242,
	229, 3, 2, 2, 2, 242, 230, 3, 2, 2, 2, 242, 237, 3, 2, 2, 2, 243, 19, 3,
	2, 2, 2, 244, 245, 7, 31, 2, 2, 245, 246, 7, 11, 2, 2, 246, 247, 7, 60,
	2, 2, 247, 253, 7, 7, 2, 2, 248, 249, 7, 60, 2, 2, 249, 250, 7, 57, 2,
	2, 250, 252, 7, 5, 2, 2, 251, 248, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253,
	251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 256, 3, 2, 2, 2, 255, 253,
	3, 2, 2, 2, 256, 257, 7, 8, 2, 2, 257, 21, 3, 2, 2, 2, 258, 263, 5, 24,
	13, 2, 259, 260, 5, 30, 16, 2, 260, 261, 7, 3, 2, 2, 261, 263, 3, 2, 2,
	2, 262, 258, 3, 2, 2, 2, 262, 259, 3, 2, 2, 2, 263, 23, 3, 2, 2, 2, 264,
	265, 7, 23, 2, 2, 265, 266, 7, 4, 2, 2, 266, 267, 5, 40, 21, 2, 267, 268,
	7, 6, 2, 2, 268, 274, 5, 28, 15, 2, 269, 272, 7, 24, 2, 2, 270, 273, 5,
	28, 15, 2, 271, 273, 5, 24, 13, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 2,
	2, 2, 273, 275, 3, 2, 2, 2, 274, 269, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2,
	275, 338, 3, 2, 2, 2, 276, 277, 7, 25, 2, 2, 277, 278, 7, 4, 2, 2, 278,
	279, 5, 40, 21, 2, 279, 280, 7, 6, 2, 2, 280, 286, 7, 7, 2, 2, 281, 282,
	7, 26, 2, 2, 282, 283, 5, 46, 24, 2, 283, 284, 7, 12, 2, 2, 284, 285, 5,
	28, 15, 2, 285, 287, 3, 2, 2, 2, 286, 281, 3, 2, 2, 2, 287, 288, 3, 2,
	2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2,
	290, 291, 7, 8, 2, 2, 291, 338, 3, 2, 2, 2, 292, 293, 7, 20, 2, 2, 293,
	294, 7, 4, 2, 2, 294, 295, 7, 60, 2, 2, 295, 296, 7, 13, 2, 2, 296, 297,
	5, 26, 14, 2, 297, 298, 7, 6, 2, 2, 298, 299, 5, 28, 15, 2, 299, 338, 3,
	2, 2, 2, 300, 301, 7, 20, 2, 2, 301, 302, 7, 4, 2, 2, 302, 303, 7, 60,
	2, 2, 303, 304, 7, 5, 2, 2, 304, 305, 7, 60, 2, 2, 305, 306, 7, 13, 2,
	2, 306, 307, 5, 26, 14, 2, 307, 308, 7, 6, 2, 2, 308, 309, 5, 28, 15, 2,
	309, 338, 3, 2, 2, 2, 310, 311, 7, 20, 2, 2, 311, 312, 7, 4, 2, 2, 312,
	313, 5, 40, 21, 2, 313, 314, 7, 3, 2, 2, 314, 315, 5, 40, 21, 2, 315, 316,
	7, 3, 2, 2, 316, 317, 5, 40, 21, 2, 317, 318, 7, 6, 2, 2, 318, 319, 5,
	28, 15, 2, 319, 338, 3, 2, 2, 2, 320, 321, 7, 21, 2, 2, 321, 338, 7, 3,
	2, 2, 322, 323, 7, 22, 2, 2, 323, 338, 7, 3, 2, 2, 324, 325, 7, 27, 2,
	2, 325, 338, 7, 3, 2, 2, 326, 327, 7, 27, 2, 2, 327, 332, 5, 40, 21, 2,
	328, 329, 7, 5, 2, 2, 329, 331, 5, 40, 21, 2, 330, 328, 3, 2, 2, 2, 331,
	334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335,
	3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 335, 336, 7, 3, 2, 2, 336, 338, 3, 2,
	2, 2, 337, 264, 3, 2, 2, 2, 337, 276, 3, 2, 2, 2, 337, 292, 3, 2, 2, 2,
	337, 300, 3, 2, 2, 2, 337, 310, 3, 2, 2, 2, 337, 320, 3, 2, 2, 2, 337,
	322, 3, 2, 2, 2, 337, 324, 3, 2, 2, 2, 337, 326, 3, 2, 2, 2, 338, 25, 3,
	2, 2, 2, 339, 340, 5, 40, 21, 2, 340, 27, 3, 2, 2, 2, 341, 346, 7, 7, 2,
	2, 342, 345, 5, 22, 12, 2, 343, 345, 5, 52, 27, 2, 344, 342, 3, 2, 2, 2,
	344, 343, 3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 349, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 350,
	7, 8, 2, 2, 350, 29, 3, 2, 2, 2, 351, 381, 5, 40, 21, 2, 352, 357, 5, 32,
	17, 2, 353, 354, 7, 5, 2, 2, 354, 356, 5, 32, 17, 2, 355, 353, 3, 2, 2,
	2, 356, 359, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358,
	360, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 360, 361, 7, 52, 2, 2, 361, 362,
	5, 48, 25, 2, 362, 381, 3, 2, 2, 2, 363, 368, 5, 32, 17, 2, 364, 365, 7,
	5, 2, 2, 365, 367, 5, 32, 17, 2, 366, 364, 3, 2, 2, 2, 367, 370, 3, 2,
	2, 2, 368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 371, 3, 2, 2, 2,
	370, 368, 3, 2, 2, 2, 371, 372, 7, 52, 2, 2, 372, 377, 5, 40, 21, 2, 373,
	374, 7, 5, 2, 2, 374, 376, 5, 40, 21, 2, 375, 373, 3, 2, 2, 2, 376, 379,
	3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 381, 3, 2,
	2, 2, 379, 377, 3, 2, 2, 2, 380, 351, 3, 2, 2, 2, 380, 352, 3, 2, 2, 2,
	380, 363, 3, 2, 2, 2, 381, 31, 3, 2, 2, 2, 382, 383, 5, 34, 18, 2, 383,
	33, 3, 2, 2, 2, 384, 385, 8, 18, 1, 2, 385, 388, 7, 60, 2, 2, 386, 388,
	7, 18, 2, 2, 387, 384, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2, 388, 430, 3, 2,
	2, 2, 389, 390, 12, 9, 2, 2, 390, 391, 7, 61, 2, 2, 391, 429, 7, 60, 2,
	2, 392, 393, 12, 8, 2, 2, 393, 394, 7, 14, 2, 2, 394, 395, 5, 36, 19, 2,
	395, 396, 7, 15, 2, 2, 396, 429, 3, 2, 2, 2, 397, 398, 12, 7, 2, 2, 398,
	399, 7, 16, 2, 2, 399, 400, 5, 40, 21, 2, 400, 401, 7, 17, 2, 2, 401, 429,
	3, 2, 2, 2, 402, 403, 12, 6, 2, 2, 403, 404, 7, 16, 2, 2, 404, 409, 5,
	38, 20, 2, 405, 406, 7, 5, 2, 2, 406, 408, 5, 38, 20, 2, 407, 405, 3, 2,
	2, 2, 408, 411, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2,
	410, 412, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 412, 413, 7, 17, 2, 2, 413,
	429, 3, 2, 2, 2, 414, 415, 12, 5, 2, 2, 415, 416, 7, 16, 2, 2, 416, 417,
	7, 16, 2, 2, 417, 422, 5, 40, 21, 2, 418, 419, 7, 5, 2, 2, 419, 421, 5,
	40, 21, 2, 420, 418, 3, 2, 2, 2, 421, 424, 3, 2, 2, 2, 422, 420, 3, 2,
	2, 2, 422, 423, 3, 2, 2, 2, 423, 425, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2,
	425, 426, 7, 17, 2, 2, 426, 427, 7, 17, 2, 2, 427, 429, 3, 2, 2, 2, 428,
	389, 3, 2, 2, 2, 428, 392, 3, 2, 2, 2, 428, 397, 3, 2, 2, 2, 428, 402,
	3, 2, 2, 2, 428, 414, 3, 2, 2, 2, 429, 432, 3, 2, 2, 2, 430, 428, 3, 2,
	2, 2, 430, 431, 3, 2, 2, 2, 431, 35, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2,
	433, 434, 5, 40, 21, 2, 434, 37, 3, 2, 2, 2, 435, 436, 5, 40, 21, 2, 436,
	437, 7, 12, 2, 2, 437, 438, 5, 40, 21, 2, 438, 439, 7, 12, 2, 2, 439, 440,
	5, 40, 21, 2, 440, 452, 3, 2, 2, 2, 441, 442, 5, 40, 21, 2, 442, 443, 7,
	12, 2, 2, 443, 444, 5, 40, 21, 2, 444, 452, 3, 2, 2, 2, 445, 446, 5, 40,
	21, 2, 446, 447, 7, 12, 2, 2, 447, 452, 3, 2, 2, 2, 448, 452, 5, 40, 21,
	2, 449, 450, 7, 12, 2, 2, 450, 452, 5, 40, 21, 2, 451, 435, 3, 2, 2, 2,
	451, 441, 3, 2, 2, 2, 451, 445, 3, 2, 2, 2, 451, 448, 3, 2, 2, 2, 451,
	449, 3, 2, 2, 2, 452, 39, 3, 2, 2, 2, 453, 454, 8, 21, 1, 2, 454, 455,
	7, 4, 2, 2, 455, 456, 5, 40, 21, 2, 456, 457, 7, 6, 2, 2, 457, 479, 3,
	2, 2, 2, 458, 459, 9, 2, 2, 2, 459, 479, 5, 34, 18, 2, 460, 461, 9, 3,
	2, 2, 461, 479, 5, 40, 21, 17, 462, 463, 5, 34, 18, 2, 463, 464, 9, 2,
	2, 2, 464, 479, 3, 2, 2, 2, 465, 466, 5, 32, 17, 2, 466, 467, 9, 4, 2,
	2, 467, 468, 5, 40, 21, 9, 468, 479, 3, 2, 2, 2, 469, 470, 5, 32, 17, 2,
	470, 471, 7, 52, 2, 2, 471, 472, 5, 42, 22, 2, 472, 479, 3, 2, 2, 2, 473,
	479, 5, 46, 24, 2, 474, 479, 5, 34, 18, 2, 475, 479, 5, 48, 25, 2, 476,
	479, 5, 50, 26, 2, 477, 479, 5, 54, 28, 2, 478, 453, 3, 2, 2, 2, 478, 458,
	3, 2, 2, 2, 478, 460, 3, 2, 2, 2, 478, 462, 3, 2, 2, 2, 478, 465, 3, 2,
	2, 2, 478, 469, 3, 2, 2, 2, 478, 473, 3, 2, 2, 2, 478, 474, 3, 2, 2, 2,
	478, 475, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 478, 477, 3, 2, 2, 2, 479,
	500, 3, 2, 2, 2, 480, 481, 12, 15, 2, 2, 481, 482, 7, 34, 2, 2, 482, 499,
	5, 40, 21, 15, 483, 484, 12, 14, 2, 2, 484, 485, 9, 5, 2, 2, 485, 499,
	5, 40, 21, 15, 486, 487, 12, 13, 2, 2, 487, 488, 9, 6, 2, 2, 488, 499,
	5, 40, 21, 14, 489, 490, 12, 12, 2, 2, 490, 491, 9, 7, 2, 2, 491, 499,
	5, 40, 21, 13, 492, 493, 12, 11, 2, 2, 493, 494, 7, 49, 2, 2, 494, 499,
	5, 40, 21, 12, 495, 496, 12, 10, 2, 2, 496, 497, 7, 50, 2, 2, 497, 499,
	5, 40, 21, 11, 498, 480, 3, 2, 2, 2, 498, 483, 3, 2, 2, 2, 498, 486, 3,
	2, 2, 2, 498, 489, 3, 2, 2, 2, 498, 492, 3, 2, 2, 2, 498, 495, 3, 2, 2,
	2, 499, 502, 3, 2, 2, 2, 500, 498, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501,
	41, 3, 2, 2, 2, 502, 500, 3, 2, 2, 2, 503, 504, 5, 44, 23, 2, 504, 43,
	3, 2, 2, 2, 505, 506, 7, 16, 2, 2, 506, 511, 5, 44, 23, 2, 507, 508, 7,
	5, 2, 2, 508, 510, 5, 44, 23, 2, 509, 507, 3, 2, 2, 2, 510, 513, 3, 2,
	2, 2, 511, 509, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 514, 3, 2, 2, 2,
	513, 511, 3, 2, 2, 2, 514, 515, 7, 17, 2, 2, 515, 552, 3, 2, 2, 2, 516,
	517, 7, 7, 2, 2, 517, 518, 7, 60, 2, 2, 518, 519, 7, 4, 2, 2, 519, 520,
	5, 44, 23, 2, 520, 529, 7, 6, 2, 2, 521, 522, 7, 5, 2, 2, 522, 523, 7,
	60, 2, 2, 523, 524, 7, 4, 2, 2, 524, 525, 5, 44, 23, 2, 525, 526, 7, 6,
	2, 2, 526, 528, 3, 2, 2, 2, 527, 521, 3, 2, 2, 2, 528, 531, 3, 2, 2, 2,
	529, 527, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 532, 3, 2, 2, 2, 531,
	529, 3, 2, 2, 2, 532, 533, 7, 8, 2, 2, 533, 552, 3, 2, 2, 2, 534, 535,
	7, 7, 2, 2, 535, 536, 5, 44, 23, 2, 536, 537, 7, 12, 2, 2, 537, 545, 5,
	44, 23, 2, 538, 539, 7, 5, 2, 2, 539, 540, 5, 44, 23, 2, 540, 541, 7, 12,
	2, 2, 541, 542, 5, 44, 23, 2, 542, 544, 3, 2, 2, 2, 543, 538, 3, 2, 2,
	2, 544, 547, 3, 2, 2, 2, 545, 543, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546,
	548, 3, 2, 2, 2, 547, 545, 3, 2, 2, 2, 548, 549, 7, 8, 2, 2, 549, 552,
	3, 2, 2, 2, 550, 552, 5, 40, 21, 2, 551, 505, 3, 2, 2, 2, 551, 516, 3,
	2, 2, 2, 551, 534, 3, 2, 2, 2, 551, 550, 3, 2, 2, 2, 552, 45, 3, 2, 2,
	2, 553, 559, 7, 57, 2, 2, 554, 559, 7, 58, 2, 2, 555, 559, 7, 32, 2, 2,
	556, 559, 7, 33, 2, 2, 557, 559, 7, 59, 2, 2, 558, 553, 3, 2, 2, 2, 558,
	554, 3, 2, 2, 2, 558, 555, 3, 2, 2, 2, 558, 556, 3, 2, 2, 2, 558, 557,
	3, 2, 2, 2, 559, 47, 3, 2, 2, 2, 560, 561, 7, 60, 2, 2, 561, 562, 7, 4,
	2, 2, 562, 567, 5, 40, 21, 2, 563, 564, 7, 5, 2, 2, 564, 566, 5, 40, 21,
	2, 565, 563, 3, 2, 2, 2, 566, 569, 3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 567,
	568, 3, 2, 2, 2, 568, 570, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 570, 571,
	7, 6, 2, 2, 571, 596, 3, 2, 2, 2, 572, 573, 7, 60, 2, 2, 573, 574, 7, 4,
	2, 2, 574, 596, 7, 6, 2, 2, 575, 576, 5, 34, 18, 2, 576, 577, 7, 61, 2,
	2, 577, 578, 7, 60, 2, 2, 578, 579, 7, 4, 2, 2, 579, 580, 7, 6, 2, 2, 580,
	596, 3, 2, 2, 2, 581, 582, 5, 34, 18, 2, 582, 583, 7, 61, 2, 2, 583, 584,
	7, 60, 2, 2, 584, 585, 7, 4, 2, 2, 585, 590, 5, 40, 21, 2, 586, 587, 7,
	5, 2, 2, 587, 589, 5, 40, 21, 2, 588, 586, 3, 2, 2, 2, 589, 592, 3, 2,
	2, 2, 590, 588, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 593, 3, 2, 2, 2,
	592, 590, 3, 2, 2, 2, 593, 594, 7, 6, 2, 2, 594, 596, 3, 2, 2, 2, 595,
	560, 3, 2, 2, 2, 595, 572, 3, 2, 2, 2, 595, 575, 3, 2, 2, 2, 595, 581,
	3, 2, 2, 2, 596, 49, 3, 2, 2, 2, 597, 598, 7, 19, 2, 2, 598, 599, 7, 60,
	2, 2, 599, 600, 7, 4, 2, 2, 600, 615, 7, 6, 2, 2, 601, 602, 7, 19, 2, 2,
	602, 603, 7, 60, 2, 2, 603, 604, 7, 4, 2, 2, 604, 609, 5, 40, 21, 2, 605,
	606, 7, 5, 2, 2, 606, 608, 5, 40, 21, 2, 607, 605, 3, 2, 2, 2, 608, 611,
	3, 2, 2, 2, 609, 607, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 612, 3, 2,
	2, 2, 611, 609, 3, 2, 2, 2, 612, 613, 7, 6, 2, 2, 613, 615, 3, 2, 2, 2,
	614, 597, 3, 2, 2, 2, 614, 601, 3, 2, 2, 2, 615, 51, 3, 2, 2, 2, 616, 617,
	7, 28, 2, 2, 617, 618, 7, 60, 2, 2, 618, 619, 7, 60, 2, 2, 619, 628, 7,
	3, 2, 2, 620, 621, 7, 28, 2, 2, 621, 622, 7, 60, 2, 2, 622, 623, 7, 60,
	2, 2, 623, 624, 7, 52, 2, 2, 624, 625, 5, 40, 21, 2, 625, 626, 7, 3, 2,
	2, 626, 628, 3, 2, 2, 2, 627, 616, 3, 2, 2, 2, 627, 620, 3, 2, 2, 2, 628,
	53, 3, 2, 2, 2, 629, 630, 7, 29, 2, 2, 630, 631, 7, 60, 2, 2, 631, 638,
	7, 60, 2, 2, 632, 633, 7, 29, 2, 2, 633, 634, 7, 60, 2, 2, 634, 635, 7,
	60, 2, 2, 635, 636, 7, 52, 2, 2, 636, 638, 5, 40, 21, 2, 637, 629, 3, 2,
	2, 2, 637, 632, 3, 2, 2, 2, 638, 55, 3, 2, 2, 2, 54, 59, 68, 78, 83, 92,
	103, 113, 129, 143, 148, 157, 168, 178, 194, 200, 223, 227, 242, 253, 262,
	272, 274, 288, 332, 337, 344, 346, 357, 368, 377, 380, 387, 409, 422, 428,
	430, 451, 478, 498, 500, 511, 529, 545, 551, 558, 567, 590, 595, 609, 614,
	627, 637,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "','", "')'", "'{'", "'}'", "'map'", "'slice'", "'enum'",
	"':'", "'in'", "'[?('", "')]'", "'['", "']'", "'@'", "'new'", "'for'",
	"'break'", "'continue'", "'if'", "'else'", "'switch'", "'case'", "'return'",
	"'var'", "'local'", "'func'", "'type'", "", "'nil'", "'**'", "'*'", "'/'",
	"'%'", "'+'", "'-'", "'++'", "'--'", "'=='", "'!='", "'>'", "'>='", "'<='",
	"'<'", "'=~'", "'&&'", "'||'", "'!'", "'='", "'+='", "'-='", "'*='", "'/='",
	"", "", "", "", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NEW",
	"FOR", "BREAK", "CONTINUE", "IF", "ELSE", "SWITCH", "CASE", "RETURN", "VAR",
	"LOCAL", "FUNCTION", "TYPEDEF", "BOOL", "NULL", "POW", "MUL", "DIV", "MOD",
	"ADD", "SUB", "UNARYADD", "UNARYSUB", "EQ", "INEQ", "GT", "GE", "LE", "LT",
	"REGEX", "AND", "OR", "NOT", "ASSIGN", "ADDEQUAL", "SUBEQUAL", "MULEQUAL",
	"DIVEQUAL", "INT", "FLOAT", "STRING", "NAME", "DOT", "WHITESPACE",
}

var ruleNames = []string{
	"program", "statement", "functiondef", "inparam", "outparam", "returntypename",
	"param", "typedef", "typename", "enumdef", "execution", "control", "collection",
	"block", "line", "lhs", "variable", "filter", "indexs", "expr", "initializationListBegin",
	"initializationList", "constant", "functions", "constructor", "vardef",
	"localdef",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type goscriptParser struct {
	*antlr.BaseParser
}

func NewgoscriptParser(input antlr.TokenStream) *goscriptParser {
	this := new(goscriptParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "goscript.g4"

	return this
}

// goscriptParser tokens.
const (
	goscriptParserEOF        = antlr.TokenEOF
	goscriptParserT__0       = 1
	goscriptParserT__1       = 2
	goscriptParserT__2       = 3
	goscriptParserT__3       = 4
	goscriptParserT__4       = 5
	goscriptParserT__5       = 6
	goscriptParserT__6       = 7
	goscriptParserT__7       = 8
	goscriptParserT__8       = 9
	goscriptParserT__9       = 10
	goscriptParserT__10      = 11
	goscriptParserT__11      = 12
	goscriptParserT__12      = 13
	goscriptParserT__13      = 14
	goscriptParserT__14      = 15
	goscriptParserT__15      = 16
	goscriptParserNEW        = 17
	goscriptParserFOR        = 18
	goscriptParserBREAK      = 19
	goscriptParserCONTINUE   = 20
	goscriptParserIF         = 21
	goscriptParserELSE       = 22
	goscriptParserSWITCH     = 23
	goscriptParserCASE       = 24
	goscriptParserRETURN     = 25
	goscriptParserVAR        = 26
	goscriptParserLOCAL      = 27
	goscriptParserFUNCTION   = 28
	goscriptParserTYPEDEF    = 29
	goscriptParserBOOL       = 30
	goscriptParserNULL       = 31
	goscriptParserPOW        = 32
	goscriptParserMUL        = 33
	goscriptParserDIV        = 34
	goscriptParserMOD        = 35
	goscriptParserADD        = 36
	goscriptParserSUB        = 37
	goscriptParserUNARYADD   = 38
	goscriptParserUNARYSUB   = 39
	goscriptParserEQ         = 40
	goscriptParserINEQ       = 41
	goscriptParserGT         = 42
	goscriptParserGE         = 43
	goscriptParserLE         = 44
	goscriptParserLT         = 45
	goscriptParserREGEX      = 46
	goscriptParserAND        = 47
	goscriptParserOR         = 48
	goscriptParserNOT        = 49
	goscriptParserASSIGN     = 50
	goscriptParserADDEQUAL   = 51
	goscriptParserSUBEQUAL   = 52
	goscriptParserMULEQUAL   = 53
	goscriptParserDIVEQUAL   = 54
	goscriptParserINT        = 55
	goscriptParserFLOAT      = 56
	goscriptParserSTRING     = 57
	goscriptParserNAME       = 58
	goscriptParserDOT        = 59
	goscriptParserWHITESPACE = 60
)

// goscriptParser rules.
const (
	goscriptParserRULE_program                 = 0
	goscriptParserRULE_statement               = 1
	goscriptParserRULE_functiondef             = 2
	goscriptParserRULE_inparam                 = 3
	goscriptParserRULE_outparam                = 4
	goscriptParserRULE_returntypename          = 5
	goscriptParserRULE_param                   = 6
	goscriptParserRULE_typedef                 = 7
	goscriptParserRULE_typename                = 8
	goscriptParserRULE_enumdef                 = 9
	goscriptParserRULE_execution               = 10
	goscriptParserRULE_control                 = 11
	goscriptParserRULE_collection              = 12
	goscriptParserRULE_block                   = 13
	goscriptParserRULE_line                    = 14
	goscriptParserRULE_lhs                     = 15
	goscriptParserRULE_variable                = 16
	goscriptParserRULE_filter                  = 17
	goscriptParserRULE_indexs                  = 18
	goscriptParserRULE_expr                    = 19
	goscriptParserRULE_initializationListBegin = 20
	goscriptParserRULE_initializationList      = 21
	goscriptParserRULE_constant                = 22
	goscriptParserRULE_functions               = 23
	goscriptParserRULE_constructor             = 24
	goscriptParserRULE_vardef                  = 25
	goscriptParserRULE_localdef                = 26
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *goscriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, goscriptParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<goscriptParserT__1)|(1<<goscriptParserT__15)|(1<<goscriptParserNEW)|(1<<goscriptParserFOR)|(1<<goscriptParserBREAK)|(1<<goscriptParserCONTINUE)|(1<<goscriptParserIF)|(1<<goscriptParserSWITCH)|(1<<goscriptParserRETURN)|(1<<goscriptParserVAR)|(1<<goscriptParserLOCAL)|(1<<goscriptParserFUNCTION)|(1<<goscriptParserTYPEDEF)|(1<<goscriptParserBOOL)|(1<<goscriptParserNULL))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(goscriptParserSUB-37))|(1<<(goscriptParserUNARYADD-37))|(1<<(goscriptParserUNARYSUB-37))|(1<<(goscriptParserNOT-37))|(1<<(goscriptParserINT-37))|(1<<(goscriptParserFLOAT-37))|(1<<(goscriptParserSTRING-37))|(1<<(goscriptParserNAME-37)))) != 0) {
		{
			p.SetState(54)
			p.Statement()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumDefContext struct {
	*StatementContext
}

func NewEnumDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumDefContext {
	var p = new(EnumDefContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) Enumdef() IEnumdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumdefContext)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

type GlobalDefContext struct {
	*StatementContext
}

func NewGlobalDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalDefContext {
	var p = new(GlobalDefContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *GlobalDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalDefContext) Vardef() IVardefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVardefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVardefContext)
}

func (s *GlobalDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterGlobalDef(s)
	}
}

func (s *GlobalDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitGlobalDef(s)
	}
}

type TypeDefContext struct {
	*StatementContext
}

func NewTypeDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefContext {
	var p = new(TypeDefContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) Typedef() ITypedefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedefContext)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

type FuncDefContext struct {
	*StatementContext
}

func NewFuncDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDefContext {
	var p = new(FuncDefContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) Functiondef() IFunctiondefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiondefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

type RunContext struct {
	*StatementContext
}

func NewRunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunContext {
	var p = new(RunContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *RunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunContext) Execution() IExecutionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutionContext)
}

func (s *RunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterRun(s)
	}
}

func (s *RunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitRun(s)
	}
}

func (p *goscriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, goscriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Functiondef()
		}

	case 2:
		localctx = NewTypeDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Typedef()
		}
		{
			p.SetState(61)
			p.Match(goscriptParserT__0)
		}

	case 3:
		localctx = NewGlobalDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Vardef()
		}

	case 4:
		localctx = NewEnumDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Enumdef()
		}

	case 5:
		localctx = NewRunContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(65)
			p.Execution()
		}

	}

	return localctx
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_functiondef
	return p
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) CopyFrom(ctx *FunctiondefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionDefContext struct {
	*FunctiondefContext
}

func NewFunctionDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefContext {
	var p = new(FunctionDefContext)

	p.FunctiondefContext = NewEmptyFunctiondefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctiondefContext))

	return p
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(goscriptParserFUNCTION, 0)
}

func (s *FunctionDefContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *FunctionDefContext) AllInparam() []IInparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInparamContext)(nil)).Elem())
	var tst = make([]IInparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInparamContext)
		}
	}

	return tst
}

func (s *FunctionDefContext) Inparam(i int) IInparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInparamContext)
}

func (s *FunctionDefContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefContext) AllReturntypename() []IReturntypenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReturntypenameContext)(nil)).Elem())
	var tst = make([]IReturntypenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReturntypenameContext)
		}
	}

	return tst
}

func (s *FunctionDefContext) Returntypename(i int) IReturntypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturntypenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReturntypenameContext)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

type FunctionDefWithReturnNameContext struct {
	*FunctiondefContext
}

func NewFunctionDefWithReturnNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefWithReturnNameContext {
	var p = new(FunctionDefWithReturnNameContext)

	p.FunctiondefContext = NewEmptyFunctiondefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctiondefContext))

	return p
}

func (s *FunctionDefWithReturnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefWithReturnNameContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(goscriptParserFUNCTION, 0)
}

func (s *FunctionDefWithReturnNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *FunctionDefWithReturnNameContext) AllInparam() []IInparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInparamContext)(nil)).Elem())
	var tst = make([]IInparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInparamContext)
		}
	}

	return tst
}

func (s *FunctionDefWithReturnNameContext) Inparam(i int) IInparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInparamContext)
}

func (s *FunctionDefWithReturnNameContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefWithReturnNameContext) AllOutparam() []IOutparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOutparamContext)(nil)).Elem())
	var tst = make([]IOutparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOutparamContext)
		}
	}

	return tst
}

func (s *FunctionDefWithReturnNameContext) Outparam(i int) IOutparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOutparamContext)
}

func (s *FunctionDefWithReturnNameContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *FunctionDefWithReturnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFunctionDefWithReturnName(s)
	}
}

func (s *FunctionDefWithReturnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFunctionDefWithReturnName(s)
	}
}

func (p *goscriptParser) Functiondef() (localctx IFunctiondefContext) {
	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, goscriptParserRULE_functiondef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(69)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(70)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(71)
			p.Inparam()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(72)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(73)
				p.Inparam()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(79)
			p.Match(goscriptParserT__3)
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserT__6 || _la == goscriptParserT__7 || _la == goscriptParserNAME {
			{
				p.SetState(80)
				p.Returntypename()
			}

		}
		{
			p.SetState(83)
			p.Block()
		}

	case 2:
		localctx = NewFunctionDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(86)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(87)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(88)
			p.Match(goscriptParserT__3)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserT__6 || _la == goscriptParserT__7 || _la == goscriptParserNAME {
			{
				p.SetState(89)
				p.Returntypename()
			}

		}
		{
			p.SetState(92)
			p.Block()
		}

	case 3:
		localctx = NewFunctionDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(94)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(95)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(96)
			p.Inparam()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(97)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(98)
				p.Inparam()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(104)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(105)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(106)
			p.Returntypename()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(107)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(108)
				p.Returntypename()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(114)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(115)
			p.Block()
		}

	case 4:
		localctx = NewFunctionDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(118)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(119)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(120)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(121)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(122)
			p.Returntypename()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(123)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(124)
				p.Returntypename()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(131)
			p.Block()
		}

	case 5:
		localctx = NewFunctionDefWithReturnNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(133)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(134)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(135)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(136)
			p.Inparam()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(137)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(138)
				p.Inparam()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.Match(goscriptParserT__3)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserNAME {
			{
				p.SetState(145)
				p.Outparam()
			}

		}
		{
			p.SetState(148)
			p.Block()
		}

	case 6:
		localctx = NewFunctionDefWithReturnNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(150)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(151)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(152)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(153)
			p.Match(goscriptParserT__3)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserT__6 || _la == goscriptParserT__7 || _la == goscriptParserNAME {
			{
				p.SetState(154)
				p.Typename()
			}

		}
		{
			p.SetState(157)
			p.Block()
		}

	case 7:
		localctx = NewFunctionDefWithReturnNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(158)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(159)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(160)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(161)
			p.Inparam()
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(162)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(163)
				p.Inparam()
			}

			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(169)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(170)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(171)
			p.Outparam()
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(172)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(173)
				p.Outparam()
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(179)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(180)
			p.Block()
		}

	case 8:
		localctx = NewFunctionDefWithReturnNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(182)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(183)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(184)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(185)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(186)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(187)
			p.Outparam()
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(188)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(189)
				p.Outparam()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(195)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(196)
			p.Block()
		}

	}

	return localctx
}

// IInparamContext is an interface to support dynamic dispatch.
type IInparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInparamContext differentiates from other interfaces.
	IsInparamContext()
}

type InparamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInparamContext() *InparamContext {
	var p = new(InparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_inparam
	return p
}

func (*InparamContext) IsInparamContext() {}

func NewInparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InparamContext {
	var p = new(InparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_inparam

	return p
}

func (s *InparamContext) GetParser() antlr.Parser { return s.parser }

func (s *InparamContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *InparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInparam(s)
	}
}

func (s *InparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInparam(s)
	}
}

func (p *goscriptParser) Inparam() (localctx IInparamContext) {
	localctx = NewInparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, goscriptParserRULE_inparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Param()
	}

	return localctx
}

// IOutparamContext is an interface to support dynamic dispatch.
type IOutparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutparamContext differentiates from other interfaces.
	IsOutparamContext()
}

type OutparamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutparamContext() *OutparamContext {
	var p = new(OutparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_outparam
	return p
}

func (*OutparamContext) IsOutparamContext() {}

func NewOutparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutparamContext {
	var p = new(OutparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_outparam

	return p
}

func (s *OutparamContext) GetParser() antlr.Parser { return s.parser }

func (s *OutparamContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *OutparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterOutparam(s)
	}
}

func (s *OutparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitOutparam(s)
	}
}

func (p *goscriptParser) Outparam() (localctx IOutparamContext) {
	localctx = NewOutparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, goscriptParserRULE_outparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Param()
	}

	return localctx
}

// IReturntypenameContext is an interface to support dynamic dispatch.
type IReturntypenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturntypenameContext differentiates from other interfaces.
	IsReturntypenameContext()
}

type ReturntypenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturntypenameContext() *ReturntypenameContext {
	var p = new(ReturntypenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_returntypename
	return p
}

func (*ReturntypenameContext) IsReturntypenameContext() {}

func NewReturntypenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturntypenameContext {
	var p = new(ReturntypenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_returntypename

	return p
}

func (s *ReturntypenameContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturntypenameContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *ReturntypenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturntypenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturntypenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterReturntypename(s)
	}
}

func (s *ReturntypenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitReturntypename(s)
	}
}

func (p *goscriptParser) Returntypename() (localctx IReturntypenameContext) {
	localctx = NewReturntypenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, goscriptParserRULE_returntypename)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Typename()
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *ParamContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *goscriptParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, goscriptParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(207)
		p.Typename()
	}

	return localctx
}

// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_typedef
	return p
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_typedef

	return p
}

func (s *TypedefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefContext) CopyFrom(ctx *TypedefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeDefComplexContext struct {
	*TypedefContext
}

func NewTypeDefComplexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefComplexContext {
	var p = new(TypeDefComplexContext)

	p.TypedefContext = NewEmptyTypedefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypedefContext))

	return p
}

func (s *TypeDefComplexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefComplexContext) TYPEDEF() antlr.TerminalNode {
	return s.GetToken(goscriptParserTYPEDEF, 0)
}

func (s *TypeDefComplexContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *TypeDefComplexContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *TypeDefComplexContext) AllTypename() []ITypenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypenameContext)(nil)).Elem())
	var tst = make([]ITypenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypenameContext)
		}
	}

	return tst
}

func (s *TypeDefComplexContext) Typename(i int) ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *TypeDefComplexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterTypeDefComplex(s)
	}
}

func (s *TypeDefComplexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitTypeDefComplex(s)
	}
}

type TypeDefAliasContext struct {
	*TypedefContext
}

func NewTypeDefAliasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDefAliasContext {
	var p = new(TypeDefAliasContext)

	p.TypedefContext = NewEmptyTypedefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypedefContext))

	return p
}

func (s *TypeDefAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefAliasContext) TYPEDEF() antlr.TerminalNode {
	return s.GetToken(goscriptParserTYPEDEF, 0)
}

func (s *TypeDefAliasContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *TypeDefAliasContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *TypeDefAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterTypeDefAlias(s)
	}
}

func (s *TypeDefAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitTypeDefAlias(s)
	}
}

func (p *goscriptParser) Typedef() (localctx ITypedefContext) {
	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, goscriptParserRULE_typedef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDefAliasContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(goscriptParserTYPEDEF)
		}
		{
			p.SetState(210)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(211)
			p.Typename()
		}

	case 2:
		localctx = NewTypeDefComplexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(goscriptParserTYPEDEF)
		}
		{
			p.SetState(213)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(214)
			p.Match(goscriptParserT__4)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserNAME {
			{
				p.SetState(215)
				p.Match(goscriptParserNAME)
			}
			{
				p.SetState(216)
				p.Typename()
			}
			{
				p.SetState(217)
				p.Match(goscriptParserT__2)
			}

			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(224)
			p.Match(goscriptParserT__5)
		}

	}

	return localctx
}

// ITypenameContext is an interface to support dynamic dispatch.
type ITypenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypenameContext differentiates from other interfaces.
	IsTypenameContext()
}

type TypenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypenameContext() *TypenameContext {
	var p = new(TypenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_typename
	return p
}

func (*TypenameContext) IsTypenameContext() {}

func NewTypenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypenameContext {
	var p = new(TypenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_typename

	return p
}

func (s *TypenameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypenameContext) CopyFrom(ctx *TypenameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MapTypeNameContext struct {
	*TypenameContext
}

func NewMapTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapTypeNameContext {
	var p = new(MapTypeNameContext)

	p.TypenameContext = NewEmptyTypenameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypenameContext))

	return p
}

func (s *MapTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeNameContext) LT() antlr.TerminalNode {
	return s.GetToken(goscriptParserLT, 0)
}

func (s *MapTypeNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *MapTypeNameContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *MapTypeNameContext) GT() antlr.TerminalNode {
	return s.GetToken(goscriptParserGT, 0)
}

func (s *MapTypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterMapTypeName(s)
	}
}

func (s *MapTypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitMapTypeName(s)
	}
}

type SimpleTypeNameContext struct {
	*TypenameContext
}

func NewSimpleTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleTypeNameContext {
	var p = new(SimpleTypeNameContext)

	p.TypenameContext = NewEmptyTypenameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypenameContext))

	return p
}

func (s *SimpleTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *SimpleTypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSimpleTypeName(s)
	}
}

func (s *SimpleTypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSimpleTypeName(s)
	}
}

type SliceTypeNameContext struct {
	*TypenameContext
}

func NewSliceTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceTypeNameContext {
	var p = new(SliceTypeNameContext)

	p.TypenameContext = NewEmptyTypenameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypenameContext))

	return p
}

func (s *SliceTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTypeNameContext) LT() antlr.TerminalNode {
	return s.GetToken(goscriptParserLT, 0)
}

func (s *SliceTypeNameContext) Typename() ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *SliceTypeNameContext) GT() antlr.TerminalNode {
	return s.GetToken(goscriptParserGT, 0)
}

func (s *SliceTypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSliceTypeName(s)
	}
}

func (s *SliceTypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSliceTypeName(s)
	}
}

func (p *goscriptParser) Typename() (localctx ITypenameContext) {
	localctx = NewTypenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, goscriptParserRULE_typename)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserNAME:
		localctx = NewSimpleTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(goscriptParserNAME)
		}

	case goscriptParserT__6:
		localctx = NewMapTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(goscriptParserT__6)
		}
		{
			p.SetState(229)
			p.Match(goscriptParserLT)
		}
		{
			p.SetState(230)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(231)
			p.Match(goscriptParserT__2)
		}
		{
			p.SetState(232)
			p.Typename()
		}
		{
			p.SetState(233)
			p.Match(goscriptParserGT)
		}

	case goscriptParserT__7:
		localctx = NewSliceTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(236)
			p.Match(goscriptParserLT)
		}
		{
			p.SetState(237)
			p.Typename()
		}
		{
			p.SetState(238)
			p.Match(goscriptParserGT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumdefContext is an interface to support dynamic dispatch.
type IEnumdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumdefContext differentiates from other interfaces.
	IsEnumdefContext()
}

type EnumdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumdefContext() *EnumdefContext {
	var p = new(EnumdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_enumdef
	return p
}

func (*EnumdefContext) IsEnumdefContext() {}

func NewEnumdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumdefContext {
	var p = new(EnumdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_enumdef

	return p
}

func (s *EnumdefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumdefContext) TYPEDEF() antlr.TerminalNode {
	return s.GetToken(goscriptParserTYPEDEF, 0)
}

func (s *EnumdefContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *EnumdefContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *EnumdefContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserINT)
}

func (s *EnumdefContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserINT, i)
}

func (s *EnumdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterEnumdef(s)
	}
}

func (s *EnumdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitEnumdef(s)
	}
}

func (p *goscriptParser) Enumdef() (localctx IEnumdefContext) {
	localctx = NewEnumdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, goscriptParserRULE_enumdef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(goscriptParserTYPEDEF)
	}
	{
		p.SetState(243)
		p.Match(goscriptParserT__8)
	}
	{
		p.SetState(244)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(245)
		p.Match(goscriptParserT__4)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == goscriptParserNAME {
		{
			p.SetState(246)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(247)
			p.Match(goscriptParserINT)
		}
		{
			p.SetState(248)
			p.Match(goscriptParserT__2)
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)
		p.Match(goscriptParserT__5)
	}

	return localctx
}

// IExecutionContext is an interface to support dynamic dispatch.
type IExecutionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutionContext differentiates from other interfaces.
	IsExecutionContext()
}

type ExecutionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutionContext() *ExecutionContext {
	var p = new(ExecutionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_execution
	return p
}

func (*ExecutionContext) IsExecutionContext() {}

func NewExecutionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutionContext {
	var p = new(ExecutionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_execution

	return p
}

func (s *ExecutionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutionContext) CopyFrom(ctx *ExecutionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExecutionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LineProgramContext struct {
	*ExecutionContext
}

func NewLineProgramContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LineProgramContext {
	var p = new(LineProgramContext)

	p.ExecutionContext = NewEmptyExecutionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExecutionContext))

	return p
}

func (s *LineProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineProgramContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *LineProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLineProgram(s)
	}
}

func (s *LineProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLineProgram(s)
	}
}

type CtrlContext struct {
	*ExecutionContext
}

func NewCtrlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CtrlContext {
	var p = new(CtrlContext)

	p.ExecutionContext = NewEmptyExecutionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExecutionContext))

	return p
}

func (s *CtrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtrlContext) Control() IControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlContext)
}

func (s *CtrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterCtrl(s)
	}
}

func (s *CtrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitCtrl(s)
	}
}

func (p *goscriptParser) Execution() (localctx IExecutionContext) {
	localctx = NewExecutionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, goscriptParserRULE_execution)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(260)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN:
		localctx = NewCtrlContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Control()
		}

	case goscriptParserT__1, goscriptParserT__15, goscriptParserNEW, goscriptParserLOCAL, goscriptParserBOOL, goscriptParserNULL, goscriptParserSUB, goscriptParserUNARYADD, goscriptParserUNARYSUB, goscriptParserNOT, goscriptParserINT, goscriptParserFLOAT, goscriptParserSTRING, goscriptParserNAME:
		localctx = NewLineProgramContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Line()
		}
		{
			p.SetState(258)
			p.Match(goscriptParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControlContext is an interface to support dynamic dispatch.
type IControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlContext differentiates from other interfaces.
	IsControlContext()
}

type ControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlContext() *ControlContext {
	var p = new(ControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_control
	return p
}

func (*ControlContext) IsControlContext() {}

func NewControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlContext {
	var p = new(ControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_control

	return p
}

func (s *ControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlContext) CopyFrom(ctx *ControlContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnVoidContext struct {
	*ControlContext
}

func NewReturnVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnVoidContext {
	var p = new(ReturnVoidContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ReturnVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnVoidContext) RETURN() antlr.TerminalNode {
	return s.GetToken(goscriptParserRETURN, 0)
}

func (s *ReturnVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterReturnVoid(s)
	}
}

func (s *ReturnVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitReturnVoid(s)
	}
}

type ForInSliceContext struct {
	*ControlContext
}

func NewForInSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInSliceContext {
	var p = new(ForInSliceContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ForInSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInSliceContext) FOR() antlr.TerminalNode {
	return s.GetToken(goscriptParserFOR, 0)
}

func (s *ForInSliceContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *ForInSliceContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ForInSliceContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterForInSlice(s)
	}
}

func (s *ForInSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitForInSlice(s)
	}
}

type SwitchContext struct {
	*ControlContext
}

func NewSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchContext {
	var p = new(SwitchContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(goscriptParserSWITCH, 0)
}

func (s *SwitchContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserCASE)
}

func (s *SwitchContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserCASE, i)
}

func (s *SwitchContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *SwitchContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *SwitchContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *SwitchContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSwitch(s)
	}
}

type ReturnValContext struct {
	*ControlContext
}

func NewReturnValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnValContext {
	var p = new(ReturnValContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ReturnValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnValContext) RETURN() antlr.TerminalNode {
	return s.GetToken(goscriptParserRETURN, 0)
}

func (s *ReturnValContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ReturnValContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterReturnVal(s)
	}
}

func (s *ReturnValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitReturnVal(s)
	}
}

type ForInMapContext struct {
	*ControlContext
}

func NewForInMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInMapContext {
	var p = new(ForInMapContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ForInMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInMapContext) FOR() antlr.TerminalNode {
	return s.GetToken(goscriptParserFOR, 0)
}

func (s *ForInMapContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *ForInMapContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *ForInMapContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ForInMapContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterForInMap(s)
	}
}

func (s *ForInMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitForInMap(s)
	}
}

type ForContext struct {
	*ControlContext
}

func NewForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForContext {
	var p = new(ForContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(goscriptParserFOR, 0)
}

func (s *ForContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ForContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFor(s)
	}
}

type BreakContext struct {
	*ControlContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(goscriptParserBREAK, 0)
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitBreak(s)
	}
}

type ContinueContext struct {
	*ControlContext
}

func NewContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueContext {
	var p = new(ContinueContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(goscriptParserCONTINUE, 0)
}

func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitContinue(s)
	}
}

type IfContext struct {
	*ControlContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ControlContext = NewEmptyControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ControlContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(goscriptParserIF, 0)
}

func (s *IfContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goscriptParserELSE, 0)
}

func (s *IfContext) Control() IControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *goscriptParser) Control() (localctx IControlContext) {
	localctx = NewControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, goscriptParserRULE_control)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(goscriptParserIF)
		}
		{
			p.SetState(263)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(264)
			p.expr(0)
		}
		{
			p.SetState(265)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(266)
			p.Block()
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserELSE {
			{
				p.SetState(267)
				p.Match(goscriptParserELSE)
			}
			p.SetState(270)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case goscriptParserT__4:
				{
					p.SetState(268)
					p.Block()
				}

			case goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN:
				{
					p.SetState(269)
					p.Control()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}

	case 2:
		localctx = NewSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(goscriptParserSWITCH)
		}
		{
			p.SetState(275)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(276)
			p.expr(0)
		}
		{
			p.SetState(277)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(278)
			p.Match(goscriptParserT__4)
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == goscriptParserCASE {
			{
				p.SetState(279)
				p.Match(goscriptParserCASE)
			}
			{
				p.SetState(280)
				p.Constant()
			}
			{
				p.SetState(281)
				p.Match(goscriptParserT__9)
			}
			{
				p.SetState(282)
				p.Block()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(288)
			p.Match(goscriptParserT__5)
		}

	case 3:
		localctx = NewForInSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(290)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(291)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(292)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(293)
			p.Match(goscriptParserT__10)
		}
		{
			p.SetState(294)
			p.Collection()
		}
		{
			p.SetState(295)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(296)
			p.Block()
		}

	case 4:
		localctx = NewForInMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(298)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(299)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(300)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(301)
			p.Match(goscriptParserT__2)
		}
		{
			p.SetState(302)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(303)
			p.Match(goscriptParserT__10)
		}
		{
			p.SetState(304)
			p.Collection()
		}
		{
			p.SetState(305)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(306)
			p.Block()
		}

	case 5:
		localctx = NewForContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(308)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(309)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(310)
			p.expr(0)
		}
		{
			p.SetState(311)
			p.Match(goscriptParserT__0)
		}
		{
			p.SetState(312)
			p.expr(0)
		}
		{
			p.SetState(313)
			p.Match(goscriptParserT__0)
		}
		{
			p.SetState(314)
			p.expr(0)
		}
		{
			p.SetState(315)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(316)
			p.Block()
		}

	case 6:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(318)
			p.Match(goscriptParserBREAK)
		}
		{
			p.SetState(319)
			p.Match(goscriptParserT__0)
		}

	case 7:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(320)
			p.Match(goscriptParserCONTINUE)
		}
		{
			p.SetState(321)
			p.Match(goscriptParserT__0)
		}

	case 8:
		localctx = NewReturnVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(322)
			p.Match(goscriptParserRETURN)
		}
		{
			p.SetState(323)
			p.Match(goscriptParserT__0)
		}

	case 9:
		localctx = NewReturnValContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(324)
			p.Match(goscriptParserRETURN)
		}
		{
			p.SetState(325)
			p.expr(0)
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(326)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(327)
				p.expr(0)
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(333)
			p.Match(goscriptParserT__0)
		}

	}

	return localctx
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_collection
	return p
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *goscriptParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, goscriptParserRULE_collection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.expr(0)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllExecution() []IExecutionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecutionContext)(nil)).Elem())
	var tst = make([]IExecutionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecutionContext)
		}
	}

	return tst
}

func (s *BlockContext) Execution(i int) IExecutionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecutionContext)
}

func (s *BlockContext) AllVardef() []IVardefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVardefContext)(nil)).Elem())
	var tst = make([]IVardefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVardefContext)
		}
	}

	return tst
}

func (s *BlockContext) Vardef(i int) IVardefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVardefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVardefContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *goscriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, goscriptParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(goscriptParserT__4)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<goscriptParserT__1)|(1<<goscriptParserT__15)|(1<<goscriptParserNEW)|(1<<goscriptParserFOR)|(1<<goscriptParserBREAK)|(1<<goscriptParserCONTINUE)|(1<<goscriptParserIF)|(1<<goscriptParserSWITCH)|(1<<goscriptParserRETURN)|(1<<goscriptParserVAR)|(1<<goscriptParserLOCAL)|(1<<goscriptParserBOOL)|(1<<goscriptParserNULL))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(goscriptParserSUB-37))|(1<<(goscriptParserUNARYADD-37))|(1<<(goscriptParserUNARYSUB-37))|(1<<(goscriptParserNOT-37))|(1<<(goscriptParserINT-37))|(1<<(goscriptParserFLOAT-37))|(1<<(goscriptParserSTRING-37))|(1<<(goscriptParserNAME-37)))) != 0) {
		p.SetState(342)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case goscriptParserT__1, goscriptParserT__15, goscriptParserNEW, goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN, goscriptParserLOCAL, goscriptParserBOOL, goscriptParserNULL, goscriptParserSUB, goscriptParserUNARYADD, goscriptParserUNARYSUB, goscriptParserNOT, goscriptParserINT, goscriptParserFLOAT, goscriptParserSTRING, goscriptParserNAME:
			{
				p.SetState(340)
				p.Execution()
			}

		case goscriptParserVAR:
			{
				p.SetState(341)
				p.Vardef()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
		p.Match(goscriptParserT__5)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) CopyFrom(ctx *LineContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprEntryContext struct {
	*LineContext
}

func NewExprEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprEntryContext {
	var p = new(ExprEntryContext)

	p.LineContext = NewEmptyLineContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LineContext))

	return p
}

func (s *ExprEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEntryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterExprEntry(s)
	}
}

func (s *ExprEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitExprEntry(s)
	}
}

type FunctionAssignContext struct {
	*LineContext
	op antlr.Token
}

func NewFunctionAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionAssignContext {
	var p = new(FunctionAssignContext)

	p.LineContext = NewEmptyLineContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LineContext))

	return p
}

func (s *FunctionAssignContext) GetOp() antlr.Token { return s.op }

func (s *FunctionAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *FunctionAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionAssignContext) AllLhs() []ILhsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILhsContext)(nil)).Elem())
	var tst = make([]ILhsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILhsContext)
		}
	}

	return tst
}

func (s *FunctionAssignContext) Lhs(i int) ILhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *FunctionAssignContext) Functions() IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *FunctionAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *FunctionAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFunctionAssign(s)
	}
}

func (s *FunctionAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFunctionAssign(s)
	}
}

type MultiAssignContext struct {
	*LineContext
	op antlr.Token
}

func NewMultiAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiAssignContext {
	var p = new(MultiAssignContext)

	p.LineContext = NewEmptyLineContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LineContext))

	return p
}

func (s *MultiAssignContext) GetOp() antlr.Token { return s.op }

func (s *MultiAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiAssignContext) AllLhs() []ILhsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILhsContext)(nil)).Elem())
	var tst = make([]ILhsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILhsContext)
		}
	}

	return tst
}

func (s *MultiAssignContext) Lhs(i int) ILhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *MultiAssignContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultiAssignContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *MultiAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterMultiAssign(s)
	}
}

func (s *MultiAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitMultiAssign(s)
	}
}

func (p *goscriptParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, goscriptParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.expr(0)
		}

	case 2:
		localctx = NewFunctionAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.Lhs()
		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(351)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(352)
				p.Lhs()
			}

			p.SetState(357)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(358)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*FunctionAssignContext).op = _m
		}
		{
			p.SetState(359)
			p.Functions()
		}

	case 3:
		localctx = NewMultiAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(361)
			p.Lhs()
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(362)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(363)
				p.Lhs()
			}

			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(369)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*MultiAssignContext).op = _m
		}
		{
			p.SetState(370)
			p.expr(0)
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(371)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(372)
				p.expr(0)
			}

			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLhsContext differentiates from other interfaces.
	IsLhsContext()
}

type LhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsContext() *LhsContext {
	var p = new(LhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_lhs
	return p
}

func (*LhsContext) IsLhsContext() {}

func NewLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsContext {
	var p = new(LhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_lhs

	return p
}

func (s *LhsContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *LhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLhs(s)
	}
}

func (s *LhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLhs(s)
	}
}

func (p *goscriptParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, goscriptParserRULE_lhs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.variable(0)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) CopyFrom(ctx *VariableContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceFilterContext struct {
	*VariableContext
}

func NewSliceFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceFilterContext {
	var p = new(SliceFilterContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *SliceFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceFilterContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *SliceFilterContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *SliceFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSliceFilter(s)
	}
}

func (s *SliceFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSliceFilter(s)
	}
}

type VariableNameContext struct {
	*VariableContext
}

func NewVariableNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableNameContext {
	var p = new(VariableNameContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *VariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *VariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterVariableName(s)
	}
}

func (s *VariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitVariableName(s)
	}
}

type SelectContext struct {
	*VariableContext
}

func NewSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectContext {
	var p = new(SelectContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *SelectContext) DOT() antlr.TerminalNode {
	return s.GetToken(goscriptParserDOT, 0)
}

func (s *SelectContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSelect(s)
	}
}

type IndexContext struct {
	*VariableContext
}

func NewIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexContext {
	var p = new(IndexContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *IndexContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndex(s)
	}
}

type MapMultiIndexContext struct {
	*VariableContext
}

func NewMapMultiIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapMultiIndexContext {
	var p = new(MapMultiIndexContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *MapMultiIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapMultiIndexContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *MapMultiIndexContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MapMultiIndexContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapMultiIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterMapMultiIndex(s)
	}
}

func (s *MapMultiIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitMapMultiIndex(s)
	}
}

type SliceMultiIndexContext struct {
	*VariableContext
}

func NewSliceMultiIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceMultiIndexContext {
	var p = new(SliceMultiIndexContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *SliceMultiIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceMultiIndexContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *SliceMultiIndexContext) AllIndexs() []IIndexsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexsContext)(nil)).Elem())
	var tst = make([]IIndexsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexsContext)
		}
	}

	return tst
}

func (s *SliceMultiIndexContext) Indexs(i int) IIndexsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexsContext)
}

func (s *SliceMultiIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterSliceMultiIndex(s)
	}
}

func (s *SliceMultiIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitSliceMultiIndex(s)
	}
}

func (p *goscriptParser) Variable() (localctx IVariableContext) {
	return p.variable(0)
}

func (p *goscriptParser) variable(_p int) (localctx IVariableContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, goscriptParserRULE_variable, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserNAME:
		localctx = NewVariableNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(383)
			p.Match(goscriptParserNAME)
		}

	case goscriptParserT__15:
		localctx = NewVariableNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(384)
			p.Match(goscriptParserT__15)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(387)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(388)
					p.Match(goscriptParserDOT)
				}
				{
					p.SetState(389)
					p.Match(goscriptParserNAME)
				}

			case 2:
				localctx = NewSliceFilterContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(390)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(391)
					p.Match(goscriptParserT__11)
				}
				{
					p.SetState(392)
					p.Filter()
				}
				{
					p.SetState(393)
					p.Match(goscriptParserT__12)
				}

			case 3:
				localctx = NewIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(395)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(396)
					p.Match(goscriptParserT__13)
				}
				{
					p.SetState(397)
					p.expr(0)
				}
				{
					p.SetState(398)
					p.Match(goscriptParserT__14)
				}

			case 4:
				localctx = NewSliceMultiIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(401)
					p.Match(goscriptParserT__13)
				}
				{
					p.SetState(402)
					p.Indexs()
				}
				p.SetState(407)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == goscriptParserT__2 {
					{
						p.SetState(403)
						p.Match(goscriptParserT__2)
					}
					{
						p.SetState(404)
						p.Indexs()
					}

					p.SetState(409)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(410)
					p.Match(goscriptParserT__14)
				}

			case 5:
				localctx = NewMapMultiIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(412)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(413)
					p.Match(goscriptParserT__13)
				}
				{
					p.SetState(414)
					p.Match(goscriptParserT__13)
				}
				{
					p.SetState(415)
					p.expr(0)
				}
				p.SetState(420)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == goscriptParserT__2 {
					{
						p.SetState(416)
						p.Match(goscriptParserT__2)
					}
					{
						p.SetState(417)
						p.expr(0)
					}

					p.SetState(422)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(423)
					p.Match(goscriptParserT__14)
				}
				{
					p.SetState(424)
					p.Match(goscriptParserT__14)
				}

			}

		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *goscriptParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, goscriptParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.expr(0)
	}

	return localctx
}

// IIndexsContext is an interface to support dynamic dispatch.
type IIndexsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexsContext differentiates from other interfaces.
	IsIndexsContext()
}

type IndexsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexsContext() *IndexsContext {
	var p = new(IndexsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_indexs
	return p
}

func (*IndexsContext) IsIndexsContext() {}

func NewIndexsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexsContext {
	var p = new(IndexsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_indexs

	return p
}

func (s *IndexsContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexsContext) CopyFrom(ctx *IndexsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IndexsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexType1Context struct {
	*IndexsContext
}

func NewIndexType1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexType1Context {
	var p = new(IndexType1Context)

	p.IndexsContext = NewEmptyIndexsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IndexsContext))

	return p
}

func (s *IndexType1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexType1Context) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IndexType1Context) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexType1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndexType1(s)
	}
}

func (s *IndexType1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndexType1(s)
	}
}

type IndexType3Context struct {
	*IndexsContext
}

func NewIndexType3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexType3Context {
	var p = new(IndexType3Context)

	p.IndexsContext = NewEmptyIndexsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IndexsContext))

	return p
}

func (s *IndexType3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexType3Context) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexType3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndexType3(s)
	}
}

func (s *IndexType3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndexType3(s)
	}
}

type IndexType2Context struct {
	*IndexsContext
}

func NewIndexType2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexType2Context {
	var p = new(IndexType2Context)

	p.IndexsContext = NewEmptyIndexsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IndexsContext))

	return p
}

func (s *IndexType2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexType2Context) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IndexType2Context) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexType2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndexType2(s)
	}
}

func (s *IndexType2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndexType2(s)
	}
}

type IndexType5Context struct {
	*IndexsContext
}

func NewIndexType5Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexType5Context {
	var p = new(IndexType5Context)

	p.IndexsContext = NewEmptyIndexsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IndexsContext))

	return p
}

func (s *IndexType5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexType5Context) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexType5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndexType5(s)
	}
}

func (s *IndexType5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndexType5(s)
	}
}

type IndexType4Context struct {
	*IndexsContext
}

func NewIndexType4Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexType4Context {
	var p = new(IndexType4Context)

	p.IndexsContext = NewEmptyIndexsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IndexsContext))

	return p
}

func (s *IndexType4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexType4Context) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexType4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterIndexType4(s)
	}
}

func (s *IndexType4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitIndexType4(s)
	}
}

func (p *goscriptParser) Indexs() (localctx IIndexsContext) {
	localctx = NewIndexsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, goscriptParserRULE_indexs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIndexType1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.expr(0)
		}
		{
			p.SetState(434)
			p.Match(goscriptParserT__9)
		}
		{
			p.SetState(435)
			p.expr(0)
		}
		{
			p.SetState(436)
			p.Match(goscriptParserT__9)
		}
		{
			p.SetState(437)
			p.expr(0)
		}

	case 2:
		localctx = NewIndexType2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)
			p.expr(0)
		}
		{
			p.SetState(440)
			p.Match(goscriptParserT__9)
		}
		{
			p.SetState(441)
			p.expr(0)
		}

	case 3:
		localctx = NewIndexType3Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.expr(0)
		}
		{
			p.SetState(444)
			p.Match(goscriptParserT__9)
		}

	case 4:
		localctx = NewIndexType4Context(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(446)
			p.expr(0)
		}

	case 5:
		localctx = NewIndexType5Context(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(447)
			p.Match(goscriptParserT__9)
		}
		{
			p.SetState(448)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassContext struct {
	*ExprContext
}

func NewPassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassContext {
	var p = new(PassContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PassContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PassContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PassContext) Functions() IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *PassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterPass(s)
	}
}

func (s *PassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitPass(s)
	}
}

type ConstructContext struct {
	*ExprContext
}

func NewConstructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstructContext {
	var p = new(ConstructContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructContext) Constructor() IConstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstructorContext)
}

func (s *ConstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstruct(s)
	}
}

func (s *ConstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstruct(s)
	}
}

type BinaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryContext) GetOp() antlr.Token { return s.op }

func (s *BinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) Lhs() ILhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *BinaryContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *BinaryContext) ADDEQUAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserADDEQUAL, 0)
}

func (s *BinaryContext) SUBEQUAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserSUBEQUAL, 0)
}

func (s *BinaryContext) MULEQUAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserMULEQUAL, 0)
}

func (s *BinaryContext) DIVEQUAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserDIVEQUAL, 0)
}

func (s *BinaryContext) POW() antlr.TerminalNode {
	return s.GetToken(goscriptParserPOW, 0)
}

func (s *BinaryContext) MUL() antlr.TerminalNode {
	return s.GetToken(goscriptParserMUL, 0)
}

func (s *BinaryContext) DIV() antlr.TerminalNode {
	return s.GetToken(goscriptParserDIV, 0)
}

func (s *BinaryContext) MOD() antlr.TerminalNode {
	return s.GetToken(goscriptParserMOD, 0)
}

func (s *BinaryContext) ADD() antlr.TerminalNode {
	return s.GetToken(goscriptParserADD, 0)
}

func (s *BinaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(goscriptParserSUB, 0)
}

func (s *BinaryContext) EQ() antlr.TerminalNode {
	return s.GetToken(goscriptParserEQ, 0)
}

func (s *BinaryContext) INEQ() antlr.TerminalNode {
	return s.GetToken(goscriptParserINEQ, 0)
}

func (s *BinaryContext) GT() antlr.TerminalNode {
	return s.GetToken(goscriptParserGT, 0)
}

func (s *BinaryContext) GE() antlr.TerminalNode {
	return s.GetToken(goscriptParserGE, 0)
}

func (s *BinaryContext) LT() antlr.TerminalNode {
	return s.GetToken(goscriptParserLT, 0)
}

func (s *BinaryContext) LE() antlr.TerminalNode {
	return s.GetToken(goscriptParserLE, 0)
}

func (s *BinaryContext) REGEX() antlr.TerminalNode {
	return s.GetToken(goscriptParserREGEX, 0)
}

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(goscriptParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(goscriptParserOR, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitBinary(s)
	}
}

type LeftUnaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewLeftUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeftUnaryContext {
	var p = new(LeftUnaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LeftUnaryContext) GetOp() antlr.Token { return s.op }

func (s *LeftUnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *LeftUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftUnaryContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *LeftUnaryContext) UNARYADD() antlr.TerminalNode {
	return s.GetToken(goscriptParserUNARYADD, 0)
}

func (s *LeftUnaryContext) UNARYSUB() antlr.TerminalNode {
	return s.GetToken(goscriptParserUNARYSUB, 0)
}

func (s *LeftUnaryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LeftUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(goscriptParserNOT, 0)
}

func (s *LeftUnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(goscriptParserSUB, 0)
}

func (s *LeftUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLeftUnary(s)
	}
}

func (s *LeftUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLeftUnary(s)
	}
}

type RightUnaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewRightUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightUnaryContext {
	var p = new(RightUnaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RightUnaryContext) GetOp() antlr.Token { return s.op }

func (s *RightUnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *RightUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightUnaryContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *RightUnaryContext) UNARYADD() antlr.TerminalNode {
	return s.GetToken(goscriptParserUNARYADD, 0)
}

func (s *RightUnaryContext) UNARYSUB() antlr.TerminalNode {
	return s.GetToken(goscriptParserUNARYSUB, 0)
}

func (s *RightUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterRightUnary(s)
	}
}

func (s *RightUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitRightUnary(s)
	}
}

type AssignInitializationlistContext struct {
	*ExprContext
	op antlr.Token
}

func NewAssignInitializationlistContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignInitializationlistContext {
	var p = new(AssignInitializationlistContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignInitializationlistContext) GetOp() antlr.Token { return s.op }

func (s *AssignInitializationlistContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignInitializationlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignInitializationlistContext) Lhs() ILhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *AssignInitializationlistContext) InitializationListBegin() IInitializationListBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializationListBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializationListBeginContext)
}

func (s *AssignInitializationlistContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *AssignInitializationlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterAssignInitializationlist(s)
	}
}

func (s *AssignInitializationlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitAssignInitializationlist(s)
	}
}

type LocalDefContext struct {
	*ExprContext
}

func NewLocalDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalDefContext {
	var p = new(LocalDefContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LocalDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalDefContext) Localdef() ILocaldefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocaldefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocaldefContext)
}

func (s *LocalDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLocalDef(s)
	}
}

func (s *LocalDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLocalDef(s)
	}
}

func (p *goscriptParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *goscriptParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, goscriptParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(452)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(453)
			p.expr(0)
		}
		{
			p.SetState(454)
			p.Match(goscriptParserT__3)
		}

	case 2:
		localctx = NewLeftUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(456)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LeftUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == goscriptParserUNARYADD || _la == goscriptParserUNARYSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LeftUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(457)
			p.variable(0)
		}

	case 3:
		localctx = NewLeftUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(458)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LeftUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == goscriptParserSUB || _la == goscriptParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LeftUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(459)
			p.expr(15)
		}

	case 4:
		localctx = NewRightUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(460)
			p.variable(0)
		}
		{
			p.SetState(461)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*RightUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == goscriptParserUNARYADD || _la == goscriptParserUNARYSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*RightUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		localctx = NewBinaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(463)
			p.Lhs()
		}
		{
			p.SetState(464)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BinaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(goscriptParserASSIGN-50))|(1<<(goscriptParserADDEQUAL-50))|(1<<(goscriptParserSUBEQUAL-50))|(1<<(goscriptParserMULEQUAL-50))|(1<<(goscriptParserDIVEQUAL-50)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BinaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(465)
			p.expr(7)
		}

	case 6:
		localctx = NewAssignInitializationlistContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(467)
			p.Lhs()
		}
		{
			p.SetState(468)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*AssignInitializationlistContext).op = _m
		}
		{
			p.SetState(469)
			p.InitializationListBegin()
		}

	case 7:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(471)
			p.Constant()
		}

	case 8:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(472)
			p.variable(0)
		}

	case 9:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(473)
			p.Functions()
		}

	case 10:
		localctx = NewConstructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(474)
			p.Constructor()
		}

	case 11:
		localctx = NewLocalDefContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(475)
			p.Localdef()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(496)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(478)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(479)

					var _m = p.Match(goscriptParserPOW)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(480)
					p.expr(13)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(481)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(482)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(goscriptParserMUL-33))|(1<<(goscriptParserDIV-33))|(1<<(goscriptParserMOD-33)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(483)
					p.expr(13)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(484)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(485)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == goscriptParserADD || _la == goscriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(486)
					p.expr(12)
				}

			case 4:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(487)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(488)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(goscriptParserEQ-40))|(1<<(goscriptParserINEQ-40))|(1<<(goscriptParserGT-40))|(1<<(goscriptParserGE-40))|(1<<(goscriptParserLE-40))|(1<<(goscriptParserLT-40))|(1<<(goscriptParserREGEX-40)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(489)
					p.expr(11)
				}

			case 5:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(490)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(491)

					var _m = p.Match(goscriptParserAND)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(492)
					p.expr(10)
				}

			case 6:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(493)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(494)

					var _m = p.Match(goscriptParserOR)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(495)
					p.expr(9)
				}

			}

		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IInitializationListBeginContext is an interface to support dynamic dispatch.
type IInitializationListBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitializationListBeginContext differentiates from other interfaces.
	IsInitializationListBeginContext()
}

type InitializationListBeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializationListBeginContext() *InitializationListBeginContext {
	var p = new(InitializationListBeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_initializationListBegin
	return p
}

func (*InitializationListBeginContext) IsInitializationListBeginContext() {}

func NewInitializationListBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializationListBeginContext {
	var p = new(InitializationListBeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_initializationListBegin

	return p
}

func (s *InitializationListBeginContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializationListBeginContext) InitializationList() IInitializationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializationListContext)
}

func (s *InitializationListBeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializationListBeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializationListBeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInitializationListBegin(s)
	}
}

func (s *InitializationListBeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInitializationListBegin(s)
	}
}

func (p *goscriptParser) InitializationListBegin() (localctx IInitializationListBeginContext) {
	localctx = NewInitializationListBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, goscriptParserRULE_initializationListBegin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.InitializationList()
	}

	return localctx
}

// IInitializationListContext is an interface to support dynamic dispatch.
type IInitializationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitializationListContext differentiates from other interfaces.
	IsInitializationListContext()
}

type InitializationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializationListContext() *InitializationListContext {
	var p = new(InitializationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_initializationList
	return p
}

func (*InitializationListContext) IsInitializationListContext() {}

func NewInitializationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializationListContext {
	var p = new(InitializationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_initializationList

	return p
}

func (s *InitializationListContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializationListContext) CopyFrom(ctx *InitializationListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InitializationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InitMessageContext struct {
	*InitializationListContext
}

func NewInitMessageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InitMessageContext {
	var p = new(InitMessageContext)

	p.InitializationListContext = NewEmptyInitializationListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InitializationListContext))

	return p
}

func (s *InitMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitMessageContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *InitMessageContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *InitMessageContext) AllInitializationList() []IInitializationListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInitializationListContext)(nil)).Elem())
	var tst = make([]IInitializationListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInitializationListContext)
		}
	}

	return tst
}

func (s *InitMessageContext) InitializationList(i int) IInitializationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializationListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInitializationListContext)
}

func (s *InitMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInitMessage(s)
	}
}

func (s *InitMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInitMessage(s)
	}
}

type InitConstantContext struct {
	*InitializationListContext
}

func NewInitConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InitConstantContext {
	var p = new(InitConstantContext)

	p.InitializationListContext = NewEmptyInitializationListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InitializationListContext))

	return p
}

func (s *InitConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitConstantContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InitConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInitConstant(s)
	}
}

func (s *InitConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInitConstant(s)
	}
}

type InitMapContext struct {
	*InitializationListContext
}

func NewInitMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InitMapContext {
	var p = new(InitMapContext)

	p.InitializationListContext = NewEmptyInitializationListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InitializationListContext))

	return p
}

func (s *InitMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitMapContext) AllInitializationList() []IInitializationListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInitializationListContext)(nil)).Elem())
	var tst = make([]IInitializationListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInitializationListContext)
		}
	}

	return tst
}

func (s *InitMapContext) InitializationList(i int) IInitializationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializationListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInitializationListContext)
}

func (s *InitMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInitMap(s)
	}
}

func (s *InitMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInitMap(s)
	}
}

type InitSliceContext struct {
	*InitializationListContext
}

func NewInitSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InitSliceContext {
	var p = new(InitSliceContext)

	p.InitializationListContext = NewEmptyInitializationListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InitializationListContext))

	return p
}

func (s *InitSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitSliceContext) AllInitializationList() []IInitializationListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInitializationListContext)(nil)).Elem())
	var tst = make([]IInitializationListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInitializationListContext)
		}
	}

	return tst
}

func (s *InitSliceContext) InitializationList(i int) IInitializationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializationListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInitializationListContext)
}

func (s *InitSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterInitSlice(s)
	}
}

func (s *InitSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitInitSlice(s)
	}
}

func (p *goscriptParser) InitializationList() (localctx IInitializationListContext) {
	localctx = NewInitializationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, goscriptParserRULE_initializationList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInitSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.Match(goscriptParserT__13)
		}
		{
			p.SetState(504)
			p.InitializationList()
		}
		p.SetState(509)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(505)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(506)
				p.InitializationList()
			}

			p.SetState(511)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(512)
			p.Match(goscriptParserT__14)
		}

	case 2:
		localctx = NewInitMessageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Match(goscriptParserT__4)
		}
		{
			p.SetState(515)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(516)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(517)
			p.InitializationList()
		}
		{
			p.SetState(518)
			p.Match(goscriptParserT__3)
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(519)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(520)
				p.Match(goscriptParserNAME)
			}
			{
				p.SetState(521)
				p.Match(goscriptParserT__1)
			}
			{
				p.SetState(522)
				p.InitializationList()
			}
			{
				p.SetState(523)
				p.Match(goscriptParserT__3)
			}

			p.SetState(529)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(530)
			p.Match(goscriptParserT__5)
		}

	case 3:
		localctx = NewInitMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(532)
			p.Match(goscriptParserT__4)
		}
		{
			p.SetState(533)
			p.InitializationList()
		}
		{
			p.SetState(534)
			p.Match(goscriptParserT__9)
		}
		{
			p.SetState(535)
			p.InitializationList()
		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(536)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(537)
				p.InitializationList()
			}
			{
				p.SetState(538)
				p.Match(goscriptParserT__9)
			}
			{
				p.SetState(539)
				p.InitializationList()
			}

			p.SetState(545)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(546)
			p.Match(goscriptParserT__5)
		}

	case 4:
		localctx = NewInitConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(548)
			p.expr(0)
		}

	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CopyFrom(ctx *ConstantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstantNilContext struct {
	*ConstantContext
}

func NewConstantNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantNilContext {
	var p = new(ConstantNilContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *ConstantNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantNilContext) NULL() antlr.TerminalNode {
	return s.GetToken(goscriptParserNULL, 0)
}

func (s *ConstantNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstantNil(s)
	}
}

func (s *ConstantNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstantNil(s)
	}
}

type ConstantIntContext struct {
	*ConstantContext
}

func NewConstantIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantIntContext {
	var p = new(ConstantIntContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *ConstantIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantIntContext) INT() antlr.TerminalNode {
	return s.GetToken(goscriptParserINT, 0)
}

func (s *ConstantIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstantInt(s)
	}
}

func (s *ConstantIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstantInt(s)
	}
}

type ConstantStringContext struct {
	*ConstantContext
}

func NewConstantStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantStringContext {
	var p = new(ConstantStringContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *ConstantStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(goscriptParserSTRING, 0)
}

func (s *ConstantStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstantString(s)
	}
}

func (s *ConstantStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstantString(s)
	}
}

type ConstantFloatContext struct {
	*ConstantContext
}

func NewConstantFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantFloatContext {
	var p = new(ConstantFloatContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *ConstantFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(goscriptParserFLOAT, 0)
}

func (s *ConstantFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstantFloat(s)
	}
}

func (s *ConstantFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstantFloat(s)
	}
}

type ConstantBoolContext struct {
	*ConstantContext
}

func NewConstantBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantBoolContext {
	var p = new(ConstantBoolContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *ConstantBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantBoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(goscriptParserBOOL, 0)
}

func (s *ConstantBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstantBool(s)
	}
}

func (s *ConstantBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstantBool(s)
	}
}

func (p *goscriptParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, goscriptParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(556)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserINT:
		localctx = NewConstantIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Match(goscriptParserINT)
		}

	case goscriptParserFLOAT:
		localctx = NewConstantFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Match(goscriptParserFLOAT)
		}

	case goscriptParserBOOL:
		localctx = NewConstantBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
			p.Match(goscriptParserBOOL)
		}

	case goscriptParserNULL:
		localctx = NewConstantNilContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(554)
			p.Match(goscriptParserNULL)
		}

	case goscriptParserSTRING:
		localctx = NewConstantStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(555)
			p.Match(goscriptParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *FunctionsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunctionsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionsContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionsContext) DOT() antlr.TerminalNode {
	return s.GetToken(goscriptParserDOT, 0)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *goscriptParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, goscriptParserRULE_functions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(558)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(559)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(560)
			p.expr(0)
		}
		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(561)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(562)
				p.expr(0)
			}

			p.SetState(567)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(568)
			p.Match(goscriptParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(570)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(571)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(572)
			p.Match(goscriptParserT__3)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(573)
			p.variable(0)
		}
		{
			p.SetState(574)
			p.Match(goscriptParserDOT)
		}
		{
			p.SetState(575)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(576)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(577)
			p.Match(goscriptParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(579)
			p.variable(0)
		}
		{
			p.SetState(580)
			p.Match(goscriptParserDOT)
		}
		{
			p.SetState(581)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(582)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(583)
			p.expr(0)
		}
		p.SetState(588)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(584)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(585)
				p.expr(0)
			}

			p.SetState(590)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(591)
			p.Match(goscriptParserT__3)
		}

	}

	return localctx
}

// IConstructorContext is an interface to support dynamic dispatch.
type IConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructorContext differentiates from other interfaces.
	IsConstructorContext()
}

type ConstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructorContext() *ConstructorContext {
	var p = new(ConstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_constructor
	return p
}

func (*ConstructorContext) IsConstructorContext() {}

func NewConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructorContext {
	var p = new(ConstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_constructor

	return p
}

func (s *ConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructorContext) NEW() antlr.TerminalNode {
	return s.GetToken(goscriptParserNEW, 0)
}

func (s *ConstructorContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *ConstructorContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ConstructorContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterConstructor(s)
	}
}

func (s *ConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitConstructor(s)
	}
}

func (p *goscriptParser) Constructor() (localctx IConstructorContext) {
	localctx = NewConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, goscriptParserRULE_constructor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(595)
			p.Match(goscriptParserNEW)
		}
		{
			p.SetState(596)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(597)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(598)
			p.Match(goscriptParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(599)
			p.Match(goscriptParserNEW)
		}
		{
			p.SetState(600)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(601)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(602)
			p.expr(0)
		}
		p.SetState(607)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(603)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(604)
				p.expr(0)
			}

			p.SetState(609)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(610)
			p.Match(goscriptParserT__3)
		}

	}

	return localctx
}

// IVardefContext is an interface to support dynamic dispatch.
type IVardefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVardefContext differentiates from other interfaces.
	IsVardefContext()
}

type VardefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVardefContext() *VardefContext {
	var p = new(VardefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_vardef
	return p
}

func (*VardefContext) IsVardefContext() {}

func NewVardefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VardefContext {
	var p = new(VardefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_vardef

	return p
}

func (s *VardefContext) GetParser() antlr.Parser { return s.parser }

func (s *VardefContext) VAR() antlr.TerminalNode {
	return s.GetToken(goscriptParserVAR, 0)
}

func (s *VardefContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *VardefContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *VardefContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *VardefContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VardefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VardefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VardefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterVardef(s)
	}
}

func (s *VardefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitVardef(s)
	}
}

func (p *goscriptParser) Vardef() (localctx IVardefContext) {
	localctx = NewVardefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, goscriptParserRULE_vardef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(614)
			p.Match(goscriptParserVAR)
		}
		{
			p.SetState(615)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(616)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(617)
			p.Match(goscriptParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(618)
			p.Match(goscriptParserVAR)
		}
		{
			p.SetState(619)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(620)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(621)
			p.Match(goscriptParserASSIGN)
		}
		{
			p.SetState(622)
			p.expr(0)
		}
		{
			p.SetState(623)
			p.Match(goscriptParserT__0)
		}

	}

	return localctx
}

// ILocaldefContext is an interface to support dynamic dispatch.
type ILocaldefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocaldefContext differentiates from other interfaces.
	IsLocaldefContext()
}

type LocaldefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocaldefContext() *LocaldefContext {
	var p = new(LocaldefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goscriptParserRULE_localdef
	return p
}

func (*LocaldefContext) IsLocaldefContext() {}

func NewLocaldefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocaldefContext {
	var p = new(LocaldefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goscriptParserRULE_localdef

	return p
}

func (s *LocaldefContext) GetParser() antlr.Parser { return s.parser }

func (s *LocaldefContext) CopyFrom(ctx *LocaldefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LocaldefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocaldefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LocalAssignContext struct {
	*LocaldefContext
}

func NewLocalAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAssignContext {
	var p = new(LocalAssignContext)

	p.LocaldefContext = NewEmptyLocaldefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LocaldefContext))

	return p
}

func (s *LocalAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalAssignContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserLOCAL, 0)
}

func (s *LocalAssignContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *LocalAssignContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *LocalAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(goscriptParserASSIGN, 0)
}

func (s *LocalAssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LocalAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLocalAssign(s)
	}
}

func (s *LocalAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLocalAssign(s)
	}
}

type LocalContext struct {
	*LocaldefContext
}

func NewLocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalContext {
	var p = new(LocalContext)

	p.LocaldefContext = NewEmptyLocaldefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LocaldefContext))

	return p
}

func (s *LocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(goscriptParserLOCAL, 0)
}

func (s *LocalContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *LocalContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *LocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterLocal(s)
	}
}

func (s *LocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitLocal(s)
	}
}

func (p *goscriptParser) Localdef() (localctx ILocaldefContext) {
	localctx = NewLocaldefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, goscriptParserRULE_localdef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(627)
			p.Match(goscriptParserLOCAL)
		}
		{
			p.SetState(628)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(629)
			p.Match(goscriptParserNAME)
		}

	case 2:
		localctx = NewLocalAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(630)
			p.Match(goscriptParserLOCAL)
		}
		{
			p.SetState(631)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(632)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(633)
			p.Match(goscriptParserASSIGN)
		}
		{
			p.SetState(634)
			p.expr(0)
		}

	}

	return localctx
}

func (p *goscriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	case 19:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *goscriptParser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *goscriptParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
