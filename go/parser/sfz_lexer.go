// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 36, 326,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 31, 6, 31, 286, 10, 31, 13, 31, 14, 31, 287, 3, 32, 3, 32,
	3, 32, 5, 32, 293, 10, 32, 3, 33, 6, 33, 296, 10, 33, 13, 33, 14, 33, 297,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 306, 10, 34, 12, 34, 14,
	34, 309, 11, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	3, 35, 7, 35, 320, 10, 35, 12, 35, 14, 35, 323, 11, 35, 3, 35, 3, 35, 3,
	307, 2, 36, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 3, 2, 6, 3, 2,
	50, 59, 8, 2, 44, 44, 47, 59, 67, 92, 94, 94, 97, 97, 99, 124, 4, 2, 11,
	11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 330, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7,
	75, 3, 2, 2, 2, 9, 82, 3, 2, 2, 2, 11, 88, 3, 2, 2, 2, 13, 96, 3, 2, 2,
	2, 15, 103, 3, 2, 2, 2, 17, 109, 3, 2, 2, 2, 19, 116, 3, 2, 2, 2, 21, 123,
	3, 2, 2, 2, 23, 128, 3, 2, 2, 2, 25, 136, 3, 2, 2, 2, 27, 138, 3, 2, 2,
	2, 29, 152, 3, 2, 2, 2, 31, 162, 3, 2, 2, 2, 33, 170, 3, 2, 2, 2, 35, 176,
	3, 2, 2, 2, 37, 182, 3, 2, 2, 2, 39, 186, 3, 2, 2, 2, 41, 192, 3, 2, 2,
	2, 43, 198, 3, 2, 2, 2, 45, 214, 3, 2, 2, 2, 47, 221, 3, 2, 2, 2, 49, 232,
	3, 2, 2, 2, 51, 245, 3, 2, 2, 2, 53, 256, 3, 2, 2, 2, 55, 265, 3, 2, 2,
	2, 57, 273, 3, 2, 2, 2, 59, 282, 3, 2, 2, 2, 61, 285, 3, 2, 2, 2, 63, 292,
	3, 2, 2, 2, 65, 295, 3, 2, 2, 2, 67, 301, 3, 2, 2, 2, 69, 315, 3, 2, 2,
	2, 71, 72, 7, 62, 2, 2, 72, 4, 3, 2, 2, 2, 73, 74, 7, 64, 2, 2, 74, 6,
	3, 2, 2, 2, 75, 76, 7, 116, 2, 2, 76, 77, 7, 103, 2, 2, 77, 78, 7, 105,
	2, 2, 78, 79, 7, 107, 2, 2, 79, 80, 7, 113, 2, 2, 80, 81, 7, 112, 2, 2,
	81, 8, 3, 2, 2, 2, 82, 83, 7, 105, 2, 2, 83, 84, 7, 116, 2, 2, 84, 85,
	7, 113, 2, 2, 85, 86, 7, 119, 2, 2, 86, 87, 7, 114, 2, 2, 87, 10, 3, 2,
	2, 2, 88, 89, 7, 101, 2, 2, 89, 90, 7, 113, 2, 2, 90, 91, 7, 112, 2, 2,
	91, 92, 7, 118, 2, 2, 92, 93, 7, 116, 2, 2, 93, 94, 7, 113, 2, 2, 94, 95,
	7, 110, 2, 2, 95, 12, 3, 2, 2, 2, 96, 97, 7, 105, 2, 2, 97, 98, 7, 110,
	2, 2, 98, 99, 7, 113, 2, 2, 99, 100, 7, 100, 2, 2, 100, 101, 7, 99, 2,
	2, 101, 102, 7, 110, 2, 2, 102, 14, 3, 2, 2, 2, 103, 104, 7, 101, 2, 2,
	104, 105, 7, 119, 2, 2, 105, 106, 7, 116, 2, 2, 106, 107, 7, 120, 2, 2,
	107, 108, 7, 103, 2, 2, 108, 16, 3, 2, 2, 2, 109, 110, 7, 103, 2, 2, 110,
	111, 7, 104, 2, 2, 111, 112, 7, 104, 2, 2, 112, 113, 7, 103, 2, 2, 113,
	114, 7, 101, 2, 2, 114, 115, 7, 118, 2, 2, 115, 18, 3, 2, 2, 2, 116, 117,
	7, 111, 2, 2, 117, 118, 7, 99, 2, 2, 118, 119, 7, 117, 2, 2, 119, 120,
	7, 118, 2, 2, 120, 121, 7, 103, 2, 2, 121, 122, 7, 116, 2, 2, 122, 20,
	3, 2, 2, 2, 123, 124, 7, 111, 2, 2, 124, 125, 7, 107, 2, 2, 125, 126, 7,
	102, 2, 2, 126, 127, 7, 107, 2, 2, 127, 22, 3, 2, 2, 2, 128, 129, 7, 117,
	2, 2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 111, 2, 2, 131, 132, 7, 114,
	2, 2, 132, 133, 7, 110, 2, 2, 133, 134, 7, 103, 2, 2, 134, 135, 7, 116,
	2, 2, 135, 24, 3, 2, 2, 2, 136, 137, 7, 63, 2, 2, 137, 26, 3, 2, 2, 2,
	138, 139, 7, 99, 2, 2, 139, 140, 7, 111, 2, 2, 140, 141, 7, 114, 2, 2,
	141, 142, 7, 103, 2, 2, 142, 143, 7, 105, 2, 2, 143, 144, 7, 97, 2, 2,
	144, 145, 7, 116, 2, 2, 145, 146, 7, 103, 2, 2, 146, 147, 7, 110, 2, 2,
	147, 148, 7, 103, 2, 2, 148, 149, 7, 99, 2, 2, 149, 150, 7, 117, 2, 2,
	150, 151, 7, 103, 2, 2, 151, 28, 3, 2, 2, 2, 152, 153, 7, 100, 2, 2, 153,
	154, 7, 103, 2, 2, 154, 155, 7, 112, 2, 2, 155, 156, 7, 102, 2, 2, 156,
	157, 7, 97, 2, 2, 157, 158, 7, 102, 2, 2, 158, 159, 7, 113, 2, 2, 159,
	160, 7, 121, 2, 2, 160, 161, 7, 112, 2, 2, 161, 30, 3, 2, 2, 2, 162, 163,
	7, 100, 2, 2, 163, 164, 7, 103, 2, 2, 164, 165, 7, 112, 2, 2, 165, 166,
	7, 102, 2, 2, 166, 167, 7, 97, 2, 2, 167, 168, 7, 119, 2, 2, 168, 169,
	7, 114, 2, 2, 169, 32, 3, 2, 2, 2, 170, 171, 7, 106, 2, 2, 171, 172, 7,
	107, 2, 2, 172, 173, 7, 109, 2, 2, 173, 174, 7, 103, 2, 2, 174, 175, 7,
	123, 2, 2, 175, 34, 3, 2, 2, 2, 176, 177, 7, 106, 2, 2, 177, 178, 7, 107,
	2, 2, 178, 179, 7, 120, 2, 2, 179, 180, 7, 103, 2, 2, 180, 181, 7, 110,
	2, 2, 181, 36, 3, 2, 2, 2, 182, 183, 7, 109, 2, 2, 183, 184, 7, 103, 2,
	2, 184, 185, 7, 123, 2, 2, 185, 38, 3, 2, 2, 2, 186, 187, 7, 110, 2, 2,
	187, 188, 7, 113, 2, 2, 188, 189, 7, 109, 2, 2, 189, 190, 7, 103, 2, 2,
	190, 191, 7, 123, 2, 2, 191, 40, 3, 2, 2, 2, 192, 193, 7, 110, 2, 2, 193,
	194, 7, 113, 2, 2, 194, 195, 7, 120, 2, 2, 195, 196, 7, 103, 2, 2, 196,
	197, 7, 110, 2, 2, 197, 42, 3, 2, 2, 2, 198, 199, 7, 114, 2, 2, 199, 200,
	7, 107, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202, 7, 101, 2, 2, 202, 203,
	7, 106, 2, 2, 203, 204, 7, 97, 2, 2, 204, 205, 7, 109, 2, 2, 205, 206,
	7, 103, 2, 2, 206, 207, 7, 123, 2, 2, 207, 208, 7, 101, 2, 2, 208, 209,
	7, 103, 2, 2, 209, 210, 7, 112, 2, 2, 210, 211, 7, 118, 2, 2, 211, 212,
	7, 103, 2, 2, 212, 213, 7, 116, 2, 2, 213, 44, 3, 2, 2, 2, 214, 215, 7,
	117, 2, 2, 215, 216, 7, 99, 2, 2, 216, 217, 7, 111, 2, 2, 217, 218, 7,
	114, 2, 2, 218, 219, 7, 110, 2, 2, 219, 220, 7, 103, 2, 2, 220, 46, 3,
	2, 2, 2, 221, 222, 7, 117, 2, 2, 222, 223, 7, 103, 2, 2, 223, 224, 7, 115,
	2, 2, 224, 225, 7, 97, 2, 2, 225, 226, 7, 110, 2, 2, 226, 227, 7, 103,
	2, 2, 227, 228, 7, 112, 2, 2, 228, 229, 7, 105, 2, 2, 229, 230, 7, 118,
	2, 2, 230, 231, 7, 106, 2, 2, 231, 48, 3, 2, 2, 2, 232, 233, 7, 117, 2,
	2, 233, 234, 7, 103, 2, 2, 234, 235, 7, 115, 2, 2, 235, 236, 7, 97, 2,
	2, 236, 237, 7, 114, 2, 2, 237, 238, 7, 113, 2, 2, 238, 239, 7, 117, 2,
	2, 239, 240, 7, 107, 2, 2, 240, 241, 7, 118, 2, 2, 241, 242, 7, 107, 2,
	2, 242, 243, 7, 113, 2, 2, 243, 244, 7, 112, 2, 2, 244, 50, 3, 2, 2, 2,
	245, 246, 7, 117, 2, 2, 246, 247, 7, 121, 2, 2, 247, 248, 7, 97, 2, 2,
	248, 249, 7, 102, 2, 2, 249, 250, 7, 103, 2, 2, 250, 251, 7, 104, 2, 2,
	251, 252, 7, 99, 2, 2, 252, 253, 7, 119, 2, 2, 253, 254, 7, 110, 2, 2,
	254, 255, 7, 118, 2, 2, 255, 52, 3, 2, 2, 2, 256, 257, 7, 117, 2, 2, 257,
	258, 7, 121, 2, 2, 258, 259, 7, 97, 2, 2, 259, 260, 7, 106, 2, 2, 260,
	261, 7, 107, 2, 2, 261, 262, 7, 109, 2, 2, 262, 263, 7, 103, 2, 2, 263,
	264, 7, 123, 2, 2, 264, 54, 3, 2, 2, 2, 265, 266, 7, 117, 2, 2, 266, 267,
	7, 121, 2, 2, 267, 268, 7, 97, 2, 2, 268, 269, 7, 110, 2, 2, 269, 270,
	7, 99, 2, 2, 270, 271, 7, 117, 2, 2, 271, 272, 7, 118, 2, 2, 272, 56, 3,
	2, 2, 2, 273, 274, 7, 117, 2, 2, 274, 275, 7, 121, 2, 2, 275, 276, 7, 97,
	2, 2, 276, 277, 7, 110, 2, 2, 277, 278, 7, 113, 2, 2, 278, 279, 7, 109,
	2, 2, 279, 280, 7, 103, 2, 2, 280, 281, 7, 123, 2, 2, 281, 58, 3, 2, 2,
	2, 282, 283, 9, 2, 2, 2, 283, 60, 3, 2, 2, 2, 284, 286, 9, 3, 2, 2, 285,
	284, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288,
	3, 2, 2, 2, 288, 62, 3, 2, 2, 2, 289, 290, 7, 15, 2, 2, 290, 293, 7, 12,
	2, 2, 291, 293, 7, 12, 2, 2, 292, 289, 3, 2, 2, 2, 292, 291, 3, 2, 2, 2,
	293, 64, 3, 2, 2, 2, 294, 296, 9, 4, 2, 2, 295, 294, 3, 2, 2, 2, 296, 297,
	3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299, 3, 2,
	2, 2, 299, 300, 8, 33, 2, 2, 300, 66, 3, 2, 2, 2, 301, 302, 7, 49, 2, 2,
	302, 303, 7, 44, 2, 2, 303, 307, 3, 2, 2, 2, 304, 306, 11, 2, 2, 2, 305,
	304, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 307, 305,
	3, 2, 2, 2, 308, 310, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311, 7, 44,
	2, 2, 311, 312, 7, 49, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 8, 34, 2,
	2, 314, 68, 3, 2, 2, 2, 315, 316, 7, 49, 2, 2, 316, 317, 7, 49, 2, 2, 317,
	321, 3, 2, 2, 2, 318, 320, 10, 5, 2, 2, 319, 318, 3, 2, 2, 2, 320, 323,
	3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 324, 3, 2,
	2, 2, 323, 321, 3, 2, 2, 2, 324, 325, 8, 35, 2, 2, 325, 70, 3, 2, 2, 2,
	8, 2, 287, 292, 297, 307, 321, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'<'", "'>'", "'region'", "'group'", "'control'", "'global'", "'curve'",
	"'effect'", "'master'", "'midi'", "'sampler'", "'='", "'ampeg_release'",
	"'bend_down'", "'bend_up'", "'hikey'", "'hivel'", "'key'", "'lokey'", "'lovel'",
	"'pitch_keycenter'", "'sample'", "'seq_length'", "'seq_position'", "'sw_default'",
	"'sw_hikey'", "'sw_last'", "'sw_lokey'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "Digit", "Text", "Newline",
	"Whitespace", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "Digit", "Text", "Newline", "Whitespace", "BlockComment",
	"LineComment",
}

type SfzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSfzLexer(input antlr.CharStream) *SfzLexer {

	l := new(SfzLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sfz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SfzLexer tokens.
const (
	SfzLexerT__0         = 1
	SfzLexerT__1         = 2
	SfzLexerT__2         = 3
	SfzLexerT__3         = 4
	SfzLexerT__4         = 5
	SfzLexerT__5         = 6
	SfzLexerT__6         = 7
	SfzLexerT__7         = 8
	SfzLexerT__8         = 9
	SfzLexerT__9         = 10
	SfzLexerT__10        = 11
	SfzLexerT__11        = 12
	SfzLexerT__12        = 13
	SfzLexerT__13        = 14
	SfzLexerT__14        = 15
	SfzLexerT__15        = 16
	SfzLexerT__16        = 17
	SfzLexerT__17        = 18
	SfzLexerT__18        = 19
	SfzLexerT__19        = 20
	SfzLexerT__20        = 21
	SfzLexerT__21        = 22
	SfzLexerT__22        = 23
	SfzLexerT__23        = 24
	SfzLexerT__24        = 25
	SfzLexerT__25        = 26
	SfzLexerT__26        = 27
	SfzLexerT__27        = 28
	SfzLexerDigit        = 29
	SfzLexerText         = 30
	SfzLexerNewline      = 31
	SfzLexerWhitespace   = 32
	SfzLexerBlockComment = 33
	SfzLexerLineComment  = 34
)
