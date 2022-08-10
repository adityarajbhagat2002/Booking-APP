package main

import (
	"fmt"
	"strings"
)



func main(){
	confrenceName :="GO confrence"
	const  confrenceTickets=50
	var remainingTickets=50
	bookings :=[]string{}

	
	fmt.Printf("Welcome to %v booking application\n ",confrenceName)
	fmt.Printf("we have total of  %v tickets and  %v are still remaining\n",confrenceTickets,remainingTickets)
	fmt.Println("Get your tickets now!!!")

	
for remainingTickets > 0 && len(bookings) < 50{

	var firstname string
	var lastname string
	var email string	
	var usertickets int tyty
	

	fmt.Println("enter tour first name  : " ) 
	fmt.Scan(&firstname)

	fmt.Printf("enter your last name :\n")
	fmt.Scan(&lastname)

	fmt.Printf("enter the email id :\n")
	fmt.Scan(&email)

	fmt.Printf("enter no of tickets you want \n")
	fmt.Scan(&usertickets)

	if usertickets < remainingTickets{
		remainingTickets= remainingTickets-usertickets
	bookings = append(bookings, firstname + " " + lastname)
	firstnames :=[]string{}
	for  _,booking := range  bookings{
		var names=strings.Fields(booking)		
		firstnames=append(firstnames,names[0])


	}
	fmt.Printf("the first names of booking are %v\n" ,firstnames)

   
	if remainingTickets == 0{
		fmt.Println("our Conference is full .come back nexy year ")
		break
		
	}

	

	}else{
		fmt.Printf("we only have %v tickets  ",usertickets)
	}
	

	

	fmt.Printf("the whole array%v\n ",bookings)
	fmt.Printf("the FIRST array%v\n ",bookings[0])
	fmt.Printf("the whole array TYPE %T\n ",bookings)
	fmt.Printf("the whole array Length%T\n ",len(bookings))

   

}
}