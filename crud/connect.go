package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"

)

type Device_Info struct{
        UUID    string `gorm:"column:UUID" json:"UUID"`
        Station_UUID string `gorm:"column:Station_UUID" json:"Station_UUID"`
        IP_Address string `gorm:"column:IP_Address" json:"IP_Address"`
        FW_Version string `gorm:"column:FW_Version" json:"FW_version"`
        Rebooted int `gorm:"column:Rebooted" json:"Rebooted"`
}

func (*Device_Info) TableName() string {
	return "Device_Info"
}

var Db *gorm.DB

func initDb(){
	var err error
	Db, err = gorm.Open("mssql", "sqlserver://SA:r0000t@Liteon@192.168.86.116:1433?database=losn-ti")
	if err != nil{
		fmt.Printf("connect error")
	}
	if Db.Error != nil{
		fmt.Printf("connect Db error")
	}

	Db.LogMode(true)
}
func main(){
	initDb()
	r :=gin.Default()

	r.GET("/find",FindDb)

	r.Run(":8080")
}


func FindDb(c *gin.Context){
	getDb_tmp := []Device_Info{}
	if err := Db.Find(&getDb_tmp).Error;err != nil{
                c.JSON(http.StatusOK, gin.H{
                        "code" : -1,
                        "message" : "Find nothing",
			"error" : err,
                })
	return
        }
        c.JSON(http.StatusOK, gin.H{
                "code" : 1,
                "data" : getDb_tmp,

        })
	return
}

