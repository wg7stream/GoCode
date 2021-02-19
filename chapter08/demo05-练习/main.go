package main
import (
	"fmt"
)

func main() {
	var score [3][5]float64

	for i := 0; i < len(scores); i ++ {
		for j := 0; j < len(scores[i]); j ++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i + 1, j + 1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalSum := 0.0
	for i := 0; i < len(scores); i ++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j ++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班的总分为%v , 平均分为%v\n", i + 1, sum,
			sum / float64(len(scores[i])))
	}

	fmt.Printf("所有班级的总分为%v , 所有班级平均分为%v\n",
		totalSum, totalSum / 1)
}