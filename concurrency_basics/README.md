> # Concurrency in golang

- enables the program to execute multiple tasks at the same time in out of order fashion without effecting other parts of the program.

> ## 1. utilizing goroutine(to spawn a concurrent process)

- go funcCall() > will spawn a new process

> ## 2. golang channels

- helps in synchronizing all of the spawned processes with goroutines, i.e communication between spawned/concurrent goroutines

                        // their are 2 goroutines in golang
                        main     task

                        # main
                        gets kicked off when the main function of golang program executes
                        # task
                        goroutine gets kicked off when we mention go keyword to spawn process

                        main        task
                         |
                         |
                         |
                         |
                         |------------| (go keyword mentioned spawn goroutine process)
                         |            |
                         |            |
                         |            |
                         |            |
                         finished     |
                                      |
                                      |(goroutine spawned process still executing while the main function has finished execution)

- The aim is to make sure the main function do not finish execution until the goroutine spawned task is completed.

> ### Way-1 (Naive-not performant) (use of shared memory)

                        main         task
                         |
                         |
                         |
                         |
                         |------------ | (go keyword mentioned spawn goroutine process)
                         |             |
                         |<---(SM)     |
                         |     |       |
                         |     |       |
                               |       |
                               |       |
                               |<------|(finished stores something(boolean value) in shared memory(SM))

then the main process constantly checks for the shared memory wheather subroutine finished

> ### Way-2 (the right way) use of Channels (based on pub-sub architecture)

                        main                 task
                         |
                         |
                         |
                         |
                         |--------------------| (go keyword mentioned spawn goroutine process)
                         |                    |
                         |<---(Channels)      |
                         |      |             |
                         |      |             |
                         |      |             |
                         |      |             |
                         |      |<------      |(finished , publishes boolean value in |                     channel)
                         |(check task finished? yes -> terminates main goroutine)

Note- their is no need to periodically check for the boolean value as the task itself is going to publish the boolean value i.e inform the main goroutine when its finished executing. # syntax
channelName := make(chan type)

                        # example- my channel with boolean type variable
                        myChannel := make(chan bool)

> **IMPORTANT- "<- works as put or get message from channel in golang" | ref: https://go.dev/tour/concurrency/2**

> ## 3. golang buffered-channels(to avoid deadlocks)

- avoiding deadlock

the channels are by default 0 capacity, so whenever a value is update in channel by task(spawned goroutine) then it is necessary that variable is received by main(goroutine).

NOTE- also its not possible to emit and receive the variable of channel in main bcz while emit variable on the channel the program is still on the line of updating value of the variable in goroutine.

- creating buffered channel (i.e increasing the capacity)

                    // a channle with string variable and capacity 1
                    channel := make(chan string, 1)

> **channels are based on FIFO(queue) rear addition & front deletion**

> ## 4. golang channels iteration & channel closing

            channel := make(chan string)
            message,open := <-channel
            if(!open){
                  // channel is closed
            }
            # close the channel explicitely
            close(channel)

> **IMPORTANT Additional Cases where deadlock can occur in channels**

- trying to access out of bound message in channel i.e their were 3 messages pushed but trying to access 4 messages during message iteration without closing the connection or using the for range method approach

- improper logic of trying to fill more than the default capacity of channel if not using the buffered channels in the first place.

> ## 5. golang channel select statements

- select statement restricts a goroutine to wait on multiple communication operations

- select blocks until one of its cases can run, then it executes that case. it chooses 1 at random if multiple cases are ready.

> ## 5. golang wait group (alternative to channel in go routines synchronization)

- the go routines main & sub task statements execution is non-deterministic( no specific order of execution)

- if main go routine gets wrapped up early then the sub go routine call wont even get a chance to get executed

- **to tackle this non-deterministic behaviour we use waitGroup that makes the go routines statements executions deterministic in main & sub task**

> ### Some drawbacks of using waitGroup

1. if we call Wait() immediately after .Add() method then it may result in deadlock as an indefinate waiting state may happen.

2. the correct sequence is .Add(), .Done() [inside subgoroutine] , Wait() [inside the main goroutine]

3. also the calling of the methods Done() & Wait() should be exactly the number specified while mentioning the Add() method for how many goroutines we are going to wait for to avoid deadlock.

> ## 6. Mutex/RWMutex (resource-control variable)

**AIM- Mutex aim is to implement locking/unlocking mechanism to make sure that resource access is democratic among all part of system during processes i.e at any given time only 1 thread/ goroutine can actually access the resource**

- suppose their is 1 resource that needs to be accessed by multiple parts of the system to avoid cases like no one getting the resource access at all or some part of the program got access while other kept on waiting indefinately we use mutex.

- this mutex is kinda control variable whichever part of the program number is equal to the mutex variable that system part will be allowed to access the resource only at that time.

- in addition to that each part of the system only will have a limited time to hold the mutex with them i.e for the time that mutex value is equal to the number of the system part.

> **Some imp terms**

- locking request the intention to hold the mutex by other part of system.

- when resource is of no use to the part of system then they can release the mutex i.e unlocking it.

> ## Important- RWMutex

                  sync.RWMutex

> Reader-Writer lock is a mutual exclusiion lock, the lock can held by an arbitary number of readers or a single writer. so only one subgoroutine allowed to write/update the variable while any number of sub goroutines can read that.

> the zero value for RWMutex is an unlocked mutex.

> also if a goroutine holds a RWMutex for reading & another goroutine might call lock, no goroutine can be expect to acquire the lock.

**ref(code): https://go.dev/src/sync/rwmutex.go**
**ref(theory):https://pkg.go.dev/sync#RWMutex.Lock**

> ## 7. "Once" [from sync] i.e DO IT ONCE

- object that perform action only once
- Once is used by multiple sub go routines so that each go routine execute exactly 1 task and it makes sure if their task overlaps then exactly 1 subroutine does the task and in scenario more than 1 go routine complete that task to avoid wasting time.

                  # A CASE STUDY
                  # FIND GOOD MANGOES

                  # a situation where multiple siblings go to market to find good quality mangoes


                  #now each of these sibling act as a sub goroutine that are sent one after the another to find the mangoes

                  #now we need to make sure if any sibling finds the mangoes then we stop sending any more siblings

- to resolve the case where multiple siblings might find good mangoes instead whenever exactly 1 sibling found mangoe they infor all other sibling so the others can take chill pill.

> ## 8. Resource Pool [sync.Pool]

- a pool is hoarding or accumalation of fixed number of things that can be used later
- pool is used to limit the things/oprns/flows that are expensive i.e requires high resource values.

                  # sync.Pool

                  instance- like sub goroutine
                  resources- method, variable, memory, storage etc...

                  1. get()
                  will check wheather their are available resources for a instance in the pool & returns i.e picks that resource and allocate to instance who made the call, if resource not available then it calls the New[func method] to create a new one.

                  2. put()
                  helps the instance to return the resource back to the pool once the instance is done doing their task

> ## 9. Signal & Broadcasting [synch.Cond]

- Cond implements a condition variable , a trigger point for go routines waiting for or announcing the occurence of an event.

- each Cond has an associated Locker L(a *Mutex or *RWMutex) which must be held when changing the condition and when calling the Wait method

- the lock and unlock of cond is based on the semaphore that is inspired from linux futex, where futex is used for blocking mech in kernel for shared-resources context.
  ref: https://man7.org/linux/man-pages/man2/futex.2.html

> ## 10. Map safety and performance (sync.Map) NOTE_USE_WITH_CAUTION

ref: sync_Map.go

- note overall syncMaps method and for loops are stable for stable keys in comparison to using regular maps with mutex.

- syncMaps have better advantage performance wise when multiple concurrent go routines are performing oprns like store,load i.e get,put or get&put together in a loop.

- for single calls or operations mutex and regular maps combo shud be enough

> ## 11. 📝 Atomic (ATOMIC in sync) NOTE_USE_WITH_CAUTION

- Atomic provides low level memory primatives used for implementing synchronization atomic operations so that race condition can be avoided.

- ref: Atomic.go

- 📝Atomic methods i.e .Load & .Store are limited to certain data types, to counter this we have atomic.values
