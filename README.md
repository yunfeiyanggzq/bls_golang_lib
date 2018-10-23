# discription
it  is  a  BLS  signature  lib for  golang 
# how  to  install  the  lib
## first :install  the  GMP
This package must be compiled using cgo. It also requires the installation of GMP and PBC. During the build process, this package will attempt to include <gmp.h> and <pbc/pbc.h>, and then dynamically link to GMP and PBC.

Most systems include a package for GMP. To install GMP in Debian / Ubuntu:

`sudo apt-get install libgmp-dev`

For an RPM installation with YUM:

`sudo yum install gmp-devel`

For installation with Fink (http://www.finkproject.org/) on Mac OS X:

`sudo fink install gmp gmp-shlibs`

For more information or to compile from source, visit https://gmplib.org/

## second:install the PBC 
To install the PBC library, download the appropriate files for your system from https://crypto.stanford.edu/pbc/download.html. PBC has three dependencies: the gcc compiler, flex (http://flex.sourceforge.net/), and bison (https://www.gnu.org/software/bison/). See the respective sites for installation instructions. Most distributions include packages for these libraries. For example, in Debian / Ubuntu:

`sudo apt-get install build-essential flex bison`

The PBC source can be compiled and installed using the usual GNU Build System:
```
./configure
make
sudo make install
```

After installing, you may need to rebuild the search path for libraries:

`sudo ldconfig`

It is possible to install the package on Windows through the use of MinGW and MSYS. MSYS is required for installing PBC, while GMP can be installed through a package. Based on your MinGW installation, you may need to add` "-I/usr/local/include"` to CPPFLAGS and `"-L/usr/local/lib" `to LDFLAGS when building PBC. Likewise, you may need to add these options to` CGO_CPPFLAGS `and` CGO_LDFLAGS` when installing this package. 
## third   install  the  bls signature  golang  lib
download the bls  signature  lib  

` go  get  go  get  github.com/yunfeiyangbuaa/bls_golang_lib/BLS`

and  imprt in your code 

`import " github.com/yunfeiyangbuaa/bls_golang_lib/BLS"`
#  how to use
##   exmple
 ```
 func main() {
    	BLS.BLS_start() 
    	privKey,pubKey:=BLS.Generate_bls_keypair()    
    	signature :=BLS.Bls_signature([]byte("hello") ,privKey)      
    	sibyte:=BLS.SetSIGIntoByte(signature)     
    	sign:=BLS.SetPubKeyFromByte(sibyte)     
    	BLS.Bls_verify([]byte("hello") ,pubKey,sign)      
    }
 ```
## function 
```
  func BLS_start()                 
    Generate_bls_keypair()(*pbc.Element,*pbc.Element)　　　　
    Bls_signature(message  []byte,privkey *pbc.Element) *pbc.Element　　
    Bls_verify(message  []byte,pubkey  *pbc.Element,signature  *pbc.Element)bool　　
    func SetPriKeyIntoByte(privkey  *pbc.Element)[]byte    
    func SetPriKeyFromByte(privkey  []byte)*pbc.Element
    func SetPubKeyIntoByte(pubkey  *pbc.Element)[]byte
    func SetPubKeyFromByte(pubkey  []byte)*pbc.Element
    func SetSIGIntoByte(sig  *pbc.Element)[]byte
    func SetSIGFromByte(sig  []byte)*pbc.Element
 ```
