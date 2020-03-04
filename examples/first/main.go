package main

import "github.com/x51ming/gomb"

func main() {
	gmb.Run("../node.dll")
	w := gmb.CreateWebWindow(0, 0, 0, 0, 800, 600)
	w.ShowWindow(true)
	w.MyBindClosing()
	w.MyBindNavi()
	w.MyBindReady()
	w.MyDisablePic()
	w.LoadURL("https://www.baidu.com")
	<-w.MyBindDestroy()
	gmb.Finalize()
}
