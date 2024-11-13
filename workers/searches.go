package workers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type Details struct {
	Index int `json:"Index"`
	FName string `json:"FirstName"`
	LName string `json:"LastName"`
	Number string `json:"Phone"`
	Email string `json:"Email"`
	Affi string `json:"Affiliation"`
}

type Node struct {
	Index int
    FirstName string
    LastName string
    Phone string
    Email string
    Affiliation string
	left *Node
    right *Node
}

type AVLNode struct {
	Det Details
	left *AVLNode
    right *AVLNode
	height int
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func NodeHeight(N *AVLNode) int {
	if N == nil {
		return 0	
	}

	return N.height
}

func GetBalanceFactor(N *AVLNode) int {
	if N == nil {
		return 0;
	}

	return NodeHeight(N.left) - NodeHeight(N.right)
}

func RightRotate (N *AVLNode) *AVLNode {
	Another := N.left
	Temp := Another.right

	Another.right = N
	N.left = Temp

	N.height = 1 + Max(NodeHeight(N.left), NodeHeight(N.right))
	Another.height = 1 + Max(NodeHeight(Another.left), NodeHeight(Another.right))

	return Another
}

func LeftRotate (N *AVLNode) *AVLNode {
	Another := N.right
	Temp := Another.left

	Another.left = N
	N.right = Temp

	N.height = 1 + Max(NodeHeight(N.left), NodeHeight(N.right))
	Another.height = 1 + Max(NodeHeight(Another.left), NodeHeight(Another.right))

	return Another
}

func RightLeftRotate(N *AVLNode) *AVLNode {
	N.right = RightRotate(N.right)
	return LeftRotate(N)
}

func LeftRightRotate(N *AVLNode) *AVLNode {
	N.left = LeftRotate(N.left)
	return RightRotate(N)
}

func AVLBalance(N *AVLNode, Data Details) *AVLNode{
	N.height = 1 + Max(NodeHeight(N.left), NodeHeight(N.right))
	BalanceFactor := GetBalanceFactor(N)

	if BalanceFactor > 1 && Data.Index < N.left.Det.Index {
		return RightRotate(N)
	} else if BalanceFactor < -1 && Data.Index > N.right.Det.Index {
		return LeftRotate(N)
	} else if BalanceFactor > 1 && Data.Index > N.left.Det.Index {
		return LeftRightRotate(N)
	} else if BalanceFactor < -1 && Data.Index < N.right.Det.Index {
		return RightLeftRotate(N)
	}

	return N
}

func NewNode(Data Details) *AVLNode {
	return &AVLNode{
		Det: Data,
		left: nil,
		right: nil,
		height: 1,
	}
}

func AVLInsert(N *AVLNode, Data Details) *AVLNode {
	if N == nil {
		return NewNode(Data)
	} 
	
	if Data.Index < N.Det.Index {
		N.left = AVLInsert(N.left, Data)
	} else if Data.Index > N.Det.Index {
		N.right = AVLInsert(N.right, Data)
	}

	return AVLBalance(N, Data)
}

func AVLSearch(N *AVLNode, Index int) *AVLNode {
	if N == nil {
		return nil
	}

	if Index == N.Det.Index {
		return N
	} else if Index < N.Det.Index {
		return AVLSearch(N.left, Index)
	} else {
		return AVLSearch(N.right, Index)
	}
}

func AVLTraversal(N *AVLNode) {
	if N != nil {
		AVLTraversal(N.left)
		fmt.Println(N.Det.Index, "")
		fmt.Println(N.Det.FName, "")
		fmt.Println(N.Det.LName, "")
		fmt.Println(N.Det.Number, "")
		fmt.Println(N.Det.Email, "")
		fmt.Println(N.Det.Affi, "")
		AVLTraversal(N.right)
	}
}

// Insert inserts a new node into the BST based on the Index.
func BSTInsert(root *Node, person *Node) *Node {
    if root == nil {
        return person
    }

    if person.Index < root.Index {
        root.left = BSTInsert(root.left, person)
    } else {
        root.right = BSTInsert(root.right, person)
    }

    return root
}

// Search searches for a node by Index in the BST.
func BSTSearch(root *Node, index int) *Node {
    if root == nil || root.Index == index {
        return root
    }

    if index < root.Index {
        return BSTSearch(root.left, index)
    }
    return BSTSearch(root.right, index)
}

// InOrderTraversal performs an in-order traversal of the BST.
func TraversalOrder(root *Node) {
    if root != nil {
        TraversalOrder(root.left)
        fmt.Printf("Index: %d, FirstName: %s, LastName: %s, Phone: %s, Email: %s, Affiliation: %s\n",
            root.Index, root.FirstName, root.LastName, root.Phone, root.Email, root.Affiliation)
			TraversalOrder(root.right)
    }
}

func SortByField(Ppl []Details, Field string) {
	sort.Slice(Ppl, func(i,j int) bool {
		IField := reflect.ValueOf(Ppl[i]).FieldByName(Field)
		JField := reflect.ValueOf(Ppl[j]).FieldByName(Field)

		if !IField.IsValid() || !JField.IsValid() {
			return false
		}

		switch IField.Kind() {
		case reflect.String:
			return IField.String() < JField.String()
		case reflect.Int:
			return IField.Int() < JField.Int()
		default:
			return false
		}
	})
}

func SortToStringArr(Input Details) []string{
	var SortedArr []string
	Val := reflect.ValueOf(Input)

	for i := 0; i < Val.NumField(); i++ {
		Field := Val.Field(i)
		// Convert the field value to string
		var FieldValue string
		switch Field.Kind() {
		case reflect.String:
			FieldValue = Field.String()
		case reflect.Int:
			FieldValue = strconv.Itoa(int(Field.Int()))
		case reflect.Float64:
			FieldValue = fmt.Sprintf("%f", Field.Float())
		case reflect.Bool:
			FieldValue = strconv.FormatBool(Field.Bool())
		default:
			FieldValue = fmt.Sprint(Field.Interface()) 
			// Convert other types to string using fmt
		}

		SortedArr = append(SortedArr, FieldValue)
	}

	sort.Strings(SortedArr)

	return SortedArr
}

func NestedBinaryLinearSearch(Conts []Details, Target interface{}) ([]*Details) {
	ResSet := []*Details{}
	SortByField(Conts, "Index")
	for i := 0; i < len(Conts); i++ {
		SortedArr := SortToStringArr(Conts[i])
		NLHS, NRHS := 0, len(SortedArr) - 1
		for NLHS <= NRHS {
			NMid := (NLHS + NRHS) / 2
			if Target == SortedArr[NMid] {
				ResSet = append(ResSet, &Conts[i])
				break;
			} else if Target.(string) < SortedArr[NMid] {
				NRHS = NMid - 1
			} else {
				NLHS = NMid + 1
			}
		}
	}

	return ResSet
}

// BinarySearch searches for a target value in a sorted slice of Person objects
func BinarySearch(Ppl []Details, Field string, Target interface{}) (*Details, bool) {
	EmptySet := &Details{}

	LHS, RHS := 0, len(Ppl) - 1
	for LHS <= RHS {
		Mid := (LHS + RHS) / 2
		// Use reflection to get the value of the field in the struct
		FieldValue := reflect.ValueOf(Ppl[Mid]).FieldByName(Field)

		// Ensure the field is valid
		if !FieldValue.IsValid() {
			return EmptySet, false
		}

		switch FieldValue.Kind() {
		case reflect.String:
			if Target.(string) == FieldValue.String() {
				return &Ppl[Mid], true
			} else if Target.(string) < FieldValue.String() {
				RHS = Mid - 1
			} else {
				LHS = Mid + 1
			}
		case reflect.Int:
			if Target.(int) == int(FieldValue.Int()) {
				return &Ppl[Mid], true
			} else if Target.(int) < int(FieldValue.Int()) {
				RHS = Mid - 1
			} else {
				LHS = Mid + 1
			}
		default:
			return EmptySet, false
		}
	}

	return EmptySet, false // Not found
}

func ParseJson(JsonInput string) ([]Details, error){
	var ResCont []Details
	Error := json.Unmarshal([]byte(JsonInput), &ResCont)
	if Error != nil{
		return nil, fmt.Errorf("error parsing JSON: %v", Error)
	}

	return ResCont, nil
}

func SortByIndex(ContactList []Details){
	sort.Slice(ContactList, func(i, j int) bool {
		return ContactList[i].Index < ContactList[j].Index
	})
}

func IndexBinarySearch(JSONInput string, TargetFind int) (int, *Details, error){
	Contacts, err := ParseJson(JSONInput)
	if err != nil{
		return -1, nil, fmt.Errorf("error parsing JSON %v", err)
	}

	SortByIndex(Contacts)
	
	LHS, RHS := 0, len(Contacts) - 1
	for LHS <= RHS {
		Mid := LHS + (RHS - LHS)/2

		if Contacts[Mid].Index == TargetFind {
			return Mid, &Contacts[Mid], nil
		} else if Contacts[Mid].Index < TargetFind {
			LHS = Mid + 1
		} else
		{
			RHS = Mid - 1
		}
	}
	
	return -1, nil, fmt.Errorf("target not found")
}

func PrintResult(detail *Details) {
	if detail != nil {
		fmt.Println("First Name: ", detail.FName)
		fmt.Println("Last Name: ", detail.LName)
		fmt.Println("Phone Number: ", detail.Number)
		fmt.Println("Email: ", detail.Email)
		fmt.Println("Affilation: ", detail.Affi)
	}
}