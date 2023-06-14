package main


func main() {
	sendToChannel := make(chan string)

	sinaaa := "sina"

	go func() {sendToChannel<-sinaaa }()

	giveFromChannel := <-sendToChannel

	print(giveFromChannel, "\n")


}