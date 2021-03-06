package phones

import "strconv"

type PhoneNumber struct {
	number      string


}

func NewPhoneNumber(number string) *PhoneNumber {
	return &PhoneNumber{number: number}
}


//implements service

func (phone *PhoneNumber) PairsA() (pairs []string) {
	numberRune := []rune(phone.number)
	pairs = append(pairs, string(numberRune[0:2]))
	pairs = append(pairs, string(numberRune[2:4]))
	pairs = append(pairs, string(numberRune[4:6]))
	pairs = append(pairs, string(numberRune[6:8]))
	pairs = append(pairs, string(numberRune[8:10]))


	return pairs
}

func (phone *PhoneNumber) PairsB() (pairs []string) {
	numberRune := []rune(phone.number)
	pairs = append(pairs, string(numberRune[0]))
	pairs = append(pairs, string(numberRune[1:3]))
	pairs = append(pairs, string(numberRune[3:5]))
	pairs = append(pairs, string(numberRune[5:7]))
	pairs = append(pairs, string(numberRune[7:9]))
	pairs = append(pairs, string(numberRune[9]))


	return pairs
}

func (phone *PhoneNumber) PairSum() string {
	var sum int

	numberRune := []rune(phone.number)

	for _, c := range numberRune {
		num, _ := strconv.Atoi(string(c))
		sum = sum + num
	}

	return strconv.Itoa(sum)
}

func (phone *PhoneNumber) PairsUnique(pairSum string, pairsA []string, pairsB []string) (pairs []string) {
	var mixPairsABS []string

	mixPairsABS = append(mixPairsABS, pairSum)

	for _, pair := range pairsA {
		mixPairsABS = append(mixPairsABS, pair)
	}

	for _, pair := range pairsB {
		mixPairsABS = append(mixPairsABS, pair)
	}
	return unique(mixPairsABS)
}
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
