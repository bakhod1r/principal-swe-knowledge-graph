---
title: "DNS Resolvers (cgo vs netgo)"
tags:
  - review
  - golang
  - environment
  - principal-swe
parent: "[[Target OS & Architecture]]"
---
# DNS Resolvers: `cgo` vs `netgo` in Go

DNS resolution in Go has an important implementation choice:

> **Should Go resolve DNS using the native Go resolver (`netgo`) or delegate to the operating system's libc resolver through cgo (`cgo`)?**

This choice affects **portability, behavior, configuration, performance, static linking, containers, and cross-compilation**.

---

## 1. The Problem

When your Go program does:

```go
ips, err := net.LookupHost("example.com")
```

or:

```go
conn, err := net.Dial("tcp", "example.com:443")
```

Go needs to convert:

```text
example.com
    ↓
DNS resolver
    ↓
93.184.216.34
```

Go can obtain this result through two major resolver implementations:

```text
                    Go application
                          │
                     net package
                          │
              ┌───────────┴───────────┐
              │                       │
          netgo resolver          cgo resolver
              │                       │
        Go implementation       libc / OS resolver
              │                       │
          DNS directly          NSS / system libraries
```

---

# 2. `netgo`

`netgo` means:

> **Use Go's pure-Go DNS resolver.**

The resolver is implemented inside Go's `net` package rather than delegating DNS resolution to libc.

Conceptually:

```text
Application
    │
    ▼
net.LookupHost()
    │
    ▼
Go DNS resolver
    │
    ▼
/etc/resolv.conf
    │
    ▼
DNS server
```

For example:

```bash
go build -tags netgo
```

You can also influence resolver selection through the environment/configuration of the build/runtime.

### Advantages

**1. Pure Go**

No libc dependency is required for DNS resolution.

This is particularly valuable for:

- static binaries
    
- minimal containers
    
- scratch images
    
- cross-compilation
    
- predictable deployments
    

---

# 3. `cgo`

With the cgo resolver, Go delegates hostname resolution to the operating system's libc facilities.

Conceptually:

```text
Application
    │
    ▼
net.LookupHost()
    │
    ▼
cgo
    │
    ▼
libc
    │
    ▼
OS name-service configuration
    │
    ├── DNS
    ├── /etc/hosts
    ├── NSS
    ├── mDNS
    └── other mechanisms
```

The important distinction is:

> `cgo` does not simply mean "DNS through libc."

It means the resolution can participate in the **OS name-service system**, which can be significantly richer than DNS alone.

---

# 4. Why does Go have two resolvers?

Because operating systems have different expectations around name resolution.

A pure DNS implementation can understand things like:

```text
/etc/resolv.conf
```

and query DNS servers.

But an OS resolver may support additional mechanisms.

For example, Linux systems can use **NSS (Name Service Switch)**.

A conceptual `/etc/nsswitch.conf` might contain:

```text
hosts: files dns
```

meaning:

```text
/etc/hosts
    ↓
DNS
```

Another system might have:

```text
hosts: files mdns4_minimal [NOTFOUND=return] dns
```

Now hostname resolution isn't simply:

```text
DNS → answer
```

It can involve multiple system-specific mechanisms.

This is one of the fundamental reasons `cgo` exists.

---

# 5. The Most Important Mental Model

Do **not** think:

```text
netgo = good
cgo   = bad
```

That's incorrect.

Think:

```text
netgo
    = Go-controlled hostname resolution

cgo
    = OS-controlled hostname resolution
```

The trade-off is:

```text
                 Control / portability
                         ▲
                         │
                    netgo │
                         │
                         │
                         │
                         │ cgo
                         ▼
                  OS integration
```

`netgo` gives Go more control.

`cgo` gives the operating system more control.

---

# 6. How Go Chooses the Resolver

The `net` package has logic that determines whether the pure-Go or cgo resolver should be used.

A simplified mental model is:

```text
                DNS lookup
                    │
                    ▼
             resolver selection
                    │
          ┌─────────┴─────────┐
          │                   │
       netgo path          cgo path
          │                   │
          ▼                   ▼
      Go resolver           libc
```

The exact selection rules are more nuanced than this diagram.

Factors can include:

- platform
    
- whether cgo is available
    
- build tags
    
- resolver configuration
    
- environment
    
- system-specific requirements
    

Therefore, don't assume:

> "If CGO_ENABLED=1, Go always uses libc DNS."

That's not necessarily true.

---

# 7. `CGO_ENABLED` and DNS

This is especially important when discussing your previous topic around `CGO_ENABLED`.

Suppose:

```bash
CGO_ENABLED=0 go build
```

Now cgo is unavailable.

The Go program therefore cannot use the cgo resolver.

For supported configurations, the pure-Go resolver becomes the relevant path.

Conceptually:

```text
CGO_ENABLED=0
      │
      ▼
No cgo
      │
      ▼
netgo resolver
```

With:

```bash
CGO_ENABLED=1
```

the cgo resolver may be available:

```text
CGO_ENABLED=1
      │
      ├── pure Go resolver
      │
      └── cgo resolver
```

Availability of cgo **does not mean cgo DNS is mandatory**.

---

# 8. Why This Matters for Containers

Consider a minimal container:

```dockerfile
FROM scratch

COPY app /app

ENTRYPOINT ["/app"]
```

Your application is a static Go binary.

There may be no:

```text
glibc
libnss_*
```

or other userspace infrastructure.

A pure-Go resolver is therefore attractive.

```text
┌─────────────────────────────┐
│ scratch container           │
│                             │
│  /app                       │
│     │                       │
│     ▼                       │
│  Go net resolver            │
│     │                       │
│     ▼                       │
│  DNS server                 │
└─────────────────────────────┘
```

This is one reason Go is popular for small container images.

---

# 9. But `/etc/resolv.conf` Still Matters

A common misconception is:

> "netgo means Go ignores the OS DNS configuration."

No.

The pure-Go resolver can use system resolver configuration, particularly:

```text
/etc/resolv.conf
```

For example:

```text
nameserver 10.96.0.10
search my-namespace.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
```

This is extremely important in Kubernetes.

A Go application using the pure-Go resolver can still resolve:

```text
postgres
```

or:

```text
postgres.default.svc.cluster.local
```

through the cluster DNS configuration.

---

# 10. Kubernetes Example

Suppose your pod receives:

```text
/etc/resolv.conf
```

containing:

```text
nameserver 10.96.0.10
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
```

Your application does:

```go
net.Dial("tcp", "postgres:5432")
```

The resolver may construct/search names according to the resolver configuration.

Conceptually:

```text
postgres
   │
   ├── postgres.default.svc.cluster.local
   ├── postgres.svc.cluster.local
   ├── postgres.cluster.local
   └── postgres
```

This is one reason DNS behavior inside Kubernetes deserves explicit testing rather than assuming "it's just DNS."

---

# 11. `netgo` vs `cgo`

|Property|`netgo`|`cgo`|
|---|---|---|
|Implementation|Go|libc / OS|
|Requires cgo|No|Yes|
|Static binary friendliness|Excellent|More complicated|
|Cross-compilation|Easier|Harder|
|OS integration|Limited|Strong|
|NSS integration|No|Yes|
|`/etc/resolv.conf`|Yes|Yes|
|Portability|High|Lower|
|Minimal containers|Excellent|Can be problematic|
|Behavior|More predictable|More OS-dependent|

The key trade-off is:

```text
netgo
  → portability + predictability

cgo
  → OS integration + system resolver behavior
```

---

# 12. A Critical Difference: NSS

This is probably the most important technical distinction.

Imagine:

```text
/etc/nsswitch.conf

hosts: files dns
```

The OS resolver can effectively do:

```text
hostname
   │
   ▼
/etc/hosts
   │
   ├── found → return
   │
   └── not found
          │
          ▼
         DNS
```

With additional NSS modules, the system can support even more resolution mechanisms.

Therefore:

```text
cgo resolver
    ↓
OS name service abstraction
```

whereas:

```text
netgo resolver
    ↓
Go's DNS-oriented implementation
```

This distinction becomes important in enterprise environments where hostname resolution may not be plain DNS.

---

# 13. `/etc/hosts`

Both approaches can interact with local host configuration, but their exact behavior and integration path differ.

For example:

```text
/etc/hosts

10.10.10.20 database.internal
```

The OS resolver can process this through its configured name-service mechanism.

This is another reason you shouldn't test DNS behavior only against public DNS names.

Production tests should include:

```text
public DNS
internal DNS
/etc/hosts
search domains
timeouts
NXDOMAIN
SERVFAIL
temporary DNS failure
```

---

# 14. Static Linking

This is where the resolver choice becomes particularly relevant.

Suppose you want:

```text
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
```

and:

```text
scratch
└── app
```

Pure Go is naturally aligned with this architecture.

With cgo:

```text
Go
 │
 ▼
cgo
 │
 ▼
libc
 │
 ├── NSS
 ├── shared libraries
 └── system resolver
```

Now your "static" deployment may no longer be as self-contained as you expected.

This is why:

> **"My Go binary is statically linked" does not automatically mean "all runtime dependencies are eliminated."**

You must reason about libc, NSS, DNS, certificates, timezone data, and other runtime resources separately.

---

# 15. Cross Compilation

This connects directly to your previous CGO cross-compilation discussion.

Pure Go:

```bash
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=arm64 \
go build
```

is straightforward.

With cgo:

```bash
CGO_ENABLED=1
```

you need a compatible C cross-compiler/toolchain.

For example:

```text
build machine
    │
    ▼
cross compiler
    │
    ▼
target libc
    │
    ▼
target binary
```

This creates additional compatibility dimensions:

```text
Go version
+
target architecture
+
target OS
+
C compiler
+
target libc
+
libc headers
+
libraries
+
linker
```

That is a significant operational cost.

---

# 16. Performance

It's tempting to say:

> "netgo is faster."

Don't.

That's an engineering smell unless you've measured it.

DNS performance depends on:

- cache behavior
    
- resolver configuration
    
- DNS server latency
    
- network RTT
    
- number of queries
    
- search domains
    
- IPv4/IPv6 behavior
    
- libc implementation
    
- Go resolver implementation
    
- connection reuse
    
- application workload
    

The DNS network round trip often dominates the actual resolver implementation overhead.

Therefore:

```text
resolver implementation overhead
             vs
       DNS network latency
```

is often not the dominant factor.

Benchmark your actual workload.

---

# 17. Operational Predictability

This is where I generally favor pure Go for cloud-native services when OS-specific name-service integration isn't required.

Imagine deploying the same binary to:

```text
Ubuntu
Alpine
Debian
Distroless
scratch
Kubernetes
VM
container
```

With `netgo`, the resolver behavior is more controlled by Go.

With cgo, behavior can depend on:

```text
OS
+
libc
+
NSS configuration
+
NSS modules
+
/etc/nsswitch.conf
+
/etc/resolv.conf
```

Therefore:

```text
cgo DNS
     ↓
more environmental variables
     ↓
more possible behavior differences
```

This is not inherently bad—but it increases the system's dependency surface.

---

# 18. Failure Modes

A production engineer should think beyond successful lookups.

### DNS timeout

```text
application
    ↓
DNS query
    ↓
no response
    ↓
timeout
```

What happens next?

If your HTTP client has no timeout, you can accumulate blocked goroutines/connections.

DNS is therefore part of your application's **failure boundary**.

---

### DNS server unavailable

```text
App
 │
 ▼
DNS
 X
```

Your service may suddenly be unable to:

```text
connect to database
call downstream service
resolve service discovery names
```

DNS failure can therefore become a cascading outage.

---

### Search-domain explosion

A configuration like:

```text
search a.example.com b.example.com c.example.com
options ndots:5
```

can cause additional DNS queries for short names.

At scale:

```text
100k requests/sec
        ×
multiple DNS attempts
        ↓
huge DNS query volume
```

This is why Kubernetes DNS behavior deserves capacity planning.

---

# 19. Security Considerations

DNS is part of your trust boundary.

Consider:

```text
service.example.com
        ↓
DNS
        ↓
10.10.10.20
```

If DNS infrastructure is compromised or misconfigured, the application can be directed somewhere unexpected.

Important controls include:

- trusted DNS infrastructure
    
- network policy
    
- DNSSEC where appropriate
    
- TLS certificate validation
    
- avoiding blind trust in resolved IPs
    
- preventing SSRF through DNS rebinding scenarios
    
- careful handling of internal/private DNS
    

Most importantly:

> **DNS resolution does not authenticate the destination.**

TLS authentication still matters.

---

# 20. A Subtle Issue: DNS Rebinding / SSRF

Suppose an application accepts:

```text
https://example.com
```

and resolves:

```text
example.com → public IP
```

Later the DNS answer changes:

```text
example.com → 169.254.x.x
```

or another internal address.

If the application uses hostname resolution naively, DNS becomes part of the SSRF attack surface.

This is not specifically a `netgo` vs `cgo` problem.

It's an **application-level DNS trust problem**.

---

# 21. How to Force `netgo`

For a build where you explicitly want the pure-Go resolver:

```bash
go build -tags netgo
```

This is useful when you want to make the deployment characteristic explicit.

For example:

```bash
CGO_ENABLED=0 \
go build -tags netgo \
  -o app .
```

The important engineering idea isn't memorizing the command.

It's recognizing why you might deliberately choose:

```text
pure Go
    ↓
fewer runtime dependencies
    ↓
more reproducible deployment
```

---

# 22. How to Force `cgo`

The corresponding build tag is:

```bash
go build -tags cgo
```

but this requires cgo support.

Conceptually:

```bash
CGO_ENABLED=1 go build -tags cgo
```

Whether you **should** do this is a separate question.

Don't enable cgo simply because "Linux uses libc."

Enable it when your application actually needs the OS resolver behavior.

---

# 23. Debugging DNS Behavior

When debugging production DNS, don't immediately change code.

Follow:

```text
Reproduce
   ↓
Observe
   ↓
Identify resolver
   ↓
Inspect configuration
   ↓
Observe DNS traffic
   ↓
Form hypothesis
   ↓
Change one variable
   ↓
Verify
```

Useful things to inspect:

```bash
cat /etc/resolv.conf
```

```bash
cat /etc/nsswitch.conf
```

and:

```bash
getent hosts example.com
```

Compare that with:

```bash
nslookup example.com
```

or:

```bash
dig example.com
```

Then test your Go program.

The key question is:

> **Does the application observe the same resolver behavior as the OS tools?**

If not, you may be dealing with `netgo` vs `cgo` differences.

---

# 24. Observability

DNS should be observable when it is operationally important.

At minimum, monitor:

```text
DNS resolution latency
DNS errors
timeouts
NXDOMAIN
SERVFAIL
application connection failures
```

For example:

```text
dns_lookup_duration_seconds
dns_lookup_errors_total
```

You should correlate:

```text
DNS latency ↑
      │
      ├── HTTP latency ↑
      ├── DB connection latency ↑
      └── request timeout ↑
```

This allows you to distinguish:

```text
application problem
```

from:

```text
DNS infrastructure problem
```

---

# 25. Production Recommendation

For a typical Go backend running in:

- Kubernetes
    
- Docker
    
- cloud VMs
    
- Linux production
    
- minimal containers
    

I would generally start with:

```text
pure Go resolver
```

unless you have a concrete requirement for OS-level name-service integration.

Why?

```text
netgo
 ├── no libc dependency for DNS
 ├── simpler cross-compilation
 ├── works well with minimal containers
 ├── predictable deployment
 └── aligns with Go's portability model
```

But for environments that depend on:

```text
NSS
LDAP-backed host resolution
mDNS
special OS resolver behavior
enterprise name-service integration
```

then:

```text
cgo resolver
```

may be the correct engineering choice.

---

# 26. Decision Tree

Use this mental model:

```text
Do I need OS-specific name-service behavior?
              │
          ┌───┴───┐
          │       │
         YES      NO
          │       │
          ▼       ▼
         cgo     netgo
```

Then:

```text
Need minimal container?
        │
       YES
        │
        ▼
Prefer netgo
```

```text
Need easy cross compilation?
        │
       YES
        │
        ▼
Prefer netgo
```

```text
Need NSS / OS resolver integration?
        │
       YES
        │
        ▼
Consider cgo
```

The key word is **consider**, not automatically **use**.

---

# 27. Principal Engineer Mental Model

The important hierarchy is:

```text
CGO_ENABLED
     │
     ├── controls whether cgo can be used
     │
     ▼
Resolver selection
     │
     ├── netgo
     │
     └── cgo
             │
             ▼
          libc / OS
             │
             ▼
        NSS / resolver
```

Don't collapse these concepts into one.

They are separate:

```text
CGO_ENABLED
     ≠
cgo DNS resolver
     ≠
libc
     ≠
DNS server
     ≠
/etc/resolv.conf
     ≠
NSS
```

Understanding those boundaries is the important part.

---

## Key Takeaways

1. **`netgo` = Go's pure-Go DNS resolver.**
    
2. **`cgo` = delegation to the OS/libc name-service machinery.**
    
3. `CGO_ENABLED=1` does **not** mean every DNS lookup necessarily uses cgo.
    
4. `CGO_ENABLED=0` removes cgo, making the pure-Go path important.
    
5. `netgo` is attractive for **static binaries, containers, and cross-compilation**.
    
6. `cgo` is valuable when you need **OS-specific resolver/NSS integration**.
    
7. `/etc/resolv.conf` remains important for pure-Go DNS.
    
8. DNS failures are **production failure modes**, not merely configuration problems.
    
9. Don't choose based on assumed performance—**measure**.
    
10. The correct question is not _"Which resolver is better?"_ but:
    

> **"Do I need OS-specific name-service behavior, or do I value a self-contained and predictable Go runtime?"**

For modern cloud-native Go services, that question often leads to **`netgo` by default, `cgo` by requirement**.
---

## 🔗 References
- ⬆️ Parent: [[Target OS & Architecture]]
- 📚 Module: `Go Environment & Commands`
