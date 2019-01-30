package algo

import "fmt"

//func main()  {
//	points := [][]int{{6,10},{-3,3},{-2,5},{0,2}}
////[[6,10],[-3,3],[-2,5],[0,2]]
//
//
//	//fmt.Println(points)
//	result := kClosest(points,2)
//
//	fmt.Println(result)
//
//
//}
func kClosest(points [][]int, K int) [][]int {

	length := make(map[int]int)
	result := make([][]int, len(points))
	result[0]=points[0]


	for i:=1;i<len(points);i++{
		length[i]=points[i][0]*points[i][0]+points[i][1]*points[i][1]
		for j:=0;j<i;j++{
				tmp := points[j][0]* points[j][0]+points[j][1]*points[j][1]
				fmt.Println(tmp)
				if length[i] < tmp  {
					t := result[j]
					result[j] = points[i]
					result[i] = t
					for k:=i+1;k<len(result);k++{
						result[k]=result[k-1]
					}
				}else{
					//t := result[j]
					result[i] = points[i]
					//result[i] = t
				}
		}
		fmt.Println(result)

	}
	//fmt.Println(result)
	return result[:K]
}

