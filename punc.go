package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func Check(e error) {
	if e!= nil {
		panic(e)
	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...) // slicein o indexteki elemanini siler
}

func HexReplacer(h string) string {
	ConvertedWord, err := strconv.ParseInt(h, 16, 64) //hexadecimal to decimal
	Check(err)	
	return fmt.Sprint(ConvertedWord) //sonucu string olarak donebilmek icin
}

func BinReplacer(b string) string {
	ConvertedWord, err := strconv.ParseInt(b, 2, 64) //binary to decimal
	Check(err)
	return fmt.Sprint(ConvertedWord) //sonucu string olarak donebilmek icin
}

func ToUpper(s string) string {
	torune := []rune(s)
	for i := range torune {
		if torune[i] >= 'a' && torune[i] <= 'z' {
			torune[i] = torune[i] - 32
		}
	}
	return string(torune)
}

func ToLower(s string) string {
	torune := []rune(s)
	for i := range torune {
		if torune[i] >= 'A' && torune[i] <= 'Z' {
			torune[i] = torune[i] + 32
		}
	}
	return string(torune)
}

func IsAlpha(s string) bool {
	gruneda := []rune(s)
	counterAllCharsof := 0
	counterAllUpLettersof := 0
	counterAllLoLettersof := 0
	counterAllNumberssof := 0
	for _, poti := range gruneda {
		counterAllCharsof++
		if poti >= 'a' && poti <= 'z' {
			counterAllLoLettersof++
		}
		if poti >= 'A' && poti <= 'Z' {
			counterAllUpLettersof++
		}
		if poti >= '0' && poti <= '9' {
			counterAllNumberssof++
		}
	}
	countalphanum := counterAllLoLettersof + counterAllUpLettersof + counterAllNumberssof
	if s == "" || (counterAllCharsof == countalphanum) {
		return true
	} else {
		return false
	}
}

func IsNumeric(s string) bool {
	toRune := []rune(s)
	countNum := 0
	countAll := 0
	for _, ind := range toRune {
		countAll++
		if (ind >= '0') && (ind <= '9') {
			countNum++
		}
	}
	if countAll == countNum {
		return true
	} else {
		return false
	}
}

func Capitalize(s string) string {
	runa := []rune(s)
	harf := ""
	for i := 0; i < len(runa); i++ {
		a := string(runa[i])
		// If is not at index 0 and the letter before is alphanumerical, change it to lowercase
		if i != 0 && IsAlpha(string(runa[i-1])) {
			harf = harf + ToLower(a)
		} else {
			harf = harf + ToUpper(a)
		}
	}
	return harf
}

func StartswithPunc(s string) bool {
	//punctList := []string{",", ".", ":", ";", "?", "!"} 
	if (s[0] == 33) || (s[0] == 44) || (s[0] == 46) || (s[0] == 58) ||(s[0] == 59) || (s[0] == 63) {
		return true
	} else {
		return false
	}
}

func IsPunc(c rune) bool {
	if c >= 33 && c <= 47 || c == 58 || c == 59 || c == 63 {
		return true
	} else {
		return false
	}
}

func controller(n []string) []string { //string slice i alip sonuc olarak gene slice donuyor
	var result []string // forun icinde donunce icinde donulen bir slicei elemanin silinemez. result ve n asagidaki removeindex. cozumu olarak bos bir slice (result)
	result = append(n) // nin hepsini resulta ekledik
	sayac := 0
	for i := 0; sayac < len(n); i++{
		if strings.Contains(n[i], "(hex)") {
			n[i-1] = HexReplacer(n[i-1])
			result = RemoveIndex(result, i)
			sayac++
		} else if strings.Contains(n[i], "(bin)") {
			n[i-1] = BinReplacer(n[i-1])
			result = RemoveIndex(result, i)
			sayac++
		} else if strings.Contains(n[i], "(up)") {
			n[i-1] = ToUpper(n[i-1])
			result = RemoveIndex(result, i)
			sayac++
		} else if strings.Contains(n[i], "(low)") {
			n[i-1] = ToLower(n[i-1])
			result = RemoveIndex(result, i)
			sayac++
		} else if strings.Contains(n[i], "(cap)") {
			n[i-1] = Capitalize(n[i-1])
			result = RemoveIndex(result, i)
			sayac++
		} else if (strings.HasPrefix(n[i], "(up,")) {			
			sayim, err := strconv.Atoi(n[i+1][:len(n[i+1])-1])
			Check(err)
			for j:=1; j<=sayim ;j++ {
					n[i-j] = ToUpper(n[i-j])
			}
			result = RemoveIndex(result, i+1)
			result = RemoveIndex(result, i)
			sayac++
		} else if (strings.HasPrefix(n[i], "(low,")) {			
			sayim, err := strconv.Atoi(n[i+1][:len(n[i+1])-1])
			Check(err)
			for j:=1; j<=sayim ;j++ {
					n[i-j] = ToLower(n[i-j])
			}
			result = RemoveIndex(result, i+1)
			result = RemoveIndex(result, i)
			sayac++
		} else if (strings.HasPrefix(n[i], "(cap,")) {			
			sayim, err := strconv.Atoi(n[i+1][:len(n[i+1])-1])
			Check(err)
			for j:=1; j<=sayim ;j++ {
					n[i-j] = Capitalize(n[i-j])
			}
			result = RemoveIndex(result, i+1)
			result = RemoveIndex(result, i)
			sayac++
		} else if StartswithPunc(n[i]) == true {
			fmt.Println(n)
			runedWord := []rune(n[i])
			//fmt.Println(runedWord)
			runedPrecedingWord := []rune(n[i-1])
			//fmt.Println(runedPrecedingWord)
			
			for j :=0; j<len(runedWord) ; j++ {
				if IsPunc(runedWord[j]) {
					runedPrecedingWord = append(runedPrecedingWord, runedWord[j])
				}
			}
			
			//fmt.Println(runedPrecedingWord)
			//fmt.Println(runedWord)
			fmt.Println(n)
			n[i-1] = string(runedPrecedingWord)
			//fmt.Println(n[i-1])
			//fmt.Println(n[i])
			n[i] = string(runedWord)
			//fmt.Println(n[i])
			wordToClean := strings.Split(n[i],"")
			//fmt.Println(wordToClean)
			for i := len(wordToClean)-1; i >= 0; i-- {
				if StartswithPunc(wordToClean[i]) {
					wordToClean = append(wordToClean[:i], wordToClean[i+1:]...)
				}
			}
			n[i] = strings.Join(wordToClean, "")
			//fmt.Println(n[i])
			for t := 0; t<len(n); t++ {
				if n[t] == "" {
					n = append(n[:t], n[t+1:]...)
				}
			}
			sayac++
			fmt.Println(n)
		
			n := strings.Join(n, " ")
			fmt.Println(n)
			sayac++
		//	fmt.Println(toString)
		//	n := strings.ReplaceAll(n, "  ", " ")
		//	fmt.Println(toString)
 		}
		sayac++
	}
	return result
}
func main() {
	fileName := os.Args[1] // dosya adini aldik
	readedText, err := os.ReadFile(fileName) // ilgili dosyanin içeriğini okuduk (variable of type []byte) dosya okumak için kullanilir
	Check(err) // ilk 3 satir dosya okuma islemi 
	kelimeler := strings.Fields(string(readedText))// (variable of type []byte) ' ı stringe dönüsturup bosluklarina gore ayirip slicea atadik.  
	sonuc := controller(kelimeler) //controller fonksiyonu calisir.kelimelere tek tek bakip duzeltmeleri yapip sonuc veriyor
	str2 := strings.Join(sonuc, " ") //sonuc string slicei idi. Join ile stringe   dönustu
	//fmt.Println(str2)
	aaa := []byte(str2) // string converted to byte slice. Write file yazabilmek icin byte slice istiyor.
	err = os.WriteFile(os.Args[2], aaa, 0644) //result,txt dosyaysini olusturdu, icine aaa yi yazdi
    Check(err)
}