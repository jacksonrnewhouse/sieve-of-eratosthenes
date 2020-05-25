import sys
import math

def sieve(primes, factor):
    for n in range(2 * factor, len(primes) , factor):
        primes[n] = 0

def main(argv):
    if len(argv) != 2:
        raise ValueError('Please enter exactly one argument, an integer limit'
                         ' to the size of the primes generated.')
    n = int(argv[1])
    if n < 2:
        raise ValueError('The integer limit must be greater than or equal to'
                         ' two.')
    primes = [i for i in range(n + 1)]
    primes[1] = 0
    for x in range(2,int(1 + math.sqrt(n))) :
        if primes[x] != 0:
            sieve(primes, x)
    print('\n'.join([str(p) for p in primes if p != 0]))


if __name__ == '__main__':
    main(sys.argv)
