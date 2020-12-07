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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 508,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 6, 2, 52, 10, 2, 13, 2, 14, 2, 53, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 7, 4, 71, 10, 4, 12, 4, 14, 4, 74, 11, 4, 3, 4, 3,
	4, 5, 4, 78, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 87,
	10, 4, 3, 4, 5, 4, 90, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 105, 10, 6, 12, 6, 14, 6, 108,
	11, 6, 3, 6, 5, 6, 111, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7,
	119, 10, 7, 12, 7, 14, 7, 122, 11, 7, 3, 7, 3, 7, 5, 7, 126, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 135, 10, 8, 12, 8, 14, 8,
	138, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 146, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 156, 10, 10, 5,
	10, 158, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 6, 10, 170, 10, 10, 13, 10, 14, 10, 171, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 214, 10, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 7, 12, 221, 10, 12, 12, 12, 14, 12, 224, 11, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 232, 10, 13, 12, 13, 14,
	13, 235, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 243,
	10, 13, 12, 13, 14, 13, 246, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	252, 10, 13, 12, 13, 14, 13, 255, 11, 13, 5, 13, 257, 10, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 5, 15, 264, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 7, 15, 284, 10, 15, 12, 15, 14, 15, 287, 11, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 297, 10,
	15, 12, 15, 14, 15, 300, 11, 15, 3, 15, 3, 15, 3, 15, 7, 15, 305, 10, 15,
	12, 15, 14, 15, 308, 11, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 5, 17, 328, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 355,
	10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 375,
	10, 18, 12, 18, 14, 18, 378, 11, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 7, 20, 386, 10, 20, 12, 20, 14, 20, 389, 11, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	7, 20, 404, 10, 20, 12, 20, 14, 20, 407, 11, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 420, 10, 20,
	12, 20, 14, 20, 423, 11, 20, 3, 20, 3, 20, 3, 20, 5, 20, 428, 10, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 435, 10, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 7, 22, 442, 10, 22, 12, 22, 14, 22, 445, 11, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 7, 22, 462, 10, 22, 12, 22, 14, 22, 465, 11, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 472, 10, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 484, 10, 23, 12,
	23, 14, 23, 487, 11, 23, 3, 23, 3, 23, 5, 23, 491, 10, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 506, 10, 25, 3, 25, 2, 4, 28, 34, 26, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2,
	8, 3, 2, 38, 39, 4, 2, 37, 37, 49, 49, 3, 2, 50, 54, 3, 2, 33, 35, 3, 2,
	36, 37, 3, 2, 40, 46, 2, 562, 2, 51, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 6,
	89, 3, 2, 2, 2, 8, 91, 3, 2, 2, 2, 10, 110, 3, 2, 2, 2, 12, 125, 3, 2,
	2, 2, 14, 127, 3, 2, 2, 2, 16, 145, 3, 2, 2, 2, 18, 213, 3, 2, 2, 2, 20,
	215, 3, 2, 2, 2, 22, 217, 3, 2, 2, 2, 24, 256, 3, 2, 2, 2, 26, 258, 3,
	2, 2, 2, 28, 263, 3, 2, 2, 2, 30, 309, 3, 2, 2, 2, 32, 327, 3, 2, 2, 2,
	34, 354, 3, 2, 2, 2, 36, 379, 3, 2, 2, 2, 38, 427, 3, 2, 2, 2, 40, 434,
	3, 2, 2, 2, 42, 471, 3, 2, 2, 2, 44, 490, 3, 2, 2, 2, 46, 492, 3, 2, 2,
	2, 48, 505, 3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 53,
	3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 3, 3, 2, 2, 2,
	55, 63, 5, 6, 4, 2, 56, 57, 5, 10, 6, 2, 57, 58, 7, 3, 2, 2, 58, 63, 3,
	2, 2, 2, 59, 63, 5, 46, 24, 2, 60, 63, 5, 14, 8, 2, 61, 63, 5, 16, 9, 2,
	62, 55, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 59, 3, 2, 2, 2, 62, 60, 3,
	2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 5, 3, 2, 2, 2, 64, 65, 7, 28, 2, 2, 65,
	66, 7, 58, 2, 2, 66, 67, 7, 4, 2, 2, 67, 72, 5, 8, 5, 2, 68, 69, 7, 5,
	2, 2, 69, 71, 5, 8, 5, 2, 70, 68, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	75, 77, 7, 6, 2, 2, 76, 78, 7, 58, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 5, 22, 12, 2, 80, 90, 3, 2, 2, 2,
	81, 82, 7, 28, 2, 2, 82, 83, 7, 58, 2, 2, 83, 84, 7, 4, 2, 2, 84, 86, 7,
	6, 2, 2, 85, 87, 7, 58, 2, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87,
	88, 3, 2, 2, 2, 88, 90, 5, 22, 12, 2, 89, 64, 3, 2, 2, 2, 89, 81, 3, 2,
	2, 2, 90, 7, 3, 2, 2, 2, 91, 92, 7, 58, 2, 2, 92, 93, 7, 58, 2, 2, 93,
	9, 3, 2, 2, 2, 94, 95, 7, 29, 2, 2, 95, 96, 7, 58, 2, 2, 96, 111, 5, 12,
	7, 2, 97, 98, 7, 29, 2, 2, 98, 99, 7, 58, 2, 2, 99, 106, 7, 7, 2, 2, 100,
	101, 7, 58, 2, 2, 101, 102, 5, 12, 7, 2, 102, 103, 7, 5, 2, 2, 103, 105,
	3, 2, 2, 2, 104, 100, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2,
	109, 111, 7, 8, 2, 2, 110, 94, 3, 2, 2, 2, 110, 97, 3, 2, 2, 2, 111, 11,
	3, 2, 2, 2, 112, 126, 7, 58, 2, 2, 113, 114, 7, 58, 2, 2, 114, 115, 7,
	45, 2, 2, 115, 120, 5, 12, 7, 2, 116, 117, 7, 5, 2, 2, 117, 119, 5, 12,
	7, 2, 118, 116, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2,
	120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123,
	124, 7, 42, 2, 2, 124, 126, 3, 2, 2, 2, 125, 112, 3, 2, 2, 2, 125, 113,
	3, 2, 2, 2, 126, 13, 3, 2, 2, 2, 127, 128, 7, 29, 2, 2, 128, 129, 7, 9,
	2, 2, 129, 130, 7, 58, 2, 2, 130, 136, 7, 7, 2, 2, 131, 132, 7, 58, 2,
	2, 132, 133, 7, 55, 2, 2, 133, 135, 7, 5, 2, 2, 134, 131, 3, 2, 2, 2, 135,
	138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 139,
	3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 140, 7, 8, 2, 2, 140, 15, 3, 2,
	2, 2, 141, 146, 5, 18, 10, 2, 142, 143, 5, 24, 13, 2, 143, 144, 7, 3, 2,
	2, 144, 146, 3, 2, 2, 2, 145, 141, 3, 2, 2, 2, 145, 142, 3, 2, 2, 2, 146,
	17, 3, 2, 2, 2, 147, 148, 7, 21, 2, 2, 148, 149, 7, 4, 2, 2, 149, 150,
	5, 34, 18, 2, 150, 151, 7, 6, 2, 2, 151, 157, 5, 22, 12, 2, 152, 155, 7,
	22, 2, 2, 153, 156, 5, 22, 12, 2, 154, 156, 5, 18, 10, 2, 155, 153, 3,
	2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157, 152, 3, 2, 2,
	2, 157, 158, 3, 2, 2, 2, 158, 214, 3, 2, 2, 2, 159, 160, 7, 23, 2, 2, 160,
	161, 7, 4, 2, 2, 161, 162, 5, 34, 18, 2, 162, 163, 7, 6, 2, 2, 163, 169,
	7, 7, 2, 2, 164, 165, 7, 24, 2, 2, 165, 166, 5, 40, 21, 2, 166, 167, 7,
	10, 2, 2, 167, 168, 5, 22, 12, 2, 168, 170, 3, 2, 2, 2, 169, 164, 3, 2,
	2, 2, 170, 171, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2,
	172, 173, 3, 2, 2, 2, 173, 174, 7, 8, 2, 2, 174, 214, 3, 2, 2, 2, 175,
	176, 7, 18, 2, 2, 176, 177, 7, 4, 2, 2, 177, 178, 7, 58, 2, 2, 178, 179,
	7, 11, 2, 2, 179, 180, 5, 20, 11, 2, 180, 181, 7, 6, 2, 2, 181, 182, 5,
	22, 12, 2, 182, 214, 3, 2, 2, 2, 183, 184, 7, 18, 2, 2, 184, 185, 7, 4,
	2, 2, 185, 186, 7, 58, 2, 2, 186, 187, 7, 5, 2, 2, 187, 188, 7, 58, 2,
	2, 188, 189, 7, 11, 2, 2, 189, 190, 5, 20, 11, 2, 190, 191, 7, 6, 2, 2,
	191, 192, 5, 22, 12, 2, 192, 214, 3, 2, 2, 2, 193, 194, 7, 18, 2, 2, 194,
	195, 7, 4, 2, 2, 195, 196, 5, 34, 18, 2, 196, 197, 7, 3, 2, 2, 197, 198,
	5, 34, 18, 2, 198, 199, 7, 3, 2, 2, 199, 200, 5, 34, 18, 2, 200, 201, 7,
	6, 2, 2, 201, 202, 5, 22, 12, 2, 202, 214, 3, 2, 2, 2, 203, 204, 7, 19,
	2, 2, 204, 214, 7, 3, 2, 2, 205, 206, 7, 20, 2, 2, 206, 214, 7, 3, 2, 2,
	207, 208, 7, 25, 2, 2, 208, 214, 7, 3, 2, 2, 209, 210, 7, 25, 2, 2, 210,
	211, 5, 34, 18, 2, 211, 212, 7, 3, 2, 2, 212, 214, 3, 2, 2, 2, 213, 147,
	3, 2, 2, 2, 213, 159, 3, 2, 2, 2, 213, 175, 3, 2, 2, 2, 213, 183, 3, 2,
	2, 2, 213, 193, 3, 2, 2, 2, 213, 203, 3, 2, 2, 2, 213, 205, 3, 2, 2, 2,
	213, 207, 3, 2, 2, 2, 213, 209, 3, 2, 2, 2, 214, 19, 3, 2, 2, 2, 215, 216,
	5, 34, 18, 2, 216, 21, 3, 2, 2, 2, 217, 222, 7, 7, 2, 2, 218, 221, 5, 16,
	9, 2, 219, 221, 5, 46, 24, 2, 220, 218, 3, 2, 2, 2, 220, 219, 3, 2, 2,
	2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223,
	225, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 226, 7, 8, 2, 2, 226, 23, 3,
	2, 2, 2, 227, 257, 5, 34, 18, 2, 228, 233, 5, 26, 14, 2, 229, 230, 7, 5,
	2, 2, 230, 232, 5, 26, 14, 2, 231, 229, 3, 2, 2, 2, 232, 235, 3, 2, 2,
	2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 236, 3, 2, 2, 2, 235,
	233, 3, 2, 2, 2, 236, 237, 7, 50, 2, 2, 237, 238, 5, 42, 22, 2, 238, 257,
	3, 2, 2, 2, 239, 244, 5, 26, 14, 2, 240, 241, 7, 5, 2, 2, 241, 243, 5,
	26, 14, 2, 242, 240, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2,
	2, 2, 244, 245, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2,
	247, 248, 7, 50, 2, 2, 248, 253, 5, 34, 18, 2, 249, 250, 7, 5, 2, 2, 250,
	252, 5, 34, 18, 2, 251, 249, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251,
	3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2,
	2, 2, 256, 227, 3, 2, 2, 2, 256, 228, 3, 2, 2, 2, 256, 239, 3, 2, 2, 2,
	257, 25, 3, 2, 2, 2, 258, 259, 5, 28, 15, 2, 259, 27, 3, 2, 2, 2, 260,
	261, 8, 15, 1, 2, 261, 264, 7, 58, 2, 2, 262, 264, 7, 16, 2, 2, 263, 260,
	3, 2, 2, 2, 263, 262, 3, 2, 2, 2, 264, 306, 3, 2, 2, 2, 265, 266, 12, 9,
	2, 2, 266, 267, 7, 59, 2, 2, 267, 305, 7, 58, 2, 2, 268, 269, 12, 8, 2,
	2, 269, 270, 7, 12, 2, 2, 270, 271, 5, 30, 16, 2, 271, 272, 7, 13, 2, 2,
	272, 305, 3, 2, 2, 2, 273, 274, 12, 7, 2, 2, 274, 275, 7, 14, 2, 2, 275,
	276, 5, 34, 18, 2, 276, 277, 7, 15, 2, 2, 277, 305, 3, 2, 2, 2, 278, 279,
	12, 6, 2, 2, 279, 280, 7, 14, 2, 2, 280, 285, 5, 32, 17, 2, 281, 282, 7,
	5, 2, 2, 282, 284, 5, 32, 17, 2, 283, 281, 3, 2, 2, 2, 284, 287, 3, 2,
	2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 288, 3, 2, 2, 2,
	287, 285, 3, 2, 2, 2, 288, 289, 7, 15, 2, 2, 289, 305, 3, 2, 2, 2, 290,
	291, 12, 5, 2, 2, 291, 292, 7, 14, 2, 2, 292, 293, 7, 14, 2, 2, 293, 298,
	5, 34, 18, 2, 294, 295, 7, 5, 2, 2, 295, 297, 5, 34, 18, 2, 296, 294, 3,
	2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2,
	2, 299, 301, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 302, 7, 15, 2, 2, 302,
	303, 7, 15, 2, 2, 303, 305, 3, 2, 2, 2, 304, 265, 3, 2, 2, 2, 304, 268,
	3, 2, 2, 2, 304, 273, 3, 2, 2, 2, 304, 278, 3, 2, 2, 2, 304, 290, 3, 2,
	2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2,
	307, 29, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310, 5, 34, 18, 2, 310,
	31, 3, 2, 2, 2, 311, 312, 5, 34, 18, 2, 312, 313, 7, 10, 2, 2, 313, 314,
	5, 34, 18, 2, 314, 315, 7, 10, 2, 2, 315, 316, 5, 34, 18, 2, 316, 328,
	3, 2, 2, 2, 317, 318, 5, 34, 18, 2, 318, 319, 7, 10, 2, 2, 319, 320, 5,
	34, 18, 2, 320, 328, 3, 2, 2, 2, 321, 322, 5, 34, 18, 2, 322, 323, 7, 10,
	2, 2, 323, 328, 3, 2, 2, 2, 324, 328, 5, 34, 18, 2, 325, 326, 7, 10, 2,
	2, 326, 328, 5, 34, 18, 2, 327, 311, 3, 2, 2, 2, 327, 317, 3, 2, 2, 2,
	327, 321, 3, 2, 2, 2, 327, 324, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 328,
	33, 3, 2, 2, 2, 329, 330, 8, 18, 1, 2, 330, 331, 7, 4, 2, 2, 331, 332,
	5, 34, 18, 2, 332, 333, 7, 6, 2, 2, 333, 355, 3, 2, 2, 2, 334, 335, 9,
	2, 2, 2, 335, 355, 5, 28, 15, 2, 336, 337, 9, 3, 2, 2, 337, 355, 5, 34,
	18, 17, 338, 339, 5, 28, 15, 2, 339, 340, 9, 2, 2, 2, 340, 355, 3, 2, 2,
	2, 341, 342, 5, 26, 14, 2, 342, 343, 9, 4, 2, 2, 343, 344, 5, 34, 18, 9,
	344, 355, 3, 2, 2, 2, 345, 346, 5, 26, 14, 2, 346, 347, 7, 50, 2, 2, 347,
	348, 5, 36, 19, 2, 348, 355, 3, 2, 2, 2, 349, 355, 5, 40, 21, 2, 350, 355,
	5, 28, 15, 2, 351, 355, 5, 42, 22, 2, 352, 355, 5, 44, 23, 2, 353, 355,
	5, 48, 25, 2, 354, 329, 3, 2, 2, 2, 354, 334, 3, 2, 2, 2, 354, 336, 3,
	2, 2, 2, 354, 338, 3, 2, 2, 2, 354, 341, 3, 2, 2, 2, 354, 345, 3, 2, 2,
	2, 354, 349, 3, 2, 2, 2, 354, 350, 3, 2, 2, 2, 354, 351, 3, 2, 2, 2, 354,
	352, 3, 2, 2, 2, 354, 353, 3, 2, 2, 2, 355, 376, 3, 2, 2, 2, 356, 357,
	12, 15, 2, 2, 357, 358, 7, 32, 2, 2, 358, 375, 5, 34, 18, 15, 359, 360,
	12, 14, 2, 2, 360, 361, 9, 5, 2, 2, 361, 375, 5, 34, 18, 15, 362, 363,
	12, 13, 2, 2, 363, 364, 9, 6, 2, 2, 364, 375, 5, 34, 18, 14, 365, 366,
	12, 12, 2, 2, 366, 367, 9, 7, 2, 2, 367, 375, 5, 34, 18, 13, 368, 369,
	12, 11, 2, 2, 369, 370, 7, 47, 2, 2, 370, 375, 5, 34, 18, 12, 371, 372,
	12, 10, 2, 2, 372, 373, 7, 48, 2, 2, 373, 375, 5, 34, 18, 11, 374, 356,
	3, 2, 2, 2, 374, 359, 3, 2, 2, 2, 374, 362, 3, 2, 2, 2, 374, 365, 3, 2,
	2, 2, 374, 368, 3, 2, 2, 2, 374, 371, 3, 2, 2, 2, 375, 378, 3, 2, 2, 2,
	376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 35, 3, 2, 2, 2, 378, 376,
	3, 2, 2, 2, 379, 380, 5, 38, 20, 2, 380, 37, 3, 2, 2, 2, 381, 382, 7, 14,
	2, 2, 382, 387, 5, 38, 20, 2, 383, 384, 7, 5, 2, 2, 384, 386, 5, 38, 20,
	2, 385, 383, 3, 2, 2, 2, 386, 389, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387,
	388, 3, 2, 2, 2, 388, 390, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 390, 391,
	7, 15, 2, 2, 391, 428, 3, 2, 2, 2, 392, 393, 7, 7, 2, 2, 393, 394, 7, 58,
	2, 2, 394, 395, 7, 4, 2, 2, 395, 396, 5, 38, 20, 2, 396, 405, 7, 6, 2,
	2, 397, 398, 7, 5, 2, 2, 398, 399, 7, 58, 2, 2, 399, 400, 7, 4, 2, 2, 400,
	401, 5, 38, 20, 2, 401, 402, 7, 6, 2, 2, 402, 404, 3, 2, 2, 2, 403, 397,
	3, 2, 2, 2, 404, 407, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 405, 406, 3, 2,
	2, 2, 406, 408, 3, 2, 2, 2, 407, 405, 3, 2, 2, 2, 408, 409, 7, 8, 2, 2,
	409, 428, 3, 2, 2, 2, 410, 411, 7, 7, 2, 2, 411, 412, 5, 38, 20, 2, 412,
	413, 7, 10, 2, 2, 413, 421, 5, 38, 20, 2, 414, 415, 7, 5, 2, 2, 415, 416,
	5, 38, 20, 2, 416, 417, 7, 10, 2, 2, 417, 418, 5, 38, 20, 2, 418, 420,
	3, 2, 2, 2, 419, 414, 3, 2, 2, 2, 420, 423, 3, 2, 2, 2, 421, 419, 3, 2,
	2, 2, 421, 422, 3, 2, 2, 2, 422, 424, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2,
	424, 425, 7, 8, 2, 2, 425, 428, 3, 2, 2, 2, 426, 428, 5, 34, 18, 2, 427,
	381, 3, 2, 2, 2, 427, 392, 3, 2, 2, 2, 427, 410, 3, 2, 2, 2, 427, 426,
	3, 2, 2, 2, 428, 39, 3, 2, 2, 2, 429, 435, 7, 55, 2, 2, 430, 435, 7, 56,
	2, 2, 431, 435, 7, 30, 2, 2, 432, 435, 7, 31, 2, 2, 433, 435, 7, 57, 2,
	2, 434, 429, 3, 2, 2, 2, 434, 430, 3, 2, 2, 2, 434, 431, 3, 2, 2, 2, 434,
	432, 3, 2, 2, 2, 434, 433, 3, 2, 2, 2, 435, 41, 3, 2, 2, 2, 436, 437, 7,
	58, 2, 2, 437, 438, 7, 4, 2, 2, 438, 443, 5, 34, 18, 2, 439, 440, 7, 5,
	2, 2, 440, 442, 5, 34, 18, 2, 441, 439, 3, 2, 2, 2, 442, 445, 3, 2, 2,
	2, 443, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 446, 3, 2, 2, 2, 445,
	443, 3, 2, 2, 2, 446, 447, 7, 6, 2, 2, 447, 472, 3, 2, 2, 2, 448, 449,
	5, 28, 15, 2, 449, 450, 7, 59, 2, 2, 450, 451, 7, 58, 2, 2, 451, 452, 7,
	4, 2, 2, 452, 453, 7, 6, 2, 2, 453, 472, 3, 2, 2, 2, 454, 455, 5, 28, 15,
	2, 455, 456, 7, 59, 2, 2, 456, 457, 7, 58, 2, 2, 457, 458, 7, 4, 2, 2,
	458, 463, 5, 34, 18, 2, 459, 460, 7, 5, 2, 2, 460, 462, 5, 34, 18, 2, 461,
	459, 3, 2, 2, 2, 462, 465, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2, 463, 464,
	3, 2, 2, 2, 464, 466, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 466, 467, 7, 6,
	2, 2, 467, 472, 3, 2, 2, 2, 468, 469, 7, 58, 2, 2, 469, 470, 7, 4, 2, 2,
	470, 472, 7, 6, 2, 2, 471, 436, 3, 2, 2, 2, 471, 448, 3, 2, 2, 2, 471,
	454, 3, 2, 2, 2, 471, 468, 3, 2, 2, 2, 472, 43, 3, 2, 2, 2, 473, 474, 7,
	17, 2, 2, 474, 475, 7, 58, 2, 2, 475, 476, 7, 4, 2, 2, 476, 491, 7, 6,
	2, 2, 477, 478, 7, 17, 2, 2, 478, 479, 7, 58, 2, 2, 479, 480, 7, 4, 2,
	2, 480, 485, 5, 34, 18, 2, 481, 482, 7, 5, 2, 2, 482, 484, 5, 34, 18, 2,
	483, 481, 3, 2, 2, 2, 484, 487, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485,
	486, 3, 2, 2, 2, 486, 488, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 488, 489,
	7, 6, 2, 2, 489, 491, 3, 2, 2, 2, 490, 473, 3, 2, 2, 2, 490, 477, 3, 2,
	2, 2, 491, 45, 3, 2, 2, 2, 492, 493, 7, 26, 2, 2, 493, 494, 7, 58, 2, 2,
	494, 495, 7, 58, 2, 2, 495, 496, 7, 3, 2, 2, 496, 47, 3, 2, 2, 2, 497,
	498, 7, 27, 2, 2, 498, 499, 7, 58, 2, 2, 499, 506, 7, 58, 2, 2, 500, 501,
	7, 27, 2, 2, 501, 502, 7, 58, 2, 2, 502, 503, 7, 58, 2, 2, 503, 504, 7,
	50, 2, 2, 504, 506, 5, 34, 18, 2, 505, 497, 3, 2, 2, 2, 505, 500, 3, 2,
	2, 2, 506, 49, 3, 2, 2, 2, 44, 53, 62, 72, 77, 86, 89, 106, 110, 120, 125,
	136, 145, 155, 157, 171, 213, 220, 222, 233, 244, 253, 256, 263, 285, 298,
	304, 306, 327, 354, 374, 376, 387, 405, 421, 427, 434, 443, 463, 471, 485,
	490, 505,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "','", "')'", "'{'", "'}'", "'enum'", "':'", "'in'",
	"'[?('", "')]'", "'['", "']'", "'@'", "'new'", "'for'", "'break'", "'continue'",
	"'if'", "'else'", "'switch'", "'case'", "'return'", "'var'", "'local'",
	"'func'", "'type'", "", "'nil'", "'**'", "'*'", "'/'", "'%'", "'+'", "'-'",
	"'++'", "'--'", "'=='", "'!='", "'>'", "'>='", "'<='", "'<'", "'=~'", "'&&'",
	"'||'", "'!'", "'='", "'+='", "'-='", "'*='", "'/='", "", "", "", "", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NEW", "FOR",
	"BREAK", "CONTINUE", "IF", "ELSE", "SWITCH", "CASE", "RETURN", "VAR", "LOCAL",
	"FUNCTION", "TYPEDEF", "BOOL", "NULL", "POW", "MUL", "DIV", "MOD", "ADD",
	"SUB", "UNARYADD", "UNARYSUB", "EQ", "INEQ", "GT", "GE", "LE", "LT", "REGEX",
	"AND", "OR", "NOT", "ASSIGN", "ADDEQUAL", "SUBEQUAL", "MULEQUAL", "DIVEQUAL",
	"INT", "FLOAT", "STRING", "NAME", "DOT", "WHITESPACE",
}

var ruleNames = []string{
	"program", "statement", "functiondef", "param", "typedef", "typename",
	"enumdef", "execution", "control", "collection", "block", "line", "lhs",
	"variable", "filter", "indexs", "expr", "initializationListBegin", "initializationList",
	"constant", "functions", "constructor", "vardef", "localdef",
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
	goscriptParserNEW        = 15
	goscriptParserFOR        = 16
	goscriptParserBREAK      = 17
	goscriptParserCONTINUE   = 18
	goscriptParserIF         = 19
	goscriptParserELSE       = 20
	goscriptParserSWITCH     = 21
	goscriptParserCASE       = 22
	goscriptParserRETURN     = 23
	goscriptParserVAR        = 24
	goscriptParserLOCAL      = 25
	goscriptParserFUNCTION   = 26
	goscriptParserTYPEDEF    = 27
	goscriptParserBOOL       = 28
	goscriptParserNULL       = 29
	goscriptParserPOW        = 30
	goscriptParserMUL        = 31
	goscriptParserDIV        = 32
	goscriptParserMOD        = 33
	goscriptParserADD        = 34
	goscriptParserSUB        = 35
	goscriptParserUNARYADD   = 36
	goscriptParserUNARYSUB   = 37
	goscriptParserEQ         = 38
	goscriptParserINEQ       = 39
	goscriptParserGT         = 40
	goscriptParserGE         = 41
	goscriptParserLE         = 42
	goscriptParserLT         = 43
	goscriptParserREGEX      = 44
	goscriptParserAND        = 45
	goscriptParserOR         = 46
	goscriptParserNOT        = 47
	goscriptParserASSIGN     = 48
	goscriptParserADDEQUAL   = 49
	goscriptParserSUBEQUAL   = 50
	goscriptParserMULEQUAL   = 51
	goscriptParserDIVEQUAL   = 52
	goscriptParserINT        = 53
	goscriptParserFLOAT      = 54
	goscriptParserSTRING     = 55
	goscriptParserNAME       = 56
	goscriptParserDOT        = 57
	goscriptParserWHITESPACE = 58
)

// goscriptParser rules.
const (
	goscriptParserRULE_program                 = 0
	goscriptParserRULE_statement               = 1
	goscriptParserRULE_functiondef             = 2
	goscriptParserRULE_param                   = 3
	goscriptParserRULE_typedef                 = 4
	goscriptParserRULE_typename                = 5
	goscriptParserRULE_enumdef                 = 6
	goscriptParserRULE_execution               = 7
	goscriptParserRULE_control                 = 8
	goscriptParserRULE_collection              = 9
	goscriptParserRULE_block                   = 10
	goscriptParserRULE_line                    = 11
	goscriptParserRULE_lhs                     = 12
	goscriptParserRULE_variable                = 13
	goscriptParserRULE_filter                  = 14
	goscriptParserRULE_indexs                  = 15
	goscriptParserRULE_expr                    = 16
	goscriptParserRULE_initializationListBegin = 17
	goscriptParserRULE_initializationList      = 18
	goscriptParserRULE_constant                = 19
	goscriptParserRULE_functions               = 20
	goscriptParserRULE_constructor             = 21
	goscriptParserRULE_vardef                  = 22
	goscriptParserRULE_localdef                = 23
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<goscriptParserT__1)|(1<<goscriptParserT__13)|(1<<goscriptParserNEW)|(1<<goscriptParserFOR)|(1<<goscriptParserBREAK)|(1<<goscriptParserCONTINUE)|(1<<goscriptParserIF)|(1<<goscriptParserSWITCH)|(1<<goscriptParserRETURN)|(1<<goscriptParserVAR)|(1<<goscriptParserLOCAL)|(1<<goscriptParserFUNCTION)|(1<<goscriptParserTYPEDEF)|(1<<goscriptParserBOOL)|(1<<goscriptParserNULL))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(goscriptParserSUB-35))|(1<<(goscriptParserUNARYADD-35))|(1<<(goscriptParserUNARYSUB-35))|(1<<(goscriptParserNOT-35))|(1<<(goscriptParserINT-35))|(1<<(goscriptParserFLOAT-35))|(1<<(goscriptParserSTRING-35))|(1<<(goscriptParserNAME-35)))) != 0) {
		{
			p.SetState(48)
			p.Statement()
		}

		p.SetState(51)
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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Functiondef()
		}

	case 2:
		localctx = NewTypeDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Typedef()
		}
		{
			p.SetState(55)
			p.Match(goscriptParserT__0)
		}

	case 3:
		localctx = NewGlobalDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Vardef()
		}

	case 4:
		localctx = NewEnumDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Enumdef()
		}

	case 5:
		localctx = NewRunContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
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

func (s *FunctiondefContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(goscriptParserFUNCTION, 0)
}

func (s *FunctiondefContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *FunctiondefContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
}

func (s *FunctiondefContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *FunctiondefContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *FunctiondefContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitFunctiondef(s)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(63)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(64)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(65)
			p.Param()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(66)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(67)
				p.Param()
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(goscriptParserT__3)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserNAME {
			{
				p.SetState(74)
				p.Match(goscriptParserNAME)
			}

		}
		{
			p.SetState(77)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(goscriptParserFUNCTION)
		}
		{
			p.SetState(80)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(81)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(82)
			p.Match(goscriptParserT__3)
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserNAME {
			{
				p.SetState(83)
				p.Match(goscriptParserNAME)
			}

		}
		{
			p.SetState(86)
			p.Block()
		}

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

func (s *ParamContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(goscriptParserNAME)
}

func (s *ParamContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, i)
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
	p.EnterRule(localctx, 6, goscriptParserRULE_param)

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
		p.SetState(89)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(90)
		p.Match(goscriptParserNAME)
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
	p.EnterRule(localctx, 8, goscriptParserRULE_typedef)
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDefAliasContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(goscriptParserTYPEDEF)
		}
		{
			p.SetState(93)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(94)
			p.Typename()
		}

	case 2:
		localctx = NewTypeDefComplexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(goscriptParserTYPEDEF)
		}
		{
			p.SetState(96)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(97)
			p.Match(goscriptParserT__4)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserNAME {
			{
				p.SetState(98)
				p.Match(goscriptParserNAME)
			}
			{
				p.SetState(99)
				p.Typename()
			}
			{
				p.SetState(100)
				p.Match(goscriptParserT__2)
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
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

type GenericTypeNameContext struct {
	*TypenameContext
}

func NewGenericTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GenericTypeNameContext {
	var p = new(GenericTypeNameContext)

	p.TypenameContext = NewEmptyTypenameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypenameContext))

	return p
}

func (s *GenericTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericTypeNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(goscriptParserNAME, 0)
}

func (s *GenericTypeNameContext) LT() antlr.TerminalNode {
	return s.GetToken(goscriptParserLT, 0)
}

func (s *GenericTypeNameContext) AllTypename() []ITypenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypenameContext)(nil)).Elem())
	var tst = make([]ITypenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypenameContext)
		}
	}

	return tst
}

func (s *GenericTypeNameContext) Typename(i int) ITypenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypenameContext)
}

func (s *GenericTypeNameContext) GT() antlr.TerminalNode {
	return s.GetToken(goscriptParserGT, 0)
}

func (s *GenericTypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.EnterGenericTypeName(s)
	}
}

func (s *GenericTypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goscriptListener); ok {
		listenerT.ExitGenericTypeName(s)
	}
}

func (p *goscriptParser) Typename() (localctx ITypenameContext) {
	localctx = NewTypenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, goscriptParserRULE_typename)
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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(goscriptParserNAME)
		}

	case 2:
		localctx = NewGenericTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(112)
			p.Match(goscriptParserLT)
		}
		{
			p.SetState(113)
			p.Typename()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(114)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(115)
				p.Typename()
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(121)
			p.Match(goscriptParserGT)
		}

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
	p.EnterRule(localctx, 12, goscriptParserRULE_enumdef)
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
		p.SetState(125)
		p.Match(goscriptParserTYPEDEF)
	}
	{
		p.SetState(126)
		p.Match(goscriptParserT__6)
	}
	{
		p.SetState(127)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(128)
		p.Match(goscriptParserT__4)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == goscriptParserNAME {
		{
			p.SetState(129)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(130)
			p.Match(goscriptParserINT)
		}
		{
			p.SetState(131)
			p.Match(goscriptParserT__2)
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 14, goscriptParserRULE_execution)

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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN:
		localctx = NewCtrlContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Control()
		}

	case goscriptParserT__1, goscriptParserT__13, goscriptParserNEW, goscriptParserLOCAL, goscriptParserBOOL, goscriptParserNULL, goscriptParserSUB, goscriptParserUNARYADD, goscriptParserUNARYSUB, goscriptParserNOT, goscriptParserINT, goscriptParserFLOAT, goscriptParserSTRING, goscriptParserNAME:
		localctx = NewLineProgramContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Line()
		}
		{
			p.SetState(141)
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

func (s *ReturnValContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 16, goscriptParserRULE_control)
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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(goscriptParserIF)
		}
		{
			p.SetState(146)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(147)
			p.expr(0)
		}
		{
			p.SetState(148)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(149)
			p.Block()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == goscriptParserELSE {
			{
				p.SetState(150)
				p.Match(goscriptParserELSE)
			}
			p.SetState(153)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case goscriptParserT__4:
				{
					p.SetState(151)
					p.Block()
				}

			case goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN:
				{
					p.SetState(152)
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
			p.SetState(157)
			p.Match(goscriptParserSWITCH)
		}
		{
			p.SetState(158)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(161)
			p.Match(goscriptParserT__4)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == goscriptParserCASE {
			{
				p.SetState(162)
				p.Match(goscriptParserCASE)
			}
			{
				p.SetState(163)
				p.Constant()
			}
			{
				p.SetState(164)
				p.Match(goscriptParserT__7)
			}
			{
				p.SetState(165)
				p.Block()
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(171)
			p.Match(goscriptParserT__5)
		}

	case 3:
		localctx = NewForInSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(174)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(175)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(176)
			p.Match(goscriptParserT__8)
		}
		{
			p.SetState(177)
			p.Collection()
		}
		{
			p.SetState(178)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(179)
			p.Block()
		}

	case 4:
		localctx = NewForInMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(182)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(183)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(184)
			p.Match(goscriptParserT__2)
		}
		{
			p.SetState(185)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(186)
			p.Match(goscriptParserT__8)
		}
		{
			p.SetState(187)
			p.Collection()
		}
		{
			p.SetState(188)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(189)
			p.Block()
		}

	case 5:
		localctx = NewForContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(191)
			p.Match(goscriptParserFOR)
		}
		{
			p.SetState(192)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(193)
			p.expr(0)
		}
		{
			p.SetState(194)
			p.Match(goscriptParserT__0)
		}
		{
			p.SetState(195)
			p.expr(0)
		}
		{
			p.SetState(196)
			p.Match(goscriptParserT__0)
		}
		{
			p.SetState(197)
			p.expr(0)
		}
		{
			p.SetState(198)
			p.Match(goscriptParserT__3)
		}
		{
			p.SetState(199)
			p.Block()
		}

	case 6:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(201)
			p.Match(goscriptParserBREAK)
		}
		{
			p.SetState(202)
			p.Match(goscriptParserT__0)
		}

	case 7:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(203)
			p.Match(goscriptParserCONTINUE)
		}
		{
			p.SetState(204)
			p.Match(goscriptParserT__0)
		}

	case 8:
		localctx = NewReturnVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(205)
			p.Match(goscriptParserRETURN)
		}
		{
			p.SetState(206)
			p.Match(goscriptParserT__0)
		}

	case 9:
		localctx = NewReturnValContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(207)
			p.Match(goscriptParserRETURN)
		}
		{
			p.SetState(208)
			p.expr(0)
		}
		{
			p.SetState(209)
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
	p.EnterRule(localctx, 18, goscriptParserRULE_collection)

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
		p.SetState(213)
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
	p.EnterRule(localctx, 20, goscriptParserRULE_block)
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
		p.SetState(215)
		p.Match(goscriptParserT__4)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<goscriptParserT__1)|(1<<goscriptParserT__13)|(1<<goscriptParserNEW)|(1<<goscriptParserFOR)|(1<<goscriptParserBREAK)|(1<<goscriptParserCONTINUE)|(1<<goscriptParserIF)|(1<<goscriptParserSWITCH)|(1<<goscriptParserRETURN)|(1<<goscriptParserVAR)|(1<<goscriptParserLOCAL)|(1<<goscriptParserBOOL)|(1<<goscriptParserNULL))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(goscriptParserSUB-35))|(1<<(goscriptParserUNARYADD-35))|(1<<(goscriptParserUNARYSUB-35))|(1<<(goscriptParserNOT-35))|(1<<(goscriptParserINT-35))|(1<<(goscriptParserFLOAT-35))|(1<<(goscriptParserSTRING-35))|(1<<(goscriptParserNAME-35)))) != 0) {
		p.SetState(218)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case goscriptParserT__1, goscriptParserT__13, goscriptParserNEW, goscriptParserFOR, goscriptParserBREAK, goscriptParserCONTINUE, goscriptParserIF, goscriptParserSWITCH, goscriptParserRETURN, goscriptParserLOCAL, goscriptParserBOOL, goscriptParserNULL, goscriptParserSUB, goscriptParserUNARYADD, goscriptParserUNARYSUB, goscriptParserNOT, goscriptParserINT, goscriptParserFLOAT, goscriptParserSTRING, goscriptParserNAME:
			{
				p.SetState(216)
				p.Execution()
			}

		case goscriptParserVAR:
			{
				p.SetState(217)
				p.Vardef()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(223)
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
	p.EnterRule(localctx, 22, goscriptParserRULE_line)
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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.expr(0)
		}

	case 2:
		localctx = NewFunctionAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Lhs()
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(227)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(228)
				p.Lhs()
			}

			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(234)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*FunctionAssignContext).op = _m
		}
		{
			p.SetState(235)
			p.Functions()
		}

	case 3:
		localctx = NewMultiAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.Lhs()
		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(238)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(239)
				p.Lhs()
			}

			p.SetState(244)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(245)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*MultiAssignContext).op = _m
		}
		{
			p.SetState(246)
			p.expr(0)
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(247)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(248)
				p.expr(0)
			}

			p.SetState(253)
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
	p.EnterRule(localctx, 24, goscriptParserRULE_lhs)

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
		p.SetState(256)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, goscriptParserRULE_variable, _p)
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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserNAME:
		localctx = NewVariableNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(259)
			p.Match(goscriptParserNAME)
		}

	case goscriptParserT__13:
		localctx = NewVariableNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(goscriptParserT__13)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(264)
					p.Match(goscriptParserDOT)
				}
				{
					p.SetState(265)
					p.Match(goscriptParserNAME)
				}

			case 2:
				localctx = NewSliceFilterContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(267)
					p.Match(goscriptParserT__9)
				}
				{
					p.SetState(268)
					p.Filter()
				}
				{
					p.SetState(269)
					p.Match(goscriptParserT__10)
				}

			case 3:
				localctx = NewIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(272)
					p.Match(goscriptParserT__11)
				}
				{
					p.SetState(273)
					p.expr(0)
				}
				{
					p.SetState(274)
					p.Match(goscriptParserT__12)
				}

			case 4:
				localctx = NewSliceMultiIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(277)
					p.Match(goscriptParserT__11)
				}
				{
					p.SetState(278)
					p.Indexs()
				}
				p.SetState(283)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == goscriptParserT__2 {
					{
						p.SetState(279)
						p.Match(goscriptParserT__2)
					}
					{
						p.SetState(280)
						p.Indexs()
					}

					p.SetState(285)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(286)
					p.Match(goscriptParserT__12)
				}

			case 5:
				localctx = NewMapMultiIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_variable)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(289)
					p.Match(goscriptParserT__11)
				}
				{
					p.SetState(290)
					p.Match(goscriptParserT__11)
				}
				{
					p.SetState(291)
					p.expr(0)
				}
				p.SetState(296)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == goscriptParserT__2 {
					{
						p.SetState(292)
						p.Match(goscriptParserT__2)
					}
					{
						p.SetState(293)
						p.expr(0)
					}

					p.SetState(298)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(299)
					p.Match(goscriptParserT__12)
				}
				{
					p.SetState(300)
					p.Match(goscriptParserT__12)
				}

			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, goscriptParserRULE_filter)

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
		p.SetState(307)
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
	p.EnterRule(localctx, 30, goscriptParserRULE_indexs)

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

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIndexType1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.expr(0)
		}
		{
			p.SetState(310)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(311)
			p.expr(0)
		}
		{
			p.SetState(312)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(313)
			p.expr(0)
		}

	case 2:
		localctx = NewIndexType2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.expr(0)
		}
		{
			p.SetState(316)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(317)
			p.expr(0)
		}

	case 3:
		localctx = NewIndexType3Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.expr(0)
		}
		{
			p.SetState(320)
			p.Match(goscriptParserT__7)
		}

	case 4:
		localctx = NewIndexType4Context(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(322)
			p.expr(0)
		}

	case 5:
		localctx = NewIndexType5Context(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(323)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(324)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, goscriptParserRULE_expr, _p)
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
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(328)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(329)
			p.expr(0)
		}
		{
			p.SetState(330)
			p.Match(goscriptParserT__3)
		}

	case 2:
		localctx = NewLeftUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(332)

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
			p.SetState(333)
			p.variable(0)
		}

	case 3:
		localctx = NewLeftUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(334)

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
			p.SetState(335)
			p.expr(15)
		}

	case 4:
		localctx = NewRightUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(336)
			p.variable(0)
		}
		{
			p.SetState(337)

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
			p.SetState(339)
			p.Lhs()
		}
		{
			p.SetState(340)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BinaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(goscriptParserASSIGN-48))|(1<<(goscriptParserADDEQUAL-48))|(1<<(goscriptParserSUBEQUAL-48))|(1<<(goscriptParserMULEQUAL-48))|(1<<(goscriptParserDIVEQUAL-48)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BinaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(341)
			p.expr(7)
		}

	case 6:
		localctx = NewAssignInitializationlistContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(343)
			p.Lhs()
		}
		{
			p.SetState(344)

			var _m = p.Match(goscriptParserASSIGN)

			localctx.(*AssignInitializationlistContext).op = _m
		}
		{
			p.SetState(345)
			p.InitializationListBegin()
		}

	case 7:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(347)
			p.Constant()
		}

	case 8:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(348)
			p.variable(0)
		}

	case 9:
		localctx = NewPassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(349)
			p.Functions()
		}

	case 10:
		localctx = NewConstructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(350)
			p.Constructor()
		}

	case 11:
		localctx = NewLocalDefContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(351)
			p.Localdef()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(372)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(355)

					var _m = p.Match(goscriptParserPOW)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(356)
					p.expr(13)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(357)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(358)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(goscriptParserMUL-31))|(1<<(goscriptParserDIV-31))|(1<<(goscriptParserMOD-31)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(359)
					p.expr(13)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(361)

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
					p.SetState(362)
					p.expr(12)
				}

			case 4:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(363)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(364)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(goscriptParserEQ-38))|(1<<(goscriptParserINEQ-38))|(1<<(goscriptParserGT-38))|(1<<(goscriptParserGE-38))|(1<<(goscriptParserLE-38))|(1<<(goscriptParserLT-38))|(1<<(goscriptParserREGEX-38)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(365)
					p.expr(11)
				}

			case 5:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(366)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(367)

					var _m = p.Match(goscriptParserAND)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(368)
					p.expr(10)
				}

			case 6:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goscriptParserRULE_expr)
				p.SetState(369)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(370)

					var _m = p.Match(goscriptParserOR)

					localctx.(*BinaryContext).op = _m
				}
				{
					p.SetState(371)
					p.expr(9)
				}

			}

		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, goscriptParserRULE_initializationListBegin)

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
		p.SetState(377)
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
	p.EnterRule(localctx, 36, goscriptParserRULE_initializationList)
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

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInitSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Match(goscriptParserT__11)
		}
		{
			p.SetState(380)
			p.InitializationList()
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(381)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(382)
				p.InitializationList()
			}

			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(388)
			p.Match(goscriptParserT__12)
		}

	case 2:
		localctx = NewInitMessageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)
			p.Match(goscriptParserT__4)
		}
		{
			p.SetState(391)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(392)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(393)
			p.InitializationList()
		}
		{
			p.SetState(394)
			p.Match(goscriptParserT__3)
		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(395)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(396)
				p.Match(goscriptParserNAME)
			}
			{
				p.SetState(397)
				p.Match(goscriptParserT__1)
			}
			{
				p.SetState(398)
				p.InitializationList()
			}
			{
				p.SetState(399)
				p.Match(goscriptParserT__3)
			}

			p.SetState(405)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(406)
			p.Match(goscriptParserT__5)
		}

	case 3:
		localctx = NewInitMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(408)
			p.Match(goscriptParserT__4)
		}
		{
			p.SetState(409)
			p.InitializationList()
		}
		{
			p.SetState(410)
			p.Match(goscriptParserT__7)
		}
		{
			p.SetState(411)
			p.InitializationList()
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(412)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(413)
				p.InitializationList()
			}
			{
				p.SetState(414)
				p.Match(goscriptParserT__7)
			}
			{
				p.SetState(415)
				p.InitializationList()
			}

			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(422)
			p.Match(goscriptParserT__5)
		}

	case 4:
		localctx = NewInitConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(424)
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
	p.EnterRule(localctx, 38, goscriptParserRULE_constant)

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

	p.SetState(432)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goscriptParserINT:
		localctx = NewConstantIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(427)
			p.Match(goscriptParserINT)
		}

	case goscriptParserFLOAT:
		localctx = NewConstantFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(goscriptParserFLOAT)
		}

	case goscriptParserBOOL:
		localctx = NewConstantBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Match(goscriptParserBOOL)
		}

	case goscriptParserNULL:
		localctx = NewConstantNilContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(430)
			p.Match(goscriptParserNULL)
		}

	case goscriptParserSTRING:
		localctx = NewConstantStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(431)
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
	p.EnterRule(localctx, 40, goscriptParserRULE_functions)
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

	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(434)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(435)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(436)
			p.expr(0)
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(437)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(438)
				p.expr(0)
			}

			p.SetState(443)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(444)
			p.Match(goscriptParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.variable(0)
		}
		{
			p.SetState(447)
			p.Match(goscriptParserDOT)
		}
		{
			p.SetState(448)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(449)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(450)
			p.Match(goscriptParserT__3)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(452)
			p.variable(0)
		}
		{
			p.SetState(453)
			p.Match(goscriptParserDOT)
		}
		{
			p.SetState(454)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(455)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(456)
			p.expr(0)
		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(457)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(458)
				p.expr(0)
			}

			p.SetState(463)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(464)
			p.Match(goscriptParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(466)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(467)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(468)
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
	p.EnterRule(localctx, 42, goscriptParserRULE_constructor)
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

	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(471)
			p.Match(goscriptParserNEW)
		}
		{
			p.SetState(472)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(473)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(474)
			p.Match(goscriptParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(475)
			p.Match(goscriptParserNEW)
		}
		{
			p.SetState(476)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(477)
			p.Match(goscriptParserT__1)
		}
		{
			p.SetState(478)
			p.expr(0)
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == goscriptParserT__2 {
			{
				p.SetState(479)
				p.Match(goscriptParserT__2)
			}
			{
				p.SetState(480)
				p.expr(0)
			}

			p.SetState(485)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(486)
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
	p.EnterRule(localctx, 44, goscriptParserRULE_vardef)

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
		p.SetState(490)
		p.Match(goscriptParserVAR)
	}
	{
		p.SetState(491)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(492)
		p.Match(goscriptParserNAME)
	}
	{
		p.SetState(493)
		p.Match(goscriptParserT__0)
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
	p.EnterRule(localctx, 46, goscriptParserRULE_localdef)

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

	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Match(goscriptParserLOCAL)
		}
		{
			p.SetState(496)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(497)
			p.Match(goscriptParserNAME)
		}

	case 2:
		localctx = NewLocalAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
			p.Match(goscriptParserLOCAL)
		}
		{
			p.SetState(499)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(500)
			p.Match(goscriptParserNAME)
		}
		{
			p.SetState(501)
			p.Match(goscriptParserASSIGN)
		}
		{
			p.SetState(502)
			p.expr(0)
		}

	}

	return localctx
}

func (p *goscriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	case 16:
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
