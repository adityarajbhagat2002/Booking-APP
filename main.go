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

	isvalidname:=len(firstfirstname) >=2 && len(lastname) >=2
	isvalidemail := string.Contains(email, "@")
	isvalidticket :=usertickets > 0 && usertickets <= remainingTickets
	

	if isvalidemail && isvalidname && isvalidticket{
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
		if !isvalidname{
			fmt.Println("first name or lastname you entered is too short")
		}
		fmt.Printf("your input data is invalid ")
	}
   

}
city : = "LONDON"
switch city {
case "New york ":

case "singapore" :

case "berlin" :

	defalut :
	fmt.Println("no valid  city selected")
	
}
}

func getusers (){
	fmt.Println("welcome to our confrence ")
}
