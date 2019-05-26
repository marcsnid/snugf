# snugf
Utility to protect, encrypt, and compress files.

### Install
To get latest version:  
`go get -u github.com/marcsj/snugf`  


To install, navigate to directory and use:  
`go install`

### Usage

At present, this allow for a simple use of decoupling and compressing files, 
and vice versa. This can be done by using it as follows:  


writing:  
`snugf write -k <key string> [input filename] [output filename]` 

reading:  
`snugf read -k <key string> [input filename] [output filename]`

Although [Decouplet](https://github.com/marcsj/decouplet) allows for variable-length keys, 
this key should be very similar to a password in that it contains a mixture of 
symbols, upper and lower letters, and numbers at a suitable length.