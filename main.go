package main

import (
	S "SearchTrees/workers"
	"fmt"
	"os"
)

func main(){
	JSONFileName := "Contacts.json"
	JSONInput, Jerr:= os.Open(JSONFileName)
	if Jerr != nil {
		fmt.Println("Error parsing json: ", Jerr)
	}
	defer JSONInput.Close()

	JSONData, _ := os.ReadFile(JSONFileName)
	JSONString := string(JSONData)

	fmt.Println("Index Binary Search")
	//Perform binary search using index
	TargetingIndex := 11
	Index, Person, Serr := S.IndexBinarySearch(JSONString, TargetingIndex)
	if Serr != nil {
		fmt.Println("Error searching from JSON")
	} else {
		if Index != -1 {
			S.PrintResult(Person)
		} else {
			fmt.Println("Contact not found")
		}
	}

	fmt.Println("\nDedicated Binary Search")
	//Perform binary search using dedicated field
	ParsedJSON, _ := S.ParseJson(JSONString)
	SortingField := "Email"
	S.SortByField(ParsedJSON, SortingField)
	TargetingField := SortingField
	TargetingValue := "NRasa@outlook.com"
	Res, Found := S.BinarySearch(ParsedJSON, TargetingField, TargetingValue)
	if Found {
		S.PrintResult(Res)
	} else {
		fmt.Println("Contact not found")
	}

	fmt.Println("\nNested Binary Search")
	//Perform binary search without defining field
	NParse, _ := S.ParseJson(JSONString)
	SearchVal := "Ali"
	NRes := S.NestedBinaryLinearSearch(NParse, SearchVal)
	if len(NRes) > 0 {
		for i := 0; i < len(NRes); i++{
			S.PrintResult(NRes[i])
		}
	} else {
		fmt.Println("Contact Not Found")
	}

	fmt.Println("\nAVL Search")
	//Perform AVL search using index
	AParse, _ := S.ParseJson(JSONString)
	var Root *S.AVLNode
	for _, Contact := range AParse {
		Root = S.AVLInsert(Root, Contact)
	}

	//S.AVLTraversal(Root)

	AVLSerachIndex := 34
	AVLRes := S.AVLSearch(Root, AVLSerachIndex)
	if AVLRes == nil {
		fmt.Println("Error searching from JSON")
	}else {
		S.PrintResult(&AVLRes.Det)
	}
}