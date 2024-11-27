package src

import (
	"strings"
	"unicode"
	"math"
)


func IsAnagram(s string, t string) bool {

	if len(t) != len(s) {
		return false
	}

	countS := make(map[rune]int)
	countT := make(map[rune]int)

	for _, char := range s {
		countS[char]++
	}
	for _, char := range t {
		countT[char]++
	}

	if len(countS) != len(countT) {
		return false
	}

	for key, count := range countS {
		if countT[key] != count {
			return false
		}
	}

	return true
}

func IsPalindrome1(s string) bool {
    sanitize := sanitizeString(s)
	for i, j := 0, len(sanitize)-1; i < j; i, j = i+1, j-1 {
		if sanitize[i] != sanitize[j] {
			return false
		}
	}
	return true
}

func IsPalindrome(s string) bool {
    sanitize := sanitizeString(s)
    i := 0
    size := len(sanitize)
    for {
		if i >= size/2 { 
			break
		}
		if sanitize[i] != sanitize[size-i-1] {
			return false
		}
		i++
	}
	return true
}

func IsPalindrome0(s string) bool {
    sanitize := sanitizeString(s)
	return sanitize == reverse(sanitize)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sanitizeString(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result.WriteRune(unicode.ToLower(char))
		}
	}
	return result.String()
}

func IsCountPrimes(arr []int) int {
	count := 0
	for _, n := range arr {
		if n > 1 && isPrime(n) {
			count++
		}
	}
	return count
}

func IsCountPrimes2(arr []int) []int {
    var primes []int
    for _, n := range arr {
        if n > 1 && isPrime(n) {
            primes = append(primes, n)
        }
    }
    return primes
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// for i := 2; i*i <= n; i++ {
	// 	if n%i == 0 {
	// 		return false
	// 	}
	// }
	// for i := 2; i < n; i++ {
    //     if n % i == 0 {
    //         return false
    //     }
    // }
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
        if n % i == 0 {
            return false
        }
    }
	return true
}

func IsTwoSum(arr []int, sum int) []int {
    seen := make(map[int]int)

    for i, v := range arr {
        search := sum - v
        if index, found := seen[search]; found {
            return []int{arr[index], arr[i]}
        }
        seen[v] = i
    }
    return nil
}

func IsLinkedList(head *ListNode) bool {
    seen := make(map[*ListNode]bool)
    current := head

    for current != nil {
        if seen[current] {
            return false
        }
        seen[current] = true
        current = current.Next
    }

    return true
}