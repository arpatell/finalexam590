# COMP 590-059 Final Exam - Aaron Patel

## Two: Go Concurrency Problem

1. **What argument can you give that there will be no deadlock among these processes?**  
   Deadlock is prevented through the use of waitgroups and closing channels between stages. Producers run concurrently and their completion is notified when the waitgroup is done. Then in a separate goroutine, inCh is closed only after the producers are done running. Now, consumers loop over inCh and square the results to outCh, and they have a separate waitgroup that notifies completion. In a separate goroutine, outCh is closed only after the consumers are done running. Finally, there is a filter stage, which reads from outCh and only takes larger values after the first one in outCh. It too has its own waitgroup and notifies completion. After the filter stage is done, the results are printed out. The structure in this program doesn't allow for goroutines to be left waiting. The use of these waitgroups and closing the channels in between stages ensures that all the senders are done, thus there is no deadlock.


2. **We get "easy" fan-out and fan-in here with Go channels, because a channel is a data structure that is not bound to any single process. How would we handle this situation in Elixir, where the mailbox semantics in actors mean that our "channel" is always attached to a single process? How do we have 2 producers creating and distributing values to 2 consumers?**  
   Go channel functionality can be recreated in Elixir with a centralized dispatcher process and message forwarding. Two producers can send messages to the dispatcher, which can forward the messages of both producers to two consumers, distributed in a systematic or even random order. This would be the fan-out part of the process. Then, two consumers receive the messages in their respective mailboxes and can concurrently send them to a separate process that would basically perform the same function as the `filter()` routine. This is the fan-in part of the process.

## Running the Program
1. **Navigation**

   ```bash
    cd two
   ```
2. **Run**

   ```bash
    go run concurrrent-pipeline.go
   ```