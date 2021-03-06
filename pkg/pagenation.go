package pkg

import (
	"BE-SISFO-KLINIK/models"
	"BE-SISFO-KLINIK/repositories"
)

type ResultData struct {
	Pages int64
	Data  []models.Pasien
}
type ResultDataPemeriksaan struct {
	Pages int64
	Data  []models.Pemeriksaan
}
type ResultDataObat struct {
	Pages int64
	Data  []repositories.ResultAllObats
}

// func CreatePagenation(data []models.Pasien, lengthData int64) [][]models.Pasien {

// 	var i, page, appendIdx int64
// 	var value []models.Pasien
// 	var result [][]models.Pasien
// 	page = 0
// 	appendIdx = 0
// 	for i = 0; i < lengthData; i++ {
// 		if appendIdx == 10 {
// 			page++
// 			appendIdx = 0
// 			result = append(result, value)
// 			value = nil
// 			value = append(value, data[i])

// 			appendIdx++
// 		} else {
// 			value = append(value, data[i])
// 			appendIdx++
// 		}
// 		// fmt.Println(appendIdx)
// 		//fmt.Println("Data ke-", i, "Masuk Ke Page-", page)
// 		// fmt.Println("---", r[page])
// 	}
// 	if value != nil {
// 		result = append(result, value)
// 	}

// 	return result
// }

func CreatePagenationV2(data []models.Pasien, lengthData int64) []ResultData {
	var i, page, appendIdx int64
	var value []models.Pasien
	var result []ResultData
	page = 0
	appendIdx = 0
	for i = 0; i < lengthData; i++ {
		if appendIdx == 10 {
			page++
			appendIdx = 0
			appendValue := ResultData{
				Pages: page,
				Data:  value,
			}
			// fmt.Println("--------------------------------", appendValue)
			result = append(result, appendValue)
			value = nil
			value = append(value, data[i])
			appendIdx++
		} else {
			value = append(value, data[i])
			appendIdx++
		}
	}
	if value != nil {
		appendValue := ResultData{
			Pages: page + 1,
			Data:  value,
		}
		// fmt.Println("--------------------------------", appendValue)
		result = append(result, appendValue)
	}

	return result
}

func CreatePagenationPemeriksaan(data []models.Pemeriksaan, lengthData int64) []ResultDataPemeriksaan {
	var i, page, appendIdx int64
	var value []models.Pemeriksaan
	var result []ResultDataPemeriksaan
	page = 0
	appendIdx = 0
	for i = 0; i < lengthData; i++ {
		if appendIdx == 10 {
			page++
			appendIdx = 0
			appendValue := ResultDataPemeriksaan{
				Pages: page,
				Data:  value,
			}
			// fmt.Println("--------------------------------", appendValue)
			result = append(result, appendValue)
			value = nil
			value = append(value, data[i])
			appendIdx++
		} else {
			value = append(value, data[i])
			appendIdx++
		}
	}
	if value != nil {
		appendValue := ResultDataPemeriksaan{
			Pages: page + 1,
			Data:  value,
		}
		// fmt.Println("--------------------------------", appendValue)
		result = append(result, appendValue)
	}

	return result
}

func CreatePagenationObat(data []repositories.ResultAllObats, lengthData int) []ResultDataObat {
	var page, appendIdx int64
	var i int
	var value []repositories.ResultAllObats
	var result []ResultDataObat
	page = 0
	appendIdx = 0
	for i = 0; i < lengthData; i++ {
		if appendIdx == 10 {
			page++
			appendIdx = 0
			appendValue := ResultDataObat{
				Pages: page,
				Data:  value,
			}
			// fmt.Println("--------------------------------", appendValue)
			result = append(result, appendValue)
			value = nil
			value = append(value, data[i])
			appendIdx++
		} else {
			value = append(value, data[i])
			appendIdx++
		}
	}
	if value != nil {
		appendValue := ResultDataObat{
			Pages: page + 1,
			Data:  value,
		}
		// fmt.Println("--------------------------------", appendValue)
		result = append(result, appendValue)
	}

	return result
}
