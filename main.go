package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

var (
	db               = make(map[string]string) // map: called in registry funct
	letters          = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	serviceparameter = make(map[string]string)
)

func randSeq(n int) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func main() {

	e := echo.New()                    //New instances
	e.POST("/iclock/cdata", cdataPOST) // Whenever this path will recv post request this function is going to handle
	e.GET("/iclock/cdata", cdataGET)
	e.GET("/iclock/ping", ping)
	e.POST("/iclock/registry", registry)
	e.GET("/iclock/push", pushGET)
	e.POST("/iclock/push", pushPOST)
	e.GET("/iclock/getrequest", getrequest)
	e.POST("/iclock/devicecmd", devicecmd)
	e.POST("/iclock/querydata", querydata)

	echo.NotFoundHandler = func(c echo.Context) error {
		uinput := c.Request().URL //  http.URL, userinput
		msg := fmt.Sprintf("not found page %s", uinput)
		fmt.Printf("\n ERROR 404 \n")
		// render your 404 page
		return c.String(http.StatusNotFound, msg)
	}
	e.Logger.Fatal(e.Start(":90")) //Start HTTP server on port 90

}

func cdataPOST(c echo.Context) error { //argument to function is an object of echo.context
	fmt.Printf("\n /cdata POST || post device's state to server \n")
	tab := c.QueryParam("table") //! check for table values bellow !
	//sn:= c.QueryParam("SN")      //Query param will extract the value from the query parameter
	//fmt.Println(sn, tab)
	var words []string
	if tab == "rtlog" {
		fmt.Printf("\n /cdata type=rtlog  || post device's event to server \n")
		var bodyBytes []byte //byte array means string converted in integers
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body) //_ means to ignore
		}
		bodyString := string(bodyBytes)    //convert array of bytes to stirng separating w.r.t white spaces
		words = strings.Fields(bodyString) // convert string to array of string
		//fmt.Println(words)
		time := words[0] + " " + words[1]
		mask := words[11]
		temp := words[12]
		fmt.Printf(" %s,  %s °C,  %s", time, temp, mask)
		ContentAsUTF8 := "OK"
		ReasonPhrase := "OK"
		StatusCode := "200"
		data := fmt.Sprintf("ContentAsUTF8:%s, ReasonPhrase:%s,StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
		fmt.Printf("\n%s\n", data)
		//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
		return c.String(http.StatusOK, data)
		//return c.String(http.StatusOK, "Message recvd") //200,oepration done

	} else if tab == "rtstate" {
		fmt.Printf("\n /cdata POST type=rtstate  || post device's state to server \n")
		var bodyBytes []byte //byte array means string converted in integers
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body) //_ means to ignore
		}
		bodyString := string(bodyBytes) //convert array of bytes to stirng separating w.r.t white spaces
		//time=2020-06-11 10:02:28    sensor=02      relay=00        alarm=0100000000000000  door=01
		words = strings.Fields(bodyString)
		time := words[0] + " " + words[1]
		sensor := words[2]
		relay := words[3]
		alarm := words[4]
		door := words[5]
		fmt.Printf(" %s,  %s,  %s,  %s,  %s", time, sensor, relay, alarm, door)
		ContentAsUTF8 := "OK"
		ReasonPhrase := "OK"
		StatusCode := "200"
		data := fmt.Sprintf("ContentAsUTF8:%s, ReasonPhrase:%s,StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
		fmt.Printf("\n%s\n", data)
		//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
		return c.String(http.StatusOK, data)
		//return c.String(http.StatusOK, "Message recvd") //200,oepration done

	} else if tab == "tabledata" {
		fmt.Printf("\n /cdata type=tabledata || post device's event to server \n")
		tabname := c.QueryParam("tablename") //! check for table values bellow !
		if tabname == "user" {
			ContentAsUTF8 := "OK"
			ReasonPhrase := "OK"
			StatusCode := "200"
			data := fmt.Sprintf("ContentAsUTF8: %s ReasonPhrase: %s StatusCode: %s\nUser=%s ", ContentAsUTF8, ReasonPhrase, StatusCode, tabname)
			//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
			return c.String(http.StatusOK, data)
		}
		rt := fmt.Sprintf("\n User not found\n")
		fmt.Printf("\n%s\n", rt)
		return c.String(http.StatusOK, rt)

	} else {
		//return following http response
		ContentAsUTF8 := "OK"
		ReasonPhrase := "OK"
		StatusCode := "200"
		data := fmt.Sprintf("ContentAsUTF8: %s ReasonPhrase: %s StatusCode: %s", ContentAsUTF8, ReasonPhrase, StatusCode)
		//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
		fmt.Printf("\n%s\n", data)
		return c.String(http.StatusOK, data)

	}

}
func cdataGET(c echo.Context) error {
	fmt.Printf("\n /cdata GET  || post device's state to server \n")
	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8: %s ReasonPhrase: %s StatusCode: %s \nServer: Echo-server\nContent-Length: 2\nDate: Wed, 17 Jun 2020 02:20:19 GMT ", ContentAsUTF8, ReasonPhrase, StatusCode)
	fmt.Printf("\n%s\n", data)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	return c.String(http.StatusOK, data)
}
func ping(c echo.Context) error {
	//sn:= c.QueryParam("SN")      //Query param will extract the value from the query parameter
	//fmt.Println(sn)
	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s ReasonPhrase:%s StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
	fmt.Printf("\nPing ok\n")
	fmt.Printf("\n%s\n", data)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	return c.String(http.StatusOK, data)
	//return c.String(http.StatusOK, "Message recvd") //200,oepration done

}

func registry(c echo.Context) error {

	fmt.Printf("\n registry  || Start to register \n")
	sn := c.QueryParam("SN") //Query param will extract the value from the query parameter
	//finding registry if present already w.r.t SN:
	var bodyBytes []byte //byte array means string converted in integers
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body) //_ means to ignore
	}
	bodyString := string(bodyBytes) //convert array of bytes to stirng separating w.r.t white spaces
	//words = strings.Fields(bodyString) //space seperated array words[]
	fmt.Printf("\n%s\n", bodyString)
	//rand.Seed(time.Now().UnixNano())
	Registrycode := randSeq(10) //random words of length 10
	_, ok := db[sn]             //check for sn in maps/db
	if ok {
		val := db[sn] //registry code from maps
		fmt.Printf("\n Device is already registered with SN: %s and registry code:%s", sn, val)
		data := fmt.Sprintf("RegistryCode=%s", val)
		fmt.Printf("\n%s\n", data)
		//return c.JSON(http.StatusOK, data) //key doesn't exist
		return c.String(http.StatusOK, data) //returns 200 ok and code
	}
	db[sn] = Registrycode //store random no and SN in db/map
	fmt.Printf("\n Device is now registered with SN: %s  against registry code:%s", sn, Registrycode)
	data := fmt.Sprintf("RegistryCode=%s", Registrycode)
	fmt.Printf("\n%s\n", data)
	return c.String(http.StatusOK, data)
}

func pushGET(c echo.Context) error {
	fmt.Printf("\n             Device push get request:            \n")
	var bodyBytes []byte //byte array means string converted in integers
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body) //_ means to ignore
	}
	bodyString := string(bodyBytes) //convert array of bytes to stirng separating w.r.t white spaces
	//words = strings.Fields(bodyString) //space seperated array words[]
	fmt.Printf("\n%s\n", bodyString)

	ServerVersion := "HTTP 1.1"
	ServerName := "Echo"
	PushVersion := "1"
	ErrorDelay := "1"
	RequestDelay := "1ms"
	TransTimes := "1"
	TransInterval := "1ms"
	TransTables := "1ms"
	Realtime := "1ms"
	SessionID := "1.1.1"
	TimeoutSec := "1.1"
	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s ReasonPhrase:%s StatusCode:%s: \nServerVersion=%s \nServerName=%s \nPushVersion=%s \nErrorDelay=%s \nRequestDelay=%s \nTransTimes=%s \nTransInterval=%s \nTransTables=%s \nRealtime=%s \nSessionID=%s \nTimeoutSec=%s", ContentAsUTF8, ReasonPhrase, StatusCode, ServerVersion, ServerName, PushVersion, ErrorDelay, RequestDelay, TransTimes, TransInterval, TransTables, Realtime, SessionID, TimeoutSec)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	fmt.Printf("\n%s\n", data)
	return c.String(http.StatusOK, data)
	//return c.String(http.StatusOK, "Message recvd") //200,oepration done
}
func pushPOST(c echo.Context) error {
	fmt.Printf("\n             Device push post request:            \n")
	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s ReasonPhrase:%s StatusCode:%s: \n Server:Echo-server \n Content-Type: text/plain;charset=ISO-8859-1 \n Content-Length: 218 \n Date: Wed, 17 Jun 2020 01:31:59 GMT \n ServerVersion = 10.2 \n ServerName = myServerName \n PushVersion = 5.6 \n ErrorDelay = 30 \n RequestDelay = 3 \n TransTimes = 00:30    13:00 \n TransInterval = 1 \n TransTables = User    Transaction \n Realtime = 1 \n SessionID = CEE878EF1C5DD54A5CB8A8DED554B3DB \n", ContentAsUTF8, ReasonPhrase, StatusCode)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	fmt.Printf("\n%s\n", data)
	return c.String(http.StatusOK, data)
	//return c.String(http.StatusOK, "Message recvd") //200,oepration done
}

func getrequest(c echo.Context) error {
	fmt.Printf("\n device say : give me instructions \n")
	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s, ReasonPhrase:%s,StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	return c.JSON(http.StatusOK, data)

}
func devicecmd(c echo.Context) error {
	fmt.Printf("\n /devicecmd  || return the result of executed command to server \n")

	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s, ReasonPhrase:%s,StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	return c.JSON(http.StatusOK, data)
}
func querydata(c echo.Context) error {
	fmt.Printf("\n/querydata  || response the server with person data that server asked\n")

	ContentAsUTF8 := "OK"
	ReasonPhrase := "OK"
	StatusCode := "200"
	data := fmt.Sprintf("ContentAsUTF8:%s, ReasonPhrase:%s,StatusCode:%s", ContentAsUTF8, ReasonPhrase, StatusCode)
	//combine := ContentAsUTF8 + " " + ReasonPhrase + " " + StatusCode
	return c.JSON(http.StatusOK, data)
}
