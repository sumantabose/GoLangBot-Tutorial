package main

import "./employee"

func main() {  
    e := employee.Employee {
        FirstName: "Sam",
        LastName: "Tarly",
        TotalLeaves: 30,
        LeavesTaken: 20,
    }
    e.LeavesRemaining()

    var f employee.Employee
    f.LeavesRemaining()
}