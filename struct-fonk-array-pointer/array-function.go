package main

import "fmt"

func main (){
	diziIslemleri();
}

func diziIslemleri(){

	fmt.Println("---------Floats array list---------------")
	var floats=[]float32{1.2,5.6,-4.7}
	for i:=0;i<len(floats);i++{
		fmt.Printf("[%d]=%f\n",i,floats[i])
		//todo clear zero numbers last of floats
	}

	fmt.Println("---------String array list---------------")
	var strings=[2]string{"Ramazan","Barda"}
	for i:=0;i<len(strings);i++ {
		fmt.Printf("[%d]=%s\n",i,strings[i])
	}

	fmt.Println("---------Matrix array list---------------")
	var matrix=[][]int{{1,2,3},{10,11,12},{20,21,22},{30,31,32}}

	for i:=0;i<len(matrix);i++  {
		for j:=0;j<len(matrix[i]);j++{
			fmt.Printf("%d,",matrix[i][j])
		}
		var a,b int=calArray(matrix[i])
		fmt.Printf("=======>%d=>%d\n",a,b)
	}
	fmt.Println(combine("Hoşgeldin","Ramazan","barda"))

}

func calArray(arr []int) (int,int){
	var plus, times int =0,1
	for i:=0;i<len(arr);i++ {
		plus+=arr[i]
		times=times*arr[i]
	}
	return plus,times
}
//değişken sayısı istenildiği kadar gönderilebilir.
func combine(texts ...string)  string{
	var a string
	for _,n:=range texts  {
		a+=n +" \n"
	}
	return a
}



