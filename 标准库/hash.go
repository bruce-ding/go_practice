package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"hash/fnv"
	"io"
	"github.com/OneOfOne/xxhash"
)

func shaHash(pwd, salt string, iterations int) string {
	data := []byte(pwd + salt)

	hash := sha512.Sum512(data)
	if iterations > 1 {
		for i := 1; i < iterations; i++ {
			hash = sha512.Sum512(hash[:])
		}
	}
	return fmt.Sprintf("%x", hash)
}

func md5Hash(pwd, salt string) {
	//方法一
	data := []byte(pwd + salt)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)

	//方法二
	w := md5.New()
	io.WriteString(w, pwd + salt)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))

	fmt.Println(md5str2)
}

func main() {
	str := shaHash("bnajks", "jjs", 2)
	fmt.Println(str)
	fmt.Println(len(str))

	md5Hash("jjjaww", "klwo")

	h := fnv.New32a().Sum([]byte("test"))
	str1 := fmt.Sprintf("%x", h) //将[]byte转成16进制
	fmt.Println(str1)

	h := xxhash.New64()
	// r, err := os.Open("......")
	// defer f.Close()
	r := strings.NewReader(F)
	io.Copy(h, r)
	fmt.Println("xxhash.Backend:", xxhash.Backend)
	fmt.Println("File checksum:", h.Sum64())

	nodeLeft := 22
	a := -1 ^ (-1 << uint64(nodeLeft))
	fmt.Println(a)
	fmt.Println(a == 1 << uint64(nodeLeft))
	fmt.Println(-1 << uint64(nodeLeft))
	fmt.Println(-1 ^ 1)
	fmt.Println(-1 ^ 2)
	fmt.Println(-1 ^ 3)
	fmt.Println(-1 ^ 4)
	fmt.Println(1 << uint64(nodeLeft))
}