// tcp/client/main.go

// �ͻ���
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
			fmt.Println("err :", err)
			return
	}
	defer conn.Close() // �ر�����
	inputReader := bufio.NewReader(os.Stdin)
	for {
			input, _ := inputReader.ReadString('\n') // ��ȡ�û�����
			inputInfo := strings.Trim(input, "\r\n")
			if strings.ToUpper(inputInfo) == "Q" { // �������q���˳�
					return
			}
			_, err = conn.Write([]byte(inputInfo)) // ��������
			if err != nil {
					return
			}
			buf := [512]byte{}
			n, err := conn.Read(buf[:])
			if err != nil {
					fmt.Println("recv failed, err:", err)
					return
			}
			fmt.Println(string(buf[:n]))
	}
}