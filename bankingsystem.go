package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	// "github.com/gotilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID          int     `json:"UserID"`
	UserAccountNo   int     `json:"UserAccountNo"`
	UserName        string  `json:"UserName"`
	UserType        string  `json:"UserType"`
	UserBalance     float64 `json:"UserBalance"`
	UserOpeinigDate string  `json:"UserOpeinigDate"`
}
type MongoField struct {
	FieldStr  string `json: "Field Str"`
	FieldInt  int    `json: "Field Int"`
	FieldBool bool   `json: "Field Bool"`
}

var (
	BooksCollection *mongo.Collection
	// AuthorsCollection   *mongo.Collection
	Ctx = context.TODO()
)

// /*Setup opens a database connection to mongodb*/
// func Setup() {
// 	host := "localhost"
// 	port := "27017"
// 	connectionURI := "mongodb://" + host + ":" + port + "/"
// 	clientOptions := options.Client().ApplyURI(connectionURI)
// 	client, err := mongo.Connect(Ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(Ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db := client.Database("lenadena")
// 	BooksCollection = db.Collection("users")
// 	// AuthorsCollection = db.Collection("authors")
// }

// var users []User = []users{}

// func addaccount(q http.ResponseWriter, r *http.Request) {
// 	var newAccount User
// 	json.NewDecoder(r.Body).Decode(&newAccount)
// 	w.Header().set("Content-Type", application/json)

// 	users = append(users, newAccount)

// 	json.NewEncoder(q).Encode(users)
// }

func main() {

	router := mux.NewRouter()
	http.ListenAndServe(":5000", router)
	router.HandleFunc("/createAccount", addaccount).Methods("POST")

	var users User // users{1,"","","",""}
	var choice int
	arr := []User{} // [users{1,"","","",""},users{1,"","","",""},users{1,"","","",""},users{1,"","","",""}]

	arr = append(arr, User{101, 20220000001, "shivam", "saving", 100, "2023-06-08 12:16:32.2969433 +0530 IST"})
	arr = append(arr, User{101, 20220000001, "shivam", "current", 1000, "2023-06-08 12:16:32.2969433 +0530 IST"})
	arr = append(arr, User{102, 20220000002, "pathak", "current", 2000, "2023-06-08 12:16:32.2969433 +0530 IST"})
	arr = append(arr, User{103, 20220000003, "somu", "fixed", 50000, "2023-06-08 12:16:32.2969433 +0530 IST"})

	fmt.Printf("\t\t\t###############################################################################")
	fmt.Printf("\n\t\t\t############                                                   ################")
	fmt.Printf("\n\t\t\t############          Welcome LenaDena Bank using GoLang   ####################")
	fmt.Printf("\n\t\t\t############                                                   ################")
	fmt.Printf("\n\t\t\t###############################################################################")

	for {

		fmt.Printf("\n\n\n\t\t\t press 1 to  create account")
		fmt.Printf("\n\t\t\t press 2 to Deposit")
		fmt.Printf("\n\t\t\t press 3 to Withdraw")
		fmt.Printf("\n\t\t\t Press 4 to show All record")
		fmt.Printf("\n\t\t\t Press 5 to search a record")
		fmt.Printf("\n\t\t\t 0. Exit")

		fmt.Printf("\n\t\t\t Enter Your choice => ")
		fmt.Scan(&choice)

		fmt.Println("\n\t\t\t Your choice is => ", choice)

		switch choice {
		case 1:
			var input int
			fmt.Printf("\n\n\t\t\t Which type of account u wanna to create ")
			fmt.Printf("\n\t\t\t press 1 to Savings Account (100 rupes bank limit)")
			fmt.Printf("\n\t\t\t press 2 to Current Account")
			fmt.Printf("\n\t\t\t press 3 to Fixed Account")
			fmt.Printf("\n\t\t\t Enter Your choice => ")
			fmt.Scan(&input)
			userData := users.CreateAccount(input, arr)
			go waiting()
			time.Sleep(3 * time.Second)
			fmt.Printf("\t\t\t--------------------------------------------------------------------------------------------")
			fmt.Println("\n\t\t\t User Detail which save in DB: ", userData)
			fmt.Printf("\t\t\t--------------------------------------------------------------------------------------------")
			if userData.UserID != 0 {
				arr = append(arr, userData)
				saveToDb(arr)
			} else {
				fmt.Println("\n\t\t\t please Try Again ")
			}
		case 2:
			var input int
			var accNo int
			var status bool = false
			var atype string
			var depoAmount float64
			fmt.Printf("\n\t\t\t Deposit From ")
			fmt.Printf("\n\t\t\t press 1 to Savings Account")
			fmt.Printf("\n\t\t\t press 2 to Current Account")
			fmt.Printf("\n\t\t\t press 3 to Fixed Account")
			fmt.Printf("\n\t\t\t Enter Your choice => ")
			fmt.Scan(&input)

			fmt.Printf("\n\t\t\t Enter Your accNo => ")
			fmt.Scan(&accNo)

			if input == 1 {
				atype = "saving"
			} else if input == 2 {
				atype = "current"
			} else {
				atype = "fixed"
			}
			DepositAmmount(arr, atype, accNo, depoAmount, status)
		case 3:
			var input int
			var accNo int
			var status bool = false
			var atype string
			var withdrawAmount float64
			fmt.Printf("\n\t\t\t Withdraw From ")
			fmt.Printf("\n\t\t\t press 1 to Savings Account")
			fmt.Printf("\n\t\t\t press 2 to Current Account")
			fmt.Printf("\n\t\t\t press 3 to Fixed Account")
			fmt.Printf("\n\t\t\t Enter Your choice => ")
			fmt.Scan(&input)

			fmt.Printf("\n\t\t\t Enter Your accNo => ")
			fmt.Scan(&accNo)

			if input == 1 {
				atype = "saving"
			} else if input == 2 {
				atype = "current"
			} else {
				atype = "fixed"
			}
			WithdrawAmmount(arr, atype, accNo, withdrawAmount, status)
		case 4:
			fmt.Printf("\n\t\t\t View All Accounts")
			ShowAllAccount(arr)
		case 5:
			var searchAccount int
			var status bool = false
			fmt.Printf("\n\t\t\tEnter a Account no. to Search :=>")
			fmt.Scan(&searchAccount)
			SearchAccount(arr, searchAccount, status)
		default:
			fmt.Printf("\n\n\n\t\t\t INVALID INPUT!!! Try again...")
		}
		if choice == 0 {
			break
		}
	}

}

func (db *User) CreateAccount(acctype int, arr []User) User {
	var id int
	var name string
	var atype string
	var balance float64
	var accNo int

	if acctype == 1 {
		atype = "saving"
	} else if acctype == 2 {
		atype = "current"
	} else {
		atype = "fixed"
	}

	fmt.Print("\n\t\t\tEnter AdharNo. : ")
	fmt.Scan(&id)

	fmt.Print("\n\t\t\tEnter Accountholdername: ")
	fmt.Scan(&name)

	fmt.Print("\n\t\t\tEnter OpeningBalance : ")
	fmt.Scanf("\n%f", &balance)
	if atype == "saving" {
		for balance >= 250000 {
			fmt.Println("\n\t\t\t In saving Account Deposit not more than 2.5 lack:")
			fmt.Print("\n\t\t\tEnter OpeningBalance less than 2.5L: ")
			fmt.Scanf("\n%f", &balance)
		}
	} else if atype == "current" {
		for balance >= 5000000 {
			fmt.Println("\n\t\t\t In current Account Deposit not more than 50 lack:")
			fmt.Print("\n\t\t\tEnter OpeningBalance less than 50L: ")
			fmt.Scanf("\n%f", &balance)
		}
	}

	accNo = rand.Intn(100000000000) // eg. 2345678837383
	currentDate := time.Now().Local().String()

	save := User{
		UserID:          id,
		UserAccountNo:   accNo,
		UserName:        name,
		UserType:        atype,
		UserBalance:     balance,
		UserOpeinigDate: currentDate,
	}
	for i := 0; i < len(arr); i++ {
		if arr[i].UserID == id {
			if arr[i].UserType == atype {
				fmt.Print("\n\t\t\t You can't open same type account ")
				save = User{
					UserID:        0,
					UserAccountNo: 0,
					UserName:      "INVALID",
					UserType:      "INVALID",
					UserBalance:   0,
				}
				return save
			}

		}
	}

	// result, _ := BooksCollection.InsertOne(Ctx, save)
	// if err != nil {
	// 	return err
	// }
	// fmt.Sprintf("%v", result.InsertedID)

	// fmt.Println("Bank: ", save)
	return save

}

func ShowAllAccount(arr []User) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("lenadena").Collection("users")
	// result, insertErr := col.InsertOne(ctx, oneDoc2)
	// filter := bson.M{}
	result, findErr := col.Find(ctx, bson.D{})

	// fmt.Println("I am func", result)
	if findErr != nil {
		fmt.Println("findErr Error:", findErr)
		os.Exit(1)
	} else {
		// fmt.Println("\n\t\t\tInsertOne() api result type: ", result)
	}
	var userResult []User
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	if err = result.All(context.TODO(), &userResult); err != nil {
		panic(err)
	}
	// fmt.Println("employee Details are: >>>", userResult)
	for i := 0; i < len(userResult); i++ {
		// fmt.Printf("hello")
		if userResult[i].UserID != 0 {
			fmt.Println("\n\t\t\tresult Default:", userResult[i])
		}
		// fmt.Println("\n\t\t\tArr Default 2:", reflect.ValueOf(arr[i].UserID).Kind())
	}
	// for i := 0; i < len(arr); i++ {
	// 	// fmt.Printf("hello")
	// 	if arr[i].UserID != 0 {
	// 		fmt.Println("\n\t\t\tArr Default:", arr[i])
	// 	}
	// 	// fmt.Println("\n\t\t\tArr Default 2:", reflect.ValueOf(arr[i].UserID).Kind())
	// }

}

func SearchAccount(arr []User, searchAccount int, status bool) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("lenadena").Collection("users")
	// result, insertErr := col.InsertOne(ctx, oneDoc2)
	// filter := bson.M{}
	result, findErr := col.Find(ctx, bson.D{})

	// fmt.Println("I am func", result)
	if findErr != nil {
		fmt.Println("findErr Error:", findErr)
		os.Exit(1)
	} else {
		// fmt.Println("\n\t\t\tInsertOne() api result type: ", result)
	}
	var userResult []User
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	if err = result.All(context.TODO(), &userResult); err != nil {
		panic(err)
	}
	// fmt.Println("employee Details are: >>>", userResult)
	for i := 0; i < len(userResult); i++ {
		// fmt.Printf("hello")
		if userResult[i].UserAccountNo == searchAccount {
			fmt.Println("\n\t\t\tresult Default:", userResult[i])
			status = !status
		}
		// fmt.Println("\n\t\t\tArr Default 2:", reflect.ValueOf(arr[i].UserID).Kind())
	}
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i].UserAccountNo == searchAccount {
	// 		fmt.Println("\n\t\t\tArr Default:", arr[i])
	// 		status = !status
	// 	}

	// }
	if status == false {
		fmt.Println("\n\t\t\tNo Data Found")
	}

}

func DepositAmmount(arr []User, atype string, accNo int, depoAmount float64, status bool) {
	for i := 0; i < len(arr); i++ {
		// fmt.Printf("hello")
		if arr[i].UserID != 0 {
			if arr[i].UserType == atype && arr[i].UserAccountNo == accNo {
				fmt.Println("\n\t\t\tYour Infomation: \n\t\t\t Your Id:", arr[i].UserID, "\n\t\t\t Your Name:", arr[i].UserName)
				fmt.Println("\n\t\t\tYou Current Balance:", arr[i].UserBalance)
				fmt.Printf("\n\t\t\tEnter Deposit Amount:")
				fmt.Scan(&depoAmount)
				if arr[i].UserType == "saving" {
					for (arr[i].UserBalance + depoAmount) > 250000 {
						fmt.Println("\n\t\t\t In saving Account Deposit not more than 2.5 lack:")
						fmt.Print("\n\t\t\tEnter  Deposit Amount less than 2.5L: ")
						fmt.Scanf("\n%f", &depoAmount)
						if (arr[i].UserBalance + depoAmount) == 250000 {
							break
						}
					}
					if depoAmount > 50000 {
						var verifyId int
						fmt.Print("\n\t\t\tEnter ID ")
						fmt.Scanf("\n%d", &verifyId)
						if arr[i].UserID == verifyId {
							fmt.Print("\n\t\t\t Id Varified ")
							arr[i].UserBalance += depoAmount
						} else {
							fmt.Print("\n\t\t\t Id Doesn't match ")
						}

					} else {
						arr[i].UserBalance += depoAmount
					}
				} else if arr[i].UserType == "current" {
					for (arr[i].UserBalance + depoAmount) > 5000000 {
						fmt.Println("\n\t\t\t In Current Account Deposit not more than 50 lack:")
						fmt.Print("\n\t\t\tEnter  Deposit Amount less than 50L: ")
						fmt.Scanf("\n%f", &depoAmount)
						if (arr[i].UserBalance + depoAmount) == 5000000 {
							break
						}
					}
					if depoAmount > 250000 {
						var verifyId int
						fmt.Print("\n\t\t\tEnter ID ")
						fmt.Scanf("\n%d", &verifyId)
						if arr[i].UserID == verifyId {
							fmt.Print("\n\t\t\t Id Varified ")
							arr[i].UserBalance += depoAmount
						} else {
							fmt.Print("\n\t\t\t Id Doesn't match ")
						}
					} else {
						arr[i].UserBalance += depoAmount
					}
				} else {
					arr[i].UserBalance += depoAmount
				}

				fmt.Println("\n\t\t\t update Current Balance:", arr[i].UserBalance)
				status = true
			}
		}

	}
	if status == false {
		fmt.Println("\n\t\t\tNo Data Found")
	}
}

func WithdrawAmmount(arr []User, atype string, accNo int, withdrawAmount float64, status bool) {
	for i := 0; i < len(arr); i++ {
		var LIMIT float64 = 200000
		var SAVINGLIMIT float64 = 100
		// fmt.Printf("hello")
		if arr[i].UserID != 0 {
			if arr[i].UserType == atype && arr[i].UserAccountNo == accNo {
				fmt.Println("\n\t\t\tYour Infomation: \n\t\t\t Your Id:", arr[i].UserID, "\n\t\t\t Your Name:", arr[i].UserName)
				fmt.Println("\n\t\t\tYou Current Balance:", arr[i].UserBalance)
				fmt.Println("\n\t\t\tEnter Withdraw Amount:")
				fmt.Scan(&withdrawAmount)
				if arr[i].UserType == "saving" {
					for arr[i].UserBalance < withdrawAmount {
						fmt.Println("\n\t\t\t Insufficient Fund :")
						fmt.Println("\n\t\t\t Please ReEnter Withdraw Amount:")
						fmt.Scan(&withdrawAmount)
					}
				} else if arr[i].UserType == "fixed" {
					for arr[i].UserBalance < withdrawAmount {
						fmt.Println("\n\t\t\t Insufficient Fund :")
						fmt.Println("\n\t\t\t Please ReEnter Withdraw Amount:")
						fmt.Scan(&withdrawAmount)
					}
				} else {
					if arr[i].UserBalance > 0 {
						for arr[i].UserBalance < withdrawAmount {
							fmt.Println("\n\t\t\t Insufficient Fund :")
							fmt.Println("\n\t\t\t  Do You Wanna Use OverDraft Limit:")
							fmt.Println("\n\t\t\t Your Limit is 2L")
							fmt.Println("\n\t\t\t Please ReEnter Withdraw Amount:")
							fmt.Scan(&withdrawAmount)
							if (arr[i].UserBalance + LIMIT) > withdrawAmount {
								break
							}
						}
					} else {
						for (arr[i].UserBalance + LIMIT) < withdrawAmount {
							fmt.Println("\n\t\t\t Insufficient Fund :")
							fmt.Println("\n\t\t\t  Do You Wanna Use OverDraft Limit:")
							fmt.Println("\n\t\t\t Your Limit is 2L")
							fmt.Println("\n\t\t\t Please ReEnter Withdraw Amount:")
							fmt.Scan(&withdrawAmount)
							if (arr[i].UserBalance + LIMIT) > withdrawAmount {
								break
							}
						}
					}

				}

				if arr[i].UserType == "saving" {
					if arr[i].UserBalance-SAVINGLIMIT >= withdrawAmount {
						arr[i].UserBalance -= withdrawAmount
						fmt.Println("\n\t\t\t update Current Balance:", arr[i].UserBalance)
					} else {
						fmt.Println("\n\t\t\t Insufficient Fund Papa se lekar aao:")
					}
				} else if arr[i].UserType == "fixed" {
					// var UserOpeinigDateLength int = stringLength(arr[i].UserOpeinigDate)
					// intput := arr[i].UserOpeinigDate
					// UserOpeinigDateArr := []string{arr[i].UserOpeinigDate}
					fmt.Println(stringLength(arr[i].UserOpeinigDate))
					fmt.Println(strings.Split(arr[i].UserOpeinigDate, " "))
					// fmt.Printf("UserOpeinigDateLength = ", UserOpeinigDateArr)

					// for i := 0; i < stringLength(arr[i].UserOpeinigDate); i++ {

					// 	char := stringLength(arr[i].UserOpeinigDate)
					// 	println(char)
					// }
					// for i := 0; i <= 9; i++ {
					// 	 := UserOpeinigDateArr[i]
					// 	fmt.Println("helo")
					// }
					// if arr[i].UserOpeinigDate == "2023-06-08 12:16:32.2969433 +0530 IST" {
					// 	// when the time perood not over
					// 	fmt.Println("\n\t\t\t  in fixed plus 2 month")
					// }
					if arr[i].UserBalance >= withdrawAmount {
						arr[i].UserBalance -= withdrawAmount
						fmt.Println("\n\t\t\t update Current Balance:", arr[i].UserBalance)
					} else {
						fmt.Println("\n\t\t\t Insufficient Fund Papa se lekar aao:")
					}
				} else {
					if (arr[i].UserBalance + LIMIT) >= withdrawAmount {
						if arr[i].UserBalance >= withdrawAmount {
							arr[i].UserBalance -= withdrawAmount
						} else {
							if arr[i].UserBalance > 0 {
								fmt.Println("\n\t\t\t Upper Limit + your balance Is", (arr[i].UserBalance + LIMIT))
								REMLIMIT := (arr[i].UserBalance + LIMIT) - withdrawAmount
								fmt.Println("\n\t\t\t Remaining LIMIT + balance", REMLIMIT)
								arr[i].UserBalance = REMLIMIT - LIMIT
								// arr[i].UserBalance = REMLIMIT - (arr[i].UserBalance + LIMIT)
							} else {
								fmt.Println("\n\t\t\t Neg  your balance Is", arr[i].UserBalance)
								fmt.Println("\n\t\t\t Neg Total Upper Limit ", LIMIT)
								REMLIMIT := arr[i].UserBalance + LIMIT
								fmt.Println("\n\t\t\t Remaining", REMLIMIT)
								if (arr[i].UserBalance - withdrawAmount) >= -200000 {
									arr[i].UserBalance = arr[i].UserBalance - withdrawAmount
									fmt.Println("\n\t\t\t Remaining Updated Limit", REMLIMIT-withdrawAmount)
								} else {
									fmt.Println("\n\t\t\t u have reached ur limit ")
								}

							}
						}
						fmt.Println("\n\t\t\t update Current Balance:", arr[i].UserBalance)
					} else {
						fmt.Println("\n\t\t\t Insufficient Fund Papa se lekar aao:")
					}
				}

				status = !status
			}
		}

	}
	if status == false {
		fmt.Println("\n\t\t\tNo Data Found")
	}
}

func waiting() {
	fmt.Println("\n\t\t\tProcessing")
}
func stringLength(str string) int {
	var length int
	for range str {
		length++
	}
	return length
}

func saveToDb(data []User) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("lenadena").Collection("users")

	// fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	// oneDoc := MongoField{
	// 	FieldStr:  "This is our first data and its very important",
	// 	FieldInt:  826482746,
	// 	FieldBool: true,
	// }
	var oneDoc2 User
	for i := 0; i < len(data); i++ {
		if data[i].UserID != 0 {
			// fmt.Println("\n\t\t\tArr Default:", data[i])
			oneDoc2 = User{
				UserID:          data[i].UserID,
				UserAccountNo:   data[i].UserAccountNo,
				UserName:        data[i].UserName,
				UserType:        data[i].UserType,
				UserBalance:     data[i].UserBalance,
				UserOpeinigDate: data[i].UserOpeinigDate,
			}
		}
	}

	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")
	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc2), "\n")
	result, insertErr := col.InsertOne(ctx, oneDoc2)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr, result)
		os.Exit(1)
	} else {
		// fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		// fmt.Println("\n\t\t\tInsertOne() api result type: ", result)

		// newID := result.InsertedID
		// fmt.Println("InsertedOne(), newID", newID)
		// fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}
}

// func updateToDb(data []User) {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("Mongo.connect() ERROR: ", err)
// 		os.Exit(1)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	col := client.Database("lenadena").Collection("users")

// 	// fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

// 	// oneDoc := MongoField{
// 	// 	FieldStr:  "This is our first data and its very important",
// 	// 	FieldInt:  826482746,
// 	// 	FieldBool: true,
// 	// }
// 	var oneDoc2 User
// 	for i := 0; i < len(data); i++ {
// 		if data[i].UserID != 0 {
// 			// fmt.Println("\n\t\t\tArr Default:", data[i])
// 			oneDoc2 = User{
// 				UserBalance:     data[i].UserBalance,
// 			}
// 		}
// 	}

// 	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")
// 	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc2), "\n")
// 	result, insertErr := col.updateOne(ctx, oneDoc2)
// 	if insertErr != nil {
// 		fmt.Println("updateOne Error:", insertErr, result)
// 		os.Exit(1)
// 	} else {
// 		// fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
// 		// fmt.Println("\n\t\t\tInsertOne() api result type: ", result)

// 		// newID := result.InsertedID
// 		// fmt.Println("InsertedOne(), newID", newID)
// 		// fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

// 	}
// }
