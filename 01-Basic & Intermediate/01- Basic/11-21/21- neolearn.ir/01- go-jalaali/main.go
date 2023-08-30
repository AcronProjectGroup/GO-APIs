package main

import (
	"fmt"
	"time"
	"github.com/jalaali/go-jalaali"
)

func main()  {
	fmt.Println(jalaali.ToJalaali(time.Now().Date()))
}

/* Output:

Starting: /home/ubuntu/go/bin/dlv dap --listen=127.0.0.1:33963 --log-dest=3 from /home/ubuntu/Desktop/Repo/Test2 YouTube
DAP server listening at: 127.0.0.1:33963
1402 مرداد 31 <nil>
Process 17976 has exited with status 0


*/