## Testing speed of various languages

### How to run this:
- first generate BIGDATA2 file `./generate-BIGDATA2.sh`
- install this if you want to run all tests `brew install php@5.6 python@3.6`
- run benchmark: `./run.sh`

### Using MacBook 15" 2015. 2.8 GHz Intel Core i7, 16GB RAM

Python 2 result=327281433600 <- Very bad...

real	**1m48.775s**
user	0m54.106s
sys	0m54.386s
-----------------------

AWK result=327281433600

real	**0m8.611s**
user	0m8.568s
sys	0m0.034s
-----------------------

Python 3 result=327281433600

real	**0m5.191s**
user	0m5.107s
sys	0m0.056s
-----------------------

Python 3 No Unicode result=327281433600

real	**0m3.897s**
user	0m3.857s
sys	0m0.034s
-----------------------

PHP 5.6 result=327281433600

real	**0m14.869s**
user	0m13.412s
sys	0m1.426s
-----------------------

PHP 7.0 result=327281433600 <- Second place

real	**0m2.198s**
user	0m1.557s
sys	0m0.617s
-----------------------

Go result=327281433600   <- Winner!

real	**0m1.301s**
user	0m1.294s
sys	0m0.034s
-----------------------

C result=327281433600.000000

real	**0m2.036s**
user	0m2.016s
sys	0m0.018s


### Using MacBook 12" 2015

Python 2 result=330616563309

real	**0m15.309s**
user	0m15.132s
sys	0m0.103s
-----------------------
AWK result=330616563309

real	**0m13.465s**
user	0m13.282s
sys	0m0.112s
-----------------------
Python 3 result=330616563309

real	**0m8.116s**
user	0m7.975s
sys	0m0.091s
-----------------------
Python 3 No Unicode result=330616563309

real	**0m6.244s**
user	0m6.136s
sys	0m0.078s
-----------------------
*PHP 5.6* result=330616563309

real    **0m52.751s**
user    0m31.319s
sys     0m15.405s
-----------------------
PHP 7.0 result=330616563309

real    **0m4.125s**
user    0m2.898s
sys     0m1.095s
-----------------------
Go result=330616563309

real    **0m2.460s**
user    0m2.369s
sys     0m0.068s
-----------------------
C result=330616563309.000000

real    **0m3.881s**
user    0m3.752s
sys     0m0.075s

