#define N 5
#define right(i) i
#define left(i) (i+1) % 5

chan forks[N] = [1] of { bool }

proctype fork(int i) {
    bool dummy
    do
    ::
        forks[i] ! true
        forks[i] ? dummy
    od
}
proctype philosopher(int i) {
    bool dummy
    do
    ::
        /* thinking */
        forks[right(i)] ? dummy
        forks[left(i)] ? dummy
        /* eating */
        forks[right(i)] ! true
        forks[left(i)] ! true
    od
}

init {
    atomic {
        byte i
        for (i : 0..4) {
            run fork(i)
            run philosopher(i)
        }
    }
}
