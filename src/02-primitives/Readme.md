## 1. Boolean Type
* By default value is `false`. So no need to assign a boolean variable. 
## 2. Numeric Types
* Integers
	* int
	* uint8 : (0-255)
	* uint16 : (0-65,535)
	* uint32 : (0-4,294,967,295)
	* Bitwise operations: And/Or/Xor/and not
	* If any int variable is not assigned any value, default would be `Zero(0)`
	 
* Floating point numbers
	* Follow IEEE-754 standard
	* Zero value is 0
	* float32/float64 bit versions
	* Literal Types:
	  * Decimal (3.14)
	  * Exponential (13c18 OR 2E10)
	  * Mixed (13.7e12)
	* Arithmatic Operations (+, -, x, /)
* Complex numbers
	* Zero value (0+0i)
	* 64 & 128 bit 
	* Built-in Functions
		* complex - make complex number from 2 floats
		* real - get real part of float32/float64
		* imag - get imaginary part of float
## 3. Text Types 
 * Strings:
		* UTF-8
		* Immutable
		* concatenate : (+) operator
		* convert to bytes : []byte
 * Runes:
		* UTF-32
		* Alias for int32
		* Special methods normally required to process : e.g `strings.Reader#ReadRune`
