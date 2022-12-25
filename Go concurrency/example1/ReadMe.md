 This is bad approach to run Go Routine or Concurrency
 go printSomething("This is my first message!")
 time.Sleep(1 * time.Second)
 printSomething("This is my second message!")


Using waitGroup in GoRoutines
1. Creating my variable for waitGroup
var wg sync.WaitGroup

2. Add number of words to wait for
wg.Add(len(words))

3. set wg.Wait() afte the go funtion of the routines
4. set a wg sync.WaitGroup pointer as parameter to the function calling the wait group 
   like func (wg *sync.WaitGroup){
    defer wg.Done(); this decrements the number of time the go routine runs 
   }


