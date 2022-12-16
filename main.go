package main

import (
	"fmt"
	"todo/activity"
	"todo/config"
	"todo/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var activMenu = activity.NewActivityMenu(conn)

	for inputMenu != 0 {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var newUser user.User
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&newUser.Password)
			res, err := authMenu.Register(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("Sukses mendaftarkan data")
			} else {
				fmt.Println("Gagal mendaftarn data")
			}
		} else if inputMenu == 2 {
			var inputNama, inputPassword string
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&inputNama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&inputPassword)
			res, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}

			if res.ID > 0 {
				isLogin := true
				loginMenu := 0
				for isLogin {
					fmt.Println("1. Tambah Aktivitas")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan menu : ")
					fmt.Scanln(&loginMenu)
					if loginMenu == 1 {
						inputActivity := activity.Activity{}
						var inputString string
						fmt.Print("Masukkan Judul Kegiatan : ")
						fmt.Scanln(&inputString)
						inputActivity.SetTitle(inputString)
						fmt.Print("Masukkan Lokasi: ")
						fmt.Scanln(&inputString)
						inputActivity.SetLocation(inputString)
						fmt.Print("Masukkan Due Date: ")
						fmt.Scanln(&inputString)
						inputActivity.SetCreateDate(inputString)
						inputActivity.SetOwner(res.ID)
						actRes, err := activMenu.Insert(inputActivity)
						if err != nil {
							fmt.Println(err.Error())
						}
						inputActivity.SetID(actRes)
						fmt.Println(inputActivity)
					} else if loginMenu == 9 {
						isLogin = false
					}
				}
			}
		}
	}
}
