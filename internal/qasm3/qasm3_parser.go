// Code generated from ./qasm3.g4 by ANTLR 4.9. DO NOT EDIT.

package qasm3 // qasm3
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 116, 732,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9,
	81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86, 9, 86,
	3, 2, 3, 2, 7, 2, 175, 10, 2, 12, 2, 14, 2, 178, 11, 2, 3, 3, 5, 3, 181,
	10, 3, 3, 3, 7, 3, 184, 10, 3, 12, 3, 14, 3, 187, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 193, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 212, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 218, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 223,
	10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	5, 11, 235, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 251, 10, 14, 12, 14,
	14, 14, 254, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 262,
	10, 15, 3, 16, 3, 16, 3, 16, 7, 16, 267, 10, 16, 12, 16, 14, 16, 270, 11,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 5, 20, 284, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 7,
	21, 291, 10, 21, 12, 21, 14, 21, 294, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 306, 10, 25, 3, 26, 3,
	26, 5, 26, 310, 10, 26, 3, 26, 3, 26, 5, 26, 314, 10, 26, 3, 26, 3, 26,
	3, 26, 5, 26, 319, 10, 26, 5, 26, 321, 10, 26, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 342, 10, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 5, 32, 348, 10, 32, 3, 33, 3, 33, 5, 33, 352, 10, 33, 3, 34,
	3, 34, 3, 34, 7, 34, 357, 10, 34, 12, 34, 14, 34, 360, 11, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 370, 10, 36, 12, 36,
	14, 36, 373, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	5, 38, 393, 10, 38, 3, 39, 3, 39, 5, 39, 397, 10, 39, 3, 39, 3, 39, 5,
	39, 401, 10, 39, 3, 39, 3, 39, 5, 39, 405, 10, 39, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 5, 41, 416, 10, 41, 3, 41, 5,
	41, 419, 10, 41, 3, 41, 3, 41, 3, 42, 3, 42, 7, 42, 425, 10, 42, 12, 42,
	14, 42, 428, 11, 42, 3, 42, 3, 42, 3, 43, 3, 43, 5, 43, 434, 10, 43, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 5, 44, 441, 10, 44, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 454, 10,
	46, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48,
	465, 10, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 5, 49, 472, 10, 49, 3,
	49, 5, 49, 475, 10, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 5, 50, 486, 10, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 5, 53, 497, 10, 53, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 507, 10, 54, 3, 54, 3, 54, 3, 54, 3,
	54, 5, 54, 513, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 54, 7, 54, 526, 10, 54, 12, 54, 14, 54, 529, 11,
	54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 537, 10, 55, 3, 56,
	3, 56, 3, 56, 7, 56, 542, 10, 56, 12, 56, 14, 56, 545, 11, 56, 3, 56, 3,
	56, 3, 57, 3, 57, 3, 57, 5, 57, 552, 10, 57, 3, 58, 3, 58, 3, 59, 3, 59,
	3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3,
	63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 574, 10, 64, 3, 65, 3, 65,
	5, 65, 578, 10, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 5,
	66, 587, 10, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67,
	596, 10, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3,
	70, 3, 70, 3, 70, 5, 70, 609, 10, 70, 3, 70, 5, 70, 612, 10, 70, 3, 70,
	5, 70, 615, 10, 70, 3, 70, 5, 70, 618, 10, 70, 3, 70, 3, 70, 3, 71, 3,
	71, 3, 71, 3, 71, 5, 71, 626, 10, 71, 3, 71, 5, 71, 629, 10, 71, 3, 71,
	5, 71, 632, 10, 71, 3, 71, 3, 71, 3, 72, 3, 72, 5, 72, 638, 10, 72, 3,
	73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 5, 75,
	650, 10, 75, 5, 75, 652, 10, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3,
	76, 3, 76, 5, 76, 661, 10, 76, 3, 77, 3, 77, 5, 77, 665, 10, 77, 3, 78,
	3, 78, 5, 78, 669, 10, 78, 3, 78, 3, 78, 3, 78, 3, 78, 5, 78, 675, 10,
	78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 5, 80, 682, 10, 80, 3, 80, 5, 80,
	685, 10, 80, 3, 80, 3, 80, 3, 80, 5, 80, 690, 10, 80, 3, 81, 3, 81, 3,
	81, 3, 81, 5, 81, 696, 10, 81, 3, 82, 3, 82, 5, 82, 700, 10, 82, 3, 83,
	3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 5, 84, 708, 10, 84, 3, 84, 3, 84, 3,
	84, 5, 84, 713, 10, 84, 3, 84, 5, 84, 716, 10, 84, 3, 84, 3, 84, 5, 84,
	720, 10, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 5,
	86, 730, 10, 86, 3, 86, 2, 3, 106, 87, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
	92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120,
	122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150,
	152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 2, 15, 3, 2, 115, 116,
	3, 2, 5, 6, 3, 2, 7, 8, 3, 2, 9, 12, 3, 2, 28, 29, 4, 2, 17, 17, 30, 47,
	3, 2, 49, 56, 3, 2, 57, 58, 4, 2, 59, 68, 105, 106, 3, 2, 74, 76, 3, 2,
	80, 84, 3, 2, 90, 91, 4, 2, 94, 94, 113, 113, 2, 752, 2, 172, 3, 2, 2,
	2, 4, 180, 3, 2, 2, 2, 6, 188, 3, 2, 2, 2, 8, 196, 3, 2, 2, 2, 10, 211,
	3, 2, 2, 2, 12, 217, 3, 2, 2, 2, 14, 222, 3, 2, 2, 2, 16, 226, 3, 2, 2,
	2, 18, 228, 3, 2, 2, 2, 20, 231, 3, 2, 2, 2, 22, 238, 3, 2, 2, 2, 24, 242,
	3, 2, 2, 2, 26, 252, 3, 2, 2, 2, 28, 257, 3, 2, 2, 2, 30, 268, 3, 2, 2,
	2, 32, 273, 3, 2, 2, 2, 34, 276, 3, 2, 2, 2, 36, 278, 3, 2, 2, 2, 38, 281,
	3, 2, 2, 2, 40, 292, 3, 2, 2, 2, 42, 297, 3, 2, 2, 2, 44, 299, 3, 2, 2,
	2, 46, 301, 3, 2, 2, 2, 48, 305, 3, 2, 2, 2, 50, 320, 3, 2, 2, 2, 52, 322,
	3, 2, 2, 2, 54, 327, 3, 2, 2, 2, 56, 331, 3, 2, 2, 2, 58, 335, 3, 2, 2,
	2, 60, 338, 3, 2, 2, 2, 62, 347, 3, 2, 2, 2, 64, 349, 3, 2, 2, 2, 66, 358,
	3, 2, 2, 2, 68, 363, 3, 2, 2, 2, 70, 371, 3, 2, 2, 2, 72, 376, 3, 2, 2,
	2, 74, 381, 3, 2, 2, 2, 76, 394, 3, 2, 2, 2, 78, 408, 3, 2, 2, 2, 80, 412,
	3, 2, 2, 2, 82, 422, 3, 2, 2, 2, 84, 433, 3, 2, 2, 2, 86, 440, 3, 2, 2,
	2, 88, 442, 3, 2, 2, 2, 90, 453, 3, 2, 2, 2, 92, 455, 3, 2, 2, 2, 94, 464,
	3, 2, 2, 2, 96, 468, 3, 2, 2, 2, 98, 485, 3, 2, 2, 2, 100, 487, 3, 2, 2,
	2, 102, 489, 3, 2, 2, 2, 104, 496, 3, 2, 2, 2, 106, 512, 3, 2, 2, 2, 108,
	536, 3, 2, 2, 2, 110, 543, 3, 2, 2, 2, 112, 551, 3, 2, 2, 2, 114, 553,
	3, 2, 2, 2, 116, 555, 3, 2, 2, 2, 118, 557, 3, 2, 2, 2, 120, 559, 3, 2,
	2, 2, 122, 562, 3, 2, 2, 2, 124, 564, 3, 2, 2, 2, 126, 573, 3, 2, 2, 2,
	128, 577, 3, 2, 2, 2, 130, 579, 3, 2, 2, 2, 132, 595, 3, 2, 2, 2, 134,
	599, 3, 2, 2, 2, 136, 602, 3, 2, 2, 2, 138, 604, 3, 2, 2, 2, 140, 621,
	3, 2, 2, 2, 142, 637, 3, 2, 2, 2, 144, 639, 3, 2, 2, 2, 146, 644, 3, 2,
	2, 2, 148, 651, 3, 2, 2, 2, 150, 660, 3, 2, 2, 2, 152, 664, 3, 2, 2, 2,
	154, 674, 3, 2, 2, 2, 156, 676, 3, 2, 2, 2, 158, 678, 3, 2, 2, 2, 160,
	695, 3, 2, 2, 2, 162, 699, 3, 2, 2, 2, 164, 701, 3, 2, 2, 2, 166, 705,
	3, 2, 2, 2, 168, 725, 3, 2, 2, 2, 170, 729, 3, 2, 2, 2, 172, 176, 5, 4,
	3, 2, 173, 175, 5, 10, 6, 2, 174, 173, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2,
	176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 3, 3, 2, 2, 2, 178, 176,
	3, 2, 2, 2, 179, 181, 5, 6, 4, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2,
	2, 2, 181, 185, 3, 2, 2, 2, 182, 184, 5, 8, 5, 2, 183, 182, 3, 2, 2, 2,
	184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	5, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 189, 7, 3, 2, 2, 189, 192, 7,
	110, 2, 2, 190, 191, 7, 103, 2, 2, 191, 193, 7, 110, 2, 2, 192, 190, 3,
	2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 7, 102,
	2, 2, 195, 7, 3, 2, 2, 2, 196, 197, 7, 4, 2, 2, 197, 198, 7, 114, 2, 2,
	198, 199, 7, 102, 2, 2, 199, 9, 3, 2, 2, 2, 200, 212, 5, 12, 7, 2, 201,
	212, 5, 104, 53, 2, 202, 212, 5, 14, 8, 2, 203, 212, 5, 130, 66, 2, 204,
	212, 5, 132, 67, 2, 205, 212, 5, 134, 68, 2, 206, 212, 5, 72, 37, 2, 207,
	212, 5, 84, 43, 2, 208, 212, 5, 160, 81, 2, 209, 212, 5, 144, 73, 2, 210,
	212, 5, 16, 9, 2, 211, 200, 3, 2, 2, 2, 211, 201, 3, 2, 2, 2, 211, 202,
	3, 2, 2, 2, 211, 203, 3, 2, 2, 2, 211, 204, 3, 2, 2, 2, 211, 205, 3, 2,
	2, 2, 211, 206, 3, 2, 2, 2, 211, 207, 3, 2, 2, 2, 211, 208, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 11, 3, 2, 2, 2, 213, 218,
	5, 140, 71, 2, 214, 218, 5, 138, 70, 2, 215, 218, 5, 78, 40, 2, 216, 218,
	5, 166, 84, 2, 217, 213, 3, 2, 2, 2, 217, 214, 3, 2, 2, 2, 217, 215, 3,
	2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 13, 3, 2, 2, 2, 219, 223, 5, 36, 19,
	2, 220, 223, 5, 64, 33, 2, 221, 223, 5, 52, 27, 2, 222, 219, 3, 2, 2, 2,
	222, 220, 3, 2, 2, 2, 222, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224,
	225, 7, 102, 2, 2, 225, 15, 3, 2, 2, 2, 226, 227, 9, 2, 2, 2, 227, 17,
	3, 2, 2, 2, 228, 229, 7, 106, 2, 2, 229, 230, 5, 64, 33, 2, 230, 19, 3,
	2, 2, 2, 231, 234, 7, 97, 2, 2, 232, 235, 5, 20, 11, 2, 233, 235, 5, 10,
	6, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2,
	236, 237, 7, 98, 2, 2, 237, 21, 3, 2, 2, 2, 238, 239, 7, 95, 2, 2, 239,
	240, 5, 106, 54, 2, 240, 241, 7, 96, 2, 2, 241, 23, 3, 2, 2, 2, 242, 243,
	7, 95, 2, 2, 243, 244, 5, 106, 54, 2, 244, 245, 7, 104, 2, 2, 245, 246,
	5, 106, 54, 2, 246, 247, 7, 96, 2, 2, 247, 25, 3, 2, 2, 2, 248, 249, 7,
	113, 2, 2, 249, 251, 7, 104, 2, 2, 250, 248, 3, 2, 2, 2, 251, 254, 3, 2,
	2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2,
	254, 252, 3, 2, 2, 2, 255, 256, 7, 113, 2, 2, 256, 27, 3, 2, 2, 2, 257,
	261, 7, 113, 2, 2, 258, 262, 5, 22, 12, 2, 259, 262, 5, 24, 13, 2, 260,
	262, 5, 76, 39, 2, 261, 258, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 260,
	3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 29, 3, 2, 2, 2, 263, 264, 5, 28,
	15, 2, 264, 265, 7, 104, 2, 2, 265, 267, 3, 2, 2, 2, 266, 263, 3, 2, 2,
	2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269,
	271, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 272, 5, 28, 15, 2, 272, 31,
	3, 2, 2, 2, 273, 274, 7, 101, 2, 2, 274, 275, 7, 113, 2, 2, 275, 33, 3,
	2, 2, 2, 276, 277, 9, 3, 2, 2, 277, 35, 3, 2, 2, 2, 278, 279, 5, 34, 18,
	2, 279, 280, 5, 30, 16, 2, 280, 37, 3, 2, 2, 2, 281, 283, 5, 34, 18, 2,
	282, 284, 5, 22, 12, 2, 283, 282, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284,
	285, 3, 2, 2, 2, 285, 286, 5, 32, 17, 2, 286, 39, 3, 2, 2, 2, 287, 288,
	5, 38, 20, 2, 288, 289, 7, 104, 2, 2, 289, 291, 3, 2, 2, 2, 290, 287, 3,
	2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2,
	2, 293, 295, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 296, 5, 38, 20, 2,
	296, 41, 3, 2, 2, 2, 297, 298, 9, 4, 2, 2, 298, 43, 3, 2, 2, 2, 299, 300,
	9, 5, 2, 2, 300, 45, 3, 2, 2, 2, 301, 302, 7, 13, 2, 2, 302, 47, 3, 2,
	2, 2, 303, 306, 7, 14, 2, 2, 304, 306, 5, 148, 75, 2, 305, 303, 3, 2, 2,
	2, 305, 304, 3, 2, 2, 2, 306, 49, 3, 2, 2, 2, 307, 309, 5, 44, 23, 2, 308,
	310, 5, 22, 12, 2, 309, 308, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 321,
	3, 2, 2, 2, 311, 313, 5, 46, 24, 2, 312, 314, 5, 24, 13, 2, 313, 312, 3,
	2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 321, 3, 2, 2, 2, 315, 321, 5, 48, 25,
	2, 316, 318, 5, 42, 22, 2, 317, 319, 5, 22, 12, 2, 318, 317, 3, 2, 2, 2,
	318, 319, 3, 2, 2, 2, 319, 321, 3, 2, 2, 2, 320, 307, 3, 2, 2, 2, 320,
	311, 3, 2, 2, 2, 320, 315, 3, 2, 2, 2, 320, 316, 3, 2, 2, 2, 321, 51, 3,
	2, 2, 2, 322, 323, 7, 15, 2, 2, 323, 324, 7, 113, 2, 2, 324, 325, 7, 105,
	2, 2, 325, 326, 5, 106, 54, 2, 326, 53, 3, 2, 2, 2, 327, 328, 5, 44, 23,
	2, 328, 329, 5, 22, 12, 2, 329, 330, 7, 113, 2, 2, 330, 55, 3, 2, 2, 2,
	331, 332, 5, 46, 24, 2, 332, 333, 5, 24, 13, 2, 333, 334, 7, 113, 2, 2,
	334, 57, 3, 2, 2, 2, 335, 336, 5, 48, 25, 2, 336, 337, 7, 113, 2, 2, 337,
	59, 3, 2, 2, 2, 338, 339, 5, 42, 22, 2, 339, 341, 7, 113, 2, 2, 340, 342,
	5, 22, 12, 2, 341, 340, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 61, 3, 2,
	2, 2, 343, 348, 5, 54, 28, 2, 344, 348, 5, 56, 29, 2, 345, 348, 5, 58,
	30, 2, 346, 348, 5, 60, 31, 2, 347, 343, 3, 2, 2, 2, 347, 344, 3, 2, 2,
	2, 347, 345, 3, 2, 2, 2, 347, 346, 3, 2, 2, 2, 348, 63, 3, 2, 2, 2, 349,
	351, 5, 62, 32, 2, 350, 352, 5, 120, 61, 2, 351, 350, 3, 2, 2, 2, 351,
	352, 3, 2, 2, 2, 352, 65, 3, 2, 2, 2, 353, 354, 5, 50, 26, 2, 354, 355,
	7, 104, 2, 2, 355, 357, 3, 2, 2, 2, 356, 353, 3, 2, 2, 2, 357, 360, 3,
	2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 361, 3, 2, 2,
	2, 360, 358, 3, 2, 2, 2, 361, 362, 5, 50, 26, 2, 362, 67, 3, 2, 2, 2, 363,
	364, 5, 50, 26, 2, 364, 365, 5, 32, 17, 2, 365, 69, 3, 2, 2, 2, 366, 367,
	5, 68, 35, 2, 367, 368, 7, 104, 2, 2, 368, 370, 3, 2, 2, 2, 369, 366, 3,
	2, 2, 2, 370, 373, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 371, 372, 3, 2, 2,
	2, 372, 374, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 374, 375, 5, 68, 35, 2,
	375, 71, 3, 2, 2, 2, 376, 377, 7, 16, 2, 2, 377, 378, 7, 113, 2, 2, 378,
	379, 7, 105, 2, 2, 379, 380, 5, 74, 38, 2, 380, 73, 3, 2, 2, 2, 381, 392,
	7, 105, 2, 2, 382, 383, 7, 113, 2, 2, 383, 393, 5, 76, 39, 2, 384, 385,
	7, 113, 2, 2, 385, 386, 7, 17, 2, 2, 386, 393, 7, 113, 2, 2, 387, 388,
	7, 113, 2, 2, 388, 389, 7, 95, 2, 2, 389, 390, 5, 110, 56, 2, 390, 391,
	7, 96, 2, 2, 391, 393, 3, 2, 2, 2, 392, 382, 3, 2, 2, 2, 392, 384, 3, 2,
	2, 2, 392, 387, 3, 2, 2, 2, 393, 75, 3, 2, 2, 2, 394, 396, 7, 95, 2, 2,
	395, 397, 5, 106, 54, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397,
	398, 3, 2, 2, 2, 398, 400, 7, 101, 2, 2, 399, 401, 5, 106, 54, 2, 400,
	399, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 404, 3, 2, 2, 2, 402, 403,
	7, 101, 2, 2, 403, 405, 5, 106, 54, 2, 404, 402, 3, 2, 2, 2, 404, 405,
	3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 407, 7, 96, 2, 2, 407, 77, 3, 2,
	2, 2, 408, 409, 7, 18, 2, 2, 409, 410, 5, 80, 41, 2, 410, 411, 5, 82, 42,
	2, 411, 79, 3, 2, 2, 2, 412, 418, 7, 113, 2, 2, 413, 415, 7, 99, 2, 2,
	414, 416, 5, 70, 36, 2, 415, 414, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416,
	417, 3, 2, 2, 2, 417, 419, 7, 100, 2, 2, 418, 413, 3, 2, 2, 2, 418, 419,
	3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 5, 26, 14, 2, 421, 81, 3, 2,
	2, 2, 422, 426, 7, 97, 2, 2, 423, 425, 5, 84, 43, 2, 424, 423, 3, 2, 2,
	2, 425, 428, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427,
	429, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 429, 430, 7, 98, 2, 2, 430, 83,
	3, 2, 2, 2, 431, 434, 5, 86, 44, 2, 432, 434, 5, 90, 46, 2, 433, 431, 3,
	2, 2, 2, 433, 432, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 436, 7, 102,
	2, 2, 436, 85, 3, 2, 2, 2, 437, 441, 5, 96, 49, 2, 438, 441, 5, 88, 45,
	2, 439, 441, 5, 92, 47, 2, 440, 437, 3, 2, 2, 2, 440, 438, 3, 2, 2, 2,
	440, 439, 3, 2, 2, 2, 441, 87, 3, 2, 2, 2, 442, 443, 7, 19, 2, 2, 443,
	444, 5, 30, 16, 2, 444, 89, 3, 2, 2, 2, 445, 446, 5, 88, 45, 2, 446, 447,
	7, 106, 2, 2, 447, 448, 5, 30, 16, 2, 448, 454, 3, 2, 2, 2, 449, 450, 5,
	30, 16, 2, 450, 451, 7, 105, 2, 2, 451, 452, 5, 88, 45, 2, 452, 454, 3,
	2, 2, 2, 453, 445, 3, 2, 2, 2, 453, 449, 3, 2, 2, 2, 454, 91, 3, 2, 2,
	2, 455, 456, 7, 20, 2, 2, 456, 457, 5, 30, 16, 2, 457, 93, 3, 2, 2, 2,
	458, 465, 7, 21, 2, 2, 459, 460, 7, 22, 2, 2, 460, 461, 7, 99, 2, 2, 461,
	462, 7, 110, 2, 2, 462, 465, 7, 100, 2, 2, 463, 465, 7, 23, 2, 2, 464,
	458, 3, 2, 2, 2, 464, 459, 3, 2, 2, 2, 464, 463, 3, 2, 2, 2, 465, 466,
	3, 2, 2, 2, 466, 467, 7, 24, 2, 2, 467, 95, 3, 2, 2, 2, 468, 474, 5, 98,
	50, 2, 469, 471, 7, 99, 2, 2, 470, 472, 5, 110, 56, 2, 471, 470, 3, 2,
	2, 2, 471, 472, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 475, 7, 100, 2,
	2, 474, 469, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2, 476,
	477, 5, 30, 16, 2, 477, 97, 3, 2, 2, 2, 478, 486, 7, 25, 2, 2, 479, 486,
	7, 26, 2, 2, 480, 486, 7, 27, 2, 2, 481, 486, 7, 113, 2, 2, 482, 483, 5,
	94, 48, 2, 483, 484, 5, 98, 50, 2, 484, 486, 3, 2, 2, 2, 485, 478, 3, 2,
	2, 2, 485, 479, 3, 2, 2, 2, 485, 480, 3, 2, 2, 2, 485, 481, 3, 2, 2, 2,
	485, 482, 3, 2, 2, 2, 486, 99, 3, 2, 2, 2, 487, 488, 9, 6, 2, 2, 488, 101,
	3, 2, 2, 2, 489, 490, 9, 7, 2, 2, 490, 103, 3, 2, 2, 2, 491, 492, 5, 106,
	54, 2, 492, 493, 7, 102, 2, 2, 493, 497, 3, 2, 2, 2, 494, 495, 7, 48, 2,
	2, 495, 497, 5, 104, 53, 2, 496, 491, 3, 2, 2, 2, 496, 494, 3, 2, 2, 2,
	497, 105, 3, 2, 2, 2, 498, 499, 8, 54, 1, 2, 499, 500, 5, 100, 51, 2, 500,
	501, 5, 106, 54, 9, 501, 513, 3, 2, 2, 2, 502, 513, 5, 124, 63, 2, 503,
	504, 5, 112, 57, 2, 504, 506, 7, 99, 2, 2, 505, 507, 5, 110, 56, 2, 506,
	505, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 509,
	7, 100, 2, 2, 509, 513, 3, 2, 2, 2, 510, 513, 5, 88, 45, 2, 511, 513, 5,
	108, 55, 2, 512, 498, 3, 2, 2, 2, 512, 502, 3, 2, 2, 2, 512, 503, 3, 2,
	2, 2, 512, 510, 3, 2, 2, 2, 512, 511, 3, 2, 2, 2, 513, 527, 3, 2, 2, 2,
	514, 515, 12, 10, 2, 2, 515, 516, 5, 102, 52, 2, 516, 517, 5, 106, 54,
	11, 517, 526, 3, 2, 2, 2, 518, 519, 12, 7, 2, 2, 519, 520, 7, 95, 2, 2,
	520, 521, 5, 106, 54, 2, 521, 522, 7, 96, 2, 2, 522, 526, 3, 2, 2, 2, 523,
	524, 12, 5, 2, 2, 524, 526, 5, 118, 60, 2, 525, 514, 3, 2, 2, 2, 525, 518,
	3, 2, 2, 2, 525, 523, 3, 2, 2, 2, 526, 529, 3, 2, 2, 2, 527, 525, 3, 2,
	2, 2, 527, 528, 3, 2, 2, 2, 528, 107, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2,
	530, 537, 7, 107, 2, 2, 531, 537, 7, 110, 2, 2, 532, 537, 7, 112, 2, 2,
	533, 537, 7, 113, 2, 2, 534, 537, 7, 114, 2, 2, 535, 537, 5, 152, 77, 2,
	536, 530, 3, 2, 2, 2, 536, 531, 3, 2, 2, 2, 536, 532, 3, 2, 2, 2, 536,
	533, 3, 2, 2, 2, 536, 534, 3, 2, 2, 2, 536, 535, 3, 2, 2, 2, 537, 109,
	3, 2, 2, 2, 538, 539, 5, 106, 54, 2, 539, 540, 7, 104, 2, 2, 540, 542,
	3, 2, 2, 2, 541, 538, 3, 2, 2, 2, 542, 545, 3, 2, 2, 2, 543, 541, 3, 2,
	2, 2, 543, 544, 3, 2, 2, 2, 544, 546, 3, 2, 2, 2, 545, 543, 3, 2, 2, 2,
	546, 547, 5, 106, 54, 2, 547, 111, 3, 2, 2, 2, 548, 552, 7, 113, 2, 2,
	549, 552, 5, 114, 58, 2, 550, 552, 5, 116, 59, 2, 551, 548, 3, 2, 2, 2,
	551, 549, 3, 2, 2, 2, 551, 550, 3, 2, 2, 2, 552, 113, 3, 2, 2, 2, 553,
	554, 9, 8, 2, 2, 554, 115, 3, 2, 2, 2, 555, 556, 5, 50, 26, 2, 556, 117,
	3, 2, 2, 2, 557, 558, 9, 9, 2, 2, 558, 119, 3, 2, 2, 2, 559, 560, 5, 122,
	62, 2, 560, 561, 5, 106, 54, 2, 561, 121, 3, 2, 2, 2, 562, 563, 9, 10,
	2, 2, 563, 123, 3, 2, 2, 2, 564, 565, 7, 113, 2, 2, 565, 566, 7, 69, 2,
	2, 566, 567, 5, 126, 64, 2, 567, 125, 3, 2, 2, 2, 568, 569, 7, 97, 2, 2,
	569, 570, 5, 110, 56, 2, 570, 571, 7, 98, 2, 2, 571, 574, 3, 2, 2, 2, 572,
	574, 5, 76, 39, 2, 573, 568, 3, 2, 2, 2, 573, 572, 3, 2, 2, 2, 574, 127,
	3, 2, 2, 2, 575, 578, 5, 10, 6, 2, 576, 578, 5, 20, 11, 2, 577, 575, 3,
	2, 2, 2, 577, 576, 3, 2, 2, 2, 578, 129, 3, 2, 2, 2, 579, 580, 7, 70, 2,
	2, 580, 581, 7, 99, 2, 2, 581, 582, 5, 106, 54, 2, 582, 583, 7, 100, 2,
	2, 583, 586, 5, 128, 65, 2, 584, 585, 7, 71, 2, 2, 585, 587, 5, 128, 65,
	2, 586, 584, 3, 2, 2, 2, 586, 587, 3, 2, 2, 2, 587, 131, 3, 2, 2, 2, 588,
	589, 7, 72, 2, 2, 589, 596, 5, 124, 63, 2, 590, 591, 7, 73, 2, 2, 591,
	592, 7, 99, 2, 2, 592, 593, 5, 106, 54, 2, 593, 594, 7, 100, 2, 2, 594,
	596, 3, 2, 2, 2, 595, 588, 3, 2, 2, 2, 595, 590, 3, 2, 2, 2, 596, 597,
	3, 2, 2, 2, 597, 598, 5, 128, 65, 2, 598, 133, 3, 2, 2, 2, 599, 600, 5,
	136, 69, 2, 600, 601, 7, 102, 2, 2, 601, 135, 3, 2, 2, 2, 602, 603, 9,
	11, 2, 2, 603, 137, 3, 2, 2, 2, 604, 605, 7, 77, 2, 2, 605, 611, 7, 113,
	2, 2, 606, 608, 7, 99, 2, 2, 607, 609, 5, 66, 34, 2, 608, 607, 3, 2, 2,
	2, 608, 609, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 612, 7, 100, 2, 2,
	611, 606, 3, 2, 2, 2, 611, 612, 3, 2, 2, 2, 612, 614, 3, 2, 2, 2, 613,
	615, 5, 18, 10, 2, 614, 613, 3, 2, 2, 2, 614, 615, 3, 2, 2, 2, 615, 617,
	3, 2, 2, 2, 616, 618, 5, 50, 26, 2, 617, 616, 3, 2, 2, 2, 617, 618, 3,
	2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 620, 7, 102, 2, 2, 620, 139, 3, 2,
	2, 2, 621, 622, 7, 78, 2, 2, 622, 628, 7, 113, 2, 2, 623, 625, 7, 99, 2,
	2, 624, 626, 5, 142, 72, 2, 625, 624, 3, 2, 2, 2, 625, 626, 3, 2, 2, 2,
	626, 627, 3, 2, 2, 2, 627, 629, 7, 100, 2, 2, 628, 623, 3, 2, 2, 2, 628,
	629, 3, 2, 2, 2, 629, 631, 3, 2, 2, 2, 630, 632, 5, 18, 10, 2, 631, 630,
	3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 633, 3, 2, 2, 2, 633, 634, 5, 20,
	11, 2, 634, 141, 3, 2, 2, 2, 635, 638, 5, 70, 36, 2, 636, 638, 5, 40, 21,
	2, 637, 635, 3, 2, 2, 2, 637, 636, 3, 2, 2, 2, 638, 143, 3, 2, 2, 2, 639,
	640, 7, 79, 2, 2, 640, 641, 7, 97, 2, 2, 641, 642, 11, 2, 2, 2, 642, 643,
	7, 98, 2, 2, 643, 145, 3, 2, 2, 2, 644, 645, 9, 12, 2, 2, 645, 147, 3,
	2, 2, 2, 646, 652, 7, 85, 2, 2, 647, 649, 7, 86, 2, 2, 648, 650, 7, 110,
	2, 2, 649, 648, 3, 2, 2, 2, 649, 650, 3, 2, 2, 2, 650, 652, 3, 2, 2, 2,
	651, 646, 3, 2, 2, 2, 651, 647, 3, 2, 2, 2, 652, 149, 3, 2, 2, 2, 653,
	654, 7, 87, 2, 2, 654, 655, 7, 113, 2, 2, 655, 661, 5, 82, 42, 2, 656,
	657, 7, 88, 2, 2, 657, 658, 5, 146, 74, 2, 658, 659, 5, 82, 42, 2, 659,
	661, 3, 2, 2, 2, 660, 653, 3, 2, 2, 2, 660, 656, 3, 2, 2, 2, 661, 151,
	3, 2, 2, 2, 662, 665, 5, 154, 78, 2, 663, 665, 7, 89, 2, 2, 664, 662, 3,
	2, 2, 2, 664, 663, 3, 2, 2, 2, 665, 153, 3, 2, 2, 2, 666, 668, 7, 113,
	2, 2, 667, 669, 5, 146, 74, 2, 668, 667, 3, 2, 2, 2, 668, 669, 3, 2, 2,
	2, 669, 675, 3, 2, 2, 2, 670, 671, 7, 56, 2, 2, 671, 672, 7, 99, 2, 2,
	672, 673, 7, 113, 2, 2, 673, 675, 7, 100, 2, 2, 674, 666, 3, 2, 2, 2, 674,
	670, 3, 2, 2, 2, 675, 155, 3, 2, 2, 2, 676, 677, 9, 13, 2, 2, 677, 157,
	3, 2, 2, 2, 678, 684, 5, 156, 79, 2, 679, 681, 7, 99, 2, 2, 680, 682, 5,
	110, 56, 2, 681, 680, 3, 2, 2, 2, 681, 682, 3, 2, 2, 2, 682, 683, 3, 2,
	2, 2, 683, 685, 7, 100, 2, 2, 684, 679, 3, 2, 2, 2, 684, 685, 3, 2, 2,
	2, 685, 686, 3, 2, 2, 2, 686, 689, 5, 22, 12, 2, 687, 690, 5, 76, 39, 2,
	688, 690, 5, 30, 16, 2, 689, 687, 3, 2, 2, 2, 689, 688, 3, 2, 2, 2, 690,
	159, 3, 2, 2, 2, 691, 692, 5, 158, 80, 2, 692, 693, 7, 102, 2, 2, 693,
	696, 3, 2, 2, 2, 694, 696, 5, 150, 76, 2, 695, 691, 3, 2, 2, 2, 695, 694,
	3, 2, 2, 2, 696, 161, 3, 2, 2, 2, 697, 700, 5, 164, 83, 2, 698, 700, 5,
	166, 84, 2, 699, 697, 3, 2, 2, 2, 699, 698, 3, 2, 2, 2, 700, 163, 3, 2,
	2, 2, 701, 702, 7, 92, 2, 2, 702, 703, 5, 168, 85, 2, 703, 704, 7, 102,
	2, 2, 704, 165, 3, 2, 2, 2, 705, 707, 7, 93, 2, 2, 706, 708, 5, 168, 85,
	2, 707, 706, 3, 2, 2, 2, 707, 708, 3, 2, 2, 2, 708, 709, 3, 2, 2, 2, 709,
	715, 7, 113, 2, 2, 710, 712, 7, 99, 2, 2, 711, 713, 5, 170, 86, 2, 712,
	711, 3, 2, 2, 2, 712, 713, 3, 2, 2, 2, 713, 714, 3, 2, 2, 2, 714, 716,
	7, 100, 2, 2, 715, 710, 3, 2, 2, 2, 715, 716, 3, 2, 2, 2, 716, 717, 3,
	2, 2, 2, 717, 719, 5, 26, 14, 2, 718, 720, 5, 18, 10, 2, 719, 718, 3, 2,
	2, 2, 719, 720, 3, 2, 2, 2, 720, 721, 3, 2, 2, 2, 721, 722, 7, 97, 2, 2,
	722, 723, 11, 2, 2, 2, 723, 724, 7, 98, 2, 2, 724, 167, 3, 2, 2, 2, 725,
	726, 9, 14, 2, 2, 726, 169, 3, 2, 2, 2, 727, 730, 5, 70, 36, 2, 728, 730,
	5, 110, 56, 2, 729, 727, 3, 2, 2, 2, 729, 728, 3, 2, 2, 2, 730, 171, 3,
	2, 2, 2, 75, 176, 180, 185, 192, 211, 217, 222, 234, 252, 261, 268, 283,
	292, 305, 309, 313, 318, 320, 341, 347, 351, 358, 371, 392, 396, 400, 404,
	415, 418, 426, 433, 440, 453, 464, 471, 474, 485, 496, 506, 512, 525, 527,
	536, 543, 551, 573, 577, 586, 595, 608, 611, 614, 617, 625, 628, 631, 637,
	649, 651, 660, 664, 668, 674, 681, 684, 689, 695, 699, 707, 712, 715, 719,
	729,
}
var literalNames = []string{
	"", "'OPENQASM'", "'include'", "'qubit'", "'qreg'", "'bit'", "'creg'",
	"'int'", "'uint'", "'float'", "'angle'", "'fixed'", "'bool'", "'const'",
	"'let'", "'||'", "'gate'", "'measure'", "'barrier'", "'inv'", "'pow'",
	"'ctrl'", "'@'", "'CX'", "'U'", "'reset'", "'~'", "'!'", "'+'", "'-'",
	"'*'", "'/'", "'<<'", "'>>'", "'rotl'", "'rotr'", "'&&'", "'&'", "'|'",
	"'^'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'return'", "'sin'",
	"'cos'", "'tan'", "'exp'", "'ln'", "'sqrt'", "'popcount'", "'lengthof'",
	"'++'", "'--'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='", "'~='",
	"'^='", "'<<='", "'>>='", "'in'", "'if'", "'else'", "'for'", "'while'",
	"'break'", "'continue'", "'end'", "'kernel'", "'def'", "'#pragma'", "'dt'",
	"'ns'", "'us'", "'ms'", "'s'", "'length'", "'stretch'", "'boxas'", "'boxto'",
	"'stretchinf'", "'delay'", "'rotary'", "'defcalgrammar'", "'defcal'", "'openpulse'",
	"'['", "']'", "'{'", "'}'", "'('", "')'", "':'", "';'", "'.'", "','", "'='",
	"'->'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "LPAREN", "RPAREN",
	"COLON", "SEMICOLON", "DOT", "COMMA", "ASSIGN", "ARROW", "Constant", "Whitespace",
	"Newline", "Integer", "Float", "RealNumber", "Identifier", "StringLiteral",
	"LineComment", "BlockComment",
}

var ruleNames = []string{
	"program", "header", "version", "include", "statement", "globalStatement",
	"declarationStatement", "comment", "returnSignature", "programBlock", "designator",
	"doubleDesignator", "identifierList", "indexIdentifier", "indexIdentifierList",
	"association", "quantumType", "quantumDeclaration", "quantumArgument",
	"quantumArgumentList", "bitType", "singleDesignatorType", "doubleDesignatorType",
	"noDesignatorType", "classicalType", "constantDeclaration", "singleDesignatorDeclaration",
	"doubleDesignatorDeclaration", "noDesignatorDeclaration", "bitDeclaration",
	"classicalVariableDeclaration", "classicalDeclaration", "classicalTypeList",
	"classicalArgument", "classicalArgumentList", "aliasStatement", "concatenateExpression",
	"rangeDefinition", "quantumGateDefinition", "quantumGateSignature", "quantumBlock",
	"quantumStatement", "quantumInstruction", "quantumMeasurement", "quantumMeasurementDeclaration",
	"quantumBarrier", "quantumGateModifier", "quantumGateCall", "quantumGateName",
	"unaryOperator", "binaryOperator", "expressionStatement", "expression",
	"expressionTerminator", "expressionList", "call", "builtInMath", "castOperator",
	"incrementor", "assignmentExpression", "assignmentOperator", "membershipTest",
	"setDeclaration", "loopBranchBlock", "branchingStatement", "loopStatement",
	"controlDirectiveStatement", "controlDirective", "kernelDeclaration", "subroutineDefinition",
	"subroutineArgumentList", "pragma", "timeUnit", "timingType", "timingBox",
	"timeTerminator", "timeIdentifier", "timeInstructionName", "timeInstruction",
	"timeStatement", "calibration", "calibrationGrammarDeclaration", "calibrationDefinition",
	"calibrationGrammar", "calibrationArgumentList",
}

type qasm3Parser struct {
	*antlr.BaseParser
}

// Newqasm3Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *qasm3Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newqasm3Parser(input antlr.TokenStream) *qasm3Parser {
	this := new(qasm3Parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "qasm3.g4"

	return this
}

// qasm3Parser tokens.
const (
	qasm3ParserEOF           = antlr.TokenEOF
	qasm3ParserT__0          = 1
	qasm3ParserT__1          = 2
	qasm3ParserT__2          = 3
	qasm3ParserT__3          = 4
	qasm3ParserT__4          = 5
	qasm3ParserT__5          = 6
	qasm3ParserT__6          = 7
	qasm3ParserT__7          = 8
	qasm3ParserT__8          = 9
	qasm3ParserT__9          = 10
	qasm3ParserT__10         = 11
	qasm3ParserT__11         = 12
	qasm3ParserT__12         = 13
	qasm3ParserT__13         = 14
	qasm3ParserT__14         = 15
	qasm3ParserT__15         = 16
	qasm3ParserT__16         = 17
	qasm3ParserT__17         = 18
	qasm3ParserT__18         = 19
	qasm3ParserT__19         = 20
	qasm3ParserT__20         = 21
	qasm3ParserT__21         = 22
	qasm3ParserT__22         = 23
	qasm3ParserT__23         = 24
	qasm3ParserT__24         = 25
	qasm3ParserT__25         = 26
	qasm3ParserT__26         = 27
	qasm3ParserT__27         = 28
	qasm3ParserT__28         = 29
	qasm3ParserT__29         = 30
	qasm3ParserT__30         = 31
	qasm3ParserT__31         = 32
	qasm3ParserT__32         = 33
	qasm3ParserT__33         = 34
	qasm3ParserT__34         = 35
	qasm3ParserT__35         = 36
	qasm3ParserT__36         = 37
	qasm3ParserT__37         = 38
	qasm3ParserT__38         = 39
	qasm3ParserT__39         = 40
	qasm3ParserT__40         = 41
	qasm3ParserT__41         = 42
	qasm3ParserT__42         = 43
	qasm3ParserT__43         = 44
	qasm3ParserT__44         = 45
	qasm3ParserT__45         = 46
	qasm3ParserT__46         = 47
	qasm3ParserT__47         = 48
	qasm3ParserT__48         = 49
	qasm3ParserT__49         = 50
	qasm3ParserT__50         = 51
	qasm3ParserT__51         = 52
	qasm3ParserT__52         = 53
	qasm3ParserT__53         = 54
	qasm3ParserT__54         = 55
	qasm3ParserT__55         = 56
	qasm3ParserT__56         = 57
	qasm3ParserT__57         = 58
	qasm3ParserT__58         = 59
	qasm3ParserT__59         = 60
	qasm3ParserT__60         = 61
	qasm3ParserT__61         = 62
	qasm3ParserT__62         = 63
	qasm3ParserT__63         = 64
	qasm3ParserT__64         = 65
	qasm3ParserT__65         = 66
	qasm3ParserT__66         = 67
	qasm3ParserT__67         = 68
	qasm3ParserT__68         = 69
	qasm3ParserT__69         = 70
	qasm3ParserT__70         = 71
	qasm3ParserT__71         = 72
	qasm3ParserT__72         = 73
	qasm3ParserT__73         = 74
	qasm3ParserT__74         = 75
	qasm3ParserT__75         = 76
	qasm3ParserT__76         = 77
	qasm3ParserT__77         = 78
	qasm3ParserT__78         = 79
	qasm3ParserT__79         = 80
	qasm3ParserT__80         = 81
	qasm3ParserT__81         = 82
	qasm3ParserT__82         = 83
	qasm3ParserT__83         = 84
	qasm3ParserT__84         = 85
	qasm3ParserT__85         = 86
	qasm3ParserT__86         = 87
	qasm3ParserT__87         = 88
	qasm3ParserT__88         = 89
	qasm3ParserT__89         = 90
	qasm3ParserT__90         = 91
	qasm3ParserT__91         = 92
	qasm3ParserLBRACKET      = 93
	qasm3ParserRBRACKET      = 94
	qasm3ParserLBRACE        = 95
	qasm3ParserRBRACE        = 96
	qasm3ParserLPAREN        = 97
	qasm3ParserRPAREN        = 98
	qasm3ParserCOLON         = 99
	qasm3ParserSEMICOLON     = 100
	qasm3ParserDOT           = 101
	qasm3ParserCOMMA         = 102
	qasm3ParserASSIGN        = 103
	qasm3ParserARROW         = 104
	qasm3ParserConstant      = 105
	qasm3ParserWhitespace    = 106
	qasm3ParserNewline       = 107
	qasm3ParserInteger       = 108
	qasm3ParserFloat         = 109
	qasm3ParserRealNumber    = 110
	qasm3ParserIdentifier    = 111
	qasm3ParserStringLiteral = 112
	qasm3ParserLineComment   = 113
	qasm3ParserBlockComment  = 114
)

// qasm3Parser rules.
const (
	qasm3ParserRULE_program                       = 0
	qasm3ParserRULE_header                        = 1
	qasm3ParserRULE_version                       = 2
	qasm3ParserRULE_include                       = 3
	qasm3ParserRULE_statement                     = 4
	qasm3ParserRULE_globalStatement               = 5
	qasm3ParserRULE_declarationStatement          = 6
	qasm3ParserRULE_comment                       = 7
	qasm3ParserRULE_returnSignature               = 8
	qasm3ParserRULE_programBlock                  = 9
	qasm3ParserRULE_designator                    = 10
	qasm3ParserRULE_doubleDesignator              = 11
	qasm3ParserRULE_identifierList                = 12
	qasm3ParserRULE_indexIdentifier               = 13
	qasm3ParserRULE_indexIdentifierList           = 14
	qasm3ParserRULE_association                   = 15
	qasm3ParserRULE_quantumType                   = 16
	qasm3ParserRULE_quantumDeclaration            = 17
	qasm3ParserRULE_quantumArgument               = 18
	qasm3ParserRULE_quantumArgumentList           = 19
	qasm3ParserRULE_bitType                       = 20
	qasm3ParserRULE_singleDesignatorType          = 21
	qasm3ParserRULE_doubleDesignatorType          = 22
	qasm3ParserRULE_noDesignatorType              = 23
	qasm3ParserRULE_classicalType                 = 24
	qasm3ParserRULE_constantDeclaration           = 25
	qasm3ParserRULE_singleDesignatorDeclaration   = 26
	qasm3ParserRULE_doubleDesignatorDeclaration   = 27
	qasm3ParserRULE_noDesignatorDeclaration       = 28
	qasm3ParserRULE_bitDeclaration                = 29
	qasm3ParserRULE_classicalVariableDeclaration  = 30
	qasm3ParserRULE_classicalDeclaration          = 31
	qasm3ParserRULE_classicalTypeList             = 32
	qasm3ParserRULE_classicalArgument             = 33
	qasm3ParserRULE_classicalArgumentList         = 34
	qasm3ParserRULE_aliasStatement                = 35
	qasm3ParserRULE_concatenateExpression         = 36
	qasm3ParserRULE_rangeDefinition               = 37
	qasm3ParserRULE_quantumGateDefinition         = 38
	qasm3ParserRULE_quantumGateSignature          = 39
	qasm3ParserRULE_quantumBlock                  = 40
	qasm3ParserRULE_quantumStatement              = 41
	qasm3ParserRULE_quantumInstruction            = 42
	qasm3ParserRULE_quantumMeasurement            = 43
	qasm3ParserRULE_quantumMeasurementDeclaration = 44
	qasm3ParserRULE_quantumBarrier                = 45
	qasm3ParserRULE_quantumGateModifier           = 46
	qasm3ParserRULE_quantumGateCall               = 47
	qasm3ParserRULE_quantumGateName               = 48
	qasm3ParserRULE_unaryOperator                 = 49
	qasm3ParserRULE_binaryOperator                = 50
	qasm3ParserRULE_expressionStatement           = 51
	qasm3ParserRULE_expression                    = 52
	qasm3ParserRULE_expressionTerminator          = 53
	qasm3ParserRULE_expressionList                = 54
	qasm3ParserRULE_call                          = 55
	qasm3ParserRULE_builtInMath                   = 56
	qasm3ParserRULE_castOperator                  = 57
	qasm3ParserRULE_incrementor                   = 58
	qasm3ParserRULE_assignmentExpression          = 59
	qasm3ParserRULE_assignmentOperator            = 60
	qasm3ParserRULE_membershipTest                = 61
	qasm3ParserRULE_setDeclaration                = 62
	qasm3ParserRULE_loopBranchBlock               = 63
	qasm3ParserRULE_branchingStatement            = 64
	qasm3ParserRULE_loopStatement                 = 65
	qasm3ParserRULE_controlDirectiveStatement     = 66
	qasm3ParserRULE_controlDirective              = 67
	qasm3ParserRULE_kernelDeclaration             = 68
	qasm3ParserRULE_subroutineDefinition          = 69
	qasm3ParserRULE_subroutineArgumentList        = 70
	qasm3ParserRULE_pragma                        = 71
	qasm3ParserRULE_timeUnit                      = 72
	qasm3ParserRULE_timingType                    = 73
	qasm3ParserRULE_timingBox                     = 74
	qasm3ParserRULE_timeTerminator                = 75
	qasm3ParserRULE_timeIdentifier                = 76
	qasm3ParserRULE_timeInstructionName           = 77
	qasm3ParserRULE_timeInstruction               = 78
	qasm3ParserRULE_timeStatement                 = 79
	qasm3ParserRULE_calibration                   = 80
	qasm3ParserRULE_calibrationGrammarDeclaration = 81
	qasm3ParserRULE_calibrationDefinition         = 82
	qasm3ParserRULE_calibrationGrammar            = 83
	qasm3ParserRULE_calibrationArgumentList       = 84
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
	p.RuleIndex = qasm3ParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

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
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *qasm3Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, qasm3ParserRULE_program)
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
		p.SetState(170)
		p.Header()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__2)|(1<<qasm3ParserT__3)|(1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__12)|(1<<qasm3ParserT__13)|(1<<qasm3ParserT__15)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__17)|(1<<qasm3ParserT__18)|(1<<qasm3ParserT__19)|(1<<qasm3ParserT__20)|(1<<qasm3ParserT__22)|(1<<qasm3ParserT__23)|(1<<qasm3ParserT__24)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(qasm3ParserT__45-46))|(1<<(qasm3ParserT__46-46))|(1<<(qasm3ParserT__47-46))|(1<<(qasm3ParserT__48-46))|(1<<(qasm3ParserT__49-46))|(1<<(qasm3ParserT__50-46))|(1<<(qasm3ParserT__51-46))|(1<<(qasm3ParserT__52-46))|(1<<(qasm3ParserT__53-46))|(1<<(qasm3ParserT__67-46))|(1<<(qasm3ParserT__69-46))|(1<<(qasm3ParserT__70-46))|(1<<(qasm3ParserT__71-46))|(1<<(qasm3ParserT__72-46))|(1<<(qasm3ParserT__73-46))|(1<<(qasm3ParserT__74-46))|(1<<(qasm3ParserT__75-46))|(1<<(qasm3ParserT__76-46)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__84-83))|(1<<(qasm3ParserT__85-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserT__87-83))|(1<<(qasm3ParserT__88-83))|(1<<(qasm3ParserT__90-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83))|(1<<(qasm3ParserLineComment-83))|(1<<(qasm3ParserBlockComment-83)))) != 0) {
		{
			p.SetState(171)
			p.Statement()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Version() IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *HeaderContext) AllInclude() []IIncludeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIncludeContext)(nil)).Elem())
	var tst = make([]IIncludeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIncludeContext)
		}
	}

	return tst
}

func (s *HeaderContext) Include(i int) IIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *qasm3Parser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, qasm3ParserRULE_header)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserT__0 {
		{
			p.SetState(177)
			p.Version()
		}

	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == qasm3ParserT__1 {
		{
			p.SetState(180)
			p.Include()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_version
	return p
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserInteger)
}

func (s *VersionContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserInteger, i)
}

func (s *VersionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *VersionContext) DOT() antlr.TerminalNode {
	return s.GetToken(qasm3ParserDOT, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *qasm3Parser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, qasm3ParserRULE_version)
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
		p.SetState(186)
		p.Match(qasm3ParserT__0)
	}
	{
		p.SetState(187)
		p.Match(qasm3ParserInteger)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserDOT {
		{
			p.SetState(188)
			p.Match(qasm3ParserDOT)
		}
		{
			p.SetState(189)
			p.Match(qasm3ParserInteger)
		}

	}
	{
		p.SetState(192)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(qasm3ParserStringLiteral, 0)
}

func (s *IncludeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitInclude(s)
	}
}

func (p *qasm3Parser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, qasm3ParserRULE_include)

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
		p.SetState(194)
		p.Match(qasm3ParserT__1)
	}
	{
		p.SetState(195)
		p.Match(qasm3ParserStringLiteral)
	}
	{
		p.SetState(196)
		p.Match(qasm3ParserSEMICOLON)
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
	p.RuleIndex = qasm3ParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GlobalStatement() IGlobalStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) DeclarationStatement() IDeclarationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationStatementContext)
}

func (s *StatementContext) BranchingStatement() IBranchingStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchingStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchingStatementContext)
}

func (s *StatementContext) LoopStatement() ILoopStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementContext) ControlDirectiveStatement() IControlDirectiveStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlDirectiveStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlDirectiveStatementContext)
}

func (s *StatementContext) AliasStatement() IAliasStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasStatementContext)
}

func (s *StatementContext) QuantumStatement() IQuantumStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumStatementContext)
}

func (s *StatementContext) TimeStatement() ITimeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeStatementContext)
}

func (s *StatementContext) Pragma() IPragmaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPragmaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPragmaContext)
}

func (s *StatementContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *qasm3Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, qasm3ParserRULE_statement)

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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.GlobalStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.ExpressionStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.DeclarationStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.BranchingStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.LoopStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(203)
			p.ControlDirectiveStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(204)
			p.AliasStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(205)
			p.QuantumStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(206)
			p.TimeStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(207)
			p.Pragma()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(208)
			p.Comment()
		}

	}

	return localctx
}

// IGlobalStatementContext is an interface to support dynamic dispatch.
type IGlobalStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalStatementContext differentiates from other interfaces.
	IsGlobalStatementContext()
}

type GlobalStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalStatementContext() *GlobalStatementContext {
	var p = new(GlobalStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_globalStatement
	return p
}

func (*GlobalStatementContext) IsGlobalStatementContext() {}

func NewGlobalStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalStatementContext {
	var p = new(GlobalStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_globalStatement

	return p
}

func (s *GlobalStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalStatementContext) SubroutineDefinition() ISubroutineDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutineDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutineDefinitionContext)
}

func (s *GlobalStatementContext) KernelDeclaration() IKernelDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKernelDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKernelDeclarationContext)
}

func (s *GlobalStatementContext) QuantumGateDefinition() IQuantumGateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateDefinitionContext)
}

func (s *GlobalStatementContext) CalibrationDefinition() ICalibrationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationDefinitionContext)
}

func (s *GlobalStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterGlobalStatement(s)
	}
}

func (s *GlobalStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitGlobalStatement(s)
	}
}

func (p *qasm3Parser) GlobalStatement() (localctx IGlobalStatementContext) {
	localctx = NewGlobalStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, qasm3ParserRULE_globalStatement)

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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__75:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.SubroutineDefinition()
		}

	case qasm3ParserT__74:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.KernelDeclaration()
		}

	case qasm3ParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(213)
			p.QuantumGateDefinition()
		}

	case qasm3ParserT__90:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(214)
			p.CalibrationDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationStatementContext is an interface to support dynamic dispatch.
type IDeclarationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationStatementContext differentiates from other interfaces.
	IsDeclarationStatementContext()
}

type DeclarationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationStatementContext() *DeclarationStatementContext {
	var p = new(DeclarationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_declarationStatement
	return p
}

func (*DeclarationStatementContext) IsDeclarationStatementContext() {}

func NewDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationStatementContext {
	var p = new(DeclarationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_declarationStatement

	return p
}

func (s *DeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *DeclarationStatementContext) QuantumDeclaration() IQuantumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumDeclarationContext)
}

func (s *DeclarationStatementContext) ClassicalDeclaration() IClassicalDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalDeclarationContext)
}

func (s *DeclarationStatementContext) ConstantDeclaration() IConstantDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantDeclarationContext)
}

func (s *DeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterDeclarationStatement(s)
	}
}

func (s *DeclarationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitDeclarationStatement(s)
	}
}

func (p *qasm3Parser) DeclarationStatement() (localctx IDeclarationStatementContext) {
	localctx = NewDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, qasm3ParserRULE_declarationStatement)

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
	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__2, qasm3ParserT__3:
		{
			p.SetState(217)
			p.QuantumDeclaration()
		}

	case qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__82, qasm3ParserT__83:
		{
			p.SetState(218)
			p.ClassicalDeclaration()
		}

	case qasm3ParserT__12:
		{
			p.SetState(219)
			p.ConstantDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(222)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LineComment() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLineComment, 0)
}

func (s *CommentContext) BlockComment() antlr.TerminalNode {
	return s.GetToken(qasm3ParserBlockComment, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *qasm3Parser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, qasm3ParserRULE_comment)
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
		p.SetState(224)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserLineComment || _la == qasm3ParserBlockComment) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReturnSignatureContext is an interface to support dynamic dispatch.
type IReturnSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnSignatureContext differentiates from other interfaces.
	IsReturnSignatureContext()
}

type ReturnSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnSignatureContext() *ReturnSignatureContext {
	var p = new(ReturnSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_returnSignature
	return p
}

func (*ReturnSignatureContext) IsReturnSignatureContext() {}

func NewReturnSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnSignatureContext {
	var p = new(ReturnSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_returnSignature

	return p
}

func (s *ReturnSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnSignatureContext) ARROW() antlr.TerminalNode {
	return s.GetToken(qasm3ParserARROW, 0)
}

func (s *ReturnSignatureContext) ClassicalDeclaration() IClassicalDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalDeclarationContext)
}

func (s *ReturnSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterReturnSignature(s)
	}
}

func (s *ReturnSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitReturnSignature(s)
	}
}

func (p *qasm3Parser) ReturnSignature() (localctx IReturnSignatureContext) {
	localctx = NewReturnSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, qasm3ParserRULE_returnSignature)

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
		p.SetState(226)
		p.Match(qasm3ParserARROW)
	}
	{
		p.SetState(227)
		p.ClassicalDeclaration()
	}

	return localctx
}

// IProgramBlockContext is an interface to support dynamic dispatch.
type IProgramBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramBlockContext differentiates from other interfaces.
	IsProgramBlockContext()
}

type ProgramBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBlockContext() *ProgramBlockContext {
	var p = new(ProgramBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_programBlock
	return p
}

func (*ProgramBlockContext) IsProgramBlockContext() {}

func NewProgramBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBlockContext {
	var p = new(ProgramBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_programBlock

	return p
}

func (s *ProgramBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACE, 0)
}

func (s *ProgramBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACE, 0)
}

func (s *ProgramBlockContext) ProgramBlock() IProgramBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramBlockContext)
}

func (s *ProgramBlockContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterProgramBlock(s)
	}
}

func (s *ProgramBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitProgramBlock(s)
	}
}

func (p *qasm3Parser) ProgramBlock() (localctx IProgramBlockContext) {
	localctx = NewProgramBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, qasm3ParserRULE_programBlock)

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
		p.SetState(229)
		p.Match(qasm3ParserLBRACE)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserLBRACE:
		{
			p.SetState(230)
			p.ProgramBlock()
		}

	case qasm3ParserT__2, qasm3ParserT__3, qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__12, qasm3ParserT__13, qasm3ParserT__15, qasm3ParserT__16, qasm3ParserT__17, qasm3ParserT__18, qasm3ParserT__19, qasm3ParserT__20, qasm3ParserT__22, qasm3ParserT__23, qasm3ParserT__24, qasm3ParserT__25, qasm3ParserT__26, qasm3ParserT__45, qasm3ParserT__46, qasm3ParserT__47, qasm3ParserT__48, qasm3ParserT__49, qasm3ParserT__50, qasm3ParserT__51, qasm3ParserT__52, qasm3ParserT__53, qasm3ParserT__67, qasm3ParserT__69, qasm3ParserT__70, qasm3ParserT__71, qasm3ParserT__72, qasm3ParserT__73, qasm3ParserT__74, qasm3ParserT__75, qasm3ParserT__76, qasm3ParserT__82, qasm3ParserT__83, qasm3ParserT__84, qasm3ParserT__85, qasm3ParserT__86, qasm3ParserT__87, qasm3ParserT__88, qasm3ParserT__90, qasm3ParserConstant, qasm3ParserInteger, qasm3ParserRealNumber, qasm3ParserIdentifier, qasm3ParserStringLiteral, qasm3ParserLineComment, qasm3ParserBlockComment:
		{
			p.SetState(231)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(234)
		p.Match(qasm3ParserRBRACE)
	}

	return localctx
}

// IDesignatorContext is an interface to support dynamic dispatch.
type IDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignatorContext differentiates from other interfaces.
	IsDesignatorContext()
}

type DesignatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorContext() *DesignatorContext {
	var p = new(DesignatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_designator
	return p
}

func (*DesignatorContext) IsDesignatorContext() {}

func NewDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorContext {
	var p = new(DesignatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_designator

	return p
}

func (s *DesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACKET, 0)
}

func (s *DesignatorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DesignatorContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACKET, 0)
}

func (s *DesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterDesignator(s)
	}
}

func (s *DesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitDesignator(s)
	}
}

func (p *qasm3Parser) Designator() (localctx IDesignatorContext) {
	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, qasm3ParserRULE_designator)

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
		p.SetState(236)
		p.Match(qasm3ParserLBRACKET)
	}
	{
		p.SetState(237)
		p.expression(0)
	}
	{
		p.SetState(238)
		p.Match(qasm3ParserRBRACKET)
	}

	return localctx
}

// IDoubleDesignatorContext is an interface to support dynamic dispatch.
type IDoubleDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleDesignatorContext differentiates from other interfaces.
	IsDoubleDesignatorContext()
}

type DoubleDesignatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleDesignatorContext() *DoubleDesignatorContext {
	var p = new(DoubleDesignatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_doubleDesignator
	return p
}

func (*DoubleDesignatorContext) IsDoubleDesignatorContext() {}

func NewDoubleDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleDesignatorContext {
	var p = new(DoubleDesignatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_doubleDesignator

	return p
}

func (s *DoubleDesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleDesignatorContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACKET, 0)
}

func (s *DoubleDesignatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DoubleDesignatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DoubleDesignatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, 0)
}

func (s *DoubleDesignatorContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACKET, 0)
}

func (s *DoubleDesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleDesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleDesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterDoubleDesignator(s)
	}
}

func (s *DoubleDesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitDoubleDesignator(s)
	}
}

func (p *qasm3Parser) DoubleDesignator() (localctx IDoubleDesignatorContext) {
	localctx = NewDoubleDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, qasm3ParserRULE_doubleDesignator)

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
		p.SetState(240)
		p.Match(qasm3ParserLBRACKET)
	}
	{
		p.SetState(241)
		p.expression(0)
	}
	{
		p.SetState(242)
		p.Match(qasm3ParserCOMMA)
	}
	{
		p.SetState(243)
		p.expression(0)
	}
	{
		p.SetState(244)
		p.Match(qasm3ParserRBRACKET)
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserIdentifier)
}

func (s *IdentifierListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (p *qasm3Parser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, qasm3ParserRULE_identifierList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)
				p.Match(qasm3ParserIdentifier)
			}
			{
				p.SetState(247)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	{
		p.SetState(253)
		p.Match(qasm3ParserIdentifier)
	}

	return localctx
}

// IIndexIdentifierContext is an interface to support dynamic dispatch.
type IIndexIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexIdentifierContext differentiates from other interfaces.
	IsIndexIdentifierContext()
}

type IndexIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexIdentifierContext() *IndexIdentifierContext {
	var p = new(IndexIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_indexIdentifier
	return p
}

func (*IndexIdentifierContext) IsIndexIdentifierContext() {}

func NewIndexIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexIdentifierContext {
	var p = new(IndexIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_indexIdentifier

	return p
}

func (s *IndexIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *IndexIdentifierContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *IndexIdentifierContext) DoubleDesignator() IDoubleDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorContext)
}

func (s *IndexIdentifierContext) RangeDefinition() IRangeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDefinitionContext)
}

func (s *IndexIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterIndexIdentifier(s)
	}
}

func (s *IndexIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitIndexIdentifier(s)
	}
}

func (p *qasm3Parser) IndexIdentifier() (localctx IIndexIdentifierContext) {
	localctx = NewIndexIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, qasm3ParserRULE_indexIdentifier)

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
		p.SetState(255)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(256)
			p.Designator()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(257)
			p.DoubleDesignator()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(258)
			p.RangeDefinition()
		}

	}

	return localctx
}

// IIndexIdentifierListContext is an interface to support dynamic dispatch.
type IIndexIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexIdentifierListContext differentiates from other interfaces.
	IsIndexIdentifierListContext()
}

type IndexIdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexIdentifierListContext() *IndexIdentifierListContext {
	var p = new(IndexIdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_indexIdentifierList
	return p
}

func (*IndexIdentifierListContext) IsIndexIdentifierListContext() {}

func NewIndexIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexIdentifierListContext {
	var p = new(IndexIdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_indexIdentifierList

	return p
}

func (s *IndexIdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexIdentifierListContext) AllIndexIdentifier() []IIndexIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexIdentifierContext)(nil)).Elem())
	var tst = make([]IIndexIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexIdentifierContext)
		}
	}

	return tst
}

func (s *IndexIdentifierListContext) IndexIdentifier(i int) IIndexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierContext)
}

func (s *IndexIdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *IndexIdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *IndexIdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexIdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexIdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterIndexIdentifierList(s)
	}
}

func (s *IndexIdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitIndexIdentifierList(s)
	}
}

func (p *qasm3Parser) IndexIdentifierList() (localctx IIndexIdentifierListContext) {
	localctx = NewIndexIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, qasm3ParserRULE_indexIdentifierList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(261)
				p.IndexIdentifier()
			}
			{
				p.SetState(262)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	{
		p.SetState(269)
		p.IndexIdentifier()
	}

	return localctx
}

// IAssociationContext is an interface to support dynamic dispatch.
type IAssociationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationContext differentiates from other interfaces.
	IsAssociationContext()
}

type AssociationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationContext() *AssociationContext {
	var p = new(AssociationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_association
	return p
}

func (*AssociationContext) IsAssociationContext() {}

func NewAssociationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationContext {
	var p = new(AssociationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_association

	return p
}

func (s *AssociationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationContext) COLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOLON, 0)
}

func (s *AssociationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *AssociationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterAssociation(s)
	}
}

func (s *AssociationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitAssociation(s)
	}
}

func (p *qasm3Parser) Association() (localctx IAssociationContext) {
	localctx = NewAssociationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, qasm3ParserRULE_association)

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
		p.SetState(271)
		p.Match(qasm3ParserCOLON)
	}
	{
		p.SetState(272)
		p.Match(qasm3ParserIdentifier)
	}

	return localctx
}

// IQuantumTypeContext is an interface to support dynamic dispatch.
type IQuantumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumTypeContext differentiates from other interfaces.
	IsQuantumTypeContext()
}

type QuantumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumTypeContext() *QuantumTypeContext {
	var p = new(QuantumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumType
	return p
}

func (*QuantumTypeContext) IsQuantumTypeContext() {}

func NewQuantumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumTypeContext {
	var p = new(QuantumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumType

	return p
}

func (s *QuantumTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *QuantumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumType(s)
	}
}

func (s *QuantumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumType(s)
	}
}

func (p *qasm3Parser) QuantumType() (localctx IQuantumTypeContext) {
	localctx = NewQuantumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, qasm3ParserRULE_quantumType)
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
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__2 || _la == qasm3ParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IQuantumDeclarationContext is an interface to support dynamic dispatch.
type IQuantumDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumDeclarationContext differentiates from other interfaces.
	IsQuantumDeclarationContext()
}

type QuantumDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumDeclarationContext() *QuantumDeclarationContext {
	var p = new(QuantumDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumDeclaration
	return p
}

func (*QuantumDeclarationContext) IsQuantumDeclarationContext() {}

func NewQuantumDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumDeclarationContext {
	var p = new(QuantumDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumDeclaration

	return p
}

func (s *QuantumDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumDeclarationContext) QuantumType() IQuantumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumTypeContext)
}

func (s *QuantumDeclarationContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *QuantumDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumDeclaration(s)
	}
}

func (s *QuantumDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumDeclaration(s)
	}
}

func (p *qasm3Parser) QuantumDeclaration() (localctx IQuantumDeclarationContext) {
	localctx = NewQuantumDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, qasm3ParserRULE_quantumDeclaration)

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
		p.SetState(276)
		p.QuantumType()
	}
	{
		p.SetState(277)
		p.IndexIdentifierList()
	}

	return localctx
}

// IQuantumArgumentContext is an interface to support dynamic dispatch.
type IQuantumArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumArgumentContext differentiates from other interfaces.
	IsQuantumArgumentContext()
}

type QuantumArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumArgumentContext() *QuantumArgumentContext {
	var p = new(QuantumArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumArgument
	return p
}

func (*QuantumArgumentContext) IsQuantumArgumentContext() {}

func NewQuantumArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumArgumentContext {
	var p = new(QuantumArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumArgument

	return p
}

func (s *QuantumArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumArgumentContext) QuantumType() IQuantumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumTypeContext)
}

func (s *QuantumArgumentContext) Association() IAssociationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationContext)
}

func (s *QuantumArgumentContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *QuantumArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumArgument(s)
	}
}

func (s *QuantumArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumArgument(s)
	}
}

func (p *qasm3Parser) QuantumArgument() (localctx IQuantumArgumentContext) {
	localctx = NewQuantumArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, qasm3ParserRULE_quantumArgument)
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
		p.SetState(279)
		p.QuantumType()
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLBRACKET {
		{
			p.SetState(280)
			p.Designator()
		}

	}
	{
		p.SetState(283)
		p.Association()
	}

	return localctx
}

// IQuantumArgumentListContext is an interface to support dynamic dispatch.
type IQuantumArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumArgumentListContext differentiates from other interfaces.
	IsQuantumArgumentListContext()
}

type QuantumArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumArgumentListContext() *QuantumArgumentListContext {
	var p = new(QuantumArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumArgumentList
	return p
}

func (*QuantumArgumentListContext) IsQuantumArgumentListContext() {}

func NewQuantumArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumArgumentListContext {
	var p = new(QuantumArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumArgumentList

	return p
}

func (s *QuantumArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumArgumentListContext) AllQuantumArgument() []IQuantumArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuantumArgumentContext)(nil)).Elem())
	var tst = make([]IQuantumArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuantumArgumentContext)
		}
	}

	return tst
}

func (s *QuantumArgumentListContext) QuantumArgument(i int) IQuantumArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuantumArgumentContext)
}

func (s *QuantumArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *QuantumArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *QuantumArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumArgumentList(s)
	}
}

func (s *QuantumArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumArgumentList(s)
	}
}

func (p *qasm3Parser) QuantumArgumentList() (localctx IQuantumArgumentListContext) {
	localctx = NewQuantumArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, qasm3ParserRULE_quantumArgumentList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(285)
				p.QuantumArgument()
			}
			{
				p.SetState(286)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	{
		p.SetState(293)
		p.QuantumArgument()
	}

	return localctx
}

// IBitTypeContext is an interface to support dynamic dispatch.
type IBitTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitTypeContext differentiates from other interfaces.
	IsBitTypeContext()
}

type BitTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitTypeContext() *BitTypeContext {
	var p = new(BitTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_bitType
	return p
}

func (*BitTypeContext) IsBitTypeContext() {}

func NewBitTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitTypeContext {
	var p = new(BitTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_bitType

	return p
}

func (s *BitTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *BitTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterBitType(s)
	}
}

func (s *BitTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitBitType(s)
	}
}

func (p *qasm3Parser) BitType() (localctx IBitTypeContext) {
	localctx = NewBitTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, qasm3ParserRULE_bitType)
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
		p.SetState(295)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__4 || _la == qasm3ParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISingleDesignatorTypeContext is an interface to support dynamic dispatch.
type ISingleDesignatorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleDesignatorTypeContext differentiates from other interfaces.
	IsSingleDesignatorTypeContext()
}

type SingleDesignatorTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleDesignatorTypeContext() *SingleDesignatorTypeContext {
	var p = new(SingleDesignatorTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_singleDesignatorType
	return p
}

func (*SingleDesignatorTypeContext) IsSingleDesignatorTypeContext() {}

func NewSingleDesignatorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleDesignatorTypeContext {
	var p = new(SingleDesignatorTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_singleDesignatorType

	return p
}

func (s *SingleDesignatorTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *SingleDesignatorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleDesignatorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleDesignatorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterSingleDesignatorType(s)
	}
}

func (s *SingleDesignatorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitSingleDesignatorType(s)
	}
}

func (p *qasm3Parser) SingleDesignatorType() (localctx ISingleDesignatorTypeContext) {
	localctx = NewSingleDesignatorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, qasm3ParserRULE_singleDesignatorType)
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
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDoubleDesignatorTypeContext is an interface to support dynamic dispatch.
type IDoubleDesignatorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleDesignatorTypeContext differentiates from other interfaces.
	IsDoubleDesignatorTypeContext()
}

type DoubleDesignatorTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleDesignatorTypeContext() *DoubleDesignatorTypeContext {
	var p = new(DoubleDesignatorTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_doubleDesignatorType
	return p
}

func (*DoubleDesignatorTypeContext) IsDoubleDesignatorTypeContext() {}

func NewDoubleDesignatorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleDesignatorTypeContext {
	var p = new(DoubleDesignatorTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_doubleDesignatorType

	return p
}

func (s *DoubleDesignatorTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *DoubleDesignatorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleDesignatorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleDesignatorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterDoubleDesignatorType(s)
	}
}

func (s *DoubleDesignatorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitDoubleDesignatorType(s)
	}
}

func (p *qasm3Parser) DoubleDesignatorType() (localctx IDoubleDesignatorTypeContext) {
	localctx = NewDoubleDesignatorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, qasm3ParserRULE_doubleDesignatorType)

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
		p.SetState(299)
		p.Match(qasm3ParserT__10)
	}

	return localctx
}

// INoDesignatorTypeContext is an interface to support dynamic dispatch.
type INoDesignatorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoDesignatorTypeContext differentiates from other interfaces.
	IsNoDesignatorTypeContext()
}

type NoDesignatorTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoDesignatorTypeContext() *NoDesignatorTypeContext {
	var p = new(NoDesignatorTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_noDesignatorType
	return p
}

func (*NoDesignatorTypeContext) IsNoDesignatorTypeContext() {}

func NewNoDesignatorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoDesignatorTypeContext {
	var p = new(NoDesignatorTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_noDesignatorType

	return p
}

func (s *NoDesignatorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NoDesignatorTypeContext) TimingType() ITimingTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimingTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimingTypeContext)
}

func (s *NoDesignatorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoDesignatorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoDesignatorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterNoDesignatorType(s)
	}
}

func (s *NoDesignatorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitNoDesignatorType(s)
	}
}

func (p *qasm3Parser) NoDesignatorType() (localctx INoDesignatorTypeContext) {
	localctx = NewNoDesignatorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, qasm3ParserRULE_noDesignatorType)

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

	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Match(qasm3ParserT__11)
		}

	case qasm3ParserT__82, qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.TimingType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassicalTypeContext is an interface to support dynamic dispatch.
type IClassicalTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalTypeContext differentiates from other interfaces.
	IsClassicalTypeContext()
}

type ClassicalTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalTypeContext() *ClassicalTypeContext {
	var p = new(ClassicalTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalType
	return p
}

func (*ClassicalTypeContext) IsClassicalTypeContext() {}

func NewClassicalTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalTypeContext {
	var p = new(ClassicalTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalType

	return p
}

func (s *ClassicalTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalTypeContext) SingleDesignatorType() ISingleDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleDesignatorTypeContext)
}

func (s *ClassicalTypeContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *ClassicalTypeContext) DoubleDesignatorType() IDoubleDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorTypeContext)
}

func (s *ClassicalTypeContext) DoubleDesignator() IDoubleDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorContext)
}

func (s *ClassicalTypeContext) NoDesignatorType() INoDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoDesignatorTypeContext)
}

func (s *ClassicalTypeContext) BitType() IBitTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitTypeContext)
}

func (s *ClassicalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalType(s)
	}
}

func (s *ClassicalTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalType(s)
	}
}

func (p *qasm3Parser) ClassicalType() (localctx IClassicalTypeContext) {
	localctx = NewClassicalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, qasm3ParserRULE_classicalType)
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

	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.SingleDesignatorType()
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == qasm3ParserLBRACKET {
			{
				p.SetState(306)
				p.Designator()
			}

		}

	case qasm3ParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.DoubleDesignatorType()
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == qasm3ParserLBRACKET {
			{
				p.SetState(310)
				p.DoubleDesignator()
			}

		}

	case qasm3ParserT__11, qasm3ParserT__82, qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.NoDesignatorType()
		}

	case qasm3ParserT__4, qasm3ParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(314)
			p.BitType()
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == qasm3ParserLBRACKET {
			{
				p.SetState(315)
				p.Designator()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantDeclarationContext is an interface to support dynamic dispatch.
type IConstantDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantDeclarationContext differentiates from other interfaces.
	IsConstantDeclarationContext()
}

type ConstantDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDeclarationContext() *ConstantDeclarationContext {
	var p = new(ConstantDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_constantDeclaration
	return p
}

func (*ConstantDeclarationContext) IsConstantDeclarationContext() {}

func NewConstantDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDeclarationContext {
	var p = new(ConstantDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_constantDeclaration

	return p
}

func (s *ConstantDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *ConstantDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserASSIGN, 0)
}

func (s *ConstantDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstantDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterConstantDeclaration(s)
	}
}

func (s *ConstantDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitConstantDeclaration(s)
	}
}

func (p *qasm3Parser) ConstantDeclaration() (localctx IConstantDeclarationContext) {
	localctx = NewConstantDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, qasm3ParserRULE_constantDeclaration)

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
		p.SetState(320)
		p.Match(qasm3ParserT__12)
	}
	{
		p.SetState(321)
		p.Match(qasm3ParserIdentifier)
	}
	{
		p.SetState(322)
		p.Match(qasm3ParserASSIGN)
	}
	{
		p.SetState(323)
		p.expression(0)
	}

	return localctx
}

// ISingleDesignatorDeclarationContext is an interface to support dynamic dispatch.
type ISingleDesignatorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleDesignatorDeclarationContext differentiates from other interfaces.
	IsSingleDesignatorDeclarationContext()
}

type SingleDesignatorDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleDesignatorDeclarationContext() *SingleDesignatorDeclarationContext {
	var p = new(SingleDesignatorDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_singleDesignatorDeclaration
	return p
}

func (*SingleDesignatorDeclarationContext) IsSingleDesignatorDeclarationContext() {}

func NewSingleDesignatorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleDesignatorDeclarationContext {
	var p = new(SingleDesignatorDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_singleDesignatorDeclaration

	return p
}

func (s *SingleDesignatorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleDesignatorDeclarationContext) SingleDesignatorType() ISingleDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleDesignatorTypeContext)
}

func (s *SingleDesignatorDeclarationContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *SingleDesignatorDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *SingleDesignatorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleDesignatorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleDesignatorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterSingleDesignatorDeclaration(s)
	}
}

func (s *SingleDesignatorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitSingleDesignatorDeclaration(s)
	}
}

func (p *qasm3Parser) SingleDesignatorDeclaration() (localctx ISingleDesignatorDeclarationContext) {
	localctx = NewSingleDesignatorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, qasm3ParserRULE_singleDesignatorDeclaration)

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
		p.SetState(325)
		p.SingleDesignatorType()
	}
	{
		p.SetState(326)
		p.Designator()
	}
	{
		p.SetState(327)
		p.Match(qasm3ParserIdentifier)
	}

	return localctx
}

// IDoubleDesignatorDeclarationContext is an interface to support dynamic dispatch.
type IDoubleDesignatorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleDesignatorDeclarationContext differentiates from other interfaces.
	IsDoubleDesignatorDeclarationContext()
}

type DoubleDesignatorDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleDesignatorDeclarationContext() *DoubleDesignatorDeclarationContext {
	var p = new(DoubleDesignatorDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_doubleDesignatorDeclaration
	return p
}

func (*DoubleDesignatorDeclarationContext) IsDoubleDesignatorDeclarationContext() {}

func NewDoubleDesignatorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleDesignatorDeclarationContext {
	var p = new(DoubleDesignatorDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_doubleDesignatorDeclaration

	return p
}

func (s *DoubleDesignatorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleDesignatorDeclarationContext) DoubleDesignatorType() IDoubleDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorTypeContext)
}

func (s *DoubleDesignatorDeclarationContext) DoubleDesignator() IDoubleDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorContext)
}

func (s *DoubleDesignatorDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *DoubleDesignatorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleDesignatorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleDesignatorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterDoubleDesignatorDeclaration(s)
	}
}

func (s *DoubleDesignatorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitDoubleDesignatorDeclaration(s)
	}
}

func (p *qasm3Parser) DoubleDesignatorDeclaration() (localctx IDoubleDesignatorDeclarationContext) {
	localctx = NewDoubleDesignatorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, qasm3ParserRULE_doubleDesignatorDeclaration)

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
		p.SetState(329)
		p.DoubleDesignatorType()
	}
	{
		p.SetState(330)
		p.DoubleDesignator()
	}
	{
		p.SetState(331)
		p.Match(qasm3ParserIdentifier)
	}

	return localctx
}

// INoDesignatorDeclarationContext is an interface to support dynamic dispatch.
type INoDesignatorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoDesignatorDeclarationContext differentiates from other interfaces.
	IsNoDesignatorDeclarationContext()
}

type NoDesignatorDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoDesignatorDeclarationContext() *NoDesignatorDeclarationContext {
	var p = new(NoDesignatorDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_noDesignatorDeclaration
	return p
}

func (*NoDesignatorDeclarationContext) IsNoDesignatorDeclarationContext() {}

func NewNoDesignatorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoDesignatorDeclarationContext {
	var p = new(NoDesignatorDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_noDesignatorDeclaration

	return p
}

func (s *NoDesignatorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NoDesignatorDeclarationContext) NoDesignatorType() INoDesignatorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoDesignatorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoDesignatorTypeContext)
}

func (s *NoDesignatorDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *NoDesignatorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoDesignatorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoDesignatorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterNoDesignatorDeclaration(s)
	}
}

func (s *NoDesignatorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitNoDesignatorDeclaration(s)
	}
}

func (p *qasm3Parser) NoDesignatorDeclaration() (localctx INoDesignatorDeclarationContext) {
	localctx = NewNoDesignatorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, qasm3ParserRULE_noDesignatorDeclaration)

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
		p.SetState(333)
		p.NoDesignatorType()
	}
	{
		p.SetState(334)
		p.Match(qasm3ParserIdentifier)
	}

	return localctx
}

// IBitDeclarationContext is an interface to support dynamic dispatch.
type IBitDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitDeclarationContext differentiates from other interfaces.
	IsBitDeclarationContext()
}

type BitDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitDeclarationContext() *BitDeclarationContext {
	var p = new(BitDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_bitDeclaration
	return p
}

func (*BitDeclarationContext) IsBitDeclarationContext() {}

func NewBitDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitDeclarationContext {
	var p = new(BitDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_bitDeclaration

	return p
}

func (s *BitDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *BitDeclarationContext) BitType() IBitTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitTypeContext)
}

func (s *BitDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *BitDeclarationContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *BitDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterBitDeclaration(s)
	}
}

func (s *BitDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitBitDeclaration(s)
	}
}

func (p *qasm3Parser) BitDeclaration() (localctx IBitDeclarationContext) {
	localctx = NewBitDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, qasm3ParserRULE_bitDeclaration)
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
		p.SetState(336)
		p.BitType()
	}
	{
		p.SetState(337)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLBRACKET {
		{
			p.SetState(338)
			p.Designator()
		}

	}

	return localctx
}

// IClassicalVariableDeclarationContext is an interface to support dynamic dispatch.
type IClassicalVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalVariableDeclarationContext differentiates from other interfaces.
	IsClassicalVariableDeclarationContext()
}

type ClassicalVariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalVariableDeclarationContext() *ClassicalVariableDeclarationContext {
	var p = new(ClassicalVariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalVariableDeclaration
	return p
}

func (*ClassicalVariableDeclarationContext) IsClassicalVariableDeclarationContext() {}

func NewClassicalVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalVariableDeclarationContext {
	var p = new(ClassicalVariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalVariableDeclaration

	return p
}

func (s *ClassicalVariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalVariableDeclarationContext) SingleDesignatorDeclaration() ISingleDesignatorDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleDesignatorDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleDesignatorDeclarationContext)
}

func (s *ClassicalVariableDeclarationContext) DoubleDesignatorDeclaration() IDoubleDesignatorDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleDesignatorDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleDesignatorDeclarationContext)
}

func (s *ClassicalVariableDeclarationContext) NoDesignatorDeclaration() INoDesignatorDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoDesignatorDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoDesignatorDeclarationContext)
}

func (s *ClassicalVariableDeclarationContext) BitDeclaration() IBitDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitDeclarationContext)
}

func (s *ClassicalVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalVariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalVariableDeclaration(s)
	}
}

func (s *ClassicalVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalVariableDeclaration(s)
	}
}

func (p *qasm3Parser) ClassicalVariableDeclaration() (localctx IClassicalVariableDeclarationContext) {
	localctx = NewClassicalVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, qasm3ParserRULE_classicalVariableDeclaration)

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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.SingleDesignatorDeclaration()
		}

	case qasm3ParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.DoubleDesignatorDeclaration()
		}

	case qasm3ParserT__11, qasm3ParserT__82, qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.NoDesignatorDeclaration()
		}

	case qasm3ParserT__4, qasm3ParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(344)
			p.BitDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassicalDeclarationContext is an interface to support dynamic dispatch.
type IClassicalDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalDeclarationContext differentiates from other interfaces.
	IsClassicalDeclarationContext()
}

type ClassicalDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalDeclarationContext() *ClassicalDeclarationContext {
	var p = new(ClassicalDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalDeclaration
	return p
}

func (*ClassicalDeclarationContext) IsClassicalDeclarationContext() {}

func NewClassicalDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalDeclarationContext {
	var p = new(ClassicalDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalDeclaration

	return p
}

func (s *ClassicalDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalDeclarationContext) ClassicalVariableDeclaration() IClassicalVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalVariableDeclarationContext)
}

func (s *ClassicalDeclarationContext) AssignmentExpression() IAssignmentExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ClassicalDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalDeclaration(s)
	}
}

func (s *ClassicalDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalDeclaration(s)
	}
}

func (p *qasm3Parser) ClassicalDeclaration() (localctx IClassicalDeclarationContext) {
	localctx = NewClassicalDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, qasm3ParserRULE_classicalDeclaration)
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
		p.SetState(347)
		p.ClassicalVariableDeclaration()
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(qasm3ParserT__56-57))|(1<<(qasm3ParserT__57-57))|(1<<(qasm3ParserT__58-57))|(1<<(qasm3ParserT__59-57))|(1<<(qasm3ParserT__60-57))|(1<<(qasm3ParserT__61-57))|(1<<(qasm3ParserT__62-57))|(1<<(qasm3ParserT__63-57))|(1<<(qasm3ParserT__64-57))|(1<<(qasm3ParserT__65-57)))) != 0) || _la == qasm3ParserASSIGN || _la == qasm3ParserARROW {
		{
			p.SetState(348)
			p.AssignmentExpression()
		}

	}

	return localctx
}

// IClassicalTypeListContext is an interface to support dynamic dispatch.
type IClassicalTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalTypeListContext differentiates from other interfaces.
	IsClassicalTypeListContext()
}

type ClassicalTypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalTypeListContext() *ClassicalTypeListContext {
	var p = new(ClassicalTypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalTypeList
	return p
}

func (*ClassicalTypeListContext) IsClassicalTypeListContext() {}

func NewClassicalTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalTypeListContext {
	var p = new(ClassicalTypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalTypeList

	return p
}

func (s *ClassicalTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalTypeListContext) AllClassicalType() []IClassicalTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassicalTypeContext)(nil)).Elem())
	var tst = make([]IClassicalTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassicalTypeContext)
		}
	}

	return tst
}

func (s *ClassicalTypeListContext) ClassicalType(i int) IClassicalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassicalTypeContext)
}

func (s *ClassicalTypeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *ClassicalTypeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *ClassicalTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalTypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalTypeList(s)
	}
}

func (s *ClassicalTypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalTypeList(s)
	}
}

func (p *qasm3Parser) ClassicalTypeList() (localctx IClassicalTypeListContext) {
	localctx = NewClassicalTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, qasm3ParserRULE_classicalTypeList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(351)
				p.ClassicalType()
			}
			{
				p.SetState(352)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	{
		p.SetState(359)
		p.ClassicalType()
	}

	return localctx
}

// IClassicalArgumentContext is an interface to support dynamic dispatch.
type IClassicalArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalArgumentContext differentiates from other interfaces.
	IsClassicalArgumentContext()
}

type ClassicalArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalArgumentContext() *ClassicalArgumentContext {
	var p = new(ClassicalArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalArgument
	return p
}

func (*ClassicalArgumentContext) IsClassicalArgumentContext() {}

func NewClassicalArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalArgumentContext {
	var p = new(ClassicalArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalArgument

	return p
}

func (s *ClassicalArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalArgumentContext) ClassicalType() IClassicalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalTypeContext)
}

func (s *ClassicalArgumentContext) Association() IAssociationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationContext)
}

func (s *ClassicalArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalArgument(s)
	}
}

func (s *ClassicalArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalArgument(s)
	}
}

func (p *qasm3Parser) ClassicalArgument() (localctx IClassicalArgumentContext) {
	localctx = NewClassicalArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, qasm3ParserRULE_classicalArgument)

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
		p.SetState(361)
		p.ClassicalType()
	}
	{
		p.SetState(362)
		p.Association()
	}

	return localctx
}

// IClassicalArgumentListContext is an interface to support dynamic dispatch.
type IClassicalArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassicalArgumentListContext differentiates from other interfaces.
	IsClassicalArgumentListContext()
}

type ClassicalArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassicalArgumentListContext() *ClassicalArgumentListContext {
	var p = new(ClassicalArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_classicalArgumentList
	return p
}

func (*ClassicalArgumentListContext) IsClassicalArgumentListContext() {}

func NewClassicalArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicalArgumentListContext {
	var p = new(ClassicalArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_classicalArgumentList

	return p
}

func (s *ClassicalArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicalArgumentListContext) AllClassicalArgument() []IClassicalArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassicalArgumentContext)(nil)).Elem())
	var tst = make([]IClassicalArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassicalArgumentContext)
		}
	}

	return tst
}

func (s *ClassicalArgumentListContext) ClassicalArgument(i int) IClassicalArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassicalArgumentContext)
}

func (s *ClassicalArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *ClassicalArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *ClassicalArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicalArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicalArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterClassicalArgumentList(s)
	}
}

func (s *ClassicalArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitClassicalArgumentList(s)
	}
}

func (p *qasm3Parser) ClassicalArgumentList() (localctx IClassicalArgumentListContext) {
	localctx = NewClassicalArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, qasm3ParserRULE_classicalArgumentList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(364)
				p.ClassicalArgument()
			}
			{
				p.SetState(365)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	{
		p.SetState(372)
		p.ClassicalArgument()
	}

	return localctx
}

// IAliasStatementContext is an interface to support dynamic dispatch.
type IAliasStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasStatementContext differentiates from other interfaces.
	IsAliasStatementContext()
}

type AliasStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasStatementContext() *AliasStatementContext {
	var p = new(AliasStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_aliasStatement
	return p
}

func (*AliasStatementContext) IsAliasStatementContext() {}

func NewAliasStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasStatementContext {
	var p = new(AliasStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_aliasStatement

	return p
}

func (s *AliasStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *AliasStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserASSIGN, 0)
}

func (s *AliasStatementContext) ConcatenateExpression() IConcatenateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenateExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenateExpressionContext)
}

func (s *AliasStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterAliasStatement(s)
	}
}

func (s *AliasStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitAliasStatement(s)
	}
}

func (p *qasm3Parser) AliasStatement() (localctx IAliasStatementContext) {
	localctx = NewAliasStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, qasm3ParserRULE_aliasStatement)

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
		p.SetState(374)
		p.Match(qasm3ParserT__13)
	}
	{
		p.SetState(375)
		p.Match(qasm3ParserIdentifier)
	}
	{
		p.SetState(376)
		p.Match(qasm3ParserASSIGN)
	}
	{
		p.SetState(377)
		p.ConcatenateExpression()
	}

	return localctx
}

// IConcatenateExpressionContext is an interface to support dynamic dispatch.
type IConcatenateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatenateExpressionContext differentiates from other interfaces.
	IsConcatenateExpressionContext()
}

type ConcatenateExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatenateExpressionContext() *ConcatenateExpressionContext {
	var p = new(ConcatenateExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_concatenateExpression
	return p
}

func (*ConcatenateExpressionContext) IsConcatenateExpressionContext() {}

func NewConcatenateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatenateExpressionContext {
	var p = new(ConcatenateExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_concatenateExpression

	return p
}

func (s *ConcatenateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatenateExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserASSIGN, 0)
}

func (s *ConcatenateExpressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserIdentifier)
}

func (s *ConcatenateExpressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, i)
}

func (s *ConcatenateExpressionContext) RangeDefinition() IRangeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDefinitionContext)
}

func (s *ConcatenateExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACKET, 0)
}

func (s *ConcatenateExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ConcatenateExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACKET, 0)
}

func (s *ConcatenateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatenateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterConcatenateExpression(s)
	}
}

func (s *ConcatenateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitConcatenateExpression(s)
	}
}

func (p *qasm3Parser) ConcatenateExpression() (localctx IConcatenateExpressionContext) {
	localctx = NewConcatenateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, qasm3ParserRULE_concatenateExpression)

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
		p.SetState(379)
		p.Match(qasm3ParserASSIGN)
	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(380)
			p.Match(qasm3ParserIdentifier)
		}
		{
			p.SetState(381)
			p.RangeDefinition()
		}

	case 2:
		{
			p.SetState(382)
			p.Match(qasm3ParserIdentifier)
		}
		{
			p.SetState(383)
			p.Match(qasm3ParserT__14)
		}
		{
			p.SetState(384)
			p.Match(qasm3ParserIdentifier)
		}

	case 3:
		{
			p.SetState(385)
			p.Match(qasm3ParserIdentifier)
		}
		{
			p.SetState(386)
			p.Match(qasm3ParserLBRACKET)
		}
		{
			p.SetState(387)
			p.ExpressionList()
		}
		{
			p.SetState(388)
			p.Match(qasm3ParserRBRACKET)
		}

	}

	return localctx
}

// IRangeDefinitionContext is an interface to support dynamic dispatch.
type IRangeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeDefinitionContext differentiates from other interfaces.
	IsRangeDefinitionContext()
}

type RangeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeDefinitionContext() *RangeDefinitionContext {
	var p = new(RangeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_rangeDefinition
	return p
}

func (*RangeDefinitionContext) IsRangeDefinitionContext() {}

func NewRangeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeDefinitionContext {
	var p = new(RangeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_rangeDefinition

	return p
}

func (s *RangeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeDefinitionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACKET, 0)
}

func (s *RangeDefinitionContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOLON)
}

func (s *RangeDefinitionContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOLON, i)
}

func (s *RangeDefinitionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACKET, 0)
}

func (s *RangeDefinitionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RangeDefinitionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RangeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterRangeDefinition(s)
	}
}

func (s *RangeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitRangeDefinition(s)
	}
}

func (p *qasm3Parser) RangeDefinition() (localctx IRangeDefinitionContext) {
	localctx = NewRangeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, qasm3ParserRULE_rangeDefinition)
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
		p.SetState(392)
		p.Match(qasm3ParserLBRACKET)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
		{
			p.SetState(393)
			p.expression(0)
		}

	}
	{
		p.SetState(396)
		p.Match(qasm3ParserCOLON)
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
		{
			p.SetState(397)
			p.expression(0)
		}

	}
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserCOLON {
		{
			p.SetState(400)
			p.Match(qasm3ParserCOLON)
		}
		{
			p.SetState(401)
			p.expression(0)
		}

	}
	{
		p.SetState(404)
		p.Match(qasm3ParserRBRACKET)
	}

	return localctx
}

// IQuantumGateDefinitionContext is an interface to support dynamic dispatch.
type IQuantumGateDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumGateDefinitionContext differentiates from other interfaces.
	IsQuantumGateDefinitionContext()
}

type QuantumGateDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumGateDefinitionContext() *QuantumGateDefinitionContext {
	var p = new(QuantumGateDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumGateDefinition
	return p
}

func (*QuantumGateDefinitionContext) IsQuantumGateDefinitionContext() {}

func NewQuantumGateDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumGateDefinitionContext {
	var p = new(QuantumGateDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumGateDefinition

	return p
}

func (s *QuantumGateDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumGateDefinitionContext) QuantumGateSignature() IQuantumGateSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateSignatureContext)
}

func (s *QuantumGateDefinitionContext) QuantumBlock() IQuantumBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumBlockContext)
}

func (s *QuantumGateDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumGateDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumGateDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumGateDefinition(s)
	}
}

func (s *QuantumGateDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumGateDefinition(s)
	}
}

func (p *qasm3Parser) QuantumGateDefinition() (localctx IQuantumGateDefinitionContext) {
	localctx = NewQuantumGateDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, qasm3ParserRULE_quantumGateDefinition)

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
		p.SetState(406)
		p.Match(qasm3ParserT__15)
	}
	{
		p.SetState(407)
		p.QuantumGateSignature()
	}
	{
		p.SetState(408)
		p.QuantumBlock()
	}

	return localctx
}

// IQuantumGateSignatureContext is an interface to support dynamic dispatch.
type IQuantumGateSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumGateSignatureContext differentiates from other interfaces.
	IsQuantumGateSignatureContext()
}

type QuantumGateSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumGateSignatureContext() *QuantumGateSignatureContext {
	var p = new(QuantumGateSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumGateSignature
	return p
}

func (*QuantumGateSignatureContext) IsQuantumGateSignatureContext() {}

func NewQuantumGateSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumGateSignatureContext {
	var p = new(QuantumGateSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumGateSignature

	return p
}

func (s *QuantumGateSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumGateSignatureContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *QuantumGateSignatureContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *QuantumGateSignatureContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *QuantumGateSignatureContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *QuantumGateSignatureContext) ClassicalArgumentList() IClassicalArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalArgumentListContext)
}

func (s *QuantumGateSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumGateSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumGateSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumGateSignature(s)
	}
}

func (s *QuantumGateSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumGateSignature(s)
	}
}

func (p *qasm3Parser) QuantumGateSignature() (localctx IQuantumGateSignatureContext) {
	localctx = NewQuantumGateSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, qasm3ParserRULE_quantumGateSignature)
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
		p.SetState(410)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(411)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11))) != 0) || _la == qasm3ParserT__82 || _la == qasm3ParserT__83 {
			{
				p.SetState(412)
				p.ClassicalArgumentList()
			}

		}
		{
			p.SetState(415)
			p.Match(qasm3ParserRPAREN)
		}

	}
	{
		p.SetState(418)
		p.IdentifierList()
	}

	return localctx
}

// IQuantumBlockContext is an interface to support dynamic dispatch.
type IQuantumBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumBlockContext differentiates from other interfaces.
	IsQuantumBlockContext()
}

type QuantumBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumBlockContext() *QuantumBlockContext {
	var p = new(QuantumBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumBlock
	return p
}

func (*QuantumBlockContext) IsQuantumBlockContext() {}

func NewQuantumBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumBlockContext {
	var p = new(QuantumBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumBlock

	return p
}

func (s *QuantumBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACE, 0)
}

func (s *QuantumBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACE, 0)
}

func (s *QuantumBlockContext) AllQuantumStatement() []IQuantumStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuantumStatementContext)(nil)).Elem())
	var tst = make([]IQuantumStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuantumStatementContext)
		}
	}

	return tst
}

func (s *QuantumBlockContext) QuantumStatement(i int) IQuantumStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuantumStatementContext)
}

func (s *QuantumBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumBlock(s)
	}
}

func (s *QuantumBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumBlock(s)
	}
}

func (p *qasm3Parser) QuantumBlock() (localctx IQuantumBlockContext) {
	localctx = NewQuantumBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, qasm3ParserRULE_quantumBlock)
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
		p.SetState(420)
		p.Match(qasm3ParserLBRACE)
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__16)|(1<<qasm3ParserT__17)|(1<<qasm3ParserT__18)|(1<<qasm3ParserT__19)|(1<<qasm3ParserT__20)|(1<<qasm3ParserT__22)|(1<<qasm3ParserT__23)|(1<<qasm3ParserT__24))) != 0) || _la == qasm3ParserIdentifier {
		{
			p.SetState(421)
			p.QuantumStatement()
		}

		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(427)
		p.Match(qasm3ParserRBRACE)
	}

	return localctx
}

// IQuantumStatementContext is an interface to support dynamic dispatch.
type IQuantumStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumStatementContext differentiates from other interfaces.
	IsQuantumStatementContext()
}

type QuantumStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumStatementContext() *QuantumStatementContext {
	var p = new(QuantumStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumStatement
	return p
}

func (*QuantumStatementContext) IsQuantumStatementContext() {}

func NewQuantumStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumStatementContext {
	var p = new(QuantumStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumStatement

	return p
}

func (s *QuantumStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *QuantumStatementContext) QuantumInstruction() IQuantumInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumInstructionContext)
}

func (s *QuantumStatementContext) QuantumMeasurementDeclaration() IQuantumMeasurementDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumMeasurementDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumMeasurementDeclarationContext)
}

func (s *QuantumStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumStatement(s)
	}
}

func (s *QuantumStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumStatement(s)
	}
}

func (p *qasm3Parser) QuantumStatement() (localctx IQuantumStatementContext) {
	localctx = NewQuantumStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, qasm3ParserRULE_quantumStatement)

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
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(429)
			p.QuantumInstruction()
		}

	case 2:
		{
			p.SetState(430)
			p.QuantumMeasurementDeclaration()
		}

	}
	{
		p.SetState(433)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// IQuantumInstructionContext is an interface to support dynamic dispatch.
type IQuantumInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumInstructionContext differentiates from other interfaces.
	IsQuantumInstructionContext()
}

type QuantumInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumInstructionContext() *QuantumInstructionContext {
	var p = new(QuantumInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumInstruction
	return p
}

func (*QuantumInstructionContext) IsQuantumInstructionContext() {}

func NewQuantumInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumInstructionContext {
	var p = new(QuantumInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumInstruction

	return p
}

func (s *QuantumInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumInstructionContext) QuantumGateCall() IQuantumGateCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateCallContext)
}

func (s *QuantumInstructionContext) QuantumMeasurement() IQuantumMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumMeasurementContext)
}

func (s *QuantumInstructionContext) QuantumBarrier() IQuantumBarrierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumBarrierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumBarrierContext)
}

func (s *QuantumInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumInstruction(s)
	}
}

func (s *QuantumInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumInstruction(s)
	}
}

func (p *qasm3Parser) QuantumInstruction() (localctx IQuantumInstructionContext) {
	localctx = NewQuantumInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, qasm3ParserRULE_quantumInstruction)

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

	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__18, qasm3ParserT__19, qasm3ParserT__20, qasm3ParserT__22, qasm3ParserT__23, qasm3ParserT__24, qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.QuantumGateCall()
		}

	case qasm3ParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.QuantumMeasurement()
		}

	case qasm3ParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.QuantumBarrier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuantumMeasurementContext is an interface to support dynamic dispatch.
type IQuantumMeasurementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumMeasurementContext differentiates from other interfaces.
	IsQuantumMeasurementContext()
}

type QuantumMeasurementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumMeasurementContext() *QuantumMeasurementContext {
	var p = new(QuantumMeasurementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumMeasurement
	return p
}

func (*QuantumMeasurementContext) IsQuantumMeasurementContext() {}

func NewQuantumMeasurementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumMeasurementContext {
	var p = new(QuantumMeasurementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumMeasurement

	return p
}

func (s *QuantumMeasurementContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumMeasurementContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *QuantumMeasurementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumMeasurementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumMeasurementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumMeasurement(s)
	}
}

func (s *QuantumMeasurementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumMeasurement(s)
	}
}

func (p *qasm3Parser) QuantumMeasurement() (localctx IQuantumMeasurementContext) {
	localctx = NewQuantumMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, qasm3ParserRULE_quantumMeasurement)

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
		p.SetState(440)
		p.Match(qasm3ParserT__16)
	}
	{
		p.SetState(441)
		p.IndexIdentifierList()
	}

	return localctx
}

// IQuantumMeasurementDeclarationContext is an interface to support dynamic dispatch.
type IQuantumMeasurementDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumMeasurementDeclarationContext differentiates from other interfaces.
	IsQuantumMeasurementDeclarationContext()
}

type QuantumMeasurementDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumMeasurementDeclarationContext() *QuantumMeasurementDeclarationContext {
	var p = new(QuantumMeasurementDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumMeasurementDeclaration
	return p
}

func (*QuantumMeasurementDeclarationContext) IsQuantumMeasurementDeclarationContext() {}

func NewQuantumMeasurementDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumMeasurementDeclarationContext {
	var p = new(QuantumMeasurementDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumMeasurementDeclaration

	return p
}

func (s *QuantumMeasurementDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumMeasurementDeclarationContext) QuantumMeasurement() IQuantumMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumMeasurementContext)
}

func (s *QuantumMeasurementDeclarationContext) ARROW() antlr.TerminalNode {
	return s.GetToken(qasm3ParserARROW, 0)
}

func (s *QuantumMeasurementDeclarationContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *QuantumMeasurementDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserASSIGN, 0)
}

func (s *QuantumMeasurementDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumMeasurementDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumMeasurementDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumMeasurementDeclaration(s)
	}
}

func (s *QuantumMeasurementDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumMeasurementDeclaration(s)
	}
}

func (p *qasm3Parser) QuantumMeasurementDeclaration() (localctx IQuantumMeasurementDeclarationContext) {
	localctx = NewQuantumMeasurementDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, qasm3ParserRULE_quantumMeasurementDeclaration)

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

	p.SetState(451)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.QuantumMeasurement()
		}
		{
			p.SetState(444)
			p.Match(qasm3ParserARROW)
		}
		{
			p.SetState(445)
			p.IndexIdentifierList()
		}

	case qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(447)
			p.IndexIdentifierList()
		}
		{
			p.SetState(448)
			p.Match(qasm3ParserASSIGN)
		}
		{
			p.SetState(449)
			p.QuantumMeasurement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuantumBarrierContext is an interface to support dynamic dispatch.
type IQuantumBarrierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumBarrierContext differentiates from other interfaces.
	IsQuantumBarrierContext()
}

type QuantumBarrierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumBarrierContext() *QuantumBarrierContext {
	var p = new(QuantumBarrierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumBarrier
	return p
}

func (*QuantumBarrierContext) IsQuantumBarrierContext() {}

func NewQuantumBarrierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumBarrierContext {
	var p = new(QuantumBarrierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumBarrier

	return p
}

func (s *QuantumBarrierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumBarrierContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *QuantumBarrierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumBarrierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumBarrierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumBarrier(s)
	}
}

func (s *QuantumBarrierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumBarrier(s)
	}
}

func (p *qasm3Parser) QuantumBarrier() (localctx IQuantumBarrierContext) {
	localctx = NewQuantumBarrierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, qasm3ParserRULE_quantumBarrier)

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
		p.SetState(453)
		p.Match(qasm3ParserT__17)
	}
	{
		p.SetState(454)
		p.IndexIdentifierList()
	}

	return localctx
}

// IQuantumGateModifierContext is an interface to support dynamic dispatch.
type IQuantumGateModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumGateModifierContext differentiates from other interfaces.
	IsQuantumGateModifierContext()
}

type QuantumGateModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumGateModifierContext() *QuantumGateModifierContext {
	var p = new(QuantumGateModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumGateModifier
	return p
}

func (*QuantumGateModifierContext) IsQuantumGateModifierContext() {}

func NewQuantumGateModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumGateModifierContext {
	var p = new(QuantumGateModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumGateModifier

	return p
}

func (s *QuantumGateModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumGateModifierContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *QuantumGateModifierContext) Integer() antlr.TerminalNode {
	return s.GetToken(qasm3ParserInteger, 0)
}

func (s *QuantumGateModifierContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *QuantumGateModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumGateModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumGateModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumGateModifier(s)
	}
}

func (s *QuantumGateModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumGateModifier(s)
	}
}

func (p *qasm3Parser) QuantumGateModifier() (localctx IQuantumGateModifierContext) {
	localctx = NewQuantumGateModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, qasm3ParserRULE_quantumGateModifier)

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
	p.SetState(462)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__18:
		{
			p.SetState(456)
			p.Match(qasm3ParserT__18)
		}

	case qasm3ParserT__19:
		{
			p.SetState(457)
			p.Match(qasm3ParserT__19)
		}
		{
			p.SetState(458)
			p.Match(qasm3ParserLPAREN)
		}
		{
			p.SetState(459)
			p.Match(qasm3ParserInteger)
		}
		{
			p.SetState(460)
			p.Match(qasm3ParserRPAREN)
		}

	case qasm3ParserT__20:
		{
			p.SetState(461)
			p.Match(qasm3ParserT__20)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(464)
		p.Match(qasm3ParserT__21)
	}

	return localctx
}

// IQuantumGateCallContext is an interface to support dynamic dispatch.
type IQuantumGateCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumGateCallContext differentiates from other interfaces.
	IsQuantumGateCallContext()
}

type QuantumGateCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumGateCallContext() *QuantumGateCallContext {
	var p = new(QuantumGateCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumGateCall
	return p
}

func (*QuantumGateCallContext) IsQuantumGateCallContext() {}

func NewQuantumGateCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumGateCallContext {
	var p = new(QuantumGateCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumGateCall

	return p
}

func (s *QuantumGateCallContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumGateCallContext) QuantumGateName() IQuantumGateNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateNameContext)
}

func (s *QuantumGateCallContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *QuantumGateCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *QuantumGateCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *QuantumGateCallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *QuantumGateCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumGateCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumGateCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumGateCall(s)
	}
}

func (s *QuantumGateCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumGateCall(s)
	}
}

func (p *qasm3Parser) QuantumGateCall() (localctx IQuantumGateCallContext) {
	localctx = NewQuantumGateCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, qasm3ParserRULE_quantumGateCall)
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
		p.SetState(466)
		p.QuantumGateName()
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(467)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
			{
				p.SetState(468)
				p.ExpressionList()
			}

		}
		{
			p.SetState(471)
			p.Match(qasm3ParserRPAREN)
		}

	}
	{
		p.SetState(474)
		p.IndexIdentifierList()
	}

	return localctx
}

// IQuantumGateNameContext is an interface to support dynamic dispatch.
type IQuantumGateNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantumGateNameContext differentiates from other interfaces.
	IsQuantumGateNameContext()
}

type QuantumGateNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantumGateNameContext() *QuantumGateNameContext {
	var p = new(QuantumGateNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_quantumGateName
	return p
}

func (*QuantumGateNameContext) IsQuantumGateNameContext() {}

func NewQuantumGateNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantumGateNameContext {
	var p = new(QuantumGateNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_quantumGateName

	return p
}

func (s *QuantumGateNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantumGateNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *QuantumGateNameContext) QuantumGateModifier() IQuantumGateModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateModifierContext)
}

func (s *QuantumGateNameContext) QuantumGateName() IQuantumGateNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumGateNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumGateNameContext)
}

func (s *QuantumGateNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantumGateNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantumGateNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterQuantumGateName(s)
	}
}

func (s *QuantumGateNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitQuantumGateName(s)
	}
}

func (p *qasm3Parser) QuantumGateName() (localctx IQuantumGateNameContext) {
	localctx = NewQuantumGateNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, qasm3ParserRULE_quantumGateName)

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

	p.SetState(483)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Match(qasm3ParserT__22)
		}

	case qasm3ParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Match(qasm3ParserT__23)
		}

	case qasm3ParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(478)
			p.Match(qasm3ParserT__24)
		}

	case qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(479)
			p.Match(qasm3ParserIdentifier)
		}

	case qasm3ParserT__18, qasm3ParserT__19, qasm3ParserT__20:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(480)
			p.QuantumGateModifier()
		}
		{
			p.SetState(481)
			p.QuantumGateName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *qasm3Parser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, qasm3ParserRULE_unaryOperator)
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
		p.SetState(485)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__25 || _la == qasm3ParserT__26) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (p *qasm3Parser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, qasm3ParserRULE_binaryOperator)
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
		p.SetState(487)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(qasm3ParserT__14-15))|(1<<(qasm3ParserT__27-15))|(1<<(qasm3ParserT__28-15))|(1<<(qasm3ParserT__29-15))|(1<<(qasm3ParserT__30-15))|(1<<(qasm3ParserT__31-15))|(1<<(qasm3ParserT__32-15))|(1<<(qasm3ParserT__33-15))|(1<<(qasm3ParserT__34-15))|(1<<(qasm3ParserT__35-15))|(1<<(qasm3ParserT__36-15))|(1<<(qasm3ParserT__37-15))|(1<<(qasm3ParserT__38-15))|(1<<(qasm3ParserT__39-15))|(1<<(qasm3ParserT__40-15))|(1<<(qasm3ParserT__41-15))|(1<<(qasm3ParserT__42-15))|(1<<(qasm3ParserT__43-15))|(1<<(qasm3ParserT__44-15)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *ExpressionStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *qasm3Parser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, qasm3ParserRULE_expressionStatement)

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

	p.SetState(494)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__16, qasm3ParserT__25, qasm3ParserT__26, qasm3ParserT__46, qasm3ParserT__47, qasm3ParserT__48, qasm3ParserT__49, qasm3ParserT__50, qasm3ParserT__51, qasm3ParserT__52, qasm3ParserT__53, qasm3ParserT__82, qasm3ParserT__83, qasm3ParserT__86, qasm3ParserConstant, qasm3ParserInteger, qasm3ParserRealNumber, qasm3ParserIdentifier, qasm3ParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.expression(0)
		}
		{
			p.SetState(490)
			p.Match(qasm3ParserSEMICOLON)
		}

	case qasm3ParserT__45:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Match(qasm3ParserT__45)
		}
		{
			p.SetState(493)
			p.ExpressionStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) MembershipTest() IMembershipTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembershipTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembershipTestContext)
}

func (s *ExpressionContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *ExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionContext) QuantumMeasurement() IQuantumMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumMeasurementContext)
}

func (s *ExpressionContext) ExpressionTerminator() IExpressionTerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionTerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionTerminatorContext)
}

func (s *ExpressionContext) BinaryOperator() IBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *ExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACKET, 0)
}

func (s *ExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACKET, 0)
}

func (s *ExpressionContext) Incrementor() IIncrementorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncrementorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncrementorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *qasm3Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *qasm3Parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 104
	p.EnterRecursionRule(localctx, 104, qasm3ParserRULE_expression, _p)
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
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(497)
			p.UnaryOperator()
		}
		{
			p.SetState(498)
			p.expression(7)
		}

	case 2:
		{
			p.SetState(500)
			p.MembershipTest()
		}

	case 3:
		{
			p.SetState(501)
			p.Call()
		}
		{
			p.SetState(502)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
			{
				p.SetState(503)
				p.ExpressionList()
			}

		}
		{
			p.SetState(506)
			p.Match(qasm3ParserRPAREN)
		}

	case 4:
		{
			p.SetState(508)
			p.QuantumMeasurement()
		}

	case 5:
		{
			p.SetState(509)
			p.ExpressionTerminator()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(523)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, qasm3ParserRULE_expression)
				p.SetState(512)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(513)
					p.BinaryOperator()
				}
				{
					p.SetState(514)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, qasm3ParserRULE_expression)
				p.SetState(516)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(517)
					p.Match(qasm3ParserLBRACKET)
				}
				{
					p.SetState(518)
					p.expression(0)
				}
				{
					p.SetState(519)
					p.Match(qasm3ParserRBRACKET)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, qasm3ParserRULE_expression)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(522)
					p.Incrementor()
				}

			}

		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionTerminatorContext is an interface to support dynamic dispatch.
type IExpressionTerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionTerminatorContext differentiates from other interfaces.
	IsExpressionTerminatorContext()
}

type ExpressionTerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionTerminatorContext() *ExpressionTerminatorContext {
	var p = new(ExpressionTerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_expressionTerminator
	return p
}

func (*ExpressionTerminatorContext) IsExpressionTerminatorContext() {}

func NewExpressionTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionTerminatorContext {
	var p = new(ExpressionTerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_expressionTerminator

	return p
}

func (s *ExpressionTerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionTerminatorContext) Constant() antlr.TerminalNode {
	return s.GetToken(qasm3ParserConstant, 0)
}

func (s *ExpressionTerminatorContext) Integer() antlr.TerminalNode {
	return s.GetToken(qasm3ParserInteger, 0)
}

func (s *ExpressionTerminatorContext) RealNumber() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRealNumber, 0)
}

func (s *ExpressionTerminatorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *ExpressionTerminatorContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(qasm3ParserStringLiteral, 0)
}

func (s *ExpressionTerminatorContext) TimeTerminator() ITimeTerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeTerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeTerminatorContext)
}

func (s *ExpressionTerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionTerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionTerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterExpressionTerminator(s)
	}
}

func (s *ExpressionTerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitExpressionTerminator(s)
	}
}

func (p *qasm3Parser) ExpressionTerminator() (localctx IExpressionTerminatorContext) {
	localctx = NewExpressionTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, qasm3ParserRULE_expressionTerminator)

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

	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(528)
			p.Match(qasm3ParserConstant)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(529)
			p.Match(qasm3ParserInteger)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(530)
			p.Match(qasm3ParserRealNumber)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(531)
			p.Match(qasm3ParserIdentifier)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(532)
			p.Match(qasm3ParserStringLiteral)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(533)
			p.TimeTerminator()
		}

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(qasm3ParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(qasm3ParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *qasm3Parser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, qasm3ParserRULE_expressionList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(536)
				p.expression(0)
			}
			{
				p.SetState(537)
				p.Match(qasm3ParserCOMMA)
			}

		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}
	{
		p.SetState(544)
		p.expression(0)
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *CallContext) BuiltInMath() IBuiltInMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltInMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltInMathContext)
}

func (s *CallContext) CastOperator() ICastOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICastOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICastOperatorContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *qasm3Parser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, qasm3ParserRULE_call)

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

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(546)
			p.Match(qasm3ParserIdentifier)
		}

	case qasm3ParserT__46, qasm3ParserT__47, qasm3ParserT__48, qasm3ParserT__49, qasm3ParserT__50, qasm3ParserT__51, qasm3ParserT__52, qasm3ParserT__53:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(547)
			p.BuiltInMath()
		}

	case qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__82, qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(548)
			p.CastOperator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuiltInMathContext is an interface to support dynamic dispatch.
type IBuiltInMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltInMathContext differentiates from other interfaces.
	IsBuiltInMathContext()
}

type BuiltInMathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltInMathContext() *BuiltInMathContext {
	var p = new(BuiltInMathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_builtInMath
	return p
}

func (*BuiltInMathContext) IsBuiltInMathContext() {}

func NewBuiltInMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltInMathContext {
	var p = new(BuiltInMathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_builtInMath

	return p
}

func (s *BuiltInMathContext) GetParser() antlr.Parser { return s.parser }
func (s *BuiltInMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltInMathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltInMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterBuiltInMath(s)
	}
}

func (s *BuiltInMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitBuiltInMath(s)
	}
}

func (p *qasm3Parser) BuiltInMath() (localctx IBuiltInMathContext) {
	localctx = NewBuiltInMathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, qasm3ParserRULE_builtInMath)
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
		p.SetState(551)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICastOperatorContext is an interface to support dynamic dispatch.
type ICastOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCastOperatorContext differentiates from other interfaces.
	IsCastOperatorContext()
}

type CastOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastOperatorContext() *CastOperatorContext {
	var p = new(CastOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_castOperator
	return p
}

func (*CastOperatorContext) IsCastOperatorContext() {}

func NewCastOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastOperatorContext {
	var p = new(CastOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_castOperator

	return p
}

func (s *CastOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CastOperatorContext) ClassicalType() IClassicalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalTypeContext)
}

func (s *CastOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCastOperator(s)
	}
}

func (s *CastOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCastOperator(s)
	}
}

func (p *qasm3Parser) CastOperator() (localctx ICastOperatorContext) {
	localctx = NewCastOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, qasm3ParserRULE_castOperator)

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
		p.SetState(553)
		p.ClassicalType()
	}

	return localctx
}

// IIncrementorContext is an interface to support dynamic dispatch.
type IIncrementorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncrementorContext differentiates from other interfaces.
	IsIncrementorContext()
}

type IncrementorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrementorContext() *IncrementorContext {
	var p = new(IncrementorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_incrementor
	return p
}

func (*IncrementorContext) IsIncrementorContext() {}

func NewIncrementorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrementorContext {
	var p = new(IncrementorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_incrementor

	return p
}

func (s *IncrementorContext) GetParser() antlr.Parser { return s.parser }
func (s *IncrementorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrementorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterIncrementor(s)
	}
}

func (s *IncrementorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitIncrementor(s)
	}
}

func (p *qasm3Parser) Incrementor() (localctx IIncrementorContext) {
	localctx = NewIncrementorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, qasm3ParserRULE_incrementor)
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
		p.SetState(555)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__54 || _la == qasm3ParserT__55) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_assignmentExpression
	return p
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (p *qasm3Parser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, qasm3ParserRULE_assignmentExpression)

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
		p.SetState(557)
		p.AssignmentOperator()
	}
	{
		p.SetState(558)
		p.expression(0)
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserASSIGN, 0)
}

func (s *AssignmentOperatorContext) ARROW() antlr.TerminalNode {
	return s.GetToken(qasm3ParserARROW, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *qasm3Parser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, qasm3ParserRULE_assignmentOperator)
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
		p.SetState(560)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(qasm3ParserT__56-57))|(1<<(qasm3ParserT__57-57))|(1<<(qasm3ParserT__58-57))|(1<<(qasm3ParserT__59-57))|(1<<(qasm3ParserT__60-57))|(1<<(qasm3ParserT__61-57))|(1<<(qasm3ParserT__62-57))|(1<<(qasm3ParserT__63-57))|(1<<(qasm3ParserT__64-57))|(1<<(qasm3ParserT__65-57)))) != 0) || _la == qasm3ParserASSIGN || _la == qasm3ParserARROW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMembershipTestContext is an interface to support dynamic dispatch.
type IMembershipTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMembershipTestContext differentiates from other interfaces.
	IsMembershipTestContext()
}

type MembershipTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembershipTestContext() *MembershipTestContext {
	var p = new(MembershipTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_membershipTest
	return p
}

func (*MembershipTestContext) IsMembershipTestContext() {}

func NewMembershipTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembershipTestContext {
	var p = new(MembershipTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_membershipTest

	return p
}

func (s *MembershipTestContext) GetParser() antlr.Parser { return s.parser }

func (s *MembershipTestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *MembershipTestContext) SetDeclaration() ISetDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetDeclarationContext)
}

func (s *MembershipTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembershipTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterMembershipTest(s)
	}
}

func (s *MembershipTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitMembershipTest(s)
	}
}

func (p *qasm3Parser) MembershipTest() (localctx IMembershipTestContext) {
	localctx = NewMembershipTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, qasm3ParserRULE_membershipTest)

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
		p.SetState(562)
		p.Match(qasm3ParserIdentifier)
	}
	{
		p.SetState(563)
		p.Match(qasm3ParserT__66)
	}
	{
		p.SetState(564)
		p.SetDeclaration()
	}

	return localctx
}

// ISetDeclarationContext is an interface to support dynamic dispatch.
type ISetDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetDeclarationContext differentiates from other interfaces.
	IsSetDeclarationContext()
}

type SetDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetDeclarationContext() *SetDeclarationContext {
	var p = new(SetDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_setDeclaration
	return p
}

func (*SetDeclarationContext) IsSetDeclarationContext() {}

func NewSetDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetDeclarationContext {
	var p = new(SetDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_setDeclaration

	return p
}

func (s *SetDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SetDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACE, 0)
}

func (s *SetDeclarationContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SetDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACE, 0)
}

func (s *SetDeclarationContext) RangeDefinition() IRangeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDefinitionContext)
}

func (s *SetDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterSetDeclaration(s)
	}
}

func (s *SetDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitSetDeclaration(s)
	}
}

func (p *qasm3Parser) SetDeclaration() (localctx ISetDeclarationContext) {
	localctx = NewSetDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, qasm3ParserRULE_setDeclaration)

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

	p.SetState(571)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(566)
			p.Match(qasm3ParserLBRACE)
		}
		{
			p.SetState(567)
			p.ExpressionList()
		}
		{
			p.SetState(568)
			p.Match(qasm3ParserRBRACE)
		}

	case qasm3ParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(570)
			p.RangeDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoopBranchBlockContext is an interface to support dynamic dispatch.
type ILoopBranchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopBranchBlockContext differentiates from other interfaces.
	IsLoopBranchBlockContext()
}

type LoopBranchBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopBranchBlockContext() *LoopBranchBlockContext {
	var p = new(LoopBranchBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_loopBranchBlock
	return p
}

func (*LoopBranchBlockContext) IsLoopBranchBlockContext() {}

func NewLoopBranchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopBranchBlockContext {
	var p = new(LoopBranchBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_loopBranchBlock

	return p
}

func (s *LoopBranchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopBranchBlockContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LoopBranchBlockContext) ProgramBlock() IProgramBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramBlockContext)
}

func (s *LoopBranchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopBranchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopBranchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterLoopBranchBlock(s)
	}
}

func (s *LoopBranchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitLoopBranchBlock(s)
	}
}

func (p *qasm3Parser) LoopBranchBlock() (localctx ILoopBranchBlockContext) {
	localctx = NewLoopBranchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, qasm3ParserRULE_loopBranchBlock)

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

	p.SetState(575)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__2, qasm3ParserT__3, qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__12, qasm3ParserT__13, qasm3ParserT__15, qasm3ParserT__16, qasm3ParserT__17, qasm3ParserT__18, qasm3ParserT__19, qasm3ParserT__20, qasm3ParserT__22, qasm3ParserT__23, qasm3ParserT__24, qasm3ParserT__25, qasm3ParserT__26, qasm3ParserT__45, qasm3ParserT__46, qasm3ParserT__47, qasm3ParserT__48, qasm3ParserT__49, qasm3ParserT__50, qasm3ParserT__51, qasm3ParserT__52, qasm3ParserT__53, qasm3ParserT__67, qasm3ParserT__69, qasm3ParserT__70, qasm3ParserT__71, qasm3ParserT__72, qasm3ParserT__73, qasm3ParserT__74, qasm3ParserT__75, qasm3ParserT__76, qasm3ParserT__82, qasm3ParserT__83, qasm3ParserT__84, qasm3ParserT__85, qasm3ParserT__86, qasm3ParserT__87, qasm3ParserT__88, qasm3ParserT__90, qasm3ParserConstant, qasm3ParserInteger, qasm3ParserRealNumber, qasm3ParserIdentifier, qasm3ParserStringLiteral, qasm3ParserLineComment, qasm3ParserBlockComment:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(573)
			p.Statement()
		}

	case qasm3ParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(574)
			p.ProgramBlock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBranchingStatementContext is an interface to support dynamic dispatch.
type IBranchingStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchingStatementContext differentiates from other interfaces.
	IsBranchingStatementContext()
}

type BranchingStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchingStatementContext() *BranchingStatementContext {
	var p = new(BranchingStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_branchingStatement
	return p
}

func (*BranchingStatementContext) IsBranchingStatementContext() {}

func NewBranchingStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchingStatementContext {
	var p = new(BranchingStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_branchingStatement

	return p
}

func (s *BranchingStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchingStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *BranchingStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BranchingStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *BranchingStatementContext) AllLoopBranchBlock() []ILoopBranchBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILoopBranchBlockContext)(nil)).Elem())
	var tst = make([]ILoopBranchBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILoopBranchBlockContext)
		}
	}

	return tst
}

func (s *BranchingStatementContext) LoopBranchBlock(i int) ILoopBranchBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopBranchBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILoopBranchBlockContext)
}

func (s *BranchingStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchingStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchingStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterBranchingStatement(s)
	}
}

func (s *BranchingStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitBranchingStatement(s)
	}
}

func (p *qasm3Parser) BranchingStatement() (localctx IBranchingStatementContext) {
	localctx = NewBranchingStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, qasm3ParserRULE_branchingStatement)

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
		p.SetState(577)
		p.Match(qasm3ParserT__67)
	}
	{
		p.SetState(578)
		p.Match(qasm3ParserLPAREN)
	}
	{
		p.SetState(579)
		p.expression(0)
	}
	{
		p.SetState(580)
		p.Match(qasm3ParserRPAREN)
	}
	{
		p.SetState(581)
		p.LoopBranchBlock()
	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(582)
			p.Match(qasm3ParserT__68)
		}
		{
			p.SetState(583)
			p.LoopBranchBlock()
		}

	}

	return localctx
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_loopStatement
	return p
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) LoopBranchBlock() ILoopBranchBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopBranchBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopBranchBlockContext)
}

func (s *LoopStatementContext) MembershipTest() IMembershipTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembershipTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembershipTestContext)
}

func (s *LoopStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *LoopStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (p *qasm3Parser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, qasm3ParserRULE_loopStatement)

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
	p.SetState(593)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__69:
		{
			p.SetState(586)
			p.Match(qasm3ParserT__69)
		}
		{
			p.SetState(587)
			p.MembershipTest()
		}

	case qasm3ParserT__70:
		{
			p.SetState(588)
			p.Match(qasm3ParserT__70)
		}
		{
			p.SetState(589)
			p.Match(qasm3ParserLPAREN)
		}
		{
			p.SetState(590)
			p.expression(0)
		}
		{
			p.SetState(591)
			p.Match(qasm3ParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(595)
		p.LoopBranchBlock()
	}

	return localctx
}

// IControlDirectiveStatementContext is an interface to support dynamic dispatch.
type IControlDirectiveStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlDirectiveStatementContext differentiates from other interfaces.
	IsControlDirectiveStatementContext()
}

type ControlDirectiveStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlDirectiveStatementContext() *ControlDirectiveStatementContext {
	var p = new(ControlDirectiveStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_controlDirectiveStatement
	return p
}

func (*ControlDirectiveStatementContext) IsControlDirectiveStatementContext() {}

func NewControlDirectiveStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlDirectiveStatementContext {
	var p = new(ControlDirectiveStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_controlDirectiveStatement

	return p
}

func (s *ControlDirectiveStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlDirectiveStatementContext) ControlDirective() IControlDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlDirectiveContext)
}

func (s *ControlDirectiveStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *ControlDirectiveStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlDirectiveStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlDirectiveStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterControlDirectiveStatement(s)
	}
}

func (s *ControlDirectiveStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitControlDirectiveStatement(s)
	}
}

func (p *qasm3Parser) ControlDirectiveStatement() (localctx IControlDirectiveStatementContext) {
	localctx = NewControlDirectiveStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, qasm3ParserRULE_controlDirectiveStatement)

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
		p.SetState(597)
		p.ControlDirective()
	}
	{
		p.SetState(598)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// IControlDirectiveContext is an interface to support dynamic dispatch.
type IControlDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlDirectiveContext differentiates from other interfaces.
	IsControlDirectiveContext()
}

type ControlDirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlDirectiveContext() *ControlDirectiveContext {
	var p = new(ControlDirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_controlDirective
	return p
}

func (*ControlDirectiveContext) IsControlDirectiveContext() {}

func NewControlDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlDirectiveContext {
	var p = new(ControlDirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_controlDirective

	return p
}

func (s *ControlDirectiveContext) GetParser() antlr.Parser { return s.parser }
func (s *ControlDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterControlDirective(s)
	}
}

func (s *ControlDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitControlDirective(s)
	}
}

func (p *qasm3Parser) ControlDirective() (localctx IControlDirectiveContext) {
	localctx = NewControlDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, qasm3ParserRULE_controlDirective)
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
		p.SetState(600)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(qasm3ParserT__71-72))|(1<<(qasm3ParserT__72-72))|(1<<(qasm3ParserT__73-72)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IKernelDeclarationContext is an interface to support dynamic dispatch.
type IKernelDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKernelDeclarationContext differentiates from other interfaces.
	IsKernelDeclarationContext()
}

type KernelDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKernelDeclarationContext() *KernelDeclarationContext {
	var p = new(KernelDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_kernelDeclaration
	return p
}

func (*KernelDeclarationContext) IsKernelDeclarationContext() {}

func NewKernelDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KernelDeclarationContext {
	var p = new(KernelDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_kernelDeclaration

	return p
}

func (s *KernelDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *KernelDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *KernelDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *KernelDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *KernelDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *KernelDeclarationContext) ReturnSignature() IReturnSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnSignatureContext)
}

func (s *KernelDeclarationContext) ClassicalType() IClassicalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalTypeContext)
}

func (s *KernelDeclarationContext) ClassicalTypeList() IClassicalTypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalTypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalTypeListContext)
}

func (s *KernelDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KernelDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KernelDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterKernelDeclaration(s)
	}
}

func (s *KernelDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitKernelDeclaration(s)
	}
}

func (p *qasm3Parser) KernelDeclaration() (localctx IKernelDeclarationContext) {
	localctx = NewKernelDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, qasm3ParserRULE_kernelDeclaration)
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
		p.SetState(602)
		p.Match(qasm3ParserT__74)
	}
	{
		p.SetState(603)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(604)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11))) != 0) || _la == qasm3ParserT__82 || _la == qasm3ParserT__83 {
			{
				p.SetState(605)
				p.ClassicalTypeList()
			}

		}
		{
			p.SetState(608)
			p.Match(qasm3ParserRPAREN)
		}

	}
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserARROW {
		{
			p.SetState(611)
			p.ReturnSignature()
		}

	}
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11))) != 0) || _la == qasm3ParserT__82 || _la == qasm3ParserT__83 {
		{
			p.SetState(614)
			p.ClassicalType()
		}

	}
	{
		p.SetState(617)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// ISubroutineDefinitionContext is an interface to support dynamic dispatch.
type ISubroutineDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutineDefinitionContext differentiates from other interfaces.
	IsSubroutineDefinitionContext()
}

type SubroutineDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutineDefinitionContext() *SubroutineDefinitionContext {
	var p = new(SubroutineDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_subroutineDefinition
	return p
}

func (*SubroutineDefinitionContext) IsSubroutineDefinitionContext() {}

func NewSubroutineDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutineDefinitionContext {
	var p = new(SubroutineDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_subroutineDefinition

	return p
}

func (s *SubroutineDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutineDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *SubroutineDefinitionContext) ProgramBlock() IProgramBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramBlockContext)
}

func (s *SubroutineDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *SubroutineDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *SubroutineDefinitionContext) ReturnSignature() IReturnSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnSignatureContext)
}

func (s *SubroutineDefinitionContext) SubroutineArgumentList() ISubroutineArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutineArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubroutineArgumentListContext)
}

func (s *SubroutineDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutineDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutineDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterSubroutineDefinition(s)
	}
}

func (s *SubroutineDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitSubroutineDefinition(s)
	}
}

func (p *qasm3Parser) SubroutineDefinition() (localctx ISubroutineDefinitionContext) {
	localctx = NewSubroutineDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, qasm3ParserRULE_subroutineDefinition)
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
		p.SetState(619)
		p.Match(qasm3ParserT__75)
	}
	{
		p.SetState(620)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(626)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(621)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(623)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__2)|(1<<qasm3ParserT__3)|(1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11))) != 0) || _la == qasm3ParserT__82 || _la == qasm3ParserT__83 {
			{
				p.SetState(622)
				p.SubroutineArgumentList()
			}

		}
		{
			p.SetState(625)
			p.Match(qasm3ParserRPAREN)
		}

	}
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserARROW {
		{
			p.SetState(628)
			p.ReturnSignature()
		}

	}
	{
		p.SetState(631)
		p.ProgramBlock()
	}

	return localctx
}

// ISubroutineArgumentListContext is an interface to support dynamic dispatch.
type ISubroutineArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutineArgumentListContext differentiates from other interfaces.
	IsSubroutineArgumentListContext()
}

type SubroutineArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutineArgumentListContext() *SubroutineArgumentListContext {
	var p = new(SubroutineArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_subroutineArgumentList
	return p
}

func (*SubroutineArgumentListContext) IsSubroutineArgumentListContext() {}

func NewSubroutineArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutineArgumentListContext {
	var p = new(SubroutineArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_subroutineArgumentList

	return p
}

func (s *SubroutineArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutineArgumentListContext) ClassicalArgumentList() IClassicalArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalArgumentListContext)
}

func (s *SubroutineArgumentListContext) QuantumArgumentList() IQuantumArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumArgumentListContext)
}

func (s *SubroutineArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutineArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutineArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterSubroutineArgumentList(s)
	}
}

func (s *SubroutineArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitSubroutineArgumentList(s)
	}
}

func (p *qasm3Parser) SubroutineArgumentList() (localctx ISubroutineArgumentListContext) {
	localctx = NewSubroutineArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, qasm3ParserRULE_subroutineArgumentList)

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

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__4, qasm3ParserT__5, qasm3ParserT__6, qasm3ParserT__7, qasm3ParserT__8, qasm3ParserT__9, qasm3ParserT__10, qasm3ParserT__11, qasm3ParserT__82, qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)
			p.ClassicalArgumentList()
		}

	case qasm3ParserT__2, qasm3ParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.QuantumArgumentList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPragmaContext is an interface to support dynamic dispatch.
type IPragmaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPragmaContext differentiates from other interfaces.
	IsPragmaContext()
}

type PragmaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPragmaContext() *PragmaContext {
	var p = new(PragmaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_pragma
	return p
}

func (*PragmaContext) IsPragmaContext() {}

func NewPragmaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PragmaContext {
	var p = new(PragmaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_pragma

	return p
}

func (s *PragmaContext) GetParser() antlr.Parser { return s.parser }

func (s *PragmaContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACE, 0)
}

func (s *PragmaContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACE, 0)
}

func (s *PragmaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PragmaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PragmaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterPragma(s)
	}
}

func (s *PragmaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitPragma(s)
	}
}

func (p *qasm3Parser) Pragma() (localctx IPragmaContext) {
	localctx = NewPragmaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, qasm3ParserRULE_pragma)

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
		p.SetState(637)
		p.Match(qasm3ParserT__76)
	}
	{
		p.SetState(638)
		p.Match(qasm3ParserLBRACE)
	}
	p.SetState(639)
	p.MatchWildcard()

	{
		p.SetState(640)
		p.Match(qasm3ParserRBRACE)
	}

	return localctx
}

// ITimeUnitContext is an interface to support dynamic dispatch.
type ITimeUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeUnitContext differentiates from other interfaces.
	IsTimeUnitContext()
}

type TimeUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeUnitContext() *TimeUnitContext {
	var p = new(TimeUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeUnit
	return p
}

func (*TimeUnitContext) IsTimeUnitContext() {}

func NewTimeUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeUnitContext {
	var p = new(TimeUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeUnit

	return p
}

func (s *TimeUnitContext) GetParser() antlr.Parser { return s.parser }
func (s *TimeUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeUnit(s)
	}
}

func (s *TimeUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeUnit(s)
	}
}

func (p *qasm3Parser) TimeUnit() (localctx ITimeUnitContext) {
	localctx = NewTimeUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, qasm3ParserRULE_timeUnit)
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
		p.SetState(642)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(qasm3ParserT__77-78))|(1<<(qasm3ParserT__78-78))|(1<<(qasm3ParserT__79-78))|(1<<(qasm3ParserT__80-78))|(1<<(qasm3ParserT__81-78)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITimingTypeContext is an interface to support dynamic dispatch.
type ITimingTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimingTypeContext differentiates from other interfaces.
	IsTimingTypeContext()
}

type TimingTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimingTypeContext() *TimingTypeContext {
	var p = new(TimingTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timingType
	return p
}

func (*TimingTypeContext) IsTimingTypeContext() {}

func NewTimingTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimingTypeContext {
	var p = new(TimingTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timingType

	return p
}

func (s *TimingTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimingTypeContext) Integer() antlr.TerminalNode {
	return s.GetToken(qasm3ParserInteger, 0)
}

func (s *TimingTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimingTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimingTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimingType(s)
	}
}

func (s *TimingTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimingType(s)
	}
}

func (p *qasm3Parser) TimingType() (localctx ITimingTypeContext) {
	localctx = NewTimingTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, qasm3ParserRULE_timingType)
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

	p.SetState(649)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__82:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(644)
			p.Match(qasm3ParserT__82)
		}

	case qasm3ParserT__83:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(645)
			p.Match(qasm3ParserT__83)
		}
		p.SetState(647)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == qasm3ParserInteger {
			{
				p.SetState(646)
				p.Match(qasm3ParserInteger)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimingBoxContext is an interface to support dynamic dispatch.
type ITimingBoxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimingBoxContext differentiates from other interfaces.
	IsTimingBoxContext()
}

type TimingBoxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimingBoxContext() *TimingBoxContext {
	var p = new(TimingBoxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timingBox
	return p
}

func (*TimingBoxContext) IsTimingBoxContext() {}

func NewTimingBoxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimingBoxContext {
	var p = new(TimingBoxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timingBox

	return p
}

func (s *TimingBoxContext) GetParser() antlr.Parser { return s.parser }

func (s *TimingBoxContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *TimingBoxContext) QuantumBlock() IQuantumBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantumBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantumBlockContext)
}

func (s *TimingBoxContext) TimeUnit() ITimeUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeUnitContext)
}

func (s *TimingBoxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimingBoxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimingBoxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimingBox(s)
	}
}

func (s *TimingBoxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimingBox(s)
	}
}

func (p *qasm3Parser) TimingBox() (localctx ITimingBoxContext) {
	localctx = NewTimingBoxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, qasm3ParserRULE_timingBox)

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

	p.SetState(658)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__84:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(651)
			p.Match(qasm3ParserT__84)
		}
		{
			p.SetState(652)
			p.Match(qasm3ParserIdentifier)
		}
		{
			p.SetState(653)
			p.QuantumBlock()
		}

	case qasm3ParserT__85:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(654)
			p.Match(qasm3ParserT__85)
		}
		{
			p.SetState(655)
			p.TimeUnit()
		}
		{
			p.SetState(656)
			p.QuantumBlock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeTerminatorContext is an interface to support dynamic dispatch.
type ITimeTerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeTerminatorContext differentiates from other interfaces.
	IsTimeTerminatorContext()
}

type TimeTerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeTerminatorContext() *TimeTerminatorContext {
	var p = new(TimeTerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeTerminator
	return p
}

func (*TimeTerminatorContext) IsTimeTerminatorContext() {}

func NewTimeTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeTerminatorContext {
	var p = new(TimeTerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeTerminator

	return p
}

func (s *TimeTerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeTerminatorContext) TimeIdentifier() ITimeIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIdentifierContext)
}

func (s *TimeTerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeTerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeTerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeTerminator(s)
	}
}

func (s *TimeTerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeTerminator(s)
	}
}

func (p *qasm3Parser) TimeTerminator() (localctx ITimeTerminatorContext) {
	localctx = NewTimeTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, qasm3ParserRULE_timeTerminator)

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

	p.SetState(662)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__53, qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(660)
			p.TimeIdentifier()
		}

	case qasm3ParserT__86:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(661)
			p.Match(qasm3ParserT__86)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeIdentifierContext is an interface to support dynamic dispatch.
type ITimeIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIdentifierContext differentiates from other interfaces.
	IsTimeIdentifierContext()
}

type TimeIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIdentifierContext() *TimeIdentifierContext {
	var p = new(TimeIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeIdentifier
	return p
}

func (*TimeIdentifierContext) IsTimeIdentifierContext() {}

func NewTimeIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIdentifierContext {
	var p = new(TimeIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeIdentifier

	return p
}

func (s *TimeIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *TimeIdentifierContext) TimeUnit() ITimeUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeUnitContext)
}

func (s *TimeIdentifierContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *TimeIdentifierContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *TimeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeIdentifier(s)
	}
}

func (s *TimeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeIdentifier(s)
	}
}

func (p *qasm3Parser) TimeIdentifier() (localctx ITimeIdentifierContext) {
	localctx = NewTimeIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, qasm3ParserRULE_timeIdentifier)

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

	p.SetState(672)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(664)
			p.Match(qasm3ParserIdentifier)
		}
		p.SetState(666)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(665)
				p.TimeUnit()
			}

		}

	case qasm3ParserT__53:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(668)
			p.Match(qasm3ParserT__53)
		}
		{
			p.SetState(669)
			p.Match(qasm3ParserLPAREN)
		}
		{
			p.SetState(670)
			p.Match(qasm3ParserIdentifier)
		}
		{
			p.SetState(671)
			p.Match(qasm3ParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeInstructionNameContext is an interface to support dynamic dispatch.
type ITimeInstructionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeInstructionNameContext differentiates from other interfaces.
	IsTimeInstructionNameContext()
}

type TimeInstructionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeInstructionNameContext() *TimeInstructionNameContext {
	var p = new(TimeInstructionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeInstructionName
	return p
}

func (*TimeInstructionNameContext) IsTimeInstructionNameContext() {}

func NewTimeInstructionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeInstructionNameContext {
	var p = new(TimeInstructionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeInstructionName

	return p
}

func (s *TimeInstructionNameContext) GetParser() antlr.Parser { return s.parser }
func (s *TimeInstructionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeInstructionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeInstructionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeInstructionName(s)
	}
}

func (s *TimeInstructionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeInstructionName(s)
	}
}

func (p *qasm3Parser) TimeInstructionName() (localctx ITimeInstructionNameContext) {
	localctx = NewTimeInstructionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, qasm3ParserRULE_timeInstructionName)
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
		p.SetState(674)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__87 || _la == qasm3ParserT__88) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITimeInstructionContext is an interface to support dynamic dispatch.
type ITimeInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeInstructionContext differentiates from other interfaces.
	IsTimeInstructionContext()
}

type TimeInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeInstructionContext() *TimeInstructionContext {
	var p = new(TimeInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeInstruction
	return p
}

func (*TimeInstructionContext) IsTimeInstructionContext() {}

func NewTimeInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeInstructionContext {
	var p = new(TimeInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeInstruction

	return p
}

func (s *TimeInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeInstructionContext) TimeInstructionName() ITimeInstructionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeInstructionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeInstructionNameContext)
}

func (s *TimeInstructionContext) Designator() IDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *TimeInstructionContext) RangeDefinition() IRangeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDefinitionContext)
}

func (s *TimeInstructionContext) IndexIdentifierList() IIndexIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexIdentifierListContext)
}

func (s *TimeInstructionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *TimeInstructionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *TimeInstructionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *TimeInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeInstruction(s)
	}
}

func (s *TimeInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeInstruction(s)
	}
}

func (p *qasm3Parser) TimeInstruction() (localctx ITimeInstructionContext) {
	localctx = NewTimeInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, qasm3ParserRULE_timeInstruction)
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
		p.SetState(676)
		p.TimeInstructionName()
	}
	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(677)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(679)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
			{
				p.SetState(678)
				p.ExpressionList()
			}

		}
		{
			p.SetState(681)
			p.Match(qasm3ParserRPAREN)
		}

	}
	{
		p.SetState(684)
		p.Designator()
	}
	p.SetState(687)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserLBRACKET:
		{
			p.SetState(685)
			p.RangeDefinition()
		}

	case qasm3ParserIdentifier:
		{
			p.SetState(686)
			p.IndexIdentifierList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeStatementContext is an interface to support dynamic dispatch.
type ITimeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeStatementContext differentiates from other interfaces.
	IsTimeStatementContext()
}

type TimeStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeStatementContext() *TimeStatementContext {
	var p = new(TimeStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_timeStatement
	return p
}

func (*TimeStatementContext) IsTimeStatementContext() {}

func NewTimeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeStatementContext {
	var p = new(TimeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_timeStatement

	return p
}

func (s *TimeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeStatementContext) TimeInstruction() ITimeInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeInstructionContext)
}

func (s *TimeStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *TimeStatementContext) TimingBox() ITimingBoxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimingBoxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimingBoxContext)
}

func (s *TimeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterTimeStatement(s)
	}
}

func (s *TimeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitTimeStatement(s)
	}
}

func (p *qasm3Parser) TimeStatement() (localctx ITimeStatementContext) {
	localctx = NewTimeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, qasm3ParserRULE_timeStatement)

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

	p.SetState(693)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__87, qasm3ParserT__88:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(689)
			p.TimeInstruction()
		}
		{
			p.SetState(690)
			p.Match(qasm3ParserSEMICOLON)
		}

	case qasm3ParserT__84, qasm3ParserT__85:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(692)
			p.TimingBox()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICalibrationContext is an interface to support dynamic dispatch.
type ICalibrationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalibrationContext differentiates from other interfaces.
	IsCalibrationContext()
}

type CalibrationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalibrationContext() *CalibrationContext {
	var p = new(CalibrationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_calibration
	return p
}

func (*CalibrationContext) IsCalibrationContext() {}

func NewCalibrationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalibrationContext {
	var p = new(CalibrationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_calibration

	return p
}

func (s *CalibrationContext) GetParser() antlr.Parser { return s.parser }

func (s *CalibrationContext) CalibrationGrammarDeclaration() ICalibrationGrammarDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationGrammarDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationGrammarDeclarationContext)
}

func (s *CalibrationContext) CalibrationDefinition() ICalibrationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationDefinitionContext)
}

func (s *CalibrationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalibrationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalibrationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCalibration(s)
	}
}

func (s *CalibrationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCalibration(s)
	}
}

func (p *qasm3Parser) Calibration() (localctx ICalibrationContext) {
	localctx = NewCalibrationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, qasm3ParserRULE_calibration)

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

	p.SetState(697)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qasm3ParserT__89:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(695)
			p.CalibrationGrammarDeclaration()
		}

	case qasm3ParserT__90:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(696)
			p.CalibrationDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICalibrationGrammarDeclarationContext is an interface to support dynamic dispatch.
type ICalibrationGrammarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalibrationGrammarDeclarationContext differentiates from other interfaces.
	IsCalibrationGrammarDeclarationContext()
}

type CalibrationGrammarDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalibrationGrammarDeclarationContext() *CalibrationGrammarDeclarationContext {
	var p = new(CalibrationGrammarDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_calibrationGrammarDeclaration
	return p
}

func (*CalibrationGrammarDeclarationContext) IsCalibrationGrammarDeclarationContext() {}

func NewCalibrationGrammarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalibrationGrammarDeclarationContext {
	var p = new(CalibrationGrammarDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_calibrationGrammarDeclaration

	return p
}

func (s *CalibrationGrammarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CalibrationGrammarDeclarationContext) CalibrationGrammar() ICalibrationGrammarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationGrammarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationGrammarContext)
}

func (s *CalibrationGrammarDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(qasm3ParserSEMICOLON, 0)
}

func (s *CalibrationGrammarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalibrationGrammarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalibrationGrammarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCalibrationGrammarDeclaration(s)
	}
}

func (s *CalibrationGrammarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCalibrationGrammarDeclaration(s)
	}
}

func (p *qasm3Parser) CalibrationGrammarDeclaration() (localctx ICalibrationGrammarDeclarationContext) {
	localctx = NewCalibrationGrammarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, qasm3ParserRULE_calibrationGrammarDeclaration)

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
		p.SetState(699)
		p.Match(qasm3ParserT__89)
	}
	{
		p.SetState(700)
		p.CalibrationGrammar()
	}
	{
		p.SetState(701)
		p.Match(qasm3ParserSEMICOLON)
	}

	return localctx
}

// ICalibrationDefinitionContext is an interface to support dynamic dispatch.
type ICalibrationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalibrationDefinitionContext differentiates from other interfaces.
	IsCalibrationDefinitionContext()
}

type CalibrationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalibrationDefinitionContext() *CalibrationDefinitionContext {
	var p = new(CalibrationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_calibrationDefinition
	return p
}

func (*CalibrationDefinitionContext) IsCalibrationDefinitionContext() {}

func NewCalibrationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalibrationDefinitionContext {
	var p = new(CalibrationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_calibrationDefinition

	return p
}

func (s *CalibrationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *CalibrationDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *CalibrationDefinitionContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *CalibrationDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLBRACE, 0)
}

func (s *CalibrationDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRBRACE, 0)
}

func (s *CalibrationDefinitionContext) CalibrationGrammar() ICalibrationGrammarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationGrammarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationGrammarContext)
}

func (s *CalibrationDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserLPAREN, 0)
}

func (s *CalibrationDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(qasm3ParserRPAREN, 0)
}

func (s *CalibrationDefinitionContext) ReturnSignature() IReturnSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnSignatureContext)
}

func (s *CalibrationDefinitionContext) CalibrationArgumentList() ICalibrationArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalibrationArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalibrationArgumentListContext)
}

func (s *CalibrationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalibrationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalibrationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCalibrationDefinition(s)
	}
}

func (s *CalibrationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCalibrationDefinition(s)
	}
}

func (p *qasm3Parser) CalibrationDefinition() (localctx ICalibrationDefinitionContext) {
	localctx = NewCalibrationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, qasm3ParserRULE_calibrationDefinition)
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
		p.SetState(703)
		p.Match(qasm3ParserT__90)
	}
	p.SetState(705)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(704)
			p.CalibrationGrammar()
		}

	}
	{
		p.SetState(707)
		p.Match(qasm3ParserIdentifier)
	}
	p.SetState(713)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserLPAREN {
		{
			p.SetState(708)
			p.Match(qasm3ParserLPAREN)
		}
		p.SetState(710)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qasm3ParserT__4)|(1<<qasm3ParserT__5)|(1<<qasm3ParserT__6)|(1<<qasm3ParserT__7)|(1<<qasm3ParserT__8)|(1<<qasm3ParserT__9)|(1<<qasm3ParserT__10)|(1<<qasm3ParserT__11)|(1<<qasm3ParserT__16)|(1<<qasm3ParserT__25)|(1<<qasm3ParserT__26))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(qasm3ParserT__46-47))|(1<<(qasm3ParserT__47-47))|(1<<(qasm3ParserT__48-47))|(1<<(qasm3ParserT__49-47))|(1<<(qasm3ParserT__50-47))|(1<<(qasm3ParserT__51-47))|(1<<(qasm3ParserT__52-47))|(1<<(qasm3ParserT__53-47)))) != 0) || (((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(qasm3ParserT__82-83))|(1<<(qasm3ParserT__83-83))|(1<<(qasm3ParserT__86-83))|(1<<(qasm3ParserConstant-83))|(1<<(qasm3ParserInteger-83))|(1<<(qasm3ParserRealNumber-83))|(1<<(qasm3ParserIdentifier-83))|(1<<(qasm3ParserStringLiteral-83)))) != 0) {
			{
				p.SetState(709)
				p.CalibrationArgumentList()
			}

		}
		{
			p.SetState(712)
			p.Match(qasm3ParserRPAREN)
		}

	}
	{
		p.SetState(715)
		p.IdentifierList()
	}
	p.SetState(717)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == qasm3ParserARROW {
		{
			p.SetState(716)
			p.ReturnSignature()
		}

	}
	{
		p.SetState(719)
		p.Match(qasm3ParserLBRACE)
	}
	p.SetState(720)
	p.MatchWildcard()

	{
		p.SetState(721)
		p.Match(qasm3ParserRBRACE)
	}

	return localctx
}

// ICalibrationGrammarContext is an interface to support dynamic dispatch.
type ICalibrationGrammarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalibrationGrammarContext differentiates from other interfaces.
	IsCalibrationGrammarContext()
}

type CalibrationGrammarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalibrationGrammarContext() *CalibrationGrammarContext {
	var p = new(CalibrationGrammarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_calibrationGrammar
	return p
}

func (*CalibrationGrammarContext) IsCalibrationGrammarContext() {}

func NewCalibrationGrammarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalibrationGrammarContext {
	var p = new(CalibrationGrammarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_calibrationGrammar

	return p
}

func (s *CalibrationGrammarContext) GetParser() antlr.Parser { return s.parser }

func (s *CalibrationGrammarContext) Identifier() antlr.TerminalNode {
	return s.GetToken(qasm3ParserIdentifier, 0)
}

func (s *CalibrationGrammarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalibrationGrammarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalibrationGrammarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCalibrationGrammar(s)
	}
}

func (s *CalibrationGrammarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCalibrationGrammar(s)
	}
}

func (p *qasm3Parser) CalibrationGrammar() (localctx ICalibrationGrammarContext) {
	localctx = NewCalibrationGrammarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, qasm3ParserRULE_calibrationGrammar)
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
		p.SetState(723)
		_la = p.GetTokenStream().LA(1)

		if !(_la == qasm3ParserT__91 || _la == qasm3ParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICalibrationArgumentListContext is an interface to support dynamic dispatch.
type ICalibrationArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalibrationArgumentListContext differentiates from other interfaces.
	IsCalibrationArgumentListContext()
}

type CalibrationArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalibrationArgumentListContext() *CalibrationArgumentListContext {
	var p = new(CalibrationArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qasm3ParserRULE_calibrationArgumentList
	return p
}

func (*CalibrationArgumentListContext) IsCalibrationArgumentListContext() {}

func NewCalibrationArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalibrationArgumentListContext {
	var p = new(CalibrationArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qasm3ParserRULE_calibrationArgumentList

	return p
}

func (s *CalibrationArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *CalibrationArgumentListContext) ClassicalArgumentList() IClassicalArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassicalArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassicalArgumentListContext)
}

func (s *CalibrationArgumentListContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CalibrationArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalibrationArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalibrationArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.EnterCalibrationArgumentList(s)
	}
}

func (s *CalibrationArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qasm3Listener); ok {
		listenerT.ExitCalibrationArgumentList(s)
	}
}

func (p *qasm3Parser) CalibrationArgumentList() (localctx ICalibrationArgumentListContext) {
	localctx = NewCalibrationArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, qasm3ParserRULE_calibrationArgumentList)

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

	p.SetState(727)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(725)
			p.ClassicalArgumentList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(726)
			p.ExpressionList()
		}

	}

	return localctx
}

func (p *qasm3Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 52:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *qasm3Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
