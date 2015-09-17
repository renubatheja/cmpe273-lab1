package main 

import ("fmt";"time")

//main Go file for solving CMPE273 - Lab 1 Problems
func main() {
	var choice int
	fmt.Println("CMPE273 - Lab1")
	fmt.Println("--------------")
	fmt.Println("Please choose among the below options")
	fmt.Println("1. Generation of Fibonacci Series")
	fmt.Println("2. Calculation of perimeter/area of a Shape")
	fmt.Println("3. Demonstration of Sleep function (using time.After)")
	fmt.Scanf("%d\n", &choice)
	switch choice {
    case 1:
    	call_fibonacci()
        
    case 2:
		call_shapeCalc()		
        
    case 3:
		call_sleep()
    }
}

//Function to give a call to a function in Fibonacci.go file. 
//This computes the Fibonacci series for a given range
func call_fibonacci() {
		var n int64
		var i int64

	    fmt.Println("Generation of Fibonacci Series for N numbers!")
	    fmt.Println("---------------------------------------------")
	    fmt.Println("Please enter a number to compute the Fibonacci Series till that range:")
	    fmt.Scanf("%d\n", &n)
	    
		if(n < 0) {
			fmt.Println("Error! Fibonacci for negative numbers can't be computed!")
		} else {
			fmt.Println("Fibonacci Series is as below:")
			for i = 0; i <= n; i++ {
				fmt.Printf("Fibonacci(%d) = %d",i,Fibonacci(i))		
				fmt.Println()
			}
		}
}

//Function to give a call to a function in Shape.go file. 
//This computes the perimeter and area of a given Shape (Circle/Rectangle)
func call_shapeCalc() {
		var num int
		fmt.Println("Calculation of perimeter/area of a Shape")
		fmt.Println("----------------------------------------")
		fmt.Println("Which would you like to calculate?")
		fmt.Println("1. Circle's perimeter")
		fmt.Println("2. Rectangle's perimeter")
		fmt.Println("3. Circle's area")
		fmt.Println("4. Rectangle's area")
		fmt.Scanf("%d\n", &num)
		
		switch num {
	    case 1:
	    	var radius float64
	    	fmt.Println("Enter the circle's radius")
	    	fmt.Scanf("%f\n", &radius)
	    	c := Circle{0, 0, radius}
	    	fmt.Println("This Circle's perimeter = ", computePerimeter(c))
	        
	    case 2:
	    	var x1,y1,x2,y2 float64
	    	fmt.Println("Enter the rectangle's co-ordinates (x1, y1, x2, y2)")
	    	fmt.Scanf("%f%f%f%f\n", &x1,&y1,&x2,&y2)
	    	r := Rectangle{x1, y1, x2, y2}	    	
			fmt.Println("This Rectangle's perimeter = " ,computePerimeter(r))		

	    case 3:
	    	var radius float64
	    	fmt.Println("Enter the circle's radius")
	    	fmt.Scanf("%f\n", &radius)
	    	c := Circle{0, 0, radius}
	    	fmt.Println("This Circle's area = ", computeArea(c))
	        
	    case 4:
	    	var x1,y1,x2,y2 float64
	    	fmt.Println("Enter the rectangle's co-ordinates (x1, y1, x2, y2)")
	    	fmt.Scanf("%f%f%f%f\n", &x1,&y1,&x2,&y2)
	    	r := Rectangle{x1, y1, x2, y2}	    	
			fmt.Println("This Rectangle's area = " ,computeArea(r))		

		}
		//fmt.Println("Circle's area using 'method' on type Circle            :" ,c.area())
		//fmt.Println("Circle's area using 'method' on type Shape             :", computeArea(c))
		//fmt.Println("Circle's perimeter using 'method' on type Circle       :" ,c.perimeter())
		//fmt.Println("Circle's perimeter using 'method' on type Shape        :", computePerimeter(c))
		//fmt.Println()	
		
		//fmt.Println("Rectangle's area using 'method' on type Rectangle      :" ,r.area())
		//fmt.Println("Rectangle's area using 'method' on type Shape          :" ,computeArea(r))
		//fmt.Println("Rectangle's perimeter using 'method' on type Rectangle :" ,r.perimeter())
		//fmt.Println("Rectangle's perimeter using 'method' on type Shape     :" ,computePerimeter(r))
		//fmt.Println()
		//fmt.Println("Total Area of Circle and Rectangle                     :",totalArea(&c, &r))
		//fmt.Println("Total Perimeter of Circle and Rectangle                :",totalPerimeter(&c, &r))
}

//Function to give a call to a function in MySleep.go file. 
//Sleep is implemented using time.After
func call_sleep() {
		var seconds int
		fmt.Println("Demonstration of Sleep function (using time.After)")
		fmt.Println("--------------------------------------------------")
		fmt.Println("For how many seconds you would like your current thread to sleep?")
		fmt.Scanf("%d\n", &seconds)

		fmt.Println("Current timestamp (only in Seconds) : ", time.Now().Second() )
		fmt.Println("This thread would sleep now for ",seconds," seconds")
		Sleep(seconds)
		fmt.Println("Thread awake now!")
		fmt.Println("Current timestamp (only in Seconds) : ", time.Now().Second() )
}