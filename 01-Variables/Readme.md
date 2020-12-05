

### 1.  Declare variables:
*    Declare a variable & initialize later
```sh
  var count int   		 
  var count int = 42 
```      
 *  Declare & initialize a variable
```sh 
  count := 42
```

###  2. Cant re-declare a variable, but can shadow the variable
###  3. All declared variables must be used. Else code will throw compilation error
###  4. Visibility:
* lower-case first letter: Package scope e.g `var i int = 32` 
* UPPER-CASE first letter: Global scope & can be exported out   e.g `var I int = 33`
* `No Provate scope`, ie to the package. However we can define `Variable block`
###  5. Naming Convensions:
* Caps for generic variables - (HTTP, URL)
* short var name - shorter lifespan
* Larger names - Longer Lifespan  e.g `var sanitizer string = "ClusterRole"` --> All along the code 
###  6. Type Conversations:
* DestinationType(var):
* `Int->Float` :- `float32(i int)`
* No implecit type Conversion. So we need to handle manually
* Internal type Conversion functions can be used. e.g  `Int <--> String` : Use Function `StringConv`
