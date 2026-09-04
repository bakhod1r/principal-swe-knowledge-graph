---
title: "Buffered Channel Ring Buffer Pointer Math"
tags:
  - review
  - golang
  - concurrency
  - principal-swe
parent: "[[Channel Architecture]]"
---
# Buffered Channel Ring Buffer Pointer Math

Go’dagi **buffered channel (`make(chan T, N)`)** ichki mexanizmini tushunishda eng muhim qismlardan biri — `sendx` va `recvx` pointer/indexlarining ring buffer ichida qanday yurishi.

---

## 1. Problem: Nega ring buffer kerak?

Buffered channel:

```go
ch := make(chan int, 4)
```

4 ta elementni receiver kutmasdan vaqtincha saqlay oladi.

Mental model:

```text
capacity = 4

        sendx
          ↓
     ┌────┬────┬────┬────┐
     │    │    │    │    │
     └────┴────┴────┴────┘
       0    1    2    3
          ↑
        recvx
```

Channel yangi elementlarni **FIFO** tartibida saqlaydi.

Ring buffer kerak, chunki `sendx` va `recvx` oxiriga yetganda yana `0` indeksiga qaytadi.

---

# 2. `hchan` dagi asosiy fields

Conceptually:

```go
type hchan struct {
    qcount   uint
    dataqsiz uint
    elemsize uint16

    buf      unsafe.Pointer
    elemsize uint16
    closed   uint32

    sendx    uint
    recvx    uint
    // ...
}
```

Buffered channel uchun ayniqsa:

```text
qcount
sendx
recvx
dataqsiz
buf
```

muhim.

### Ma'nosi

|Field|Ma'nosi|
|---|---|
|`qcount`|buffer ichidagi hozirgi elementlar soni|
|`dataqsiz`|buffer capacity|
|`sendx`|keyingi send qilinadigan slot|
|`recvx`|keyingi receive qilinadigan slot|
|`buf`|ring buffer memory boshlanish address'i|

Masalan:

```go
ch := make(chan int, 4)
```

initial:

```text
qcount = 0
capacity = 4

sendx = 0
recvx = 0
```

---

# 3. Pointer emas, index

Muhim nuance:

`sendx` va `recvx` **actual memory pointer emas**.

Ular:

```text
index / offset position
```

sifatida ishlaydi.

Masalan:

```text
capacity = 4

sendx = 2
recvx = 1
```

bu:

```text
sendx → slot #2
recvx → slot #1
```

degani.

Actual memory address esa `buf` va element size orqali hisoblanadi.

---

# 4. Pointer arithmetic

Faraz qilaylik:

```go
type T struct {
    A int64
}
```

64-bit architecture'da:

```text
sizeof(T) = 8 bytes
```

Buffer:

```text
buf = 0x1000
```

Capacity:

```text
4
```

Shunda memory:

```text
slot 0 → 0x1000
slot 1 → 0x1008
slot 2 → 0x1010
slot 3 → 0x1018
```

Formula:

```text
address = buf + index * elemsize
```

Masalan:

```text
sendx = 2

address
= 0x1000 + 2 * 8
= 0x1010
```

Shuning uchun channel runtime slot'ni topa oladi.

---

# 5. Send qanday ishlaydi?

Boshlang'ich:

```text
capacity = 4
qcount   = 0
sendx    = 0
recvx    = 0
```

Buffer:

```text
        sendx
          ↓
     ┌────┬────┬────┬────┐
     │ ∅  │ ∅  │ ∅  │ ∅  │
     └────┴────┴────┴────┘
       0    1    2    3
```

### `ch <- 10`

Element `sendx=0` slot'iga yoziladi:

```text
     ┌────┬────┬────┬────┐
     │ 10 │ ∅  │ ∅  │ ∅  │
     └────┴────┴────┴────┘
       0    1    2    3
        ↑
      recvx

             ↑
            sendx
```

Keyin:

```text
qcount = 1
sendx  = 1
```

---

### `ch <- 20`

```text
     ┌────┬────┬────┬────┐
     │ 10 │ 20 │ ∅  │ ∅  │
     └────┴────┴────┴────┘
       0    1    2    3
        ↑
      recvx

                  ↑
                sendx
```

State:

```text
qcount = 2
sendx  = 2
recvx  = 0
```

---

### `ch <- 30`

```text
     ┌────┬────┬────┬────┐
     │ 10 │ 20 │ 30 │ ∅  │
     └────┴────┴────┴────┘
       0    1    2    3
        ↑
      recvx

                       ↑
                      sendx
```

```text
sendx = 3
qcount = 3
```

---

### `ch <- 40`

```text
     ┌────┬────┬────┬────┐
     │ 10 │ 20 │ 30 │ 40 │
     └────┴────┴────┴────┘
       0    1    2    3
        ↑
      recvx

                       sendx
                       ↓
```

End:

```text
qcount = 4
sendx = 4
```

But capacity is only `4`.

So `sendx=4` is not a valid slot.

---

# 6. The wrap-around

Here is the key ring-buffer operation.

After using slot `3`:

```text
sendx++
```

gives:

```text
sendx = 4
```

Runtime wraps it back:

```text
if sendx == capacity {
    sendx = 0
}
```

Therefore:

```text
3 → 0
```

Conceptually:

```text
0 → 1 → 2 → 3
↑           ↓
└───────────┘
```

This is the **ring**.

---

# 7. Receive pointer behaves identically

Suppose:

```text
buffer:

┌────┬────┬────┬────┐
│ 10 │ 20 │ 30 │ 40 │
└────┴────┴────┴────┘
  ↑
recvx
```

Receive:

```go
x := <-ch
```

reads slot `0`:

```text
x = 10
```

Then:

```text
recvx = 1
qcount = 3
```

Now:

```text
┌────┬────┬────┬────┐
│ ∅  │ 20 │ 30 │ 40 │
└────┴────┴────┴────┘
       ↑
     recvx
```

Next receive:

```text
20
```

and:

```text
recvx = 2
```

Eventually:

```text
recvx = 3
```

then:

```text
recvx = 0
```

again.

---

# 8. The interesting case: wrapping while buffer contains data

Bu eng muhim mental model.

Faraz qilaylik:

```text
capacity = 4
```

Initially:

```text
sendx = 0
recvx = 0
```

Send:

```text
10
20
30
```

Result:

```text
┌────┬────┬────┬────┐
│ 10 │ 20 │ 30 │ ∅  │
└────┴────┴────┴────┘
  ↑              ↑
recvx           sendx
```

State:

```text
qcount = 3
recvx  = 0
sendx  = 3
```

Endi receive:

```go
<-ch
```

returns:

```text
10
```

Now:

```text
recvx = 1
qcount = 2
```

Buffer logically:

```text
┌────┬────┬────┬────┐
│ ∅  │ 20 │ 30 │ ∅  │
└────┴────┴────┴────┘
       ↑
     recvx

                       ↑
                      sendx
```

Now send:

```go
ch <- 40
```

writes slot `3`:

```text
┌────┬────┬────┬────┐
│ ∅  │ 20 │ 30 │ 40 │
└────┴────┴────┴────┘
       ↑         ↑
     recvx      sendx
```

`sendx` wraps:

```text
sendx = 0
```

Now send:

```go
ch <- 50
```

writes **slot 0**:

```text
┌────┬────┬────┬────┐
│ 50 │ 20 │ 30 │ 40 │
└────┴────┴────┴────┘
  ↑    ↑
sendx recvx
```

State:

```text
qcount = 4
sendx  = 1
recvx  = 1
```

Notice something subtle:

```text
sendx == recvx
```

yet the buffer is **full**, not empty.

That's why `qcount` is important.

---

# 9. Why `sendx == recvx` is ambiguous

In a ring buffer, this state:

```text
sendx == recvx
```

can mean:

### Empty

```text
qcount = 0
```

or:

### Full

```text
qcount = capacity
```

Therefore you cannot determine buffer state only from:

```text
sendx
recvx
```

You also need:

```text
qcount
```

or another full/empty representation.

Go uses `qcount`.

---

# 10. The core pointer math

Conceptually:

### Send slot

```text
slot = sendx

address = buf + slot * elemsize
```

then:

```text
sendx++

if sendx == capacity {
    sendx = 0
}
```

### Receive slot

```text
slot = recvx

address = buf + slot * elemsize
```

then:

```text
recvx++

if recvx == capacity {
    recvx = 0
}
```

So:

```text
sendx = (sendx + 1) % capacity
recvx = (recvx + 1) % capacity
```

is the conceptual model.

---

# 11. Lekin runtime aynan `%` ishlatishi shart emas

Bu production-level nuance.

Matematik jihatdan:

```go
sendx = (sendx + 1) % capacity
```

to'g'ri.

Lekin runtime hot path'da:

```go
if sendx == capacity {
    sendx = 0
}
```

kabi branch ishlatishi mumkin.

Nega?

Chunki:

```text
division/modulo
```

oddiy increment + compare + reset'dan qimmatroq bo'lishi mumkin.

Shuning uchun mental model:

```text
index = (index + 1) % N
```

lekin implementation:

```text
index++

if index == N {
    index = 0
}
```

bo'lishi mumkin.

---

# 12. Power-of-two capacity haqida muhim nuance

Ko'p ring-buffer implementationlar:

```text
capacity = 2^n
```

bo'lsa:

```text
index = (index + 1) & (capacity - 1)
```

ishlatishi mumkin.

Masalan:

```text
capacity = 8
mask = 7
```

```text
index = (index + 1) & 7
```

Bu modulo:

```text
(index + 1) % 8
```

bilan ekvivalent.

Ammo **Go channel capacity power-of-two bo'lishi shart emas**.

Masalan:

```go
make(chan int, 7)
```

valid.

Shuning uchun Go channel ring buffer uchun umumiy mental model:

```text
if index == capacity {
    index = 0
}
```

deb qarash yaxshiroq.

---

# 13. Pointer math va element size

Bu ham juda muhim.

Agar:

```go
ch := make(chan uint64, 4)
```

bo'lsa:

```text
elemsize = 8
```

Memory:

```text
buf
 ↓
+--------+--------+--------+--------+
| uint64 | uint64 | uint64 | uint64 |
+--------+--------+--------+--------+
    0        1        2        3
```

Address:

```text
slot 0 = buf + 0 * 8
slot 1 = buf + 1 * 8
slot 2 = buf + 2 * 8
slot 3 = buf + 3 * 8
```

Agar:

```go
type Item struct {
    A [64]byte
}
```

bo'lsa:

```text
elemsize = 64
```

va:

```text
slot 0 = buf + 0
slot 1 = buf + 64
slot 2 = buf + 128
slot 3 = buf + 192
```

Shuning uchun:

```text
index
```

va

```text
byte offset
```

bir xil narsa emas.

Formula:

```text
byte_offset = index × elemsize
```

---

# 14. Full picture

Buffered channel'ni shunday tasavvur qilish eng foydali:

```text
                     hchan
                ┌───────────────┐
                │ qcount        │
                │ dataqsiz      │
                │ sendx         │
                │ recvx         │
                │ buf ──────────┼─────┐
                └───────────────┘     │
                                      ▼
             contiguous memory
        ┌───────┬───────┬───────┬───────┐
        │ slot0 │ slot1 │ slot2 │ slot3 │
        └───────┴───────┴───────┴───────┘
             ↑                       ↑
           recvx                    sendx
```

`sendx`:

```text
write position
```

`recvx`:

```text
read position
```

`qcount`:

```text
how many valid elements exist
```

`dataqsiz`:

```text
how many slots exist
```

---

# 15. Principal-level mental model

Buffered channelni **queue + circular array + synchronization mechanism** deb o'ylash foydali:

```text
                 Buffered Channel
                       │
        ┌──────────────┼──────────────┐
        │              │              │
        ▼              ▼              ▼
    Ring Buffer     Wait Queues      Lock
        │              │
   ┌────┴────┐      ┌──┴───┐
   │ sendx   │      │sudog  │
   │ recvx   │      │queues │
   │ qcount  │      └───────┘
   └─────────┘
```

Shuning uchun channel faqat:

```text
"goroutine-safe queue"
```

emas.

U:

```text
ring-buffer storage
+
sender/receiver synchronization
+
blocking/wakeup
+
FIFO semantics
+
memory synchronization
```

ni birlashtiradi.

Va `sendx`/`recvx` pointer math — faqat **storage layer**ning bir qismi.

---

## Eng muhim 5 ta takeaway

1. **`sendx` va `recvx` actual pointers emas — ring-buffer indexlari.**
    
2. Element address'i conceptually:
    
    ```text
    buf + index × elemsize
    ```
    
3. Index capacity chegarasiga yetganda `0` ga wrap qiladi.
    
4. `sendx == recvx` ham **empty**, ham **full** bo'lishi mumkin; `qcount` ambiguity'ni yechadi.
    
5. `sendx`/`recvx` ring-buffer mexanizmi; **sudog/wait queues esa blocked goroutine synchronization layer**.
    

Keyingi natural qadam — **`hchan.buf`, `sendx`, `recvx`, `qcount`, `sendq`, `recvq` birgalikda `ch <- x` va `<-ch` vaqtida qanday state transition qilishi**. Bu channel runtime'ning eng muhim mental modelidir.
---

## 🔗 References
- ⬆️ Parent: [[Channel Architecture]]
- 📚 Module: `Concurrency & Synchronization`
