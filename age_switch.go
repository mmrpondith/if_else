package main
import(
	"fmt"
)
func main(){
	fmt.Print("enter your age: ")
	var age int 
	fmt.Scanf("%d",&age)

	/*if age <3{
		fmt.Println("infant")
	}else if age >2 && age<13 {
		fmt.Println("children")
	}else if age >12 && age <= 19{
		fmt.Println("teen age")
	}else{
		fmt.Println("adult")
	}

	fmt.Println(age)*/
	switch age {
	case 2:
		fmt.Println("infant")
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")
	case 13,14,15,16,17,18,19:
		fmt.Println("Teen age")
	default:
		fmt.Println("adult")
	}
}