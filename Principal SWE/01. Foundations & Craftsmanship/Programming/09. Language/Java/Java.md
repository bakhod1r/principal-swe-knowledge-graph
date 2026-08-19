---
title: Java
tags:
  - programming
  - java
  - language-mastery
  - principal-swe
parent: "[[Language]]"
---

# 💻 Java (Production Mastery & Engineering Deep Dive)

Enterprise Java 21+ platform engineering: HotSpot JVM internals, JIT compilation (C1/C2), garbage collection (ZGC, G1), Java Memory Model (JMM), Project Loom Virtual Threads, Collections Framework, Generics type erasure, Spring Boot ecosystem, and high-throughput concurrent architectures.

```text
Java
│
├── [[Learn the Basics in Java|01. Learn the Basics]]
├── [[Object Oriented Programming in Java|02. Object Oriented Programming]]
├── [[Java Generics, Wildcards, and Type Erasure|03. Generics and Type Erasure]]
├── [[Java Exception Handling and Try with Resources|04. Exception Handling Architecture]]
├── [[Java Lambdas, Method References, and Streams API|05. Lambda Expressions and Streams API]]
├── [[Java Annotations, Reflection, and Metaprogramming|06. Annotations and Reflection]]
├── [[Java Module System (jpms) and Strong Encapsulation|07. Java Module System JPMS]]
├── [[Java Optionals and Modern Null Safe Design|08. Optionals and Null Safety]]
├── [[Java Collections Framework and Data Structures|09. Java Collections Framework]]
├── [[Dependency Injection and Inversion of Control in Java|10. Dependency Injection in Java]]
├── [[Java Concurrency, Locks, and Virtual Threads (loom)|11. Concurrency and Virtual Threads]]
├── [[Java Cryptography Architecture (jca) and Security|12. Cryptography and Security JCA]]
├── [[Java Modern Date and Time API (java.time)|13. Modern Date and Time API]]
├── [[Java Networking, Sockets, and Http Client|14. Networking and Socket Programming]]
├── [[Java Regular Expressions (java.util.regex)|15. Regular Expressions in Java]]
├── [[Java Io and Modern NIO2 File Operations|16. Io and NIO2 File Operations]]
├── [[Functional Programming Patterns and Immutability in Java|17. Functional Programming Patterns]]
├── [[Java Build Tools (maven and Gradle Ecosystem)|18. Build Tools Maven and Gradle]]
├── [[Enterprise Spring Boot and Microservices|19. Enterprise Web Frameworks Spring Boot]]
├── [[Database Access (jdbc, Hikaricp, and JPA Hibernate)|20. Database Access JDBC and Hibernate JPA]]
├── [[Testing and Mocking in Java (junit 5, Mockito)|21. Testing and Mocking in Java]]
└── [[Java Logging Architecture (slf4j, Logback, Log4j2)|22. Logging Frameworks Slf4j and Logback]]
```

---

## 🗂️ Core Knowledge Pillars

- 📂 [[Learn the Basics in Java|01. Learn the Basics]] — Java source compilation, JVM bytecode lifecycle, primitive types, control flow, and variable scopes.
- 📂 [[Object Oriented Programming in Java|02. Object Oriented Programming]] — Encapsulation, abstract classes, interfaces, inheritance, polymorphism, and records.
- 📂 [[Java Generics, Wildcards, and Type Erasure|03. Generics and Type Erasure]] — Generic classes/methods, bounded wildcards (PECS: Producer Extends Consumer Super), and type erasure.
- 📂 [[Java Exception Handling and Try with Resources|04. Exception Handling Architecture]] — Checked vs Unchecked exceptions, try-with-resources with AutoCloseable, and custom exceptions.
- 📂 [[Java Lambdas, Method References, and Streams API|05. Lambda Expressions and Streams API]] — Functional interfaces (Function, Predicate, Consumer), Stream pipelines, intermediate/terminal operations, and collectors.
- 📂 [[Java Annotations, Reflection, and Metaprogramming|06. Annotations and Reflection]] — Runtime vs compile-time annotations, java.lang.reflect, dynamic proxies, and framework metadata processing.
- 📂 [[Java Module System (jpms) and Strong Encapsulation|07. Java Module System JPMS]] — module-info.java, requires, exports, open modules, and service provider interfaces (SPI).
- 📂 [[Java Optionals and Modern Null Safe Design|08. Optionals and Null Safety]] — Optional<T> best practices, avoiding null returns, and modern pattern matching with switch expressions.
- 📂 [[Java Collections Framework and Data Structures|09. Java Collections Framework]] — List (ArrayList), Set (HashSet), Map (HashMap, ConcurrentHashMap), Deque, and internal bucket structures.
- 📂 [[Dependency Injection and Inversion of Control in Java|10. Dependency Injection in Java]] — Jakarta CDI, Spring IoC container, bean lifecycles, and component scanning.
- 📂 [[Java Concurrency, Locks, and Virtual Threads (loom)|11. Concurrency and Virtual Threads]] — synchronized, ReentrantLock, java.util.concurrent (Executors, ConcurrentHashMap), and Project Loom virtual threads.
- 📂 [[Java Cryptography Architecture (jca) and Security|12. Cryptography and Security JCA]] — Cipher encryption (AES), MessageDigest, KeyStore, SSL/TLS contexts, and secure random generators.
- 📂 [[Java Modern Date and Time API (java.time)|13. Modern Date and Time API]] — Instant, LocalDate, LocalDateTime, ZonedDateTime, Duration, Period, and timezone handling.
- 📂 [[Java Networking, Sockets, and Http Client|14. Networking and Socket Programming]] — Socket, ServerSocket, java.net.http.HttpClient, non-blocking NIO channels, and Netty.
- 📂 [[Java Regular Expressions (java.util.regex)|15. Regular Expressions in Java]] — Pattern compilation, Matcher methods, regex optimization, and multi-line capture groups.
- 📂 [[Java Io and Modern NIO2 File Operations|16. Io and NIO2 File Operations]] — InputStream, OutputStream, java.nio.file.Files, Path, FileChannel, and memory-mapped ByteBuffers.
- 📂 [[Functional Programming Patterns and Immutability in Java|17. Functional Programming Patterns]] — Pure functions, persistent immutable data structures, Vavr library, and Monadic patterns.
- 📂 [[Java Build Tools (maven and Gradle Ecosystem)|18. Build Tools Maven and Gradle]] — Maven pom.xml lifecycles, Gradle build scripts, dependency resolution, and multi-module builds.
- 📂 [[Enterprise Spring Boot and Microservices|19. Enterprise Web Frameworks Spring Boot]] — Spring Boot auto-configuration, Spring Web, Spring Data, Actuator metrics, and microservice patterns.
- 📂 [[Database Access (jdbc, Hikaricp, and JPA Hibernate)|20. Database Access JDBC and Hibernate JPA]] — Direct JDBC, HikariCP connection pooling, Hibernate ORM entity mappings, dirty checking, and caching.
- 📂 [[Testing and Mocking in Java (junit 5, Mockito)|21. Testing and Mocking in Java]] — JUnit 5 lifecycle, Mockito mocking, AssertJ fluent assertions, and Testcontainers Docker testing.
- 📂 [[Java Logging Architecture (slf4j, Logback, Log4j2)|22. Logging Frameworks Slf4j and Logback]] — Logging facades, log levels, asynchronous appenders, and Mapped Diagnostic Context (MDC) tracing.

---

## 🔗 References
- ⬆️ Parent: [[Language]]

