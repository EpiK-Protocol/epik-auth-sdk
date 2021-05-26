EpiK bls signature SDK for third-part auth

### Usage

#### Sing
```
func Sign(plain []byte, epikWalletPrivateKey string) (signature []byte, addr address.Address, err error) 
```
#### Verify
```
func Verify(plain, signature []byte, addr address.Address) (err error) 
```