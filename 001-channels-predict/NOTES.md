# 001 — My predictions

## Expected output (before running)
1."read 1: vm=cow-1 cpus=2 ok=true\n"
2."read 2: vm=cow-2 cpus=2 ok=true\n"
3."read 3: vm=cow-3 cpus=4 ok=true\n"
4."read 4: vm="" cpus="" ok=false\n"
5."read 5: vm="" cpus="" ok=false\n"

## Why lines 4 and 5
When we try to print the 4th iteration the channel is already close so the application, than it returns an OK false.


## Bonus: with `for j := range jobs`
It would just iterate as many times as there are jobs in the channel and immediately break out of the loop. 

## Surprises after running
