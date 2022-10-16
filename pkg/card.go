package pkg

import (
	"fmt"
	"strings"
)

var (
	Cards = []Card{
		{
			No:   1,
			Name: "ﾋｰﾛｰｼｭｰﾀｰ",
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM},
				{NM, NM, NM, SP, NM},
				{__, NM, __, __, __},
				{NM, __, __, __, __},
			},
		},
		{
			No:   2,
			Name: "ﾎﾞｰﾙﾄﾞﾏｰｶｰ",
			Blocks: [][]Block{
				{NM, SP, __, NM},
				{NM, NM, NM, NM},
				{__, NM, __, NM},
			},
		},
		{
			No:   3,
			Name: "わかばｼｭｰﾀｰ",
			Blocks: [][]Block{
				{NM, SP, NM},
				{NM, __, __},
			},
		},
		{
			No:   4,
			Name: "ｼｬｰﾌﾟﾏｰｶｰ",
			Blocks: [][]Block{
				{__, __, NM, __},
				{NM, NM, NM, SP},
				{__, NM, NM, __},
			},
		},
		{
			No:   5,
			Name: "ﾌﾟﾛﾓﾃﾞﾗｰMG",
			Blocks: [][]Block{
				{NM, SP, __, __},
				{NM, NM, NM, NM},
				{__, NM, __, __},
			},
		},
		{
			No:   6,
			Name: "ｽﾌﾟﾗｼｭｰﾀｰ",
			Blocks: [][]Block{
				{NM, NM, SP, __},
				{NM, NM, NM, SP},
				{NM, __, __, __},
			},
		},
		{
			No:   7,
			Name: ".52ｶﾞﾛﾝ",
			Blocks: [][]Block{
				{NM, NM, NM, SP},
				{__, NM, NM, __},
				{NM, __, __, __},
			},
		},
		{
			No:   8,
			Name: "",
			Blocks: [][]Block{
				{NM, NM, SP, NM},
				{NM, __, __, __},
			},
		},
		{
			No:   9,
			Name: "",
			Blocks: [][]Block{
				{__, __, __, NM, SP, __},
				{NM, NM, NM, NM, NM, NM},
				{__, __, NM, __, __, __},
			},
		},
		{
			No:   10,
			Name: "",
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, SP},
				{__, NM, NM, NM, __, __},
				{NM, __, __, __, __, __},
			},
		},
		{
			No: 11,
			Blocks: [][]Block{
				{NM, SP, NM, NM, NM, NM},
				{NM, __, __, __, __, __},
				{NM, NM, NM, __, __, __},
			},
		},
		{
			No: 12,
			Blocks: [][]Block{
				{__, __, __, SP},
				{NM, NM, NM, __},
				{__, NM, NM, __},
				{__, NM, __, __},
			},
		},
		{
			No: 13,
			Blocks: [][]Block{
				{NM, __, __, NM, SP},
				{__, NM, NM, NM, NM},
				{__, NM, __, __, __},
			},
		},
		{
			No: 14,
			Blocks: [][]Block{
				{NM, __, __, __, NM, SP},
				{__, NM, NM, NM, NM, NM},
				{__, NM, __, __, __, __},
			},
		},
		{
			No: 15,
			Blocks: [][]Block{
				{__, NM, NM, NM},
				{NM, NM, NM, SP},
				{NM, __, __, __},
			},
		},
		{
			No: 16,
			Blocks: [][]Block{
				{NM, NM, SP, NM},
				{NM, NM, NM, __},
				{NM, __, __, __},
			},
		},
		{
			No: 17,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM},
				{NM, NM, SP, __, __},
				{NM, __, __, __, __},
			},
		},
		{
			No: 18,
			Blocks: [][]Block{
				{NM, SP, NM, NM},
				{NM, NM, __, __},
				{NM, __, __, __},
			},
		},
		{
			No: 19,
			Blocks: [][]Block{
				{NM, NM, NM, SP, NM},
				{NM, NM, __, __, __},
				{NM, __, __, __, __},
			},
		},
		{
			No: 20,
			Blocks: [][]Block{
				{__, __, __, SP},
				{__, __, NM, __},
				{NM, NM, __, __},
				{NM, NM, __, __},
				{NM, __, __, __},
			},
		},
		{
			No: 21,
			Blocks: [][]Block{
				{NM, NM, NM, NM},
				{__, NM, SP, __},
				{__, __, NM, __},
			},
		},
		{
			No: 22,
			Blocks: [][]Block{
				{NM, NM, SP, NM, NM},
				{__, __, NM, NM, NM},
				{__, __, NM, __, __},
			},
		},
		{
			No: 23,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM, NM},
				{__, __, __, SP, NM, NM, NM},
				{__, __, __, NM, __, __, __},
				{__, __, __, NM, __, __, __},
			},
		},
		{
			No: 24,
			Blocks: [][]Block{
				{NM, NM, SP, NM, NM},
				{__, __, NM, NM, __},
				{__, __, NM, NM, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 25,
			Blocks: [][]Block{
				{__, __, __, SP, NM},
				{__, __, __, NM, NM},
				{__, __, NM, __, __},
				{__, NM, __, __, __},
				{NM, __, __, __, __},
			},
		},
		{
			No: 26,
			Blocks: [][]Block{
				{__, __, __, __, NM, __},
				{__, __, __, SP, NM, NM},
				{__, __, __, NM, NM, __},
				{__, __, NM, __, __, __},
				{__, NM, __, __, __, __},
				{NM, __, __, __, __, __},
			},
		},
		{
			No: 27,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM},
				{NM, SP, __, __, __, __},
			},
		},
		{
			No: 28,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM, NM},
				{__, __, SP, __, __, __, __},
			},
		},
		{
			No: 29,
			Blocks: [][]Block{
				{__, __, NM, SP, __, __, __},
				{NM, NM, NM, NM, NM, NM, NM},
				{__, __, NM, __, __, __, __},
			},
		},
		{
			No: 30,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM, NM},
				{NM, NM, NM, __, __, SP, __},
			},
		},
		{
			No: 31,
			Blocks: [][]Block{
				{__, __, __, NM, SP, __, __},
				{NM, NM, NM, NM, NM, NM, NM},
				{NM, NM, NM, __, __, NM, __},
			},
		},
		{
			No: 32,
			Blocks: [][]Block{
				{NM, __, __, __, NM},
				{NM, SP, NM, NM, NM},
			},
		},
		{
			No: 33,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM},
				{NM, NM, SP, __, NM, __},
				{__, __, __, NM, __, __},
			},
		},
		{
			No: 34,
			Blocks: [][]Block{
				{NM, __, __},
				{__, SP, NM},
				{NM, NM, NM},
			},
		},
		{
			No: 35,
			Blocks: [][]Block{
				{NM, __},
				{NM, SP},
				{NM, NM},
			},
		},
		{
			No: 36,
			Blocks: [][]Block{
				{NM, SP, NM},
				{__, NM, NM},
				{__, NM, NM},
			},
		},
		{
			No: 37,
			Blocks: [][]Block{
				{NM, __, __, __, __},
				{SP, NM, NM, NM, NM},
				{NM, NM, NM, NM, __},
			},
		},
		{
			No: 38,
			Blocks: [][]Block{
				{NM, __, __, NM, __},
				{NM, NM, NM, NM, __},
				{NM, NM, NM, NM, SP},
			},
		},
		{
			No: 39,
			Blocks: [][]Block{
				{NM, NM, NM, NM, SP},
				{NM, NM, __, __, __},
			},
		},
		{
			No: 40,
			Blocks: [][]Block{
				{NM, NM, __, __, __},
				{NM, NM, __, __, __},
				{NM, NM, __, __, __},
				{__, NM, NM, __, __},
				{__, __, NM, SP, __},
				{__, __, __, NM, NM},
			},
		},
		{
			No: 41,
			Blocks: [][]Block{
				{NM, NM, __, __, __, __},
				{NM, NM, __, __, __, __},
				{NM, NM, NM, NM, NM, SP},
				{__, NM, NM, NM, __, __},
			},
		},
		{
			No: 42,
			Blocks: [][]Block{
				{__, NM, __, __, __},
				{NM, NM, NM, NM, __},
				{NM, NM, NM, SP, NM},
			},
		},
		{
			No: 43,
			Blocks: [][]Block{
				{NM, NM, __, __, __},
				{NM, NM, __, __, __},
				{NM, NM, NM, NM, SP},
			},
		},
		{
			No: 44,
			Blocks: [][]Block{
				{NM, NM, NM},
				{NM, __, SP},
				{__, NM, NM},
			},
		},
		{
			No: 45,
			Blocks: [][]Block{
				{__, NM, NM, NM, NM},
				{__, NM, SP, __, __},
				{NM, NM, __, __, __},
			},
		},
		{
			No: 46,
			Blocks: [][]Block{
				{__, __, NM, __},
				{NM, NM, NM, SP},
				{__, NM, NM, __},
				{__, NM, NM, __},
			},
		},
		{
			No: 47,
			Blocks: [][]Block{
				{NM, NM, __, __, __},
				{NM, NM, NM, NM, NM},
				{SP, __, __, __, __},
				{NM, NM, __, __, __},
			},
		},
		{
			No: 48,
			Blocks: [][]Block{
				{__, NM, SP, NM},
				{NM, __, NM, NM},
				{__, NM, NM, __},
			},
		},
		{
			No: 49,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, __, NM, NM, __},
				{__, NM, SP, NM, NM},
				{NM, __, NM, NM, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 50,
			Blocks: [][]Block{
				{__, NM, NM, __},
				{__, __, NM, NM},
				{__, NM, SP, NM},
				{NM, __, NM, NM},
				{__, NM, NM, __},
			},
		},
		{
			No: 51,
			Blocks: [][]Block{
				{__, NM, __, __},
				{__, __, NM, __},
				{NM, NM, __, SP},
				{NM, __, NM, __},
				{__, NM, __, __},
			},
		},
		{
			No: 52,
			Blocks: [][]Block{
				{NM, SP, NM, NM, NM},
				{NM, __, NM, __, __},
				{NM, NM, __, __, __},
				{NM, __, __, __, __},
				{NM, __, __, __, __},
			},
		},
		{
			No: 53,
			Blocks: [][]Block{
				{NM, __, NM, NM},
				{__, SP, NM, __},
				{NM, NM, NM, __},
				{NM, __, __, __},
			},
		},
		{
			No: 54,
			Blocks: [][]Block{
				{SP},
				{NM},
				{NM},
				{NM},
				{NM},
				{NM},
			},
		},
		{
			No: 55,
			Blocks: [][]Block{
				{SP},
				{NM},
				{NM},
				{NM},
				{NM},
			},
		},
		{
			No: 56,
			Blocks: [][]Block{
				{__, SP},
				{NM, NM},
			},
		},
		{
			No: 57,
			Blocks: [][]Block{
				{NM},
				{NM},
				{SP},
			},
		},
		{
			No: 58,
			Blocks: [][]Block{
				{SP},
			},
		},
		{
			No: 59,
			Blocks: [][]Block{
				{NM, __, NM},
				{__, SP, __},
			},
		},
		{
			No: 60,
			Blocks: [][]Block{
				{NM, NM, NM},
				{__, SP, __},
				{__, NM, __},
			},
		},
		{
			No: 61,
			Blocks: [][]Block{
				{NM},
				{SP},
				{NM},
			},
		},
		{
			No: 62,
			Blocks: [][]Block{
				{__, SP, __},
				{NM, NM, NM},
			},
		},
		{
			No: 63,
			Blocks: [][]Block{
				{__, NM, SP},
				{NM, NM, NM},
				{__, NM, __},
			},
		},
		{
			No: 64,
			Blocks: [][]Block{
				{NM, NM, NM},
				{NM, SP, NM},
				{__, NM, __},
				{NM, __, NM},
			},
		},
		{
			No: 65,
			Blocks: [][]Block{
				{NM, __, NM},
				{__, SP, __},
				{NM, __, NM},
			},
		},
		{
			No: 66,
			Blocks: [][]Block{
				{__, NM, __},
				{NM, __, NM},
				{__, SP, __},
			},
		},
		{
			No: 67,
			Blocks: [][]Block{
				{NM, __, __},
				{__, SP, NM},
				{__, NM, NM},
			},
		},
		{
			No: 68,
			Blocks: [][]Block{
				{__, __, SP},
				{__, NM, __},
				{NM, __, __},
				{__, NM, __},
				{__, __, NM},
			},
		},
		{
			No: 69,
			Blocks: [][]Block{
				{__, NM, __},
				{__, NM, NM},
				{SP, __, __},
			},
		},
		{
			No: 70,
			Blocks: [][]Block{
				{__, __, __, __, NM, NM},
				{__, __, NM, __, NM, NM},
				{__, NM, __, NM, __, __},
				{NM, __, NM, __, __, __},
				{__, NM, NM, __, __, __},
				{__, __, NM, __, __, __},
			},
		},
		{
			No: 71,
			Blocks: [][]Block{
				{__, __, NM, NM, __, __},
				{__, NM, __, __, NM, __},
				{NM, __, __, __, __, NM},
				{NM, __, __, __, __, NM},
				{__, NM, __, __, NM, __},
				{__, __, NM, NM, __, __},
			},
		},
		{
			No: 72,
			Blocks: [][]Block{
				{__, __, __, __, __, NM},
				{__, __, __, __, NM, __},
				{NM, NM, __, NM, __, __},
				{NM, NM, NM, __, __, __},
				{__, NM, NM, NM, __, __},
				{__, __, NM, __, __, __},
			},
		},
		{
			No: 73,
			Blocks: [][]Block{
				{NM, NM, __, __, NM, NM},
				{NM, NM, NM, NM, NM, NM},
				{__, NM, __, __, NM, __},
			},
		},
		{
			No: 74,
			Blocks: [][]Block{
				{NM, __, NM, __, __},
				{__, NM, NM, NM, NM},
				{__, NM, NM, NM, NM},
				{NM, __, NM, __, __},
			},
		},
		{
			No: 75,
			Blocks: [][]Block{
				{__, NM, NM, __},
				{NM, NM, NM, NM},
				{NM, NM, NM, NM},
				{__, NM, NM, __},
			},
		},
		{
			No: 76,
			Blocks: [][]Block{
				{NM, __, NM, __, NM},
				{__, NM, __, NM, __},
				{NM, __, __, __, NM},
				{__, NM, __, NM, __},
				{NM, __, NM, __, NM},
			},
		},
		{
			No: 77,
			Blocks: [][]Block{
				{NM, NM, __, __, __, NM},
				{NM, NM, NM, __, __, NM},
				{__, NM, NM, __, __, NM},
				{__, __, __, NM, NM, __},
			},
		},
		{
			No: 78,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{NM, __, NM, __, NM},
				{__, NM, __, NM, __},
				{__, NM, __, NM, __},
				{NM, __, NM, __, NM},
				{__, __, NM, __, __},
			},
		},
		{
			No: 79,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM, NM},
				{NM, __, NM, NM, __, NM},
				{__, __, NM, NM, __, __},
			},
		},
		{
			No: 80,
			Blocks: [][]Block{
				{__, __, __, NM, NM, __},
				{NM, NM, NM, NM, NM, NM},
				{__, __, NM, NM, NM, NM},
			},
		},
		{
			No: 81,
			Blocks: [][]Block{
				{__, NM, __, __, NM, __},
				{__, __, NM, NM, __, __},
				{__, NM, __, __, NM, __},
				{NM, __, __, __, __, NM},
				{NM, NM, __, __, NM, NM},
			},
		},
		{
			No: 82,
			Blocks: [][]Block{
				{__, __, NM, __, __, __},
				{__, NM, NM, NM, __, NM},
				{NM, NM, NM, NM, NM, NM},
				{__, __, NM, __, __, __},
			},
		},
		{
			No: 83,
			Blocks: [][]Block{
				{__, __, NM, NM, __, __},
				{__, __, NM, NM, __, __},
				{NM, NM, __, __, NM, NM},
				{NM, NM, __, __, NM, NM},
			},
		},
		{
			No: 84,
			Blocks: [][]Block{
				{__, NM, NM, NM, __},
				{__, NM, __, NM, __},
				{NM, __, NM, __, NM},
				{__, NM, NM, NM, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 85,
			Blocks: [][]Block{
				{__, NM, NM, NM, __},
				{NM, NM, NM, NM, NM},
				{NM, __, NM, __, NM},
				{__, SP, __, NM, __},
			},
		},
		{
			No: 86,
			Blocks: [][]Block{
				{__, NM, NM, SP, __},
				{NM, __, NM, __, NM},
				{NM, NM, NM, NM, NM},
				{NM, __, NM, __, __},
				{__, NM, NM, __, __},
			},
		},
		{
			No: 87,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{__, NM, NM, NM, __},
				{NM, __, SP, __, NM},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 88,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{__, NM, NM, NM, __},
				{NM, NM, NM, NM, NM},
				{SP, __, NM, __, NM},
			},
		},
		{
			No: 89,
			Blocks: [][]Block{
				{__, NM, NM, NM, __, __},
				{__, NM, NM, NM, NM, __},
				{__, __, NM, SP, __, __},
				{__, NM, __, __, NM, __},
				{__, NM, __, __, NM, __},
				{NM, __, __, __, __, NM},
			},
		},
		{
			No: 90,
			Blocks: [][]Block{
				{__, NM, __, NM, __},
				{NM, __, NM, __, NM},
				{__, NM, __, NM, __},
				{NM, __, SP, __, NM},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 91,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, SP, __, __, __},
				{NM, __, NM, __, __},
				{NM, NM, NM, __, __},
				{NM, NM, NM, NM, NM},
				{__, __, NM, __, NM},
			},
		},
		{
			No: 92,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{NM, __, NM, __, NM},
				{NM, NM, SP, NM, NM},
				{__, NM, NM, NM, __},
			},
		},
		{
			No: 93,
			Blocks: [][]Block{
				{SP, __},
				{NM, NM},
			},
		},
		{
			No: 94,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{__, NM, NM, NM, __},
				{NM, __, SP, __, NM},
			},
		},
		{
			No: 95,
			Blocks: [][]Block{
				{__, NM, NM, NM, __, __, __},
				{__, NM, NM, NM, NM, __, __},
				{NM, NM, SP, NM, __, NM, __},
				{__, __, NM, __, __, __, NM},
				{__, __, __, NM, __, __, __},
				{__, __, __, __, NM, __, __},
			},
		},
		{
			No: 96,
			Blocks: [][]Block{
				{NM, NM, __, __, __},
				{NM, NM, NM, NM, __},
				{__, NM, NM, SP, __},
				{__, __, NM, __, NM},
				{NM, __, NM, __, __},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 97,
			Blocks: [][]Block{
				{__, __, __, NM, NM},
				{__, NM, NM, NM, NM},
				{__, SP, NM, NM, __},
				{NM, __, NM, __, __},
				{__, __, NM, __, NM},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 98,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{SP, __, NM, __, NM},
				{NM, NM, __, NM, NM},
				{__, NM, NM, NM, __},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 99,
			Blocks: [][]Block{
				{__, NM, __, NM, __},
				{__, NM, NM, NM, __},
				{NM, __, NM, __, SP},
				{NM, NM, __, NM, NM},
				{__, NM, __, NM, __},
				{__, NM, __, NM, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 100,
			Blocks: [][]Block{
				{__, NM, __, __, NM, __},
				{NM, __, NM, NM, NM, __},
				{__, NM, NM, NM, __, __},
				{__, NM, NM, NM, __, NM},
				{NM, NM, __, __, SP, __},
				{__, __, __, NM, __, __},
			},
		},
		{
			No: 101,
			Blocks: [][]Block{
				{__, NM, __, NM, __},
				{NM, NM, NM, NM, NM},
				{NM, __, NM, __, NM},
				{NM, NM, __, SP, NM},
			},
		},
		{
			No: 102,
			Blocks: [][]Block{
				{NM, __, NM},
				{NM, NM, NM},
				{NM, NM, SP},
			},
		},
		{
			No: 103,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, __, NM, __, __},
				{NM, NM, NM, NM, NM},
				{__, __, NM, SP, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 104,
			Blocks: [][]Block{
				{__, NM, NM, NM, NM},
				{__, __, NM, NM, NM},
				{__, NM, __, NM, NM},
				{NM, __, SP, __, NM},
				{__, NM, __, __, __},
			},
		},
		{
			No: 105,
			Blocks: [][]Block{
				{__, __, SP, NM, __, __},
				{__, NM, NM, NM, NM, NM},
				{NM, NM, NM, NM, NM, __},
				{__, __, NM, NM, __, __},
			},
		},
		{
			No: 106,
			Blocks: [][]Block{
				{NM, __, NM},
				{NM, SP, __},
				{NM, __, NM},
				{NM, NM, NM},
			},
		},
		{
			No: 107,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{NM, NM, SP, NM, NM},
				{NM, NM, NM, NM, NM},
				{__, __, NM, __, __},
			},
		},
		{
			No: 108,
			Blocks: [][]Block{
				{__, NM, NM, NM},
				{NM, NM, NM, __},
				{SP, NM, __, __},
				{NM, __, __, __},
			},
		},
		{
			No: 109,
			Blocks: [][]Block{
				{__, __, SP, __},
				{__, __, __, NM},
				{__, NM, NM, NM},
				{NM, NM, NM, NM},
			},
		},
		{
			No: 110,
			Blocks: [][]Block{
				{__, SP, __, NM, __},
				{NM, NM, NM, NM, NM},
				{NM, NM, NM, NM, NM},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 111,
			Blocks: [][]Block{
				{NM, NM, SP, NM, __, NM},
				{NM, NM, __, NM, NM, NM},
			},
		},
		{
			No: 112,
			Blocks: [][]Block{
				{NM, NM, SP, NM, NM},
				{NM, NM, __, NM, NM},
				{NM, __, __, __, NM},
			},
		},
		{
			No: 113,
			Blocks: [][]Block{
				{__, SP},
				{NM, __},
			},
		},
		{
			No: 114,
			Blocks: [][]Block{
				{NM, NM, SP, __, NM, __},
				{__, NM, NM, NM, NM, NM},
				{NM, NM, NM, __, NM, __},
			},
		},
		{
			No: 115,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{__, NM, NM, NM, __},
				{__, NM, SP, NM, __},
				{NM, __, __, __, NM},
			},
		},
		{
			No: 116,
			Blocks: [][]Block{
				{__, __, __, NM, __},
				{NM, __, SP, NM, NM},
				{__, NM, NM, NM, __},
				{NM, NM, NM, NM, NM},
			},
		},
		{
			No: 117,
			Blocks: [][]Block{
				{NM, __, __},
				{NM, NM, NM},
				{NM, __, __},
				{NM, SP, NM},
			},
		},
		{
			No: 118,
			Blocks: [][]Block{
				{__, __, __, NM},
				{__, __, NM, NM},
				{__, NM, __, SP},
				{NM, NM, NM, NM},
			},
		},
		{
			No: 119,
			Blocks: [][]Block{
				{__, NM, __, NM, __},
				{NM, __, __, __, NM},
				{__, NM, __, NM, __},
				{__, __, SP, __, __},
				{__, NM, __, NM, __},
			},
		},
		{
			No: 120,
			Blocks: [][]Block{
				{__, NM, __},
				{__, SP, NM},
				{NM, NM, NM},
				{__, NM, NM},
			},
		},
		{
			No: 121,
			Blocks: [][]Block{
				{__, __, NM, __},
				{NM, __, SP, NM},
				{NM, NM, NM, NM},
				{NM, __, NM, NM},
			},
		},
		{
			No: 122,
			Blocks: [][]Block{
				{NM, NM, NM, __, NM},
				{__, __, SP, NM, __},
				{__, NM, NM, NM, __},
				{__, __, NM, NM, __},
			},
		},
		{
			No: 123,
			Blocks: [][]Block{
				{__, NM, __},
				{__, SP, NM},
				{NM, NM, NM},
				{__, NM, NM},
				{__, NM, __},
				{__, NM, __},
			},
		},
		{
			No: 124,
			Blocks: [][]Block{
				{NM, __, NM},
				{__, SP, __},
				{__, NM, __},
			},
		},
		{
			No: 125,
			Blocks: [][]Block{
				{__, __, NM, __, __, __},
				{__, NM, NM, NM, __, __},
				{NM, NM, NM, NM, NM, __},
				{__, NM, NM, NM, __, SP},
			},
		},
		{
			No: 126,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{NM, NM, NM, NM, NM},
				{NM, __, __, __, NM},
				{SP, __, __, __, NM},
				{NM, __, __, __, NM},
			},
		},
		{
			No: 127,
			Blocks: [][]Block{
				{SP, NM},
			},
		},
		{
			No: 128,
			Blocks: [][]Block{
				{NM, NM, NM, NM},
				{NM, NM, NM, NM},
				{__, SP, __, __},
				{NM, __, NM, __},
			},
		},
		{
			No: 129,
			Blocks: [][]Block{
				{__, __, __, NM},
				{__, __, NM, NM},
				{NM, SP, __, __},
				{NM, NM, __, __},
			},
		},
		{
			No: 130,
			Blocks: [][]Block{
				{__, __, NM, NM},
				{__, __, NM, __},
				{__, __, NM, SP},
				{NM, NM, NM, NM},
				{__, __, NM, __},
			},
		},
		{
			No: 131,
			Blocks: [][]Block{
				{__, __, NM, NM, __},
				{__, __, NM, __, __},
				{__, __, NM, SP, __},
				{NM, NM, NM, NM, NM},
				{__, __, NM, __, NM},
			},
		},
		{
			No: 131,
			Blocks: [][]Block{
				{__, __, NM, NM, __},
				{__, __, NM, __, __},
				{__, __, NM, SP, __},
				{NM, NM, NM, NM, NM},
				{__, __, NM, __, NM},
			},
		},
		{
			No: 132,
			Blocks: [][]Block{
				{__, NM, __, __},
				{SP, NM, NM, NM},
			},
		},
		{
			No: 133,
			Blocks: [][]Block{
				{__, NM, __, __},
				{SP, NM, NM, NM},
				{NM, NM, NM, NM},
			},
		},
		{
			No: 134,
			Blocks: [][]Block{
				{NM, NM, NM},
				{NM, NM, NM},
				{__, SP, __},
				{NM, __, __},
			},
		},
		{
			No: 135,
			Blocks: [][]Block{
				{NM, NM},
				{SP, NM},
			},
		},
		{
			No: 136,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, __, NM, __},
				{NM, __, SP, __, NM},
				{__, NM, __, NM, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 137,
			Blocks: [][]Block{
				{NM, SP, NM, NM, NM},
				{NM, __, NM, __, NM},
				{NM, __, NM, __, NM},
				{NM, __, NM, __, NM},
			},
		},
		{
			No: 138,
			Blocks: [][]Block{
				{NM, __, __, __},
				{__, NM, NM, __},
				{NM, NM, NM, NM},
				{NM, NM, NM, NM},
				{__, NM, SP, __},
			},
		},
		{
			No: 139,
			Blocks: [][]Block{
				{__, NM, __, NM, __},
				{__, NM, NM, NM, __},
				{__, NM, SP, NM, __},
				{NM, __, NM, __, NM},
			},
		},
		{
			No: 140,
			Blocks: [][]Block{
				{__, NM, __, __, __, __},
				{NM, NM, NM, NM, __, __},
				{__, NM, NM, NM, __, NM},
				{__, NM, NM, SP, NM, __},
				{__, __, __, NM, __, NM},
				{__, __, NM, __, NM, __},
			},
		},
		{
			No: 141,
			Blocks: [][]Block{
				{__, NM, __},
				{SP, NM, __},
				{__, NM, NM},
			},
		},
		{
			No: 142,
			Blocks: [][]Block{
				{__, __, NM, __},
				{__, NM, NM, __},
				{SP, NM, NM, __},
				{__, NM, NM, NM},
			},
		},
		{
			No: 143,
			Blocks: [][]Block{
				{__, NM, __},
				{NM, NM, __},
				{__, NM, __},
				{NM, __, SP},
			},
		},
		{
			No: 144,
			Blocks: [][]Block{
				{__, SP, __, __, __},
				{NM, NM, NM, __, __},
				{__, NM, NM, NM, __},
				{__, __, NM, NM, __},
				{__, __, NM, NM, NM},
			},
		},
		{
			No: 145,
			Blocks: [][]Block{
				{NM, SP, NM, NM, NM, NM, __},
				{NM, NM, __, __, __, NM, NM},
			},
		},
		{
			No: 146,
			Blocks: [][]Block{
				{__, __, NM, SP, __, __},
				{NM, __, NM, NM, __, NM},
				{__, NM, NM, NM, NM, __},
			},
		},
		{
			No: 147,
			Blocks: [][]Block{
				{SP, NM},
				{__, NM},
				{__, NM},
				{__, NM},
				{__, NM},
				{__, NM},
			},
		},
		{
			No: 148,
			Blocks: [][]Block{
				{__, NM, __, __, NM, __},
				{NM, NM, __, __, NM, NM},
				{NM, NM, NM, NM, SP, NM},
			},
		},
		{
			No: 149,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, NM, NM, __},
				{NM, NM, NM, NM, NM},
				{NM, __, SP, __, NM},
			},
		},
		{
			No: 150,
			Blocks: [][]Block{
				{NM, NM, __, NM, NM},
				{NM, NM, SP, NM, NM},
				{NM, NM, __, NM, NM},
			},
		},
		{
			No: 151,
			Blocks: [][]Block{
				{NM, NM, __, NM, NM},
				{__, __, NM, __, __},
				{__, __, NM, __, __},
				{__, __, NM, __, __},
				{__, __, SP, __, __},
				{__, __, NM, __, __},
			},
		},
		{
			No: 152,
			Blocks: [][]Block{
				{__, NM, __, NM},
				{__, __, NM, __},
				{NM, NM, NM, __},
				{__, NM, NM, NM},
				{__, SP, NM, __},
				{__, __, NM, __},
			},
		},
		{
			No: 153,
			Blocks: [][]Block{
				{__, __, SP, __, __},
				{NM, NM, NM, NM, NM},
			},
		},
		{
			No: 154,
			Blocks: [][]Block{
				{__, __, __, NM, __, __},
				{__, __, NM, NM, NM, __},
				{SP, NM, __, NM, NM, __},
				{NM, NM, __, NM, NM, NM},
			},
		},
		{
			No: 155,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{NM, NM, __, __, __},
				{__, NM, __, NM, __},
				{NM, NM, SP, __, NM},
				{__, NM, NM, NM, __},
			},
		},
		{
			No: 156,
			Blocks: [][]Block{
				{__, NM, NM, NM, NM},
				{SP, NM, NM, NM, NM},
				{__, __, NM, NM, __},
				{__, NM, __, __, NM},
				{NM, NM, SP, __, NM},
				{__, NM, NM, NM, __},
			},
		},
		{
			No: 157,
			Blocks: [][]Block{
				{NM, NM, NM, NM, NM},
				{NM, NM, NM, NM, NM},
				{__, __, NM, __, __},
				{__, __, SP, __, __},
			},
		},
		{
			No: 158,
			Blocks: [][]Block{
				{__, NM, __},
				{NM, __, NM},
				{NM, __, NM},
				{NM, __, NM},
				{NM, SP, NM},
				{NM, NM, NM},
			},
		},
		{
			No: 159,
			Blocks: [][]Block{
				{__, __, __, NM},
				{__, NM, NM, __},
				{__, NM, SP, NM},
				{NM, __, NM, NM},
			},
		},
		{
			No: 160,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, __, NM, __, __},
				{NM, NM, SP, NM, NM},
				{__, NM, NM, NM, __},
				{__, NM, NM, NM, __},
			},
		},
		{
			No: 161,
			Blocks: [][]Block{
				{__, __, NM, __, __},
				{__, NM, __, __, __},
				{__, NM, SP, NM, NM},
				{NM, NM, NM, NM, __},
				{__, __, __, NM, NM},
			},
		},
		{
			No: 162,
			Blocks: [][]Block{
				{__, SP, NM, NM, __},
				{NM, NM, NM, NM, NM},
				{__, NM, NM, NM, __},
			},
		},
	}

	NoCardMap = func() map[int]Card {
		m := make(map[int]Card, len(Cards))
		for _, c := range Cards {
			m[c.No] = c
		}
		return m
	}()
)

type Card struct {
	No     int
	Name   string
	Blocks [][]Block
}

func (p Card) Turn(d Dir) Card {
	dst := p
	for i := 0; i < int(d); i++ {
		src := dst
		dst.Blocks = make([][]Block, 0, src.Width())
		for x := src.Width() - 1; x >= 0; x-- {
			bb := make([]Block, 0, src.Height())
			for y := 0; y < src.Height(); y++ {
				bb = append(bb, src.Blocks[y][x])
			}
			dst.Blocks = append(dst.Blocks, bb)
		}
	}
	return dst
}

func (p Card) GetFilledPoint() []Pos {
	ret := []Pos{}
	for y, bb := range p.Blocks {
		for x, b := range bb {
			if b != __ {
				ret = append(ret, Pos{X: x, Y: y})
			}
		}
	}
	return ret
}

func (c Card) Width() int {
	max := 0
	for _, bb := range c.Blocks {
		l := len(bb)
		if l > max {
			max = l
		}
	}
	return max
}

func (c Card) Height() int {
	return len(c.Blocks)
}

func (c Card) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("No.%d\n", c.No))
	b.WriteString(fmt.Sprintf("%s\n\n", c.Name))
	offsetY := (8 - c.Height()) / 2
	offsetX := (8 - c.Width()) / 2
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if y < offsetY || y >= offsetY+c.Height() ||
				x < offsetX || x >= offsetX+c.Width() {
				b.WriteString(__.String())
				continue
			}
			b.WriteString(c.Blocks[y-offsetY][x-offsetX].String())
		}
		b.WriteRune('\n')
	}
	return b.String()
}
