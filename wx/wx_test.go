package wx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"testing"
	"fmt"
)

var (
	ErrAppIDNotMatch       = errors.New("app id not match")
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

type WXBizDataCrypt struct {
	appID, sessionKey string
}

func NewWXBizDataCrypt(appID, sessionKey string) *WXBizDataCrypt {
	return &WXBizDataCrypt{
		appID:      appID,
		sessionKey: sessionKey,
	}
}

// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

func (w *WXBizDataCrypt) Decrypt(encryptedData, iv string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(w.sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

var data="fNAlX+z0guJPBcfBhnOmrhQDJldvCQ9LWQCSXWNKT7I79o4DKiBrLBPgNYrFtpctoLXaFaM2Dok9rN6UwuEHgAc59h3R8f7XK2R+xEocPVWomOFOPzIKl1zO86cVUUXu3lrs1OnpZk2KoTN/2/v/6oRb4DpGJa0OU9/Q2umR33aGaa3WR/BfjPWLpHNrQ8UA6vpISItsziWisdvTvG6Jf638t2/xL6Lf6lJKssfNkxJ8xjRQnijkUL7pqb+2b00qLIe8Jx4fvbhcJNPwxK3ugXB/94KjIJ6HHuRDXG11Jzt4M4w+xtcgN9vs7uGp2lsOOInLB54QLnCQEdedfRuz3fuISBhx9yVlpjBqnNzCbyAyZa+FiMTtQAf377A/LxmtBv4qp/3pJkcTDrOie7Ad+A3vD5AAUriMQY0Yioe+nCYxSenhER9OqaGST16QZb+hBpuHeYFB/Y5RNGWiy2Eq77w0ZBDtXmZeS3aN+HDvuN/SBQ10ce6DEAHYEKNJiVEmDWJmc/LlXCviO2zgLdl3GGM6jwECoe8TE/4lGFxNVp0PAxxaMSVTS73ZG8k8pTpVZxsSjySPs8IY2Q9zUMjWBygyA46wjzTyC66/tkm15NnxWA/2nWQwF/OmWRqF3STScUJgX7W2Dm+tP5hxvh24f49WiOLbzSZ21b8hKI2gg1gqkovnIOdB+GCSDwmT4Xo8s7z3AehOJ3ioN6m1st7fEOGubtGtj9Lmb77Pku+ZOTp78TeSxxsammzg8Tyji9/8m3UIeHBTBrdototcrjX6dV3W19b9rF/vKQ4uNCPMnqVqxdiFlmd3182p9mzVK1JdU1c70mgUBT5hhW2CR/jnpbf99CxHnumlZ9K8XjwUurk0jXCp0fVRmBt7WO/cyFdPMxXmjuawGUYL249HiPsAPRSRNc4DGwbHiPpVEr+F+Gn+UyumK11x7EAwvjbnn5Z7KPmQlnISM3zItxTRW9ID+Rv4fjGAM+TlK3jGvKh7AWg5FaPlscOJnVB/Fubhc0qoRrtjkNWHlgxy4GV9acBZPJno3OTry8M9EEyP6G7UFvXCc62x7gkNyh3Fy5s389tpBuIUMsLan/0cd+daxc8PipfsfPs9flaYDeWJAxalmdFSb5ooz/0kqmY2bUtkF/d4bxFv/jiz4idw3StMM6f83mSx3gMeRzXJYmZx1x0LVFL+m5wOoOexzy92H0fgQreHpuM+yeOxUsEHiZIBa/16sz3tmV2Q8CPO5wlHpbHcZ976zwTEyL75Qpc3JdKhN+Qh0J0iTuPxwm7AtVAjms1vz1sAmSzLUEnQnH1coOiSiNJw+kKEPhBf+f833vWjHE/Bg9tK+vscSlDIv070RDAxuPvThleIICyVFcM44MrZsbsaTSaFCiV4zpmDr5d6GjMqPUZrKTDRcK14PuJKxaeyuxDiZXxvPwHcC7T8MhtMtphAj/MlrIYdEw6I2jwL74Q3qbhBYCvq8XJvjCqtTRK3OAHUiYA8in8KXn3yYw12+GTZCU+OPu17LjkguG/06V2WDFgo7Ba1LKz+kfhFsnPUXFqqUqnrHq6M5zocvojNvhcryPNy6D3fAngP4UKrxP/X//uqdxtdQtYrLk6haJFH+aE4FQDf+2q5SakiWQEthXZRTSzrIsRwjBcWo8eITqHa"
var iv="GxU5zfW7l7aZSJ2k20SbnA=="


func TestWX(t *testing.T){
	bytes,err:=NewWXBizDataCrypt("wxe8633ed42df5ea99","rWIG3jSAcFLGCXCHPmpkEA==").Decrypt(data,iv)
	if err!=nil{
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
