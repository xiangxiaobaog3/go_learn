package main

import "fmt"

func main() {
	//统计3个班级成绩情况，每个班有5名同学，
	//求出各个班的平均分和所有班级的平均分【学生的成绩从键盘输入】
	var totalSum float64 = 0.0
	for j := 1; j <= 3; j++ {
		sum := 0.0
		for i := 1; i <= 5; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩 \n", j, i)
			fmt.Scanln(&score)
			//累计总分xx`
			sum += score
		}

		fmt.Printf("第%d个班级的平均分是%v \n", j, sum/5)
		totalSum += sum
	}

	fmt.Printf("各个班的总成绩%v 所有班级平均分是 \n", totalSum, totalSum/(3*5))
}
