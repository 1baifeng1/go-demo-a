package main
// https://www.nowcoder.com/practice/f2fbd8f61c564ca0b5feaa63ab42dae5?tpId=90&&tqId=30984&rp=1&ru=/activity/oj&qru=/ta/2018test/question-ranking
import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var hasLetter bool = false
		var hasCapitalLetter bool = false
		var hasDigital bool = false
		var rule1, rule2, rule3, rule4 bool
		password, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		for j := 0; j < len(password) - 1; j++ {
			//fmt.Println(password[j])
			if password[j] >= 97 && password[j] <= 122 {
				hasLetter = true
			} else if password[j] >= 65 && password[j] <= 90 {
				hasCapitalLetter = true
			} else if password[j] >= 48 && password[j] <= 57 {
				if 0 == j {
					rule2 = true
					break
				} else {
					hasDigital = true
				}
			} else {
				rule1 = true
			}
		}
		if len(password) < 8 {
			rule4 = true
		}
		if (!hasDigital && !hasCapitalLetter) || (!hasDigital && !hasLetter) || (!hasCapitalLetter && !hasLetter) {
			rule3 = true
		}

		if rule1 || rule2 || rule3 || rule4 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
