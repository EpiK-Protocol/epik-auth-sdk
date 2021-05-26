EpiK bls signature SDK for third-part auth

### Usage

#### Sign
```
func Sign(plain []byte, epikWalletPrivateKey string) (signature []byte, addr address.Address, err error) 
```
#### Verify
```
func Verify(plain, signature []byte, addr address.Address) (err error) 
```

![EpiK_Wallet_Remote_Auth](https://user-images.githubusercontent.com/7609140/119651539-b56a3200-be57-11eb-80b3-31725c15ef1f.jpeg)
