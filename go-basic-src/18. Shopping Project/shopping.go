// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

type item struct {
	name   string
	price  int
	amount int
}

type buyer struct {
	point          int
	shoppingBucket map[string]int
}

type delivery struct {
	status      string
	onedelivery map[string]int
}

func newBuyer() *buyer {
	person := buyer{}
	person.point = 1000000
	person.shoppingBucket = map[string]int{}
	return &person
}

func newDelivery() delivery {
	truck := delivery{}
	truck.onedelivery = map[string]int{}
	return truck
}

func buying(items []item, man *buyer, selectNum int, num *int, deliveryStart chan bool, temp map[string]int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	buyingAmount := 0

	fmt.Print("구매할 수량을 입력하시오: ")
	fmt.Scanln(&buyingAmount)
	fmt.Println()

	if buyingAmount <= 0 {
		panic("올바른 수량을 입력하세요.")
	}

	if man.point < items[selectNum-1].price*buyingAmount || items[selectNum-1].amount < buyingAmount {
		panic("주문이 불가능합니다.")
	} else {
		for {
			buyingWay := 0
			fmt.Println("1. 바로 구매")
			fmt.Println("2. 장바구니에 담기")
			fmt.Print("구매 방법을 선택하시오: ")
			fmt.Scanln(&buyingWay)
			fmt.Println()

			switch buyingWay {
			case 1: // 바로 구매
				if *num < 5 {
					items[selectNum-1].amount -= buyingAmount
					man.point -= items[selectNum-1].price * buyingAmount
					temp[items[selectNum-1].name] = buyingAmount
					fmt.Println("상품 주문이 접수되었습니다.")
					deliveryStart <- true

					*num++
				} else {
					fmt.Println("배송 한도를 초과하였습니다. 배송이 완료되면 주문해주세요.")
				}
			case 2: // 장바구니 담기
				_, checkBucket := man.shoppingBucket[items[selectNum-1].name]
				if checkBucket == false {
					man.shoppingBucket[items[selectNum-1].name] = buyingAmount
				} else {
					if man.shoppingBucket[items[selectNum-1].name]+buyingAmount > items[selectNum-1].amount {
						panic("물품의 잔여 수량을 초과하였습니다.")
						break
					} else {
						man.shoppingBucket[items[selectNum-1].name] += buyingAmount
					}
				}
				fmt.Println("상품이 장바구니에 담겼습니다.")
			default:
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			}
			if buyingWay == 1 || buyingWay == 2 {
				break
			}
		}
	}
}

func showBucket(man *buyer) {
	if len(man.shoppingBucket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for index, val := range man.shoppingBucket {
			fmt.Printf("%s, 수량: %d\n", index, val)
		}
	}
}

func requirePoint(items []item, man *buyer) (canBuy bool) {
	bucketPoint := 0
	for index, val := range man.shoppingBucket {
		for i := 0; i < len(items); i++ {
			if items[i].name == index {
				bucketPoint += items[i].price * val
			}
		}
	}
	fmt.Println("필요 마일리지: ", bucketPoint)
	fmt.Println("보유 마일리지: ", man.point)
	fmt.Println()

	if man.point >= bucketPoint {
		return true
	} else {
		fmt.Printf("마일리지가 %d점 부족합니다.\n", bucketPoint-man.point)
		return false
	}
}

func excessAmount(items []item, man *buyer) (canBuy bool) {
	for index, val := range man.shoppingBucket {
		for i := 0; i < len(items); i++ {
			if items[i].name == index {
				if items[i].amount < val {
					fmt.Printf("%s %d개 초과", index, val-items[i].amount)
					return false
				}
			}
		}
	}
	return true
}

func bucketBuy(items []item, man *buyer, num *int, deliveryStart chan bool, temp map[string]int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if len(man.shoppingBucket) == 0 {
		panic("주문 가능한 목록이 없습니다.")
	} else {
		if *num < 5 {
			for index, val := range man.shoppingBucket {
				temp[index] = val
				for i := 0; i < len(items); i++ {
					if items[i].name == index {
						items[i].amount -= val
						man.point -= items[i].price * val
					}
				}
			}
			man.shoppingBucket = map[string]int{}
			deliveryStart <- true
			*num++
			fmt.Println("장바구니 주문이 접수되었습니다.")
		} else {
			fmt.Println("배송 한도를 초과하였습니다. 배송이 완료되면 주문해주세요.")
		}
	}
}

func deliveryStatus(deliveryStart chan bool, i int, deliveries []delivery, numBuy *int, temp *map[string]int) {
	for {
		if <-deliveryStart {
			for index, val := range *temp {
				deliveries[i].onedelivery[index] = val
			}
			deliveries[i].status = "주문접수"
			time.Sleep(time.Second * 10)

			deliveries[i].status = "배송중"
			time.Sleep(time.Second * 30)

			deliveries[i].status = "배송완료"
			time.Sleep(time.Second * 10)

			deliveries[i].status = ""
			*numBuy--

			deliveries[i].onedelivery = map[string]int{}
		}
	}
}

func main() {
	numBuy := 0
	items := make([]item, 5)
	buyer := newBuyer()
	deliveries := make([]delivery, 5)
	deliveryStart := make(chan bool)
	tempDelivery := make(map[string]int)

	for i := 0; i < 5; i++ {
		deliveries[i] = newDelivery()
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		go deliveryStatus(deliveryStart, i, deliveries, &numBuy, &tempDelivery)
	}

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"롱패딩", 500000, 20}
	items[2] = item{"투미 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menu := 0

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오: ")
		fmt.Scanln(&menu)
		fmt.Println()

		switch menu {
		case 1: // 구매
			for {
				itemNum := 0

				fmt.Println("구매 가능 제품 목록")
				for i := 0; i < 5; i++ {
					fmt.Printf("제품%d. %s(%d원) - 잔여수량 %d개\n", i+1, items[i].name, items[i].price, items[i].amount)
				}

				fmt.Print("구매하려는 제품 번호를 입력하시오(취소는 0번): ")
				fmt.Scanln(&itemNum)
				fmt.Println()

				switch itemNum {
				case 0:
					fmt.Println("제품 구매를 취소합니다.")
				case 1, 2, 3, 4, 5:
					buying(items, buyer, itemNum, &numBuy, deliveryStart, tempDelivery)
				default:
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
				}
				if itemNum == 0 {
					break
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
			fmt.Println("=======================================")
		case 2: // 잔여 수량 확인
			for i := 0; i < 5; i++ {
				fmt.Printf("%s - 잔여수량 %d개\n", items[i].name, items[i].amount)
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
			fmt.Println("=======================================")
		case 3: // 잔여 마일리지 확인
			fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", buyer.point)
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
			fmt.Println("=======================================")
		case 4: // 배송 상태 확인
			total := 0
			for i := 0; i < 5; i++ {
				total += len(deliveries[i].onedelivery)
			}
			if total == 0 {
				fmt.Println("배송중인 상품이 없습니다.")
			} else {
				for i := 0; i < len(deliveries); i++ {
					if len(deliveries[i].onedelivery) != 0 {
						for index, val := range deliveries[i].onedelivery {
							fmt.Printf("%s %d개 / ", index, val)
						}
						fmt.Printf("%s\n", deliveries[i].status)
					}
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
			fmt.Println("=======================================")
		case 5: // 장바구니 확인
			bucketMenu := 0
			for {
				showBucket(buyer)
				canBuy := requirePoint(items, buyer)
				canBuy = excessAmount(items, buyer)

				fmt.Println("1. 장바구니 물품 주문")
				fmt.Println("2. 장바구니 초기화")
				fmt.Println("3. 메뉴로 돌아가기")
				fmt.Print("실행할 메뉴를 입력하시오: ")
				fmt.Scanln(&bucketMenu)

				if bucketMenu == 1 {
					if canBuy {
						bucketBuy(items, buyer, &numBuy, deliveryStart, tempDelivery)
						break
					} else {
						fmt.Println("주문할 수 없습니다.")
						break
					}
				} else if bucketMenu == 2 {
					buyer.shoppingBucket = map[string]int{}
					fmt.Println("장바구니를 초기화했습니다.")
					break
				} else if bucketMenu == 3 {
					fmt.Println()
					break
				} else {
					fmt.Println("잘못된 입력입니다.")
				}
				fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
				fmt.Scanln()
				fmt.Println("=======================================")
			}
		case 6: // 프로그램 종료
			fmt.Println("프로그램을 종료합니다.")
			return
		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
		}
	}
}
