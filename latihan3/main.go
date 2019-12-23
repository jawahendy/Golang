package main

import (
	"fmt"
)

type DoorContract interface {
	open() string
	close() string
}

type TrDoor struct {
	DoorHandler
	DoorLock
	WoodenDoor
}

type DoorHandler struct {
	material string
}

type DoorLock struct {
	password string
}

type WoodenDoor struct {
	material string
	size     string
}

func (dh DoorHandler) openhan() string {
	return fmt.Sprintf(" %s", dh.material)
}

func (w WoodenDoor) openwood() string {
	return fmt.Sprintf("%s", w.material)
}

func (d DoorLock) openlock() string {
	return fmt.Sprintf(" %s", d.password)
}

func (t TrDoor) TypeDoor() string {
	return fmt.Sprintf("%s", t.WoodenDoor)
}

type Door struct {
	color    string
	Material string
	TrDoor
}

func (d Door) open() {
	// fmt.Println("pintu sedang dibuka : pintu dengan tipe", d.TypeDoor(),"dengan keamanan dengan ",d.DoorLock," warnanya ", d.color, "materialnya ", d.Material)

	fmt.Println("selamat datang di Training room : pintu kebuka dengan", d.openlock(), "dan materialnya dari ", d.openwood(), "dengan size ", d.WoodenDoor.size, ".")
}

func (d Door) close() {
	fmt.Println("pintu sedang ditutup : pintu dengan tipe", d.TypeDoor(), "dengan keamanannya ", d.DoorLock, " warnanya ", d.color, "materialnya ", d.Material)
}

func main() {
	hand := DoorHandler{
		"Iron",
		// "Password",
		// "Jati ",
	}
	lock := DoorLock{
		"hash password",
	}
	wood := WoodenDoor{
		"jati",
		"1,5m",
	}
	open := TrDoor{
		hand,
		lock,
		wood,
	}
	Material := Door{
		"brown",
		"full furnish",
		open,
	}
	// close := TrDoor{
	// 	"Iron",
	// 	"kunci",
	// 	"Jati ",
	// }
	// Material2 := Door{
	// 	"Black",
	// 	"full furnish",
	// 	close,
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	// if text == "open"{
	//   (Material.open())
	// }else if text == "close"{
	//   Material2.close()
	// }
	Material.open()
	// Material2.close()
}

// type Trainer struct {
// 	skill      string
// 	experiance string

// }

// func (t Trainer) hardskill() string {
// 	return fmt.Sprintf("%s %s", t.skill, t.experiance)
// }

// type jobdesk struct {
// 	Job string
// 	Trainer
// }

// func (d jobdesk) jobdesk() {
// 	fmt.Println("desk :", d.Job)
// 	fmt.Println("keahlian :", d.hardskill())
// }

// func main() {
// 	Train := Trainer{
// 		"Programing",
// 		"15 years",
// 	}
// 	desk2 := jobdesk{
// 		"Teach programing",
// 		Train,
// 	}
// 	desk2.jobdesk()
// }
