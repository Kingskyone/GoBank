package files

import (
	"flag"
)

var (
	timeout int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 0
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16 //两个字节
	ID          uint16 //两个字节
	SequenceNum uint16 //两个字节
}

//func main4() {
//	//var a uint16 = 33
//	//var b uint16 = 33
//	getCommonArgs()
//	desIp := os.Args[len(os.Args)-1]
//	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	defer conn.Close()
//	fmt.Println("正在向ip  "+desIp+" [", conn.RemoteAddr(), "] 发送具有"+strconv.Itoa(size)+"字节的数据数据")
//	for i := 0; i < count; i++ {
//		t1 := time.Now()
//		var icmp *ICMP = &ICMP{
//			Type:        typ,
//			Code:        code,
//			CheckSum:    0,
//			ID:          1,
//			SequenceNum: 1,
//		}
//
//		data := make([]byte, size)
//		var buffer bytes.Buffer
//		binary.Write(&buffer, binary.BigEndian, icmp) //大端写入，从左往右，针对每个变量（变量内正反）  把struct写进去
//		buffer.Write(data)                            //把data写进去
//		data = buffer.Bytes()                         //把写好的buffer变为bytes
//		checksum := checkSum(data)
//		data[2] = byte(checksum >> 8)
//		data[3] = byte(checksum)
//		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond)) //设置超时
//		n, err := conn.Write(data)
//		if err != nil {
//			log.Println(err)
//			continue
//		}
//		buf := make([]byte, 65535)
//		n, err = conn.Read(buf)
//		if err != nil {
//			log.Println(err)
//			continue
//		}
//		t2 := time.Since(t1).Milliseconds()
//		fmt.Printf("来自 %d.%d.%d.%d 的回复: 字节=%d 时间=%dms TTL=%dms\n", buf[12], buf[13], buf[14], buf[15], n-28, t2, buf[8])
//	}
//
//}

func getCommonArgs() {
	flag.Int64Var(&timeout, "w", 1000, "超时") // 定义参数
	flag.IntVar(&size, "l", 64, "缓冲区")
	flag.IntVar(&count, "n", 4, "发送次数")
	flag.Parse() // 读参数
}

// 签名   校验和
func checkSum(data []byte) uint16 {
	lenn := len(data)
	ind := 0
	sum := uint32(0)
	for lenn > 1 {
		sum += uint32(data[ind])<<8 + uint32(data[ind+1])
		lenn -= 2
		ind += 2
	}
	if lenn != 0 {
		sum += uint32(data[ind])
	}
	high16 := sum >> 16
	for high16 != 0 {
		sum = high16 + uint32(uint16(sum))
		high16 = sum >> 16
	}
	return uint16(^sum)
}
