## Testing speed of various languages

### Using MacBook 12" 2015

```
 ./run.sh
Python 2 result=330616563309

real	0m15.309s
user	0m15.132s
sys	0m0.103s
-----------------------
AWK result=330616563309

real	0m13.465s
user	0m13.282s
sys	0m0.112s
-----------------------
Python 3 result=330616563309

real	0m8.116s
user	0m7.975s
sys	0m0.091s
-----------------------
Python 3 No Unicode result=330616563309

real	0m6.244s
user	0m6.136s
sys	0m0.078s
-----------------------
*PHP 5.6* result=330616563309

real    *0m52.751s*
user    0m31.319s
sys     0m15.405s
-----------------------
PHP 7.0 result=330616563309

real    0m4.125s
user    0m2.898s
sys     0m1.095s
-----------------------
Go result=330616563309

real    0m2.460s
user    0m2.369s
sys     0m0.068s
-----------------------
C result=330616563309.000000

real    0m3.881s
user    0m3.752s
sys     0m0.075s
```
